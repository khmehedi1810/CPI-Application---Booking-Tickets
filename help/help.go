package help

import "strings"

func UserinVal(fname string, lname string, email string, tik uint, uptik uint) (bool, bool, bool) {
	namVal := len(fname) >= 2 && len(lname) >= 2
	emVal := strings.Contains(email, "@")
	tickVal := tik <= uptik

	return namVal, emVal, tickVal
}
