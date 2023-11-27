package bundlecounseling

type BundleCounseling struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Sessions     uint   `json:"sessions"`
	Type         string `json:"type"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	ActivePriode string `json:"active_priode"`
}

type BundleCounselingInfo struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Sessions     uint   `json:"sessions"`
	Type         string `json:"type"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	ActivePriode string `json:"active_priode"`
}

type BundleCounselingDataInterface interface {
	GetAll() ([]BundleCounselingInfo, error)
}

type BundleCounselingHandlerInterface interface {
}

type BundleCounselingServiceInterface interface {
}
