// Package auth implements JWT authentication, RBAC authorization,
// and middleware for the search engine API.
//
// Architecture (per SRS Section 14):
//   - JWT with RS256 signing
//   - Access token TTL: 15 minutes
//   - Refresh token TTL: 7 days
//   - Roles: ADMIN, USER, READONLY, API_USER
package auth

// JWTManager handles JWT token creation, validation, and refresh.
type JWTManager struct {
	AccessTokenTTL  int // seconds (default: 900 = 15 minutes)
	RefreshTokenTTL int // seconds (default: 604800 = 7 days)
	Algorithm       string
}

// NewJWTManager creates a new JWT manager with default settings.
func NewJWTManager() *JWTManager {
	return &JWTManager{
		AccessTokenTTL:  900,
		RefreshTokenTTL: 604800,
		Algorithm:       "RS256",
	}
}

// Role represents a user's authorization level.
type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleUser     Role = "USER"
	RoleReadonly Role = "READONLY"
	RoleAPIUser  Role = "API_USER"
)

// RBACPolicy defines role-based access control rules.
type RBACPolicy struct {
	Role        Role
	Permissions []string
}
