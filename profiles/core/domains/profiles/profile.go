package profiles

type Profile struct {
	ID        string        `json:"id" bson:"_id,omitempty"`
	AccountID string        `json:"account_id" bson:"account_id,omitempty"`
	FirstName string        `json:"firstname" bson:"firstname,omitempty"`
	LastName  string        `json:"lastname" bson:"lastname,omitempty"`
	CreatedAt int64         `json:"created_at" bson:"created_at,omitempty"`
	Contacts  []ContactInfo `json:"contacts" bson:"contacts,omitempty"`
	Address   []Address     `json:"address" bson:"address,omitempty"`
}
