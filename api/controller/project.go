package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rogeriofontes/api-go-gin/api/service"
	"github.com/rogeriofontes/api-go-gin/models"
	"github.com/rogeriofontes/api-go-gin/util"
	_ "github.com/swaggo/swag/example/celler/httputil"

	"github.com/gin-gonic/gin"
)

//ProjectController -> ProjectController
type ProjectController struct {
	service service.ProjectService
}

//NewProjectController : NewProjectController
func NewProjectController(s service.ProjectService) ProjectController {
	return ProjectController{
		service: s,
	}
}

// GetProjects : GetProjects controller

// GetProjects godoc
// @Summary get all projects
// @Schemes
// @Description get all projects
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {Response} Response
// @Router /projects [get]
func (p ProjectController) GetProjects(ctx *gin.Context) {
	var projects models.Project

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(projects, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Project result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddProject : AddProject controller

// AddProject godoc
// @Summary create a new project
// @Schemes
// @Description get all projects
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param project body models.Project true "Projects model"
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /projects [post]
func (p *ProjectController) AddProject(ctx *gin.Context) {
	var project models.Project
	fmt.Println(project)
	ctx.ShouldBindJSON(&project)

	if project.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if project.Content == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	err := p.service.Save(project)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create project")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Project")
}

//GetProject : get project by id

// GetProjects godoc
// @Summary get project by id
// @Schemes
// @Description get project by id
// @Tags projects
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search project by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /projects/{id} [get]
func (p *ProjectController) GetProject(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var project models.Project
	project.ID = uint(id)
	foundProject, err := p.service.Find(project)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Project")
		return
	}

	response := foundProject.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Project",
		Data:    &response})

}

//DeleteProject : Deletes Project

// DeleteProject godoc
// @Summary delete project by id
// @Schemes
// @Description delete project by id
// @Tags projects
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "delete project by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /projects/{id} [delete]
func (p *ProjectController) DeleteProject(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(uint(id))

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Project")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateProject : get update by id

// UpdateProject godoc
// @Summary get project by id
// @Schemes
// @Description get project by id
// @Tags projects
// @Accept json
// @Param id path int true "search project by id"
// @Param project body models.Project true "Projects model"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /projects/{id} [put]
func (p ProjectController) UpdateProject(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var project models.Project
	project.ID = uint(id)

	projectRecord, err := p.service.Find(project)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Project with given id not found")
		return
	}
	ctx.ShouldBindJSON(&projectRecord)

	if projectRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if projectRecord.Content == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(projectRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Project")
		return
	}
	response := projectRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Project",
		Data:    response,
	})
}
