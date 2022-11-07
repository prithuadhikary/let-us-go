package auth

import (
	"github.com/prithuadhikary/let-us-go/common/errors"
	"github.com/prithuadhikary/let-us-go/common/model"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"time"
)

type AppClaims struct {
	jwt.Claims
	Role string `json:"role,omitempty"`
}

type TokenService interface {
	Generate(user *model.User) (string, errors.ServiceError)
	Validate(token string) (*AppClaims, errors.ServiceError)
}

type tokenService struct {
	signingKey jose.SigningKey
}

func (service *tokenService) Generate(user *model.User) (string, errors.ServiceError) {
	signer, err := jose.NewSigner(service.signingKey, (&jose.SignerOptions{}).WithType("JWT"))
	builder := jwt.Signed(signer)
	appClaims := &AppClaims{
		Claims: jwt.Claims{
			Issuer:    "auth-service",
			Subject:   user.Username,
			Expiry:    jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role: user.Role,
	}
	token, err := builder.Claims(appClaims).CompactSerialize()
	if err != nil {
		return "", errors.InternalServerError()
	}
	return token, nil
}

func (service *tokenService) Validate(token string) (*AppClaims, errors.ServiceError) {
	jwt, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, errors.NewServiceError(
			"invalid.token",
			"invalid bearer token ecountered",
			http.StatusBadRequest,
		)
	}
	appClaims := &AppClaims{}
	err = jwt.Claims(service.signingKey.Key, appClaims)
	if err != nil {
		return nil, errors.NewServiceError(
			"invalid.signature",
			"Invalid token signature.",
			http.StatusBadRequest,
		)
	}
	return appClaims, nil
}

func NewTokenService(signingKeyBytes []byte) TokenService {
	signingKey := jose.SigningKey{
		Algorithm: jose.HS256,
		Key:       signingKeyBytes,
	}
	return &tokenService{
		signingKey: signingKey,
	}
}
