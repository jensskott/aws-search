package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("us-west-1")
	assert.IsType(t, Ec2Implementation{}, c)
	assert.IsType(t, session.Session{}, *c.Session)
}
