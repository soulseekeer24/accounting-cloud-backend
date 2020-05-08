package user

type User struct {
	ID        string `json:"id" bson:"id"`
	Username  string `json:"username" `
	CreatedAt int64  `json:"created_at" bson:"created_at" `
}
