package utils

type ResponseBody struct {
	Username         string  `json:"username"`
	Filename         string  `json:"file_name"`
	TotalLines       int     `json:"Total_lines"`
	TotalWords       int     `json:"Total_words"`
	TotalPuncuations int     `json:"Total_puncuations"`
	TotalVowels      int     `json:"Total_vowels"`
	ExecutionTime    float64 `json:"Execution_Time"`
	Routines         int     `json:"No_of_Routines"`

	// Id               int    `json:"id"`

}

type ProcessesResponse struct {
	Id               int     `json:"id"`
	Username         string  `json:"username"`
	Filename         string  `json:"file_name"`
	TotalLines       int     `json:"Total_lines"`
	TotalWords       int     `json:"Total_words"`
	TotalPuncuations int     `json:"Total_puncuations"`
	TotalVowels      int     `json:"Total_vowels"`
	ExecutionTime    float64 `json:"Execution_Time"`
	Routines         int     `json:"No_of_Routines"`
}

type ExecutionData struct {
	AveragTime float64 `json:"average_execution_seconds"`
}

type Token struct {
	AccessToken string `json:"acessToken"`
	RefreshToken string `json:"refreshToken"`

}
