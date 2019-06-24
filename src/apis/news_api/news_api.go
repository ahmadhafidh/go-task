package news_api

import (
	"config"
	"encoding/json"
	"models"
	"net/http"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		newsModel := models.NewsModel{
			Db: db,
		}
		newsnew, err2 := newsModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, newsnew)
		}
	}
}

// func Search(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	keyword := vars["keyword"]
// 	db, err := config.GetDB()
// 	if err != nil {
// 		respondWithError(response, http.StatusBadRequest, err.Error())
// 	} else {
// 		newsModel := models.NewsModel{
// 			Db: db,
// 		}
// 		newsnew, err2 := newsModel.Search(keyword)
// 		if err2 != nil {
// 			respondWithError(response, http.StatusBadRequest, err2.Error())
// 		} else {
// 			respondWithJson(response, http.StatusOK, newsnew)
// 		}
// 	}
// }

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
