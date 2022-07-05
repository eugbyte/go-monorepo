package models

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	Endpoint   string `json:"endpoint"`
	Message    string `json:"message"`
	Expiration string `json:"expiration_time"`
	Keys       Keys   `json:"keys"`
}
