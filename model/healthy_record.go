package model

type BodyStatus struct {
	Id                int     `json:"id"`
	UserId            int     `json:"user_id"`
	Weight            int     `json:"weight"`
	Height            int     `json:"height"`
	BodyFatPercentage float64 `json:"body_fat_percentage"`
	MusclePercentage  float64 `json:"muscle_percentage"`
}
