package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["academica/controllers:FacultadController"] = append(beego.GlobalControllerRouter["academica/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:FacultadController"] = append(beego.GlobalControllerRouter["academica/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:FacultadController"] = append(beego.GlobalControllerRouter["academica/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:FacultadController"] = append(beego.GlobalControllerRouter["academica/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:FacultadController"] = append(beego.GlobalControllerRouter["academica/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["academica/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"] = append(beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"] = append(beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"] = append(beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"] = append(beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"] = append(beego.GlobalControllerRouter["academica/controllers:TipoCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
