package security

import (
	"fmt"
	"log"
	"time"

	"github.com/VictorzllDev/ZorpPay/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	GetUserIDFromToken(tokenString string) (string, error)
}

type jwtService struct {
	secretKey  string
	issuer     string
	expiration time.Duration
}

type JWTClaims struct {
	sub string
	jwt.RegisteredClaims
}

func NewJWT() JWTService {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	return &jwtService{
		secretKey:  cfg.JWTSecretKey,
		issuer:     cfg.JWTIssuer,
		expiration: time.Hour * 1, // 1 hour
	}
}

func (j *jwtService) GenerateToken(userID string) (string, error) {
	claims := &JWTClaims{
		sub: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signature method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failure to decode the token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid CLAMS TYPE")
	}

	return claims, nil
}

func (j *jwtService) GetUserIDFromToken(tokenString string) (string, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.sub, nil
}
