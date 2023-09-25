package dto

import (
	"strconv"
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type CommunityDto struct {
	CommunityID   string `json:"community_id"`
	CommunityName string `json:"community_name"`
	Introduction  string `json:"introduction"`
	CreatedAt     time.Time
}

func ToCommunityDto(community *models.Community) *CommunityDto {
	return &CommunityDto{
		CommunityID:   strconv.Itoa(community.CommunityID),
		CommunityName: community.CommunityName,
		Introduction:  community.Introduction,
		CreatedAt:     community.CreatedAt,
	}
}

func ToCommunityDtoList(CommunityList []models.Community) (communityDtoList []CommunityDto) {
	for _, community := range CommunityList {
		communityDtoList = append(communityDtoList, *ToCommunityDto(&community))
	}
	return communityDtoList
}
