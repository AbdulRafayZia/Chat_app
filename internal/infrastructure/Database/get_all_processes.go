package database

import (
	"log"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func GetAllProcesses() []utils.ProcessesResponse {
	db := dbinit.OpenDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	record := make([]utils.ProcessesResponse, 0)
	for rows.Next() {
		var response utils.ProcessesResponse

		err := rows.Scan(&response.Id, &response.Filename, &response.TotalWords, &response.TotalLines, &response.TotalPuncuations, &response.TotalVowels, &response.Routines, &response.Username, &response.ExecutionTime)
		if err != nil {
			log.Fatal(err)
		}
		record = append(record, response)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return record

}
