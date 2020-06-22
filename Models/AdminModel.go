//Models/AdminModel.go
package Models

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *Admin) TableName() string {
	return "admin"
}