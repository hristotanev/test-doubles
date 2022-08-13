package main

import "fmt"

func GetUserDetails(session SessionChecker, user UserDetailsRetriever) error {
	if session.hasExpired() {
		return fmt.Errorf("Session has expired.")
	}

	user.getDetails()
	return nil
}
