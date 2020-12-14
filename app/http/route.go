package http

import (
	"{{hade_project_name}}/app/http/module/demo"

	"github.com/gohade/hade/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
