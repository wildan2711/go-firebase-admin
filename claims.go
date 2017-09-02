package firebase

// Claims is the firebase authentication token claims
type Claims struct {
	Issuer    string      `json:"iss,omitempty"`
	Subject   string      `json:"sub,omitempty"`
	Audience  string      `json:"aud,omitempty"`
	IssuedAt  int64       `json:"iat,omitempty"`
	ExpiresAt int64       `json:"exp,omitempty"`
	UserID    string      `json:"uid,omitempty"`
	Claims    interface{} `json:"claims,omitempty"`
}

func (c *Claims) verifyExpiresAt(now int64) bool {
	return now <= c.ExpiresAt
}

func (c *Claims) verifyIssuedAt(now int64) bool {
	return now >= c.IssuedAt
}

func (c *Claims) verifyAudience(aud string) bool {
	return c.Audience == aud
}

func (c *Claims) verifyIssuer(iss string) bool {
	return c.Issuer == iss
}
