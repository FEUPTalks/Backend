package model

//PictureDTO picture dto for upload
type PictureDTO struct {
	PictureID	  int64  `json:"pictureID,omitempty"`
	SpeakerPicture    string `json:"speakerPicture"`
}
