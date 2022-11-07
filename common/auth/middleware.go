package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/prithuadhikary/let-us-go/common/errors"
	"github.com/prithuadhikary/let-us-go/common/util"
	"net/http"
	"strings"
)

const ClaimsContextVar = "claims"

// TokenValidator returns a gin.HanlerFunc which validates the JWT signature
// and sets a context variable containing the claims.
func TokenValidator(service TokenService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.GetHeader("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			splits := strings.Split(authorization, " ")
			if len(splits) != 2 {
				util.RenderServiceError(context, errors.NewServiceError(
					"invalid.authorization.header",
					"Invalid authorisation header encountered.",
					http.StatusNotAcceptable),
				)
				return
			}
			claims, err := service.Validate(splits[1])
			if err != nil {
				util.RenderServiceError(context, err)
				return
			}
			context.Set(ClaimsContextVar, claims)
		}
		context.Next()
	}
}
