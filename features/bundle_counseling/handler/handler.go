package handler

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/helper"
)

type BundleCounselingHandler struct {
	s   bundlecounseling.BundleCounselingServiceInterface
	jwt helper.JWTInterface
}

func New(s bundlecounseling.BundleCounselingServiceInterface, jwt helper.JWTInterface) bundlecounseling.BundleCounselingHandlerInterface {
	return &BundleCounselingHandler{
		s:   s,
		jwt: jwt,
	}
}
