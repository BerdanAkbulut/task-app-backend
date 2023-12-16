package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedBy   string    `gorm:"type:varchar(50);not null" json:"created_by"`
	UpdatedBy   string    `gorm:"type:varchar(50);not null" json:"updated_by"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	Status      string    `gorm:"type:varchar(1);not null" json:"status"`
}

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(100);unique; not null" json:"email"`
	Password  string    `gorm:"type:varchar(120);not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	IsActive  bool      `gorm:"type:bool;default:false;not null" json:"is_active"`
	IsDeleted bool      `gorm:"type:bool;default:false;not null" json:"is_deleted"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return result == nil
}
