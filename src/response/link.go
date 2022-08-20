package response

import "time"

// UserResponse
type UserResponse struct {
	ID       string    `json:"id,omitempty" example:"62fbfaa5f79e97a5501979f3"`             // ObjectID represented as a string
	Url      string    `json:"url" example:"http://chupe.ba"`                               // Full URL
	Code     string    `json:"code" example:"a1b2c3"`                                       // Short alphanumeric 6 letter code that is used for redirection
	HitCount int       `json:"hitCount" example:"42"`                                       // Number of times the redirection took place
	Created  time.Time `json:"created" example:"2021-05-25T00:00:00.0Z" format:"date-time"` // Date the User was stored
} // @name UserResponse
