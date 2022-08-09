package campaign

import "strings"

type CampaignFormatter struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	ImageUrl         string `json:"image_url"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Id = campaign.Id
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Title = campaign.Title
	campaignFormatter.ShortDescription = campaign.ShortDesc
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].Image
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormater := []CampaignFormatter{}

	for _, campaign := range campaigns {
		CampaignFormatter := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, CampaignFormatter)
	}

	return campaignsFormater
}

type CampaignUserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CampaignImageFormatter struct {
	Image     string `json:"image"`
	IsPrimary bool   `json:"is_primary"`
}

type CampaignDetailFormatter struct {
	Id            int                      `json:"id"`
	Title         string                   `json:"title"`
	ShortDesc     string                   `json:"short_desc"`
	Description   string                   `json:"description"`
	GoalAmount    int                      `json:"goal_amount"`
	CurrentAmount int                      `json:"current_amount"`
	BackerCount   int                      `json:"backer_count"`
	ImageUrl      string                   `json:"image_url"`
	Slug          string                   `json:"slug"`
	UserId        int                      `json:"user_id"`
	Perks         []string                 `json:"perks"`
	User          CampaignUserFormatter    `json:"user"`
	Images        []CampaignImageFormatter `json:"images"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignFormatter := CampaignDetailFormatter{}
	campaignFormatter.Id = campaign.Id
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Title = campaign.Title
	campaignFormatter.ShortDesc = campaign.ShortDesc
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.BackerCount = campaign.BackerCount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].Image
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormatter.Perks = perks
	campaignUserFormatter := CampaignUserFormatter{}
	user := campaign.User
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.Avatar = user.Avatar

	campaignFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.Image = image.Image
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}

	campaignFormatter.Images = images

	return campaignFormatter
}
