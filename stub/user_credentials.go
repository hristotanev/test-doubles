package main

type UserCredentials struct {
	username string
	password string
}

func (credentials *UserCredentials) GetUsername() string {
	return credentials.username
}

func (credentials *UserCredentials) GetPassword() string {
	return credentials.password
}
