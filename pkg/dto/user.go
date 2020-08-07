package dto

// User ...
type User struct {
	ID           string   `json:"id" bson:"id"`
	Role         string   `json:"role" bson:"role"`
	Name         string   `json:"name" bson:"name"`
	PhoneNumber  string   `json:"phoneNumber" bson:"phoneNumber"`
	Email        string   `json:"email" bson:"email"`
	Password     string   `json:"password" bson:"password"`
	Disabled     bool     `json:"disabled" bson:"disabled"`
	AccessToken  string   `json:"accessToken" bson:"-"`
	RefreshToken string   `json:"refreshToken" bson:"-"`
	ResetToken   string   `json:"resetToken" bson:"-"`
	AccessUuid   string   `json:"accessUuid" bson:"-"`
	RefreshUuid  string   `json:"refreshUuId" bson:"-"`
	AtExpires    int64    `json:"atExpires" bson:"-"`
	RtExpires    int64    `json:"rtExpires" bson:"-"`
	ResetExpires int64    `json:"resetExpires" bson:"-"`
	BlockList    []string `json:"blockList" bson:"blockList"`
}
