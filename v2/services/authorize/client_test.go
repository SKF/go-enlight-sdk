package authorize_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize"
	authMock "github.com/SKF/go-enlight-sdk/v2/services/authorize/mock"
	"github.com/SKF/go-utility/v2/log"
	"github.com/SKF/proto/v2/common"
	"go.uber.org/zap"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func clientForFakeHostName(t *testing.T, authority, host, port string, servers ...*authMock.AuthorizeServer) authorize.AuthorizeClient {
	hostnameWithAuthority := fmt.Sprintf("dns://%s/%s", authority, host)

	for i := range servers {
		servers[i].On("LogClientState", mock.Anything, mock.Anything).Return(&common.Void{}, nil).Maybe()
	}

	client := authorize.CreateClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	require.NoError(t, client.DialWithContext(ctx, hostnameWithAuthority, port, grpc.WithInsecure()))

	return client
}

type dnsServer struct {
	dns.Server
	domain  string
	records map[string][]string
	host    string
	port    string
}

func NewDNS(t *testing.T, domain string) *dnsServer {
	host := "127.0.0.53"
	port := getPort(t)

	return &dnsServer{
		Server: dns.Server{
			Addr: net.JoinHostPort(host, port),
			Net:  "udp",
		},
		domain:  domain,
		records: make(map[string][]string),
		host:    host,
		port:    port,
	}
}

func (d *dnsServer) Start(t *testing.T) {
	dns.HandleFunc(d.domain+".", d.handleDnsRequest)
	go func() {
		require.NoError(t, d.Server.ListenAndServe())
	}()
	time.Sleep(50 * time.Millisecond)
}

func (d *dnsServer) Stop(t *testing.T) {
	require.NoError(t, d.Server.Shutdown())
}

func (d *dnsServer) AddEndpoint(domain, ip string) {
	key := domain + "."
	if _, ok := d.records[key]; !ok {
		d.records[key] = []string{}
	}
	d.records[key] = append(d.records[key], ip)
}

func (d *dnsServer) Authority() string {
	return net.JoinHostPort(d.host, d.port)
}

func (d *dnsServer) handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range m.Question {
			switch q.Qtype {
			case dns.TypeA:
				ips := d.records[q.Name]
				log.WithFields(log.Fields{
					zap.String("name", q.Name),
					zap.Strings("records", ips),
				}).Info("dns returning")
				for ip := range ips {
					rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ips[ip]))
					if err == nil {
						m.Answer = append(m.Answer, rr)
					}
				}
			}
		}
	}

	if err := w.WriteMsg(m); err != nil {
		log.WithError(err).Error("dns server failed to write reply")
	}
}

func getPort(t *testing.T) string {
	l, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	_, port, err := net.SplitHostPort(l.Addr().String())
	require.NoError(t, err)
	require.NoError(t, l.Close())
	return port
}

func Test_Loadbalancing(t *testing.T) {
	backends := 5

	domain := "foobar.com"
	dnsServer := NewDNS(t, domain)
	dnsServer.Start(t)
	defer dnsServer.Stop(t)

	serverPort := getPort(t)

	servers := make([]*authMock.AuthorizeServer, backends)
	for i := range servers {
		ip := fmt.Sprintf("127.0.0.%d", i+1)
		server, err := authMock.NewServerOnHostPort(ip, serverPort)
		require.NoError(t, err)
		dnsServer.AddEndpoint(domain, ip)
		server.On("DeepPing", mock.Anything, mock.Anything).Return(&common.PrimitiveString{Value: ""}, nil).Once()
		servers[i] = server
	}

	client := clientForFakeHostName(t, dnsServer.Authority(), domain, serverPort, servers...)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	for i := 0; i < backends; i++ {
		err := client.DeepPingWithContext(ctx)
		require.NoError(t, err)
	}

	for i := range servers {
		servers[i].AssertExpectations(t)
	}
}

func init() {
	grpclog.SetLoggerV2(Logger{Logger: log.Base()})
}

type Logger struct {
	log.Logger
}

func (l Logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l Logger) Infoln(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...))
}

func (l Logger) Warning(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l Logger) Warningln(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l Logger) Warningf(format string, args ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, args...))
}

func (l Logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l Logger) Errorln(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...))
}

func (l Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l Logger) Fatalln(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, args...))
}

func (l Logger) V(level int) bool {
	return true
}
