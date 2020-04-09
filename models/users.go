package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nafisfaysal/goapi/hash"
	"golang.org/x/crypto/bcrypt"
)

var userPassHash = "wow_wow_random_string"
var hmac = hash.NewHMAC("very-key")

type User struct {
	gorm.Model

	Email string `gorm:"not null;unique_index"`

	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

type UserService interface {
	Create(*User) error

	ByID(uint) (*User, error)
	ByEmail(string) (*User, error)

	Authenticate(email, password string) (*User, error)

	Update(*User) error

	Delete(uint) error
}

type UserGORM struct {
	*gorm.DB
}

func NewUserGORM(db *gorm.DB) *UserGORM {
	return &UserGORM{db}
}

func (us UserGORM) Create(user *User) error {
	if user.Password == "" {
		return nil
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password+userPassHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	r := us.DB.Create(user)
	return r.Error
}

func (us UserGORM) byQuery(q *gorm.DB) (*User, error) {
	u := User{}

	err := q.First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (us UserGORM) ByID(id uint) (*User, error) {
	return us.byQuery(us.DB.Where("id = ?", id))
}

func (us UserGORM) ByEmail(email string) (*User, error) {
	return us.byQuery(us.DB.Where("email = ?", email))
}

func (us UserGORM) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if foundUser == nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPassHash))
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (us UserGORM) Update(user *User) error {
	r := us.DB.Save(user)
	return r.Error
}

func (us UserGORM) Delete(id uint) error {
	u := &User{
		Model: gorm.Model{ID: id},
	}
	r := us.DB.Delete(u)
	return r.Error
}

func (us UserGORM) DestructiveReset() {
	us.DropTableIfExists(&User{})
	us.AutoMigrate()
}

func (us UserGORM) AutoMigrate() {
	us.DB.AutoMigrate(&User{})
}
