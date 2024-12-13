package request

// LoginRequest represents the data required for user login.
// @Summary Login Data
// @Description Structure containing the necessary fields for user login.
type LoginRequest struct {
	// User's email (required and must be a valid email address).
	Email string `json:"email" binding:"required,email" example:"test@test.com"`

	// User's password (required, minimum of 8 characters, and must contain at least one of the characters: !@#$%&*()_-=+).
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_-=+" example:"password#@#@!2121"`
}