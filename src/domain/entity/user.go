package entity

import (
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//User : Struct
type User struct {
	ID          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Username    string     `gorm:"size:100;not null;" json:"username"`
	FirstName   string     `gorm:"size:100;not null;" json:"first_name"`
	LastName    string     `gorm:"size:100;not null;" json:"last_name"`
	Email       string     `gorm:"size:100;not null;unique" json:"email"`
	Pdam        string     `gorm:"size:100;not null;unique" json:"pdam"`
	Name        string     `gorm:"size:100;null;unique" json:"name"`
	Type        string     `gorm:"size:100;null;unique" json:"type"`
	Password    string     `gorm:"size:100;not null;" json:"password"`
	PhotoURI    string     `gorm:"size:100;null;" json:"photo_uri"`
	PhotoBase64 string     `gorm:"size:100;not null;" json:"photo_base64"`
	Pin         int        `gorm:"size:255;not null;" json:"pin"`
	IsActive    bool       `gorm:"size:255;null;" json:"is_active"`
	RoleID      int        `gorm:"size:10;not null;" json:"role_id"`
	LimitSaldo  float64    `json:"limit_saldo" gorm:"null"`
	Limit       float64    `json:"limit" gorm:"null"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	LastLogin   *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_login"`
	LastLogout  *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_logout"`
	LoginStatus string     `gorm:"size:255;null;" json:"login_status"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

//PublicUser : So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:         u.ID,
		FirstName:  u.FirstName,
		Type:       u.Type,
		LastName:   u.LastName,
		Name:       u.Name,
		Email:      u.Email,
		Pdam:       u.Pdam,
		Pin:        u.Pin,
		IsActive:   u.IsActive,
		Username:   u.Username,
		RoleID:     u.RoleID,
		Password:   u.Password,
		LimitSaldo: u.LimitSaldo,
	}

}

//PublicUser : Struct
type PublicUser struct {
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Username   string  `gorm:"size:100;not null;" json:"username"`
	FirstName  string  `gorm:"size:100;not null;" json:"first_name"`
	LastName   string  `gorm:"size:100;not null;" json:"last_name"`
	Name       string  `gorm:"size:100;not null;" json:"name"`
	Type       string  `gorm:"size:100;not null;" json:"type"`
	Email      string  `gorm:"size:100;not null;unique" json:"email"`
	Pdam       string  `gorm:"size:100;not null;unique" json:"pdam"`
	RoleID     int     `gorm:"size:10;not null;" json:"role_id"`
	Password   string  `gorm:"size:100;not null;" json:"password"`
	Pin        int     `gorm:"size:255;not null;" json:"pin"`
	LimitSaldo float64 `json:"limit_saldo" gorm:"null"`
	IsActive   bool    `gorm:"size:255;null;" json:"is_active"`
}

//Validate : user validation by Action
func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "email email"
			}
		}

	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	case "login_username":
		if u.Password == "" {
			errorMessages["message"] = "password is required"
		}
		if u.Username == "" {
			errorMessages["message"] = "username is required"
		}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["firstname_required"] = "first name is required"
		}
		if u.LastName == "" {
			errorMessages["lastname_required"] = "last name is required"
		}
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["invalid_password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}
	return errorMessages
}
