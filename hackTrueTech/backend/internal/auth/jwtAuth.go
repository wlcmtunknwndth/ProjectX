package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wlcmtunknwndth/hackBPA/internal/lib/httpResponse"
	"net/http"
	"os"
	"time"
)

type Info struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

const (
	access             = "access"
	ttlToken           = 4 * time.Minute
	statusUnauthorized = "Unauthorized"
	statusBadRequest   = "Bad request"
	authEnv            = "auth_key"
)

func checkRequest(r *http.Request) (*Info, error) {
	const op = "auth.jwtAuth.checkRequest"
	cookie, err := r.Cookie(access)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, fmt.Errorf("%s: No cookie: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var info Info

	token, err := jwt.ParseWithClaims(cookie.Value, &info, func(token *jwt.Token) (any, error) {
		key, ok := os.LookupEnv(authEnv)
		if !ok {
			return nil, fmt.Errorf("no secret key found")
		}
		return []byte(key), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("%s: Invalid jwt signature: %w", op, err)
		}
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("%s: Invalid token", op)
	}

	return &info, err
}

func IsAdmin(r *http.Request) (bool, error) {
	const op = "auth.jwtAuth.IsAdmin"
	inf, err := checkRequest(r)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return inf.IsAdmin, nil
}

func Access(r *http.Request) (bool, error) {
	const op = "auth.jwtAuth.Access"
	_, err := checkRequest(r)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, err
}

func Refresh(w http.ResponseWriter, r *http.Request) error {
	const op = "auth.jwtAuth.Refresh"

	info, err := checkRequest(r)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	expiresAt := time.Now().Add(ttlToken)
	info.ExpiresAt = jwt.NewNumericDate(expiresAt)

	//token := jwt.NewWithClaims(jwt.SigningMethodHS512, info)
	key, ok := os.LookupEnv(authEnv)
	if !ok {
		return fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, info).SignedString([]byte(key))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    access,
		Value:   token,
		Expires: expiresAt,
	})
	return nil
}

func WriteNewToken(w http.ResponseWriter, usr User) {
	var expiresAt = time.Now().Add(ttlToken)

	inf := &Info{
		Username: usr.Username,
		IsAdmin:  usr.isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	key, ok := os.LookupEnv(authEnv)
	if !ok {
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, inf).SignedString([]byte(key))
	if err != nil {
		httpResponse.Write(w, http.StatusInternalServerError, internalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    access,
		Value:   token,
		Expires: expiresAt,
	})
}
