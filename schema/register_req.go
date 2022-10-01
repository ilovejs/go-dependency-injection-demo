package schema

// RegisterReq request
type RegisterReq struct {
	Username string
	Password string
}

// RegisterRes response
type RegisterRes struct {
	Msg string
}
