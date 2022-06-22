package usrdata

type User struct {
	Id         string `json:"userid"`
	Name       string `json:"username"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
}
