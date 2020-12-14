package provider

import (
	"{{hade_project_name}}/app/provider/demo"

	"github.com/gohade/hade/framework"
)

// customer provider
func RegisterCustomProvider(c framework.Container) {
	c.Bind(&demo.DemoProvider{}, true)
}
