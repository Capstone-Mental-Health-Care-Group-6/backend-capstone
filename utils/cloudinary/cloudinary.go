package cloudinary

import (
	"FinalProject/configs"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"time"
)

type CloudinaryInterface interface {
	UploadImageHelper(input interface{}) (string, error)
}

type Cloudinary struct {
	cfg configs.ProgrammingConfig
}

func InitCloud(config configs.ProgrammingConfig) CloudinaryInterface {
	return &Cloudinary{
		cfg: config,
	}
}

func (cld *Cloudinary) UploadImageHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := cloudinary.NewFromURL(cld.cfg.CloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadParam, err := cl.Upload.Upload(ctx, input, uploader.UploadParams{Folder: "Avatar"})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
