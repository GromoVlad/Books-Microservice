package headers

type Headers struct {
	BearerToken string `header:"Authorization"`
}
