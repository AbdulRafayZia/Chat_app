package database

import (
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func InsertData(responseBody utils.ResponseBody) error {
	db := dbinit.OpenDB()
	_, err := db.Exec("INSERT INTO file_processing_data( filename, words, lines, punctuations, vowels, execution_time, routines , username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", responseBody.Filename, responseBody.TotalWords, responseBody.TotalLines, responseBody.TotalPuncuations, responseBody.TotalVowels, responseBody.ExecutionTime, responseBody.Routines, responseBody.Username)
	if err != nil {

		fmt.Println(err)
		return err
	}
	return nil

}
