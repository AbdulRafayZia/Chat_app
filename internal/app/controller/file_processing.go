package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	filehandle "github.com/AbdulRafayZia/Gorilla-mux/pkg/fileHandle"
	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
)

func ProcessFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return
	}

	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "could not Get Claims")
		return
	}
	responseBody, err := filehandle.GetFormData(w, r, claims)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	//Insert response into Database
	err = database.InsertData(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "failed to INSERT file Data")
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

}
