package main

type AuthoriserStub struct{}

func (stub *AuthoriserStub) getUserCredentials() *UserCredentials {
	return &UserCredentials{
		username: "test username",
		password: "test password",
	}
}
