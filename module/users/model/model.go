package usermodel

import (
	"errors"
	"h5travelotobackend/common"
	"h5travelotobackend/component/tokenprovider"
)

const EntityName = "User"

type User struct {
	common.SqlModel `json:",inline"`
	Email           string            `json:"email" gorm:"column:email;"`
	Password        string            `json:"-" gorm:"column:password;"`
	Salt            string            `json:"-" gorm:"column:salt;"`
	LastName        string            `json:"last_name" gorm:"column:last_name;"`
	Firstname       string            `json:"first_name" gorm:"column:first_name;"`
	Phone           string            `json:"phone" gorm:"column:phone;"`
	Role            string            `json:"role" gorm:"column:role;"`
	Avatar          *common.Image     `json:"avatar,omitempty" gorm:"column:avatar;type:json;"`
	Gender          string            `json:"gender" gorm:"column:gender"`
	DateOfBirth     *common.CivilDate `json:"date_of_birth" gorm:"column:date_of_birth"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SqlModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Firstname       string        `json:"first_name" gorm:"column:first_name;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (u *UserCreate) Validate() error {
	if !common.IsEmail(u.Email) {
		return ErrInvalidEmail
	}
	if !common.IsValidPassword(u.Password) {
		return ErrPasswordNotStrong
	}

	if common.IsEmpty(u.Firstname) || common.IsEmpty(u.LastName) {
		return ErrNameIsEmpty
	}

	return nil
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	LastName    string            `json:"last_name" gorm:"column:last_name;"`
	Firstname   string            `json:"first_name" gorm:"column:first_name;"`
	Phone       string            `json:"phone" gorm:"column:phone;"`
	Avatar      *common.Image     `json:"avatar,omitempty" gorm:"column:avatar;type:json;"`
	Gender      string            `json:"gender" gorm:"column:gender"`
	DateOfBirth *common.CivilDate `json:"date_of_birth" gorm:"column:date_of_birth"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (u *UserUpdate) Validate() error {
	if common.IsEmpty(u.Firstname) || common.IsEmpty(u.LastName) {
		return ErrNameIsEmpty
	}

	if !common.IsEmpty(u.Phone) && !common.IsPhoneNumber(u.Phone) {
		return ErrInvalidPhone
	}
	return nil
}

type UserChangePassword struct {
	OldPassword string `json:"old_password" gorm:"-"`
	Password    string `json:"password" gorm:"column:password;"`
}

func (UserChangePassword) TableName() string {
	return User{}.TableName()
}

func (u *UserChangePassword) Validate() error {
	if !common.IsValidPassword(u.Password) {
		return ErrPasswordNotStrong
	}
	return nil
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token""`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrInvalidPhone = common.NewCustomError(
		errors.New("invalid phone"),
		"invalid phone",
		"ErrInvalidPhone",
	)

	ErrNameIsEmpty = common.NewCustomError(
		errors.New("name is empty"),
		"name is empty",
		"ErrNameIsEmpty",
	)

	ErrInvalidEmail = common.NewCustomError(
		errors.New("invalid email"),
		"invalid email",
		"ErrInvalidEmail",
	)

	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUserHasDeletedOrDisabled = common.NewCustomError(
		errors.New("user has been deleted or disabled"),
		"user has been deleted or disabled",
		"ErrUserHasDeletedOrDisabled",
	)

	ErrPasswordNotStrong = common.NewCustomError(
		errors.New("password is not strong enough"),
		"password is not strong enough",
		"ErrPasswordNotStrong",
	)

	ErrUserBanned = common.NewCustomError(
		errors.New("user has been banned"),
		"user has been banned",
		"ErrUserBanned",
	)
	ErrPasswordMatchedWithPast = common.NewCustomError(
		errors.New("password matched with past"),
		"password matched with past",
		"ErrPasswordMatchedWithPast",
	)
	ErrOldPasswordNotMatch = common.NewCustomError(
		errors.New("old password not match"),
		"old password not match",
		"ErrOldPasswordNotMatch",
	)
)
