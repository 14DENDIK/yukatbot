package methods

// Method ...
type Method struct {
	token string
}

// New ...
func New(token string) *Method {
	return &Method {
		token: token,
	}
}