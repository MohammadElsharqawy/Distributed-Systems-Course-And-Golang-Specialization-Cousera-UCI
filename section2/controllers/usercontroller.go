package controllers

import (
	"encoding/json"
	"net/http"
	"section2/packages/models"
	"strconv"
)

type usercontroller struct {
}

func (uc *usercontroller) ParseRequest(req *http.Request) (models.User, error) {
	dec := json.NewDecoder(req.Body)
	user := models.User{}
	err := dec.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (uc *usercontroller) post(req *http.Request, res http.ResponseWriter) {
	user, _ := uc.ParseRequest(req)
	user, _ = models.AddUser(user)

	encodeResponseAsJSON(user, res)
}

func (*usercontroller) get(id int, res http.ResponseWriter) {
	user, _ := models.GetUserById(id)
	encodeResponseAsJSON(user, res)
}

func (uc usercontroller) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/users" {
		switch req.Method {
		case http.MethodGet:
			userId, _ := strconv.Atoi(req.URL.Query()["id"][0])
			uc.get(userId, res)
		case http.MethodPost:
			uc.post(req, res)
		}

	} else {
		res.WriteHeader(http.StatusNotImplemented)
	}
}

func newUserController() *usercontroller {
	return &usercontroller{}
}
