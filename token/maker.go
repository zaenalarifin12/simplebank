package token

import "time"

// Maker is an interface to manage token
type Maker interface {

	// createToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// verifyToken check if the token is valid or nor
	VerifyToken(token string) (*Payload, error)
}
