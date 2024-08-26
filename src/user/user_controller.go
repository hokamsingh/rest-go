package user

import (
	"net/http"

	LessGo "github.com/hokamsingh/lessgo/pkg/lessgo"
)

// UserController is the controller for handling user-related requests.
type UserController struct {
	Path    string
	Service UserService
}

// NewUserController creates a new instance of UserController.
func NewUserController(service *UserService, path string) *UserController {
	return &UserController{
		Service: *service,
		Path:    path,
	}
}

// RegisterRoutes registers the user routes with the given router.
func (uc *UserController) RegisterRoutes(r *LessGo.Router) {
	tr := r.SubRouter(uc.Path)

	//	@Summary		Ping the server
	//	@Description	Simple ping to check server status
	//	@Tags			User
	//	@Produce		plain
	//	@Success		200	{string}	string	"pong"
	//	@Router			/user/ping [get]
	tr.Get("/ping", func(ctx *LessGo.Context) {
		ctx.Send("pong")
	})

	//	@Summary		Get user by ID
	//	@Description	Fetch user information based on ID
	//	@Tags			User
	//	@Param			id	path	string	true	"User ID"
	//	@Produce		json
	//	@Success		200	{object}	map[string]interface{}
	//	@Failure		400	{object}	map[string]string
	//	@Router			/user/user/{id} [get]
	tr.Get("/user/{id}", func(ctx *LessGo.Context) {
		params, ok := ctx.GetAllParams()
		id := params["id"]
		if !ok {
			ctx.Error(400, "no params found")
			return
		}
		queryParams, _ := ctx.GetAllQuery()
		ctx.SetHeader("X-Custom-Header", "MyValue")
		cookie, ok := ctx.GetCookie("auth_token")
		if !ok {
			ctx.SetCookie("auth_token", "0xc000013a", 60, "", true, false, http.SameSiteDefaultMode)
		}
		ctx.JSON(200, map[string]interface{}{
			"params":      params,
			"queryParams": queryParams,
			"id":          id,
			"cookie":      cookie,
		})
	})

	//	@Summary		Submit a new user
	//	@Description	Submit user data
	//	@Tags			User
	//	@Accept			json
	//	@Produce		json
	//	@Param			user	body		User	true	"User Data"
	//	@Success		200		{object}	User
	//	@Router			/user/submit [post]
	tr.Post("/submit", func(ctx *LessGo.Context) {
		var body User
		ctx.Body(&body)
		ctx.JSON(200, body)
	})

	//	@Summary		Delete a user by ID
	//	@Description	Delete user based on ID
	//	@Tags			User
	//	@Param			id	path		string	true	"User ID"
	//	@Failure		400	{object}	map[string]string
	//	@Router			/user/{id} [delete]
	tr.Delete("/{id}", func(ctx *LessGo.Context) {
		id, _ := ctx.GetParam("id")
		ctx.Error(400, id)
	})
}
