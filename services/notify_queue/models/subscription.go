package models

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	Endpoint   string `json:"endpoint"`
	Expiration int    `json:"expirationTime"`
	Keys       Keys   `json:"keys"`
}
