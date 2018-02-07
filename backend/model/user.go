package model

type User struct {
	ID       uint32    `db:"id" bson:"id,omitempty"`
	Location geo.Point `db:"point" bson:"point"`
}
