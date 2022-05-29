package user_constant

type Permission int64

const (
	Permission_Master Permission = 1
	Permission_Friend Permission = 2
	Permission_Normal Permission = 3
)

type Gender int64

const (
	Gender_Male   Gender = 1
	Gender_Female Gender = 2
)
