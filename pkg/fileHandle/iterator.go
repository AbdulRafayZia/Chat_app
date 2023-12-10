package filehandle

import 	 "github.com/AbdulRafayZia/Gorilla-mux/utils"




func Counts(data string, channal chan  utils.Summary) {
	DocCounts := utils.Summary{}

	for _, char := range data {
		switch {
		case char == '\n':
			DocCounts.LineCount++
		case char == 32:
			DocCounts.WordsCount++
		case char == 65 || char == 69 || char == 73 || char == 79 || char == 85 || char == 97 || char == 101 || char == 105 || char == 111 || char == 117:
			DocCounts.VowelsCount++
		case (char >= 33 && char <= 47) || (char >= 58 && char <= 64) || (char >= 91 && char <= 96) || (char >= 123 && char <= 126):
			DocCounts.PuncuationsCount++
		}
	}

	channal <- DocCounts
}