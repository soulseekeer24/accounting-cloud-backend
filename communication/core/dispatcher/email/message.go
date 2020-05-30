package email

type Message struct {
	ID        string   `json:"id" bson:"_id,omitempty"`
	Recipient []string `json:"recipient" bson:"recipient"`
	Emitter   string   `json:"emitter" bson:"recipient"`
	Body      string   `json:"body`

	ExternalID string `json:"external_id", bson:"external_id,omitempty"`
	Provider   string `json:"provider", bson:"provider,omitempty"`
	// write record
	CreatedAt int64 `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at,omitempty"`
}
