package controllers

import (
	"encoding/json"
	"strconv"
	
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
	c.Mapping("GetAllByEnte", c.GetAllByEnte)
	c.Mapping("Delete",c.Delete)
	c.Mapping("Put",c.Put)
}


// GetAllByEnte ...
// @Title Get All By Ente
// @Description get TrProduccionAcademicaController
// @Param	id		path 	string	true		"Ente"
// @Success 200 {object} models.TrProduccionAcademicaController
// @Failure 404 not found resource
// @router /:ente [get]
func (c *TrProduccionAcademicaController) GetAllByEnte() {
	idEnteStr := c.Ctx.Input.Param(":ente")
	idEnte, _ := strconv.Atoi(idEnteStr)
	l, err := models.GetProduccionesAcademicasByEnte(idEnte)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
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

// Put ...
// @Title Put
// @Description update the TrProduccionAcademica
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrProduccionAcademica	true		"body for TrProduccionAcademica content"
// @Success 200 {object} models.TrProduccionAcademica
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TrProduccionAcademicaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var v models.TrProduccionAcademica
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.ProduccionAcademica.Id = id
		if err := models.UpdateTransaccionProduccionAcademica(&v); err == nil {
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

// Delete ...
// @Title Delete
// @Description delete the ProduccionAcademica
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TrProduccionAcademicaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.TrDeleteProduccionAcademica(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
