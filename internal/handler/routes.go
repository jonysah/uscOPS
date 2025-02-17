// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	User "uscOPS/internal/handler/User"
	"uscOPS/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: User.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
