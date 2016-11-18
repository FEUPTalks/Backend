package controllers

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"net/http"

	"errors"

	"image"

	"strings"

	"github.com/FEUPTalks/Backend/util"
)

const (
	filepath string = "C:\\Users\\Renato\\Desktop\\"
)

var allowedTypes = [...]string{"image/jpeg", "image/jpg", "image/png"}

//PictureController used to upload files
type PictureController struct {
}

func getFileName(speakerName string) (string, error) {
	if len(speakerName) < 1 {
		return "", errors.New("Speaker name can not be empty")
	}

	speakerName = strings.Replace(speakerName, " ", "", -1)

	return filepath + speakerName, nil
}

func okContentType(fileType string) bool {
	for _, allowed := range allowedTypes {
		if allowed == fileType {
			return true
		}
	}
	return false
}

//Upload upload files to the server
func (*PictureController) Upload(writer http.ResponseWriter, request *http.Request) {
	file, info, err := request.FormFile("picture")

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	defer file.Close()

	contentType := info.Header.Get("Content-Type")

	if !okContentType(contentType) {
		util.ErrHandler(errors.New("Invalid file type. Use jpeg or png"), writer, http.StatusUnsupportedMediaType)
		return
	}

	buffer, err := ioutil.ReadAll(file)

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	_, _, err = image.Decode(bytes.NewReader(buffer))
	if err != nil {
		util.ErrHandler(err, writer, http.StatusUnsupportedMediaType)
		return
	}

	filename, err := getFileName(request.FormValue("speakerName"))
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filename, buffer, 0600)

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

//Download download files from the server
func (*PictureController) Download(writer http.ResponseWriter, request *http.Request) {
	filename := ""

	buffer, err := ioutil.ReadFile(filename)

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(buffer))
	data := new(bytes.Buffer)

	err = png.Encode(data, img)

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	buffer = data.Bytes()

	writer.Header().Set("Content-Type", allowedTypes[2])
	writer.Write(buffer)
}
