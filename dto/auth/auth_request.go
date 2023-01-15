package authdto

type Register struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Gender string `json:"gender"`
	Address string `json:"address"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}