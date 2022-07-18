package models

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	ID             string `json:"id" bson:"_id"` // e.g. lowercase(`${company}__${username}`), e.g. fakepanda__abc@m.com
	Endpoint       string `json:"endpoint"`
	ExpirationTime string `json:"expirationTime"`
	Keys           Keys   `json:"keys"`
	Username       string `json:"username"`
	Company        string `json:"company"`
}
