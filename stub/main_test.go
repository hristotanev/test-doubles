package main

import "testing"

func TestAuthorizerReturnsUserCredentials(t *testing.T) {
	authoriser := &AuthoriserStub{}
	credentials := GetUserCredentials(authoriser)

	if credentials.GetUsername() != "test username" {
		t.FailNow()
	}

	if credentials.GetPassword() != "test password" {
		t.FailNow()
	}
}
