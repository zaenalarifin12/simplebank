package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// Different types of error returned by the verifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	//TODO implement me
	return jwt.NewNumericDate(payload.IssueAt), nil
}

func (payload *Payload) GetIssuer() (string, error) {
	//TODO implement me
	return "your_issuer_string", nil
}

func (payload *Payload) GetSubject() (string, error) {
	//TODO implement me
	return "your_subject_string", nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	//TODO implement me
	audience := jwt.ClaimStrings{"audience1", "audience2"} // Replace with your desired audiences
	return audience, nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssueAt), nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil

}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
