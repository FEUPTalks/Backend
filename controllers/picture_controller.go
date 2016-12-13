package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"io/ioutil"
	"net/http"

	"errors"

	"image"

	"strings"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/util"
)

const (
	filepath string = "Users/Pedro/Documents/"
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
	/*file, info, err := request.FormFile("picture")

	log.Println(request.FormFile("picture"));

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	defer file.Close()

	contentType := info.Header.Get("Content-Type")

	if !okContentType(contentType) {
		util.ErrHandler(errors.New("Invalid file type. Use jpeg or png"), writer, http.StatusUnsupportedMediaType)
		return
	}*/

	pictureDTO := model.PictureDTO{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&pictureDTO)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	sEnc, err := base64.StdEncoding.DecodeString(pictureDTO.Picture)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	buffer := []byte(sEnc)

	_, _, err = image.Decode(bytes.NewReader(buffer))
	if err != nil {
		util.ErrHandler(err, writer, http.StatusUnsupportedMediaType)
		return
	}

	filename, err := getFileName(pictureDTO.SpeakerName)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filename, buffer, 0600)

	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	id, err := instance.SavePicture(filename)
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	util.SendJSON(writer, request, id, http.StatusCreated)
}

//Download download files from the server
func (*PictureController) Download(writer http.ResponseWriter, request *http.Request) {
	id := request.FormValue("pictureID")
	if id == "" {
		util.ErrHandler(errors.New("Picture not found"), writer, http.StatusNotFound)
		return
	}

	instance, err := database.GetTalkDatabaseManagerInstance()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusNotFound)
		return
	}

	filepath, err := instance.GetPicture(id)

	buffer, err := ioutil.ReadFile(filepath)

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

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", allowedTypes[2])
	writer.Write(buffer)
}
