package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
)

func Authorize(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the cookie
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Decode/Validate
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["aig"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Find the user
			var user models.User
			db.DB.First(&user, claims["sub"])

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Attach to req
			c.Set("user", user.Username)
			// c.Set("uid", user.NikID)
			c.Set("urole", user.Role)

			val, existed := c.Get("user")
			if !existed {
				c.AbortWithStatusJSON(401, "user hasn't logged in yet")
				return
			}

			// casbin enforces policy
			ok, err := enforce(val.(string), obj, act)
			if err != nil {
				c.AbortWithStatusJSON(500, "error occurred when authorizing user")
				return
			}
			if !ok {
				c.AbortWithStatusJSON(403, "forbidden")
				return
			}

			// Continue
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func enforce(sub string, obj string, act string) (bool, error) {
	// db.Enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to create casbin enforcer %w", err)
	// }

	err := db.Enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	ok, err := db.Enforcer.Enforce(sub, obj, act)
	return ok, err
}
