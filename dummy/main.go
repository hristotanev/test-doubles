package main

import "fmt"

func GetUserDetails(session Session, user User) error {
	if session.hasExpired() {
		return fmt.Errorf("Session has expired.")
	}

	user.getDetails()
	return nil
}
