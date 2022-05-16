package model

type Bodystatus struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Weight            int     `json:"weight"`
	Height            int     `json:"height"`
	BodyFatPercentage float64 `json:"body_fat_percentage"`
	MusclePercentage  float64 `json:"muscle_percentage"`
}
