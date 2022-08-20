package response

type ErrorResponse struct {
	Status  int    `json:"status,omitempty" example:"404"`           // Http status
	Message string `json:"message" example:"Review input"`           // User friendly message
	Error   string `json:"error" example:"No Users found in the DB"` // Actual error thrown
} // @name ErrorResponse
