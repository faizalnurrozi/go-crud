package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtPkg "github.com/faizalnurrozi/go-crud/pkg/jwt"
	"github.com/faizalnurrozi/go-crud/pkg/messages"
	"github.com/faizalnurrozi/go-crud/pkg/str"
	"github.com/faizalnurrozi/go-crud/server/http/handlers"
	"github.com/faizalnurrozi/go-crud/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func (jwtMiddleware JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}
	handler := handlers.Handler{UcContract: jwtMiddleware.Contract}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		fmt.Println(err)
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			fmt.Println(err)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(jwtMiddleware.JwtCredential.TokenSecret)
		return secret, nil
	})
	if err != nil {
		fmt.Println(err)
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		fmt.Println(err)
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//jwe roll back encrypted id
	jweRes, err := jwtMiddleware.JweCredential.Rollback(claims.Payload)
	if err != nil {
		fmt.Println(err)
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}
	if jweRes == nil {
		fmt.Println(err)
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//set id to uce case contract
	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	roleID := fmt.Sprintf("%v", jweRes["roleID"])
	jwtMiddleware.Contract.UserID = claims.Id
	jwtMiddleware.Contract.RoleID = str.StringToInt(roleID)

	return ctx.Next()
}
