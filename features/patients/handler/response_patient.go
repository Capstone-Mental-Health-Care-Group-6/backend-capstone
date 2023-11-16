package handlerimport

type PatientResponse struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	Name           string `json:"name"`
	DateOfBirth    string `json:"date_of_birth"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Gender         string `json:"gender"`
	MarriageStatus string `json:"marriage_status"`
	Avatar         string `json:"avatar"`
	Address        string `json:"address"`
}
