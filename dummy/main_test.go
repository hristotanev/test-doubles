package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfSessionHasExpiredAnErrorIsReturned(t *testing.T) {
	expiredSession := &ExpiredSessionStub{}
	dummyUser := &DummyUser{}

	err := GetUserDetails(expiredSession, dummyUser)
	assert.NotNil(t, err)
}
