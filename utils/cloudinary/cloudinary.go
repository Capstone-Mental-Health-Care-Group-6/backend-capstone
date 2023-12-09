package cloudinary

import (
	"FinalProject/configs"
	"context"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sirupsen/logrus"
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
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	cl, err := cloudinary.NewFromURL(cld.cfg.CloudinaryURL)
	//2TFjfPK1yqcIZ1hEKTZ_cZQsLlc
	logrus.Info("Ini input: ", input)

	//DEV MODE
	if err != nil {
		return "", err
	}

	uploadParam, err := cl.Upload.Upload(ctx, input, uploader.UploadParams{Folder: "Avatar"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("Sampe sini 2")
	logrus.Info("Ini payment proof: ", uploadParam.SecureURL)
	// logrus.Info("Ini map payment: ", uploadParam)

	return uploadParam.SecureURL, nil
}
