package profiles

type Profile struct {
	ID        string        `json:"id" bson:"_id,omitempty"`
	AccountID string        `json:"account_id" bson:"account_id,omitempty"`
	FirstName string        `json:"first_name" bson:"first_name,omitempty"`
	LastName  string        `json:"last_name" bson:"last_name,omitempty"`
	Contacts  []ContactInfo `json:"contacts" bson:"contacts,omitempty"`
	Address   []Address     `json:"address" bson:"address,omitempty"`
	// write record
	CreatedAt int64         `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64         `json:"updated_at" bson:"updated_at,omitempty"`
}
