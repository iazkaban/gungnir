package entities

type User struct {
	ID       string `xorm:"id"`
	UserName string `xorm:"username"`
}
