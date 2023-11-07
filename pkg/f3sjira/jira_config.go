package f3sjira

type JiraConfig struct {
	Url  string
	Auth *BasicAuth
}

type BasicAuth struct {
	Usr string
	Pwd []byte
}
