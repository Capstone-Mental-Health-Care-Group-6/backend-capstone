package bundlecounseling

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type BundleCounseling struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Sessions     uint   `json:"sessions"`
	Type         string `json:"type"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	ActivePriode uint   `json:"active_priode"`
	Avatar       string `json:"avatar"`
}

type BundleCounselingInfo struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Sessions     uint   `json:"sessions"`
	Type         string `json:"type"`
	Price        uint   `json:"price"`
	Description  string `json:"description"`
	ActivePriode uint   `json:"active_priode"`
	Avatar       string `json:"avatar"`
}

type BundleCounselingFile struct {
	Avatar multipart.File `json:"avatar"`
}

type BundleCounselingDataInterface interface {
	GetAll() ([]BundleCounselingInfo, error)
	Create(input BundleCounseling) (*BundleCounseling, error)
}

type BundleCounselingHandlerInterface interface {
	GetAllBundle() echo.HandlerFunc
	CreateBundle() echo.HandlerFunc
}

type BundleCounselingServiceInterface interface {
	GetAllBundle() ([]BundleCounselingInfo, error)
	CreateBundle(input BundleCounseling, file BundleCounselingFile) (*BundleCounseling, error)
}
