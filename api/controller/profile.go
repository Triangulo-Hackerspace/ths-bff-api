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

//ProfileController -> ProfileController
type ProfileController struct {
	service service.ProfileService
}

//NewProfileController : NewProfileController
func NewProfileController(s service.ProfileService) ProfileController {
	return ProfileController{
		service: s,
	}
}

// GetProfiles : GetProfiles controller

// GetProfiles godoc
// @Summary get all Profiles
// @Schemes
// @Description get all Profiles
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags Profiles
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Router /Profiles [get]
func (p ProfileController) GetProfiles(ctx *gin.Context) {
	var Profiles models.Profile

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(Profiles, keyword)

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
		Message: "Profile result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddProfile : AddProfile controller

// AddProfile godoc
// @Summary create a new Profile
// @Schemes
// @Description get all Profiles
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Profile body models.Profile true "Profiles model"
// @Tags Profiles
// @Accept json
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /Profiles [post]
func (p *ProfileController) AddProfile(ctx *gin.Context) {
	var Profile models.Profile
	fmt.Println(Profile)
	ctx.ShouldBindJSON(&Profile)

	if Profile.Bio == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Bio is required")
		return
	}
	if Profile.Image == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Image is required")
		return
	}

	err := p.service.Save(Profile)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create Profile")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Profile")
}

//GetProfile : get Profile by id

// GetProfiles godoc
// @Summary get Profile by id
// @Schemes
// @Description get Profile by id
// @Tags Profiles
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search Profile by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /Profiles/{id} [get]
func (p *ProfileController) GetProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var Profile models.Profile
	Profile.ID = uint(id)
	foundProfile, err := p.service.Find(Profile)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Profile")
		return
	}

	response := foundProfile.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Profile",
		Data:    &response})

}

//GetProfile : get Profile by User id

// GetProfiles godoc
// @Summary get Profile by User id
// @Schemes
// @Description get Profile by User id
// @Tags Profiles
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "search Profile by User id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /Profiles/user/{id} [get]
func (p *ProfileController) GetProfileByUserId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}

	var userId = uint64(id)
	foundProfile, err := p.service.FindByUserId(userId)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Profile")
		return
	}

	response := foundProfile.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Profile",
		Data:    &response})

}

//DeleteProfile : Deletes Profile

// DeleteProfile godoc
// @Summary delete Profile by id
// @Schemes
// @Description delete Profile by id
// @Tags Profiles
// @Accept json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "delete Profile by id"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /Profiles/{id} [delete]
func (p *ProfileController) DeleteProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(uint(id))

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Profile")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateProfile : get update by id

// UpdateProfile godoc
// @Summary get Profile by id
// @Schemes
// @Description get Profile by id
// @Tags Profiles
// @Accept json
// @Param id path int true "search Profile by id"
// @Param Profile body models.Profile true "Profiles model"
// @Produce json
// @Success 200 {object} util.ResponseLogin
// @Failure 400 {object} httputil.HTTPError
// @Router /Profiles/{id} [put]
func (p ProfileController) UpdateProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var Profile models.Profile
	Profile.ID = uint(id)

	ProfileRecord, err := p.service.Find(Profile)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Profile with given id not found")
		return
	}
	ctx.ShouldBindJSON(&ProfileRecord)

	if ProfileRecord.Bio == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Bio is required")
		return
	}
	if ProfileRecord.Image == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Image is required")
		return
	}

	if err := p.service.Update(ProfileRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Profile")
		return
	}
	response := ProfileRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Profile",
		Data:    response,
	})
}
