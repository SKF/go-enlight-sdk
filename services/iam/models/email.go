package models

import (
	"fmt"
	"regexp"
)

type Email string

//Validate email
func (email Email) Validate() error {
	pattern := `^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	match, _ := regexp.MatchString(pattern, string(email))

	if !match {
		return fmt.Errorf("Invalid Email: %s", email)
	}

	return nil
}

//String converts email to string
func (email Email) String() string {
	return string(email)
}

//StringPtr converts email to string pointer
func (email Email) StringPtr() *string {
	emailstr := string(email)
	return &emailstr
}
