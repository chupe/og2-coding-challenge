package response

import "time"

type Factory struct {
	Type              string    `json:"type,omitempty" validate:"required,alpha" example:"iron"`                    // Factory type
	Level             int       `json:"level,omitempty" validate:"num" example:"3"`                                 // Factory level
	RatePerMinute     int       `json:"ratePerMinute,omitempty" validate:"num" example:"20"`                        // Ore production rate per minute
	UnderConstruction bool      `json:"underConstruction,omitempty" validate:"bool" example:"false"`                // Factory is under construction
	TimeToFinish      time.Time `json:"timeToFinish,omitempty" example:"2021-05-25T00:00:00.0Z" format:"date-time"` // Time of finishing the latest update
} // @name Factory
