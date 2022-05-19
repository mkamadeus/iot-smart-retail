package storefront

import (
	"encoding/json"
	"github.com/mkamadeus/iot-smart-retail/models"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (s *Service) UploadStoreLayout(fh *multipart.FileHeader) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileLoc := filepath.Join(dir, "layout", "storefrontlayout.json")
	targetLoc, err := os.OpenFile(fileLoc, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer targetLoc.Close()

	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(targetLoc, file); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStoreLayout() (models.StorefrontLayout, error) {
	var result models.StorefrontLayout
	dir, err := os.Getwd()
	if err != nil {
		return result, err
	}

	fileLoc := filepath.Join(dir, "layout", "storefrontlayout.json")

	content, err := os.ReadFile(fileLoc)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(content, &result); err != nil {
		return result, err
	}

	return result, nil
}
