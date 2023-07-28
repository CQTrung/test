package entity

type User struct {
	Id       int    `gorm:"primary_key:auto_increment" json:"id"`
	UserName string `gorm:"username" json:"name"`
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"password" json:"password"`
}

type UserRepository interface {
	Save(user User) error
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
}

type UserUsecase interface {
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user User) error
	HashedPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
	LogIn(user User) error
	Register(user User) error
}
