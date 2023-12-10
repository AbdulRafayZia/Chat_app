package filehandle

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func GetFormData(w http.ResponseWriter, r *http.Request, claims *utils.MyClaims) (utils.ResponseBody, error) {
	startTime := time.Now()

	err := r.ParseMultipartForm(10000 << 20) // 10000 MB max file size
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}

	stringRoutines := r.FormValue("routines")
	routines, err := strconv.Atoi(stringRoutines)

	if err != nil {
		http.Error(w, "Invalid routines value", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}
	fmt.Printf("routienes :%d\n", routines)

	// Get file from form data
	file, FileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}
	defer file.Close()

	response, err := ReadFile(claims.Username, file, routines)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not getting response")
		return utils.ResponseBody{}, err

	}

	endTime := time.Now()
	// Calculate the execution time
	executionTime := endTime.Sub(startTime)

	responseBody := utils.ResponseBody{
		TotalLines:       response.LineCount,
		TotalWords:       response.WordsCount,
		TotalVowels:      response.VowelsCount,
		TotalPuncuations: response.PuncuationsCount,
		ExecutionTime:    executionTime.Seconds(),
		Routines:         routines,
		Filename:         FileHeader.Filename,
		Username:         claims.Username,
	}
	return responseBody, nil
}
