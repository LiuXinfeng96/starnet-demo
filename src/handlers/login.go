package handlers

import (
	"starnet-demo/src/models"
	"starnet-demo/src/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func JWTAuthMiddleware(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			var req models.TokenReq
			if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
				ParamsMissingJSONResp("the token not found", c)
				c.Abort()
				return
			}
			token = req.Token
		}

		claims, err := s.ParseToken(token)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			c.Abort()
			return
		}

		c.Set("token", claims)
		c.Next()
	}
}
