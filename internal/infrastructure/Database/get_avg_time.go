package database

import (
	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func GetAvergeExeTime(request utils.StatsRequest) (utils.ExecutionData, error) {
	db := dbinit.OpenDB()

	defer db.Close()

	rows, err := db.Query("SELECT AVG(execution_time) FROM file_processing_data WHERE filename = $1", request.Filename)

	if err != nil {

		return utils.ExecutionData{}, err

	}

	defer rows.Close()
	var avgExecutionTime utils.ExecutionData

	// Check for rows
	for rows.Next() {

		// Scan the average execution time value into the variable
		if err := rows.Scan(&avgExecutionTime.AveragTime); err != nil {

			return utils.ExecutionData{}, err

		}
	}
	if err := rows.Err(); err != nil {

		return utils.ExecutionData{}, err

	}
	return avgExecutionTime, nil
}
