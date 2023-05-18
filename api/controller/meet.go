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

//MeetController -> MeetController
type MeetController struct {
	service service.MeetService
}

//NewMeetController : NewMeetController
func NewMeetController(s service.MeetService) MeetController {
	return MeetController{
		service: s,
	}
}

// GetMeets : GetMeets controller

// GetMeets godoc
// @Summary get all meets
// @Schemes
// @Description get all meets
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags meets
// @Accept json
// @Produce json
// @Success 200 {Response} Response
// @Router /meets [get]
func (p MeetController) GetMeets(ctx *gin.Context) {
	var meets models.Meet

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(meets, keyword)

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
		Message: "Meet result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddMeet : AddMeet controller

// AddMeet godoc
// @Summary create a new meet
// @Schemes
// @Description get all meets
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param meet body models.Meet true "Meets model"
// @Tags meets
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /meets [post]
func (p *MeetController) AddMeet(ctx *gin.Context) {
	var meet models.Meet
	fmt.Println(meet)
	ctx.ShouldBindJSON(&meet)

	if meet.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	if meet.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
		return
	}

	err := p.service.Save(meet)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create meet")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Meet")
}

//GetMeet : get meet by id

// GetMeets godoc
// @Summary get meet by id
// @Schemes
// @Description get meet by id
// @Tags meets
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search meet by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /meets/{id} [get]
func (p *MeetController) GetMeet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var meet models.Meet
	meet.ID = uint(id)
	foundMeet, err := p.service.Find(meet)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Meet")
		return
	}

	response := foundMeet.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Meet",
		Data:    &response})

}

//DeleteMeet : Deletes Meet

// DeleteMeet godoc
// @Summary delete meet by id
// @Schemes
// @Description delete meet by id
// @Tags meets
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "delete meet by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /meets/{id} [delete]
func (p *MeetController) DeleteMeet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(uint(id))

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Meet")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateMeet : get update by id

// UpdateMeet godoc
// @Summary get meet by id
// @Schemes
// @Description get meet by id
// @Tags meets
// @Accept json
// @Param id path int true "search meet by id"
// @Param meet body models.Meet true "Meets model"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /meets/{id} [put]
func (p MeetController) UpdateMeet(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var meet models.Meet
	meet.ID = uint(id)

	meetRecord, err := p.service.Find(meet)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Meet with given id not found")
		return
	}
	ctx.ShouldBindJSON(&meetRecord)

	if meetRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if meetRecord.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(meetRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Meet")
		return
	}
	response := meetRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Meet",
		Data:    response,
	})
}
