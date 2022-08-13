package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizerReturnsUserCredentials(t *testing.T) {
	authoriser := &AuthoriserStub{}
	credentials := GetUserCredentials(authoriser)
	assert.Equal(t, credentials.GetUsername(), "test username")
	assert.Equal(t, credentials.GetPassword(), "test password")
}
