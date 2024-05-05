package util

import (
	"time"

	"loanManagement/constant"
	databaseModel "loanManagement/database/model"
	"loanManagement/logger"

	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Jwt interface {
	GenerateToken(user databaseModel.User) (string, error)
}

type jwt struct {
	signingKey string
}

// customClaims is not exported as this is only used inside this util
type customClaims struct {
	Type constant.UserType `json:"type"`
	jwtLib.RegisteredClaims
}

func (j *jwt) GenerateToken(user databaseModel.User) (string, error) {
	currentTime := time.Now()
	claims := customClaims{
		Type: user.Type,
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: jwtLib.NewNumericDate(currentTime.Add(time.Hour * 24)),
			IssuedAt:  jwtLib.NewNumericDate(currentTime),
			ID:        user.ID.String(),
		},
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		logger.Log.Error(err.Error(), logger.String("userID", user.ID.String()))
		return tokenString, errors.Wrap(err, "failed to generate token")
	}
	logger.Log.Info("token generated successfully", logger.String("userID", user.ID.String()))
	return tokenString, nil
}

func (j *jwt) ValidateToken(tokenString string) (databaseModel.User, error) {
	user := databaseModel.User{}
	token, err := jwtLib.ParseWithClaims(tokenString, &customClaims{}, func(token *jwtLib.Token) (interface{}, error) {
		return []byte(j.signingKey), nil
	})
	if err != nil {
		logger.Log.Error(err.Error())
		return user, errors.Wrap(err, "failed to parse token")
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		logger.Log.Error("failed to parse claims")
		return user, errors.New("failed to parse claims")
	}

	user.ID, err = uuid.Parse(claims.RegisteredClaims.ID)
	if err != nil {
		logger.Log.Error(err.Error())
		return user, errors.Wrap(err, "failed to parse userID")
	}
	user.Type = claims.Type

	return user, nil
}

func NewJwt(signingKey string) Jwt {
	return &jwt{
		signingKey: signingKey,
	}
}
