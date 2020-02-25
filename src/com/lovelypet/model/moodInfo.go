package model

import (
	"github.com/jinzhu/gorm"
	"lovelypet/src/com/lovelypet/cache"
)

type MoodInfo struct {
	UserId uint `json:"userId" binding:"required"`
	MoodText string `json:"moodText" binding:"required"`
	gorm.Model
}

func NewMoodInfo(userId uint,moodText string) *MoodInfo {
	return &MoodInfo{UserId:userId,MoodText:moodText}
}

func (mood *MoodInfo) Insert()bool {
	return cache.Insert(mood)
}