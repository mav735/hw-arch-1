package server

func Address(getenv func(string) string) string {
	if address := getenv("HTTP_ADDR"); address != "" {
		return address
	}
	return ":8080"
}
