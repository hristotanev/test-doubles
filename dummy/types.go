package main

type SessionChecker interface {
	hasExpired() bool
}

type UserDetailsRetriever interface {
	getDetails()
}
