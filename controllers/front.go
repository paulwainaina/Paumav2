package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterContorllers() {
	userController := NewUserController()
	http.HandleFunc("/users",userController.ServeHTTP)
	http.HandleFunc("/users/",userController.ServeHTTP)
}

func encodeResponseAsJson(data interface{},w io.Writer){
	enc:=json.NewEncoder(w)
	enc.Encode(data)
}