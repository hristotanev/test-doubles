package main

type Authoriser interface {
	getUserCredentials() *UserCredentials
}
