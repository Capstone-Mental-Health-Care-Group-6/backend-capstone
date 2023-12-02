package handler

type BundleCounselingResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Sessions     uint   `json:"sessions"`
	Type         string `json:"type"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	ActivePriode string `json:"active_priode"`
	Avatar       string `json:"avatar"`
}
