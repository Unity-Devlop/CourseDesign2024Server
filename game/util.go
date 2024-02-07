package game

func GetStableHash(str string) int {
	hash := 23
	for _, c := range str {
		hash = hash*31 + int(c)
	}
	return hash
}

func PasswordCheck(str string, reason *string) bool {
	if len(str) < 6 {
		*reason = "Password length must be greater than 6"
		return false
	}
	if len(str) > 20 {
		*reason = "Password length must be less than 20"
		return false
	}
	return true
}

func CharacterNameCheck(str string, reason *string) bool {
	if len(str) < 1 {
		*reason = "Character name length must be greater than 1"
		return false
	}
	if len(str) > 7 {
		*reason = "Character name length must be less than 7"
		return false
	}
	return true
}
