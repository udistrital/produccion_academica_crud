package controllers

import (
	"encoding/json"
	
	"github.com/udistrital/produccion_academica_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// operations for TrProduccionAcademica
type TrProduccionAcademicaController struct {
	beego.Controller
}

func (c *TrProduccionAcademicaController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrProduccionAcademica
// @Description create the TrProduccionAcademica
// @Param	body		body 	models.TrProduccionAcademica	true	"body for TrProduccionAcademica content"
// @Success 201 {int} models.TrProduccionAcademica
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrProduccionAcademicaController) Post() {
	var v models.TrProduccionAcademica
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTransaccionProduccionAcademica(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
