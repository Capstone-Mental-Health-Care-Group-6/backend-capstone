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

type BundleCounselingHandlerInterface interface {
	GetAllBundle() echo.HandlerFunc
	GetAllBundleFilter() echo.HandlerFunc
	CreateBundle() echo.HandlerFunc
	GetBundle() echo.HandlerFunc
	UpdateBundle() echo.HandlerFunc
	DeleteBundle() echo.HandlerFunc
}

type BundleCounselingServiceInterface interface {
	GetAllBundle() ([]BundleCounselingInfo, error)
	GetAllBundleFilter(jenis string, metode int, durasi int) ([]BundleCounselingInfo, error)
	CreateBundle(input BundleCounseling, file BundleCounselingFile) (*BundleCounseling, error)
	GetBundle(id int) (*BundleCounseling, error)
	UpdateBundle(id int, input BundleCounseling, file BundleCounselingFile) (bool, error)
	DeleteBundle(id int) (bool, error)
}

type BundleCounselingDataInterface interface {
	GetAll() ([]BundleCounselingInfo, error)
	GetAllFilter(jenis string) ([]BundleCounselingInfo, error)
	Create(input BundleCounseling) (*BundleCounseling, error)
	GetById(id int) (*BundleCounseling, error)
	Update(id int, newData BundleCounseling) (bool, error)
	Delete(id int) (bool, error)
	HargaMetode(id int) (uint, error)
	HargaDurasi(id int) (uint, error)
}
