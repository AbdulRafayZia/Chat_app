package utils

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Role     string `json:"role"`

}

type StatsRequest struct {
	Filename string `json:"filename"`
}
type Username struct{
	Username string `json:"username"`
}
