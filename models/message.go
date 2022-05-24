package models

type Message struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	UserToID     uint   `json:"userToID"`
	UserFromID     uint   `json:"userFromID"`
	Message  string `json:"message"`
}
