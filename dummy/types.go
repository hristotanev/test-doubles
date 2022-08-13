package main

type Session interface {
	hasExpired() bool
}

type User interface {
	getDetails()
}
