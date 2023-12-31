package dto

import (
	"strconv"
	"time"

	"github.com/chuxin0816/bluebell/models"
)

type CommunityDto struct {
	CommunityID   string    `json:"id"`
	CommunityName string    `json:"name"`
	Introduction  string    `json:"introduction"`
	CreatedAt     time.Time `json:"created_time"`
}

func ToCommunityDto(community *models.Community) *CommunityDto {
	return &CommunityDto{
		CommunityID:   strconv.Itoa(community.CommunityID),
		CommunityName: community.CommunityName,
		Introduction:  community.Introduction,
		CreatedAt:     community.CreatedAt,
	}
}

func ToCommunityDtoList(CommunityList []*models.Community) (communityDtoList []*CommunityDto) {
	communityDtoList = make([]*CommunityDto, 0, len(CommunityList))
	for _, community := range CommunityList {
		communityDtoList = append(communityDtoList, ToCommunityDto(community))
	}
	return communityDtoList
}
