package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}

	campaign.Title = input.Title
	campaign.ShortDesc = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.Id

	stringSlug := fmt.Sprintf("%s %d", input.Title, input.User.Id)
	campaign.Slug = slug.Make(stringSlug)

	savedCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return savedCampaign, err
	}

	return savedCampaign, nil

}

func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputId.Id)

	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Title = inputData.Title
	campaign.ShortDesc = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {

	campaign, err := s.repository.FindById(input.CampaigId)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserId != input.User.Id {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0

	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.ChangeImageAsNonPrimary(input.CampaigId)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}

	campaignImage.CampaignId = input.CampaigId
	campaignImage.IsPrimary = isPrimary
	campaignImage.Image = fileLocation

	saveImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {
		return saveImage, err
	}

	return saveImage, nil
}
