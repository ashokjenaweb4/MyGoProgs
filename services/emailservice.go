package emailservice

//GetServerIP returns server IP address
func GetServerIP(servername string) string {
	if servername == "localhost" {
		return "127.0.0.1"
	}
	return "no ip found";
}