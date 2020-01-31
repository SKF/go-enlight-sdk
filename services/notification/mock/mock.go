package mock

import (
	"context"
	"time"

	"github.com/SKF/go-enlight-sdk/services/notification"
	"github.com/SKF/proto/common"
	proto "github.com/SKF/proto/notification"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type client struct {
	mock.Mock
}

func Create() *client { // nolint: golint
	return new(client)
}

var _ notification.NotificationClient = &client{}

func (mock *client) SetRequestTimeout(d time.Duration) {
	mock.Called(d)
}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}

func (mock *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *client) DialUsingCredentials(sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *client) DialUsingCredentialsWithContext(ctx context.Context, sess *session.Session, host, port, secretKey string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, sess, host, port, secretKey, opts)
	return args.Error(0)
}

func (mock *client) Close() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) SetNotificationType(notificationType proto.NotificationType) error {
args := mock.Called(notificationType)
return args.Error(0)
}
func (mock *client) SetNotificationTypeWithContext(ctx context.Context, notificationType proto.NotificationType) error {
	args := mock.Called(ctx, notificationType)
	return args.Error(0)
}

func (mock *client) GetNotificationType(name string) (proto.NotificationType, error) {
	args := mock.Called(name)
	return args.Get(0).(proto.NotificationType), args.Error(1)
}
func (mock *client) GetNotificationTypeWithContext(ctx context.Context, name string) (proto.NotificationType, error) {
	args := mock.Called(ctx, name)
	return args.Get(0).(proto.NotificationType), args.Error(1)
}

func (mock *client) InitiateNotification(notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	args := mock.Called(notificationType, resource, header, body, createdBy)
	return args.String(0), args.Error(1)
}
func (mock *client) InitiateNotificationWithContext(ctx context.Context, notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	args := mock.Called(ctx, notificationType, resource, header, body, createdBy)
	return args.String(0), args.Error(1)
}

func (mock *client) SetUserPreferences(prefs []proto.UserPreference) error {
	args := mock.Called(prefs)
	return args.Error(0)
}
func (mock *client) SetUserPreferencesWithContext(ctx context.Context, prefs []proto.UserPreference) error {
	args := mock.Called(ctx, prefs)
	return args.Error(0)
}

func (mock *client) GetUserPreferences(userID string) ([]proto.UserPreference, error) {
	args := mock.Called(userID)
	return args.Get(0).([]proto.UserPreference), args.Error(1)
}
func (mock *client) GetUserPreferencesWithContext(ctx context.Context, userID string) ([]proto.UserPreference, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]proto.UserPreference), args.Error(1)
}

func (mock *client) GetUserNotifications(userID string, limit int32) ([]proto.UserNotification, error) {
	args := mock.Called(userID, limit)
	return args.Get(0).([]proto.UserNotification), args.Error(1)
}

func (mock *client) GetUserNotificationsWithContext(ctx context.Context, userID string, limit int32) ([]proto.UserNotification, error) {
	args := mock.Called(ctx, userID, limit)
	return args.Get(0).([]proto.UserNotification), args.Error(1)
}
