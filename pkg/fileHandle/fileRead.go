package filehandle

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func ReadFile(username string, file multipart.File, routines int) (utils.Summary, error) {

	var fileContent bytes.Buffer
	_, err := io.Copy(&fileContent, file)
	if err != nil {
		return utils.Summary{}, err
	}

	// Process file
	result := ProcessFile(fileContent.String(), routines)

	return result, nil

}
