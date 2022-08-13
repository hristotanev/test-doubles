package main

import "testing"

func TestIfSessionHasExpiredAnErrorIsReturned(t *testing.T) {
	expiredSession := &ExpiredSessionStub{}
	dummyUser := &DummyUser{}

	err := GetUserDetails(expiredSession, dummyUser)
	if err == nil {
		t.FailNow()
	}
}
