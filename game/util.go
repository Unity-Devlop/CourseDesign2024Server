package game

func GetStableHash(str string) int {
	hash := 23
	for _, c := range str {
		hash = hash*31 + int(c)
	}
	return hash
}

func PasswordCheck(str string) bool {
	if len(str) < 6 || len(str) > 20 {
		return false
	}
	return true
}
