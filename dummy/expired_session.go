package main

type ExpiredSession struct{}

func (session *ExpiredSession) hasExpired() bool {
	return true
}
