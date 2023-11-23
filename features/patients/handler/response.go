package handler

type PatientResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Avatar      string `json:"avatar"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}

type PatientLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token any    `json:"token"`
}
