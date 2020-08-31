package pixie

//NewClient returns a client used for requests to the Pixie API
func NewClient(username string, password string) Client {
	return Client{
		Username: username,
		Password: password,
	}
}
