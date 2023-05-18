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

//EventController -> EventController
type EventController struct {
	service service.EventService
}

//NewEventController : NewEventController
func NewEventController(s service.EventService) EventController {
	return EventController{
		service: s,
	}
}

// GetEvents : GetEvents controller

// GetEvents godoc
// @Summary get all events
// @Schemes
// @Description get all events
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {Response} Response
// @Router /events [get]
func (p EventController) GetEvents(ctx *gin.Context) {
	var events models.Event

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(events, keyword)

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
		Message: "Event result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddEvent : AddEvent controller

// AddEvent godoc
// @Summary create a new event
// @Schemes
// @Description get all events
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param event body models.Event true "Events model"
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /events [post]
func (p *EventController) AddEvent(ctx *gin.Context) {
	var event models.Event
	fmt.Println(event)
	ctx.ShouldBindJSON(&event)

	if event.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	if event.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Description is required")
		return
	}

	err := p.service.Save(event)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create event")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Event")
}

//GetEvent : get event by id

// GetEvents godoc
// @Summary get event by id
// @Schemes
// @Description get event by id
// @Tags events
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search event by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /events/{id} [get]
func (p *EventController) GetEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var event models.Event
	event.ID = uint(id)
	foundEvent, err := p.service.Find(event)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Event")
		return
	}

	response := foundEvent.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Event",
		Data:    &response})

}

//DeleteEvent : Deletes Event

// DeleteEvent godoc
// @Summary delete event by id
// @Schemes
// @Description delete event by id
// @Tags events
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "delete event by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /events/{id} [delete]
func (p *EventController) DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(uint(id))

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Event")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateEvent : get update by id

// UpdateEvent godoc
// @Summary get event by id
// @Schemes
// @Description get event by id
// @Tags events
// @Accept json
// @Param id path int true "search event by id"
// @Param event body models.Event true "Events model"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /events/{id} [put]
func (p EventController) UpdateEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var event models.Event
	event.ID = uint(id)

	eventRecord, err := p.service.Find(event)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Event with given id not found")
		return
	}
	ctx.ShouldBindJSON(&eventRecord)

	if eventRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if eventRecord.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(eventRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Event")
		return
	}
	response := eventRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Event",
		Data:    response,
	})
}
