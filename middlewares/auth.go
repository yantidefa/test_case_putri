package middlewares

import (
	"net/http"
	"test_case_putri/constants"
	userrespository "test_case_putri/repositories/user_respository"
	"test_case_putri/utilities"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrEmptyAuthHeader, nil)
			c.Abort()
			return
		}
		err := utilities.ValidateToken(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
			c.Abort()
			return
		}

		getUser, err := userrespository.GetUserByIsTokenRepository(tokenString)
		if err != nil {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

		if !getUser.IsLogin {
			utilities.SetResponseJSON(c, http.StatusNotFound, nil, constants.ErrLogout, err)
		}

	}
}
