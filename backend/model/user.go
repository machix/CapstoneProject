package model

type User struct {
	ID uint32 `db:"id" bson:"id,omitempty"`
}
