package models

type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	// The ID field is automatically indexed by CosmosDB
	// Made up by combining the Username field and the Company field. e.g. `lowercase({company}__{username})` => fakepanda__abc@m.com
	ID             string `json:"id" bson:"_id"`
	Company        string `json:"company"`
	UserID         string `json:"userID" bson:"userID"`
	Endpoint       string `json:"endpoint"`
	ExpirationTime int    `json:"expirationTime" bson:"expirationTime"`
	Keys           Keys   `json:"keys"`
}
