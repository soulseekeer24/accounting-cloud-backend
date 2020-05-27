package bill

type Bill struct {
	ID          string `json:"id" bson:"id"`
	Description string `json:"description" `
	CreatedAt   int64  `json:"created_at" bson:"created_at" `
}
