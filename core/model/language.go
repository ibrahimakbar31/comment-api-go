package model

//Language struct. Define language struct
type Language struct {
	ID   string `gorm:"primary_key;size:5" json:"id" groups:"error,member,membersOrganization"`
	Name string `gorm:"size:100" json:"name" groups:"error,member,membersOrganization"`
}
