package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prithuadhikary/let-us-go/authorisation-service/service"
	"github.com/prithuadhikary/let-us-go/common/model"
	"github.com/prithuadhikary/let-us-go/common/util"
	"net/http"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
}

func (controller *loginController) Login(ctx *gin.Context) {
	request := &model.LoginRequest{}
	var validationError validator.ValidationErrors
	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validationError) {
		util.RenderBindingErrors(ctx, validationError)
		return
	}
	response, err := controller.loginService.Login(request)
	if err != nil {
		util.RenderServiceError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func NewLoginController(engine *gin.Engine, loginService service.LoginService) {
	controller := &loginController{
		loginService: loginService,
	}
	api := engine.Group("login")
	{
		api.POST("", controller.Login)
	}
}
