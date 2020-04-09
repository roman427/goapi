package models

import (
	"github.com/jinzhu/gorm"
)

type PhoneBook struct {
	gorm.Model

	UserID uint `gorm:"not null"`

	Name  string `gorm:"not null"`
	Phone string `gorm:"not null"`
}

type PhoneBookService interface {
	Create(*PhoneBook) error

	ByID(uint) (*PhoneBook, error)

	ListByUserID(uint) ([]PhoneBook, error)

	Update(*PhoneBook) error

	Delete(uint) error
}

type PhoneBookGORM struct {
	*gorm.DB
}

func NewPhoneBookGORM(db *gorm.DB) *PhoneBookGORM {
	return &PhoneBookGORM{db}
}

func (s PhoneBookGORM) Create(p *PhoneBook) error {
	r := s.DB.Create(p)
	return r.Error
}

func (s PhoneBookGORM) byQuery(q *gorm.DB) (*PhoneBook, error) {
	pb := PhoneBook{}

	err := q.First(&pb).Error
	if err != nil {
		return nil, err
	}

	return &pb, nil
}

func (s PhoneBookGORM) listByQuery(q *gorm.DB) ([]PhoneBook, error) {
	pbs := []PhoneBook{}

	err := q.Find(&pbs).Error
	if err != nil {
		return nil, err
	}

	return pbs, nil
}

func (s PhoneBookGORM) ByID(id uint) (*PhoneBook, error) {
	return s.byQuery(s.DB.Where("id = ?", id))
}

func (s PhoneBookGORM) ListByUserID(userID uint) ([]PhoneBook, error) {
	return s.listByQuery(s.DB.Where("user_id = ?", userID))
}

func (s PhoneBookGORM) Update(p *PhoneBook) error {
	r := s.DB.Save(p)
	return r.Error
}

func (s PhoneBookGORM) Delete(id uint) error {
	pb := PhoneBook{
		Model: gorm.Model{ID: id},
	}
	r := s.DB.Delete(pb)
	return r.Error
}

func (s PhoneBookGORM) DestructiveReset() {
	s.DropTableIfExists(&PhoneBook{})
	s.AutoMigrate()
}

func (s PhoneBookGORM) AutoMigrate() {
	s.DB.AutoMigrate(&PhoneBook{})
}
