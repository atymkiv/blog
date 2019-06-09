package mock

// Secure mock
type Secure struct {
	TokenFn func(string) string
}

// Token mock
func (s *Secure) Token(token string) string {
	return s.TokenFn(token)
}
