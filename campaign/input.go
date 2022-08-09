package campaign

import "bwastartup/user"

type GetCampaignDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaigId int  `form:"campaign_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}

type FormCreateCampaignInput struct {
	Title            string `form:"title" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	UserId           int    `form:"user_id" binding:"required"`
	Users            []user.User
	Error            error
}

type FormUpdateCampaignInput struct {
	Id               int
	Title            string `form:"title" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	User             user.User
	Error            error
}
