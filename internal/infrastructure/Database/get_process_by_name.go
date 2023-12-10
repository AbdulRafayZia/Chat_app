package database

import (
	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func GetProcessesByUserName( username string) ([]utils.ProcessesResponse , error) {
	db := dbinit.OpenDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data WHERE username = $1 ", username)
	if err != nil {
		return nil , err
	}

	defer rows.Close()
	record := make([]utils.ProcessesResponse, 0)
	for rows.Next() {
		var response utils.ProcessesResponse


		err := rows.Scan(&response.Id, &response.Filename, &response.TotalWords, &response.TotalLines, &response.TotalPuncuations, &response.TotalVowels, &response.Routines, &response.Username, &response.ExecutionTime)
		if err != nil {
			return nil , err
		}
		record = append(record, response)

	}
	if err := rows.Err(); err != nil {
		return nil , err
	}
	return record,nil

}
