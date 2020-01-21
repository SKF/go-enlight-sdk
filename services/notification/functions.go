package notification

import (
	"context"

	"github.com/SKF/proto/common"
	proto "github.com/SKF/proto/notification"
)

func (c *client) SendNotification(notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.SendNotificationWithContext(ctx, notificationType, resource, header, body, createdBy)
}

func (c *client) SendNotificationWithContext(ctx context.Context, notificationType proto.NotificationType, resource common.Origin, header, body, createdBy string) (string, error) {
	input := proto.SendNotificationInput{
		Type:      notificationType,
		Resource:  &resource,
		Header:    header,
		Body:      body,
		CreatedBy: createdBy,
	}
	output, err := c.api.SendNotification(ctx, &input)
	if err != nil {
		return "", err
	}

	return output.Id, nil
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

func (c *client) GetUserNotifications(userID string, limit int) ([]proto.NotificationMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserNotificationsWithContext(ctx, userID, limit)
}

func (c *client) GetUserNotificationsWithContext(ctx context.Context, userID string, limit int) ([]proto.NotificationMessage, error) {
	input := proto.GetUserNotificationsInput{
		UserId: userID,
		Limit:  "limit",
	}

	result := []proto.NotificationMessage{}
	output, err := c.api.GetUserNotifications(ctx, &input)
	if err != nil {
		return result, err
	}

	for _, el := range output.Notifications {
		result = append(result, *el)
	}

	return result, nil
}
