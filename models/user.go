package models

type User struct {
	Model

	Account  string `gorm:"column:account;unique"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Sex      string `gorm:"column:sex"`
	Age      int    `gorm:"column:age"`
	Address  string `gorm:"column:address"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`

	Avatar string `gorm:"column:avatar"`
	Status string `grom:"column:status"`
}

func (u User) TableName() string {
	return "users"
}
