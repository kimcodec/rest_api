package lib

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func ValidatePicture(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return err
	}

	if img.Height > 1080 || img.Width > 1920 {
		return errors.New("invalid picture size")
	}
	return nil
}
