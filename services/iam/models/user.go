package models

import (
	"fmt"
	"reflect"

	"github.com/SKF/go-utility/uuid"
)

//User struct
type User struct {
	ID             uuid.UUID `json:"id"`
	Email          Email     `json:"email"`
	UserRoles      []string  `json:"userRoles"`
	Username       string    `json:"username"`
	UserStatus     string    `json:"userStatus"`
	EulaAgreedDate string    `json:"eulaAgreedDate"`
	ValidEula      bool      `json:"validEula"`
}

func (user *User) Validate() error {
	if err := UserStatus.Validate(user.UserStatus); err != nil {
		return fmt.Errorf("User status invalid: %+v", err)
	}

	if user.Email == "" {
		return fmt.Errorf("No email provided")
	}

	if len(user.UserRoles) == 0 {
		return fmt.Errorf("No roles provided")
	}

	for _, role := range user.UserRoles {
		if err := UserRoles.Validate(role); err != nil {
			return fmt.Errorf("Invalid UserRoles provided: %+v", err)
		}
	}

	if user.Username == "" {
		return fmt.Errorf("No username provided")
	}

	return nil
}

//UserCredentials struct
type UserCredentials struct {
	Username Email
	Email    string
	Password string
}

//UserAgreement struct
type UserAgreement struct {
	Email      Email
	EULAAgreed string
}

// UserWithCompanies struct
type UserWithCompanies struct {
	ID         string
	Email      Email
	UserRoles  []string
	Username   string
	UserStatus string
	Companies  []string
}

type userStatus struct {
	Active   string
	Inactive string
}

type userRoles struct {
	UserAdmin        string
	HierarchyManager string
	DeviceManager    string
	Inspector        string
}

var (
	//UserStatus variables
	UserStatus = userStatus{Active: "active", Inactive: "inactive"}

	//UserRoles variables
	UserRoles = userRoles{
		UserAdmin:        "user_admin",
		HierarchyManager: "hierarchy_manager",
		DeviceManager:    "device_manager",
		Inspector:        "inspector",
	}
)

func (userStatus) Validate(status string) (err error) {
	v := reflect.ValueOf(UserStatus)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == status {
			return
		}
	}
	err = fmt.Errorf("Unknown status: %s", status)
	return
}

func (userRoles) Validate(role string) (err error) {
	v := reflect.ValueOf(UserRoles)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == role {
			return
		}
	}
	err = fmt.Errorf("Unknown role: %s", role)
	return
}

func (userRoles) AreEqual(a []string, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
