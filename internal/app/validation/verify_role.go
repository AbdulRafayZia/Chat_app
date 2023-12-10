package validation

func CheckStaffRole(role string) bool {

	if role == "admin" {
		return true
	} else {
		return false
	}

}
func CheckUserRole(role string) bool {

	if role == "user" {
		return true
	} else {
		return false
	}

}
