package web_push

import (
	webpush "github.com/SherClockHolmes/webpush-go"
)

type WebPusher interface {
	SendNotification(message string, endpoint string, auth string, p256dh string, ttl int) error
}

type webPush struct {
	vapidPrivateKey string
	VapidPublicKey  string
	SenderEmail     string
}

func NewWebPush(VAPIDPrivateKey string, VAPIDPublicKey string, senderEmail string) webPush {
	wp := webPush{}
	wp.vapidPrivateKey = VAPIDPrivateKey
	wp.VapidPublicKey = VAPIDPublicKey
	wp.SenderEmail = senderEmail
	return wp
}

func (wp *webPush) SendNotification(message string, endpoint string, auth string, p256dh string, ttl int) error {
	sub := &webpush.Subscription{
		Endpoint: endpoint,
		Keys: webpush.Keys{
			Auth:   auth,
			P256dh: p256dh,
		},
	}

	// Send Notification
	resp, err := webpush.SendNotification([]byte(message), sub, &webpush.Options{
		Subscriber:      wp.SenderEmail,
		VAPIDPublicKey:  wp.VapidPublicKey,
		VAPIDPrivateKey: wp.vapidPrivateKey,
		TTL:             ttl,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
