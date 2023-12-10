package validation

import "log"

func VerifyPassword(dbPassword, password string) bool {

	if dbPassword == password {
		return true
	} else {
		log.Printf("error in verify hashed password:")
		return false
	}

}
