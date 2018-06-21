package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"github.com/SKF/go-utility/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func WithBlock() grpc.DialOption {
	return grpc.WithBlock()
}

func FailOnNonTempDialError(f bool) grpc.DialOption {
	return grpc.FailOnNonTempDialError(f)
}

func WithTransportCredentials(serverName, clientCert, clientKey, caCert string) (opt grpc.DialOption, err error) {
	certificate, err := tls.LoadX509KeyPair(
		clientCert,
		clientKey,
	)

	if err != nil {
		log.WithField("error", err).
			Error("Failed to load client certs")
		return
	}

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.WithField("error", err).
			Error("Failed to read ca cert")
		return
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		err = errors.New("failed to append certs")
		log.Error(err.Error())
		return
	}

	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	opt = grpc.WithTransportCredentials(transportCreds)
	return
}
