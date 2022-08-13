package main

import "testing"

func TestIfSessionHasExpiredAnErrorIsReturned(t *testing.T) {
	expiredSession := &ExpiredSession{}
	dummyUser := &DummyUser{}

	err := GetUserDetails(expiredSession, dummyUser)
	if err == nil {
		t.FailNow()
	}
}
