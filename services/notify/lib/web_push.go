package lib

import (
	webpush "github.com/SherClockHolmes/webpush-go"
)

type NotifyServiceImpl interface {
	Init(VAPIDPrivateKey string, VAPIDPublicKey string, senderEmail string)
	SendNotification(message string, endpoint string, auth string, p256dh string, ttl int) error
}

type NotifyService struct {
	vapidPrivateKey string
	VapidPublicKey  string
	SenderEmail     string
}

// Set the Vapid details
func (ns *NotifyService) Init(VAPIDPrivateKey string, VAPIDPublicKey string, senderEmail string) {
	ns.vapidPrivateKey = VAPIDPrivateKey
	ns.VapidPublicKey = VAPIDPublicKey
	ns.SenderEmail = senderEmail
}

func (ns *NotifyService) SendNotification(message string, endpoint string, auth string, p256dh string, ttl int) error {
	sub := &webpush.Subscription{
		Endpoint: endpoint,
		Keys: webpush.Keys{
			Auth:   auth,
			P256dh: p256dh,
		},
	}

	// Send Notification
	resp, err := webpush.SendNotification([]byte(message), sub, &webpush.Options{
		Subscriber:      ns.SenderEmail,
		VAPIDPublicKey:  ns.VapidPublicKey,
		VAPIDPrivateKey: ns.vapidPrivateKey,
		TTL:             ttl,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
