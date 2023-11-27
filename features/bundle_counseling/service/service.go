package service

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
)

type BundleCounselingService struct {
	d bundlecounseling.BundleCounselingDataInterface
}

func New(data bundlecounseling.BundleCounselingDataInterface) bundlecounseling.BundleCounselingServiceInterface {
	return &BundleCounselingService{
		d: data,
	}
}
