package notification

import (
	"context"
	"net"
	"time"

	"github.com/SKF/go-enlight-sdk/v2/interceptors/reconnect"
	"github.com/SKF/go-utility/v2/log"
	"github.com/SKF/proto/v2/common"
	proto "github.com/SKF/proto/v2/notification"
	"github.com/aws/aws-sdk-go/aws/session"
	"google.golang.org/grpc"
)

type client struct {
	conn           *grpc.ClientConn
	api            proto.NotificationClient
	requestTimeout time.Duration
}

type NotificationClient interface {
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error)
	DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error
	Close() error
	SetRequestTimeout(d time.Duration)

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	SetNotificationType(notificationType proto.NotificationType) error
	SetNotificationTypeWithContext(ctx context.Context, notificationType proto.NotificationType) error

	GetNotificationType(name string) (proto.NotificationType, error)
	GetNotificationTypeWithContext(ctx context.Context, name string) (proto.NotificationType, error)

	RemoveNotificationType(name string) error
	RemoveNotificationTypeWithContext(ctx context.Context, name string) error

	InitiateNotification(notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error)
	InitiateNotificationWithContext(ctx context.Context, notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error)

	GetInitiatedNotification(externalId string) (proto.InitiatedNotification, error)
	GetInitiatedNotificationWithContext(ctx context.Context, externalId string) (proto.InitiatedNotification, error)

	RemoveInitiatedNotification(externalId string) error
	RemoveInitiatedNotificationWithContext(ctx context.Context, externalId string) error

	SetUserPreferences(prefs []proto.UserPreference) error
	SetUserPreferencesWithContext(ctx context.Context, prefs []proto.UserPreference) error

	GetUserPreferences(userID string) ([]proto.UserPreference, error)
	GetUserPreferencesWithContext(ctx context.Context, userID string) ([]proto.UserPreference, error)

	RemoveUserPreferences(userID string, notificationTypeExtId string) error
	RemoveUserPreferencesWithContext(ctx context.Context, userID string, notificationTypeExtId string) error

	GetUserNotifications(userID string, limit int32) ([]proto.UserNotification, error)
	GetUserNotificationsWithContext(ctx context.Context, userID string, limit int32) ([]proto.UserNotification, error)

	RemoveUserNotifications(userID string, initatedNotificationExtId string) error
	RemoveUserNotificationsWithContext(ctx context.Context, userID string, initatedNotificationExtId string) error

	UpdateUserNotificationStatus (notificationId string, state proto.SendState) error
	UpdateUserNotificationStatusWithContext (ctx context.Context, notificationId string, state proto.SendState) error
}

func CreateClient() NotificationClient {
	return &client{
		requestTimeout: 60 * time.Second,
	}
}

func (c *client) SetRequestTimeout(d time.Duration) {
	c.requestTimeout = d
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(host, port string, opts ...grpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialWithContext(ctx, host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = proto.NewNotificationClient(conn)
	return
}

// DialUsingCredentials creates a client connection to the given host with background context and no timeout
func (c *client) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DialUsingCredentialsWithContext(ctx, sess, host, port, secretKey, opts...)
}

// DialUsingCredentialsWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	var newClientConn reconnect.NewConnectionFunc
	newClientConn = func(invokerCtx context.Context, invokerConn *grpc.ClientConn, invokerOptions ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
		credOpt, err := getCredentialOption(invokerCtx, sess, host, secretKey)
		if err != nil {
			log.WithTracing(invokerCtx).WithError(err).Error("Failed to get credential options")
			return invokerCtx, invokerConn, invokerOptions, err
		}

		reconnectOpts := grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithNewConnection(newClientConn),
		))

		dialOpts := append(opts, credOpt, reconnectOpts, grpc.WithBlock())
		newConn, err := grpc.DialContext(invokerCtx, net.JoinHostPort(host, port), dialOpts...)
		if err != nil {
			log.WithTracing(invokerCtx).WithError(err).Error("Failed to dial context")
			return invokerCtx, invokerConn, invokerOptions, err
		}
		_ = invokerConn.Close()

		c.conn = newConn
		c.api = proto.NewNotificationClient(c.conn)
		return invokerCtx, c.conn, invokerOptions, err
	}

	opt, err := getCredentialOption(ctx, sess, host, secretKey)
	if err != nil {
		return err
	}

	reconnectOpts := grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
		reconnect.WithNewConnection(newClientConn),
	))
	newOpts := append(opts, opt, reconnectOpts)

	conn, err := grpc.DialContext(ctx, host+":"+port, newOpts...)
	if err != nil {
		return err
	}

	c.conn = conn
	c.api = proto.NewNotificationClient(c.conn)

	return nil
}

func (c *client) Close() (err error) {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}
