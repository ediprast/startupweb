package handler

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
	userService     user.Service
}

func NewCampaignHandler(campaignService campaign.Service, userService user.Service) *campaignHandler {
	return &campaignHandler{campaignService, userService}
}

func (h *campaignHandler) Index(c *gin.Context) {
	campaigns, err := h.campaignService.GetCampaigns(0)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}

func (h *campaignHandler) Create(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := campaign.FormCreateCampaignInput{}
	input.Users = users
	c.HTML(http.StatusOK, "campaign_create.html", input)
}

func (h *campaignHandler) Save(c *gin.Context) {
	var input campaign.FormCreateCampaignInput

	err := c.ShouldBind(&input)

	if err != nil {
		users, e := h.userService.GetAllUsers()

		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		input.Users = users
		input.Error = err
		c.HTML(http.StatusOK, "campaign_create.html", input)
		return
	}

	user, err := h.userService.GetUserById(input.UserId)

	if err != nil {
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
	}

	campaignInput := campaign.CreateCampaignInput{}
	campaignInput.Title = input.Title
	campaignInput.ShortDescription = input.ShortDescription
	campaignInput.Description = input.Description
	campaignInput.GoalAmount = input.GoalAmount
	campaignInput.Perks = input.Perks
	campaignInput.User = user

	_, err = h.campaignService.CreateCampaign(campaignInput)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) NewImage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	c.HTML(http.StatusOK, "campaign_image.html", gin.H{"Id": id})
}

func (h *campaignHandler) CreateImage(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("file")

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	currentCampaign, err := h.campaignService.GetCampaign(campaign.GetCampaignDetailInput{Id: id})

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userId := currentCampaign.UserId

	path := fmt.Sprintf("images/campaign_images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	campaignImageInput := campaign.CreateCampaignImageInput{}

	campaignImageInput.CampaigId = id
	campaignImageInput.IsPrimary = true

	userCampaign, _ := h.userService.GetUserById(userId)

	campaignImageInput.User = userCampaign

	_, err = h.campaignService.SaveCampaignImage(campaignImageInput, path)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	editCampaign, err := h.campaignService.GetCampaign(campaign.GetCampaignDetailInput{Id: id})

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	input := campaign.FormUpdateCampaignInput{}
	input.Id = editCampaign.Id
	input.Title = editCampaign.Title
	input.ShortDescription = editCampaign.ShortDesc
	input.Description = editCampaign.Description
	input.GoalAmount = editCampaign.GoalAmount
	input.Perks = editCampaign.Perks

	c.HTML(http.StatusOK, "campaign_edit.html", input)
}

func (h *campaignHandler) Update(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input campaign.FormUpdateCampaignInput
	err := c.ShouldBind(&input)

	if err != nil {
		input.Error = err
		input.Id = id
		c.Redirect(http.StatusFound, "/campaigns/edit/"+idParam)
		// c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	currentCampaign, err := h.campaignService.GetCampaign(campaign.GetCampaignDetailInput{Id: id})

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userId := currentCampaign.UserId

	userCampaign, _ := h.userService.GetUserById(userId)

	updateInput := campaign.CreateCampaignInput{}

	updateInput.Title = input.Title
	updateInput.ShortDescription = input.ShortDescription
	updateInput.Description = input.Description
	updateInput.GoalAmount = input.GoalAmount
	updateInput.Perks = input.Perks
	updateInput.User = userCampaign

	_, err = h.campaignService.UpdateCampaign(campaign.GetCampaignDetailInput{Id: id}, updateInput)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	getCampaign, err := h.campaignService.GetCampaign(campaign.GetCampaignDetailInput{Id: id})

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_show.html", getCampaign)
}
