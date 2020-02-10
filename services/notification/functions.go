package notification

import (
	"context"

	"github.com/SKF/proto/common"
	proto "github.com/SKF/proto/notification"
)


func (c *client) SetNotificationType(notificationType proto.NotificationType) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.SetNotificationTypeWithContext(ctx, notificationType)
}
func (c *client) SetNotificationTypeWithContext(ctx context.Context, notificationType proto.NotificationType) error {
	input := proto.SetNotificationTypeInput{
		NotificationType: &notificationType,
	}
	_, err := c.api.SetNotificationType(ctx, &input)

	return err
}

func (c *client) GetNotificationType(name string) (proto.NotificationType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetNotificationTypeWithContext(ctx, name)
}
func (c *client) GetNotificationTypeWithContext(ctx context.Context, name string) (proto.NotificationType, error) {
	input := proto.GetNotificationTypeInput{
		Name: name,
	}
	out, err := c.api.GetNotificationType(ctx, &input)

	return *out.NotificationType, err
}

func (c *client) RemoveNotificationType(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveNotificationTypeWithContext(ctx, name)
}
func (c *client) RemoveNotificationTypeWithContext(ctx context.Context, name string) error {
	input := proto.RemoveNotificationTypeInput{
		Name: name,
	}
	_, err := c.api.RemoveNotificationType(ctx, &input)

	return err
}


func (c *client) InitiateNotification(notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.InitiateNotificationWithContext(ctx, notificationType, resource, header, body, createdBy)
}
func (c *client) InitiateNotificationWithContext(ctx context.Context, notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	input := proto.InitiateNotificationInput{
		Type:      &notificationType,
		Resource:  &resource,
		Header:    header,
		Body:      body,
		CreatedBy: createdBy,
	}
	output, err := c.api.InitiateNotification(ctx, &input)
	if err != nil {
		return "", err
	}

	return output.ExternalId, nil
}

func (c *client) GetInitiatedNotification(externalId string) (proto.InitiatedNotification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetInitiatedNotificationWithContext(ctx, externalId)
}
func (c *client) GetInitiatedNotificationWithContext(ctx context.Context, externalId string) (proto.InitiatedNotification, error) {
	input := proto.GetInitiatedNotificationInput{
		ExternalId: externalId,
	}
	output, err := c.api.GetInitiatedNotification(ctx, &input)
	if err != nil {
		return proto.InitiatedNotification{}, err
	}

	return *output.InitiatedNotification, nil
}

func (c *client) RemoveInitiatedNotification(externalId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveInitiatedNotificationWithContext(ctx, externalId)
}
func (c *client) RemoveInitiatedNotificationWithContext(ctx context.Context, externalId string) error {
	input := proto.RemoveInitiatedNotificationInput{
		ExternalId: externalId,
	}
	_, err := c.api.RemoveInitiatedNotification(ctx, &input)

	return err
}


func (c *client) SetUserPreferences(prefs []proto.UserPreference) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.SetUserPreferencesWithContext(ctx, prefs)
}
func (c *client) SetUserPreferencesWithContext(ctx context.Context, prefs []proto.UserPreference) error {
	input := proto.SetUserPreferencesInput{
		Preferences: []*proto.UserPreference{},
	}

	for _, pref := range prefs {
		input.Preferences = append(input.Preferences, &pref)
	}

	_, err := c.api.SetUserPreferences(ctx, &input)
	return err
}

func (c *client) GetUserPreferences(userID string) ([]proto.UserPreference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserPreferencesWithContext(ctx, userID)
}
func (c *client) GetUserPreferencesWithContext(ctx context.Context, userID string) ([]proto.UserPreference, error) {
	result := []proto.UserPreference{}
	input := proto.GetUserPreferencesInput{UserId: userID}
	output, err := c.api.GetUserPreferences(ctx, &input)
	if err != nil {
		return result, err
	}

	for _, pref := range output.Preferences {
		result = append(result, *pref)
	}
	return result, nil
}

func (c *client) RemoveUserPreferences(userID, notificationTypeExtId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveUserPreferencesWithContext(ctx, userID, notificationTypeExtId)
}
func (c *client) RemoveUserPreferencesWithContext(ctx context.Context, userID, notificationTypeExtId string) error {
	input := proto.RemoveUserPreferencesInput{
		UserId: userID,
		NotificationTypeExtId:notificationTypeExtId,
	}
	_, err := c.api.RemoveUserPreferences(ctx, &input)

	return err
}


func (c *client) GetUserNotifications(userID string, limit int32) ([]proto.UserNotification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserNotificationsWithContext(ctx, userID, limit)
}
func (c *client) GetUserNotificationsWithContext(ctx context.Context, userID string, limit int32) ([]proto.UserNotification, error) {
	input := proto.GetUserNotificationsInput{
		UserId: userID,
		Limit:  limit,
	}

	result := []proto.UserNotification{}
	output, err := c.api.GetUserNotifications(ctx, &input)
	if err != nil {
		return result, err
	}

	for _, el := range output.Notifications {
		result = append(result, *el)
	}

	return result, nil
}

func (c *client) RemoveUserNotifications(userID, initiatedNotificationExternalId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveUserNotificationsWithContext(ctx, userID, initiatedNotificationExternalId)
}
func (c *client) RemoveUserNotificationsWithContext(ctx context.Context, userID, initiatedNotificationExternalId string) error {
	input := proto.RemoveUserNotificationsInput{
		UserId:                          userID,
		InitiatedNotificationExternalId: initiatedNotificationExternalId,
	}
	_, err := c.api.RemoveUserNotifications(ctx, &input)

	return err
}
