package response

import "time"

type UserResponse struct {
	ID        string `json:"id,omitempty" example:"62fbfaa5f79e97a5501979f3"`            // ObjectID represented as a string
	Username  string `json:"username" validate:"required,alphanum" example:"example123"` // Full URL
	Iron      int    `json:"iron" validate:"numeric" example:"42"`                       // Short alphanumeric 6 letter code that is used for redirection
	Copper    int    `json:"copper" validate:"numeric" example:"42"`                     // Number of times the redirection took place
	Gold      int    `json:"gold" validate:"numeric" example:"42"`
	Factories []Factories
	Created   time.Time `json:"created" example:"2021-05-25T00:00:00.0Z" format:"date-time"` // Date the User was stored
} // @name UserResponse
