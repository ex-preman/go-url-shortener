package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ExPreman/url-shortener-go/helper"
	"github.com/ExPreman/url-shortener-go/storage"
)

func EncodeHandler(storage storage.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			response(nil, http.StatusNotFound, "Method not allowed", w)
			return
		}
		if r.URL.Path != "/shorten" {
			response(nil, http.StatusNotFound, "URL not found", w)
			return
		}

		// get body as object
		var payload EncodePayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response(nil, http.StatusBadRequest, err.Error(), w)
			return
		}

		// validate body
		_, err = url.ParseRequestURI(payload.URL)
		if err != nil {
			response(nil, http.StatusBadRequest, "Invalid URL, please input correct URL", w)
			return
		}

		// generate code
		code, err := helper.GenerateRandomString(7)
		if err != nil {
			response(nil, http.StatusBadRequest, err.Error(), w)
			return
		}

		// save to storage
		err = storage.Save(payload.URL, code)
		if err != nil {
			response(nil, http.StatusBadRequest, err.Error(), w)
			return
		}

		respBody := EncodeResponse{
			ShortenURL: getURLOrigin(r) + "/" + code,
		}
		response(respBody, http.StatusOK, "", w)
	}

	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storage.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			response(nil, http.StatusNotFound, "Method not allowed", w)
			return
		}

		// get code from path
		code := r.URL.Path[len("/"):]

		// get url from storage
		url, err := storage.Load(code)
		if err != nil {
			response(nil, http.StatusNotFound, "URL not found", w)
			return
		}

		// redirecting
		log.Print("redirecting to: ", url)
		http.Redirect(w, r, string(url), 301)
	}

	return http.HandlerFunc(handleFunc)
}

func response(data interface{}, statusCode int, errorMessage string, w http.ResponseWriter) {
	type Status struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Set Body
	json.NewEncoder(w).Encode(
		struct {
			Status Status      `json:"status"`
			Data   interface{} `json:"data"`
		}{
			Status: Status{
				Code:  statusCode,
				Error: errorMessage,
			},
			Data: data,
		},
	)
}

func getURLOrigin(r *http.Request) string {
	protocol := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s", protocol, r.Host)
}
