package webpush

import (
	"encoding/json"
	"errors"
	"io"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type WebPushServicer interface {
	SendNotification(message interface{}, endpoint string, auth string, p256dh string, ttl int) error
}

type webPushService struct {
	vapidPrivateKey string
	VapidPublicKey  string
	SenderEmail     string
}

func NewWebPush(VAPIDPrivateKey string, VAPIDPublicKey string, senderEmail string) WebPushServicer {
	wp := webPushService{
		vapidPrivateKey: VAPIDPrivateKey,
		VapidPublicKey:  VAPIDPublicKey,
		SenderEmail:     senderEmail,
	}
	return &wp
}

func (wp *webPushService) SendNotification(message interface{}, endpoint string, auth string, p256dh string, ttl int) error {
	sub := &webpush.Subscription{
		Endpoint: endpoint,
		Keys: webpush.Keys{
			Auth:   auth,
			P256dh: p256dh,
		},
	}

	objBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Send Notification
	resp, err := webpush.SendNotification(objBytes, sub, &webpush.Options{
		Subscriber:      wp.SenderEmail,
		VAPIDPublicKey:  wp.VapidPublicKey,
		VAPIDPrivateKey: wp.vapidPrivateKey,
		TTL:             ttl,
	})
	if err != nil {
		formats.Trace("cannot send notification")
		return err
	}

	defer resp.Body.Close()

	// https://web.dev/push-notifications-web-push-protocol/#response-from-push-service
	formats.Trace(resp.Status, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	str := string(body)

	if resp.StatusCode != 201 {
		return errors.New(str)
	}

	return nil
}
