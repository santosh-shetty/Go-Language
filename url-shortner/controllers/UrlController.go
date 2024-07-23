package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type URLStruct struct {
	Url string `json:"url"`
}

var UrlDB map[string]string

func init() {
	UrlDB = make(map[string]string)
}
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "view/index.html")
}
func Shortner(res http.ResponseWriter, req *http.Request) {
	var url URLStruct

	json.NewDecoder(req.Body).Decode(&url)

	shortUrl := generateShortKey(url.Url)
	UrlDB[shortUrl] = url.Url

	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(shortUrl)
}

func generateShortKey(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[8:])
}

func RedirectUrl(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortUrl := vars["url"]
	redirectUrl := UrlDB[shortUrl]

	http.Redirect(res, req, redirectUrl, http.StatusFound)
}
