package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hollson/kendo/infrastructure/errorext"
)

func Cors() gin.HandlerFunc {
	conf := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowOriginFunc:  func(str string) bool { return true },
		MaxAge:           12 * time.Hour,
	}
	return cors.New(conf)
}

func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only run if there are some errors to handle
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			if err != nil {
				var listErr *errorext.ListMsgError
				if errors.As(err, &listErr) {
					c.JSON(http.StatusBadRequest, gin.H{"error": listErr.MsgMap})
					c.Abort()
					return
				}

				var validatErr validator.ValidationErrors
				if errors.As(err, &validatErr) {
					list := make(map[string]string)
					for _, err := range validatErr {
						list[strings.ToLower(errorext.ToSnakeCase(err.Field()))] = errorext.ValidationErrorToText(err)
					}
					c.JSON(http.StatusBadRequest, gin.H{"error": list})
					c.Abort()
					return
				}

				var codeErr *errorext.CodeError
				if errors.As(err, &codeErr) {
					c.JSON(http.StatusConflict, gin.H{"code": codeErr.Code, "msg": codeErr.Msg})
					c.Abort()
					return
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
}
