package main

type ExpiredSessionStub struct{}

func (session *ExpiredSessionStub) hasExpired() bool {
	return true
}
