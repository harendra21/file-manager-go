package models

// Detail represents phone number and address
// of a user. A user can have multiple numbers and addresses.
// But only one primary addresses
type Detail struct {
	Id          int
	PhoneNumber int
	Primary     bool
	User        *User `orm:"rel(fk)"`
}
