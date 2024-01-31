package game

func GetStableHash(str string) int {
	hash := 23
	for _, c := range str {
		hash = hash*31 + int(c)
	}
	return hash
}
