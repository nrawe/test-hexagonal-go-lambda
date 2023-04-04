package auth_test

import (
	"testing"

	"github.com/nrawe/tasks-kata/internal/auth"
	"github.com/stretchr/testify/assert"
)

func validateAuth(t *testing.T, a auth.Session) {
	assert.Equal(t, "user-id", a.GetUserId(), "user id set correctly")
	assert.True(t, a.IsAuthenticated(), "user has authenticated")
	assert.True(t, a.HasClaim("claim1"), "has claim1")
	assert.True(t, a.HasClaim("claim2"), "has claim2")
	assert.False(t, a.HasClaim("claim3"), "does not have claim3")
}

func TestCreateSimpleAuth(t *testing.T) {
	a := auth.TestSession("user-id", []string{"claim1", "claim2"})

	validateAuth(t, a)
}
