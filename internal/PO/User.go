package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;type:varchar(255); primaryKey; autoIncrement ;not null; index:idx_user_id; unique" json:"id"`
	Username string    `gorm:"not null" json:"user_name"`
	Email    string    `gorm:"unique;not null" json:"email"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
