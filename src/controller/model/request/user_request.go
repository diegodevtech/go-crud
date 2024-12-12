package request

// UserRequest represents the input data for creating a new user.
// @Summary User Input Data
// @Description Structure containing the required fields for creating a new user.
type UserRequest struct {
	// User's email (required and must be a valid email address).
	// Example: user@example.com
	// @json
	// @jsonTag email
	// @jsonExample user@example.com
	// @binding required,email
	Email string `json:"email" binding:"required,email" example:"test@test.com"`

	// User's password (required, minimum of 8 characters, and must contain at least one of the characters: !@#$%&*()_-=+).
	// @json
	// @jsonTag password
	// @jsonExample P@ssw0rd!
	// @binding required,min=8,containsany=!@#$%&*()_-=+
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_-=+" example:"password#@#@!2121"`

	// User's name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	// @json
	// @jsonTag name
	// @jsonExample John Doe
	// @binding required,min=4,max=100
	Name string `json:"name" binding:"required,min=4,max=100" example:"John Doe"`

	// User's age (required, must be between 18 and 100).
	// @json
	// @jsonTag age
	// @jsonExample 30
	// @binding required,min=18,max=100
	Age int8 `json:"age" binding:"required,min=18,max=100" example:"30"`
}

type UserUpdateRequest struct {
	// User's name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	// @json
	// @jsonTag name
	// @jsonExample John Doe
	// @binding required,min=4,max=100
	Name string `json:"name" binding:"omitempty,min=4,max=100" example:"John Doe"`

	// User's age (required, must be between 18 and 100).
	// @json
	// @jsonTag age
	// @jsonExample 30
	// @binding required,min=18,max=100
	Age int8 `json:"age" binding:"omitempty,min=18,max=100" example:"30"`
}