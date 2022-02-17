package data

var serialNumber []string = []string{}

func CheckKey(userKey string) bool {
	for _, value := range serialNumber {
		if value == userKey {
			return true
		}
	}
	return false
}

func AddKey(userKey string) {
	serialNumber = append(serialNumber, userKey)
}
