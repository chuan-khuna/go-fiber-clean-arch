package entities

type Ping struct {
	// just for mocking purpose
	// not save to db
	Message string `json:"message"`
}
