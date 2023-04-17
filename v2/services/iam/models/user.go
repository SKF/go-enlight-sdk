package models

import (
	"github.com/SKF/go-utility/v2/uuid"
)

// User struct
type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	UserRoles      []string  `json:"userRoles"`
	Username       string    `json:"username"`
	UserStatus     string    `json:"userStatus"`
	EulaAgreedDate string    `json:"eulaAgreedDate"`
	ValidEula      bool      `json:"validEula"`
}

// UserCredentials struct
type UserCredentials struct {
	Username string
	Email    string
	Password string
}

// UserAgreement struct
type UserAgreement struct {
	Email      string
	EULAAgreed string
}

// UserWithCompanies struct
type UserWithCompanies struct {
	ID         string
	Email      string
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
	// UserStatus variables
	UserStatus = userStatus{Active: "active", Inactive: "inactive"}

	// UserRoles variables
	UserRoles = userRoles{
		UserAdmin:        "user_admin",
		HierarchyManager: "hierarchy_manager",
		DeviceManager:    "device_manager",
		Inspector:        "inspector",
	}
)
