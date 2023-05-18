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

//CommentController -> CommentController
type CommentController struct {
	service service.CommentService
}

//NewCommentController : NewCommentController
func NewCommentController(s service.CommentService) CommentController {
	return CommentController{
		service: s,
	}
}

// GetComments : GetComments controller

// GetComments godoc
// @Summary get all comments
// @Schemes
// @Description get all comments
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Router /comments [get]
func (p CommentController) GetComments(ctx *gin.Context) {
	var comments models.Comment

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(comments, keyword)

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
		Message: "Comment result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddComment : AddComment controller

// AddComment godoc
// @Summary create a new comment
// @Schemes
// @Description get all comments
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param comment body models.Comment true "Comments model"
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /comments [post]
func (p *CommentController) AddComment(ctx *gin.Context) {
	var comment models.Comment
	fmt.Println(comment)
	ctx.ShouldBindJSON(&comment)

	if comment.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if comment.Comment == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	err := p.service.Save(comment)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create comment")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Comment")
}

//GetComment : get comment by id

// GetComments godoc
// @Summary get comment by id
// @Schemes
// @Description get comment by id
// @Tags comments
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search comment by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /comments/{id} [get]
func (p *CommentController) GetComment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var comment models.Comment
	comment.ID = uint(id)
	foundComment, err := p.service.Find(comment)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Comment")
		return
	}

	response := foundComment.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Comment",
		Data:    &response})

}

//DeleteComment : Deletes Comment

// DeleteComment godoc
// @Summary delete comment by id
// @Schemes
// @Description delete comment by id
// @Tags comments
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "delete comment by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /comments/{id} [delete]
func (p *CommentController) DeleteComment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(uint(id))

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Comment")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateComment : get update by id

// UpdateComment godoc
// @Summary get comment by id
// @Schemes
// @Description get comment by id
// @Tags comments
// @Accept json
// @Param id path int true "search comment by id"
// @Param comment body models.Comment true "Comments model"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /comments/{id} [put]
func (p CommentController) UpdateComment(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var comment models.Comment
	comment.ID = uint(id)

	commentRecord, err := p.service.Find(comment)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Comment with given id not found")
		return
	}
	ctx.ShouldBindJSON(&commentRecord)

	if commentRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if commentRecord.Comment == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(commentRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Comment")
		return
	}
	response := commentRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Comment",
		Data:    response,
	})
}
