// auth provides a common interface for accessing details of the user executing
// commands and queries on the service.
package auth

// Session represents an individual users' session on the service.
type Session interface {
	GetUserId() string
	HasClaim(claim string) bool
	IsAuthenticated() bool
}

// Basic implementation for testing.
type auth struct {
	userId string
	claims []string
}

func (a auth) GetUserId() string {
	return a.userId
}

func (a auth) HasClaim(claim string) bool {
	for i := range a.claims {
		if claim == a.claims[i] {
			return true
		}
	}

	return false
}

func (a auth) IsAuthenticated() bool {
	return a.userId != ""
}

// Returns a new, simple session for use in tests.
func TestSession(userId string, claims []string) Session {
	return auth{userId: userId, claims: claims}
}
