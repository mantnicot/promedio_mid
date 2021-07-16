package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// EstudianteController operations for Estudiante
type EstudianteController struct {
	beego.Controller
}

// URLMapping ...
func (c *EstudianteController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Estudiante
// @Param	body		body 	models.Estudiante	true		"body for Estudiante content"
// @Success 201 {object} models.Estudiante
// @Failure 403 body is empty
// @router / [post]
func (c *EstudianteController) Post() {

	var estudi map[string]interface{}
	var res map[string]interface{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &estudi)

	if err := request.SendJson(beego.AppConfig.String("UrlCrud")+"/estudiante", "POST", &res, &estudi); err == nil {
		c.Data["json"] = res
	} else {
		logs.Error(err)
		c.Abort("404 fallo")
	}
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description get Estudiante by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Estudiante
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EstudianteController) GetOne() {
	Idestudi := c.Ctx.Input.Param(":id")
	var res map[string]interface{}

	if err := request.GetJson(beego.AppConfig.String("UrlCrud")+"/estudiante/"+Idestudi, &res); err == nil {
		c.Data["json"] = res
	} else {
		logs.Error(err)
		c.Abort("404 fallo")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Estudiante
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Estudiante
// @Failure 403
// @router / [get]
func (c *EstudianteController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Estudiante
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Estudiante	true		"body for Estudiante content"
// @Success 200 {object} models.Estudiante
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EstudianteController) Put() {
	Idestudi := c.Ctx.Input.Param(":id")
	var res map[string]interface{}

	if err := request.GetJson(beego.AppConfig.String("UrlCrud")+"/estudiante/"+Idestudi, &res); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "10", "Message": "Peticion completa", "Data": res}
	} else {
		logs.Error(err)
		c.Abort("404 fallo")
	}
	/*var def = helpers.CalcularDefinitiva(res.nota_1, res.nota_2, res.nota_3)
	valor := models.Estudiante{Id: id, nota_def: def}
	if err := request.SendJson(beego.AppConfig.String("UrlCrud")+"/estudiante/"+Idestudi, &res, valor); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "10", "Message": "Peticion completa", "Data": res}
	} else {
		logs.Error(err)
		c.Abort("400")
	}*/
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Estudiante
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EstudianteController) Delete() {

}
