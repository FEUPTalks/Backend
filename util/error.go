package util

import (
	"log"
	"net/http"
)

//ErrHandler func used to handle errors and writing http responses
func ErrHandler(err error, writer http.ResponseWriter, httpStatus int) {
	log.Println(err)
	http.Error(writer, err.Error(), httpStatus)
}
