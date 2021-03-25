package tests

import (
	"github.com/ktrysmt/go-bitbucket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase(t *testing.T) {

	user := getUsername()
	pass := getPassword()

	c := bitbucket.NewBasicAuth(user, pass)

	res, err := c.Base.Get("my-url")

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
