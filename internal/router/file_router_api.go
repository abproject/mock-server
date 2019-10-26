package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/abproject/mock-server/internal/models"
)

var fileURL = "/_api/files"
var maxFileSize int64 = 20 << 20

// RouteFileAPI Rest API
func RouteFileAPI(c models.AppContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == fileURL {
		switch r.Method {
		case "GET":
			getAllFileHandlers(c, w, r)
			return
		case "POST":
			postFileHandler(c, w, r)
			return
		case "DELETE":
			deleteAllFileHandlers(c, w, r)
			return
		}
	} else if strings.HasPrefix(r.RequestURI, fileURL) {
		reg, _ := regexp.Compile(fileURL + "/(.*)")
		groups := reg.FindStringSubmatch(r.RequestURI)
		if len(groups) == 2 {
			id := groups[1]
			switch r.Method {
			case "GET":
				getFileHandler(c, w, r, id)
				return
			case "PUT":
				putFileHandler(c, w, r, id)
				return
			case "DELETE":
				deleteFileHandler(c, w, r, id)
				return
			}
		}
	}
	notFoundHandler(c, w, r)
}

func getAllFileHandlers(c models.AppContext, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.FileStorage).GetAll())
}

func getFileHandler(c models.AppContext, w http.ResponseWriter, r *http.Request, id string) {
	fileDto, err := (*c.FileStorage).Get(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	json.NewEncoder(w).Encode(fileDto)
}

func postFileHandler(c models.AppContext, w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		errorHandler(w, err)
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		errorHandler(w, err)
		return
	}
	fileDto := (*c.FileStorage).Add(handler.Filename, bytes)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fileDto)
}

func putFileHandler(c models.AppContext, w http.ResponseWriter, r *http.Request, id string) {
	r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		errorHandler(w, err)
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		errorHandler(w, err)
		return
	}
	fileDto := (*c.FileStorage).Put(id, handler.Filename, bytes)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fileDto)
}

func deleteFileHandler(c models.AppContext, w http.ResponseWriter, r *http.Request, id string) {
	err := (*c.FileStorage).Delete(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAllFileHandlers(c models.AppContext, w http.ResponseWriter, r *http.Request) {
	(*c.FileStorage).DeleteAll()
	w.WriteHeader(http.StatusNoContent)
}
