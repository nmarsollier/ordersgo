package security

import (
	"time"

	"github.com/nmarsollier/ordersgo/log"
	"github.com/nmarsollier/ordersgo/tools/errs"
	gocache "github.com/patrickmn/go-cache"
)

var cache = gocache.New(60*time.Minute, 10*time.Minute)

// User es el usuario logueado
type User struct {
	ID          string   `json:"id"  validate:"required"`
	Name        string   `json:"name"  validate:"required"`
	Permissions []string `json:"permissions"`
	Login       string   `json:"login"  validate:"required"`
}

// Validate valida si el token es valido
func Validate(token string) (*User, error) {
	// Si esta en cache, retornamos el cache
	if found, ok := cache.Get(token); ok {
		if user, ok := found.(*User); ok {
			return user, nil
		}
	}

	user, err := getRemoteToken(token)
	if err != nil {
		return nil, errs.Unauthorized
	}

	// Todo bien, se agrega al cache y se retorna
	cache.Set(token, user, gocache.DefaultExpiration)

	return user, nil
}

// Invalidate invalida un token del cache
func Invalidate(token string, ctx ...interface{}) {
	if len(token) <= 7 {
		log.Get(ctx...).Info("Token no valido: ", token)
		return
	}

	cache.Delete(token[7:])
	log.Get(ctx...).Info("Token invalidado: ", token)
}
