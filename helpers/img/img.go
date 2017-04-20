package img

import (
	"errors"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	maxUploadSize    = int64(1 * 1024) // 1 mb
	ErrImageNotValid = errors.New("not a valid image")
	ErrImageToLarge  = errors.New("to large") // TODO: Include size in mb
)

func Save(photo multipart.File, handler *multipart.FileHeader, path string) error {
	defer photo.Close()
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, photo)
	if err != nil {
		return err
	}
	return nil
}

func Validate(header textproto.MIMEHeader) error {
	for _, mime := range mimes {
		if header.Get("Content-Type") == mime {
			size, err := strconv.ParseInt(header.Get("Content-Length"), 10, 64)

			if err != nil {
				return err
			}

			if size > maxUploadSize {
				return ErrImageToLarge
			}
			return nil
		}
	}
	return ErrImageNotValid
}
