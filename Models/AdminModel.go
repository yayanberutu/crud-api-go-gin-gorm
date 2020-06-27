//Models/AdminModel.go
package models

type Admin struct {
	IPKA     string `json:"username"`
	Password string `json:"password"`
}

func (b *Admin) TableName() string {
	return "admin"
}
