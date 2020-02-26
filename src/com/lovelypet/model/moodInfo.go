package model

import (
	"fmt"
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

func Delete(id string) bool  {
	return cache.Delete(&MoodInfo{},"id = ?",id)
}

func (mood *MoodInfo)IsMoodExist() bool {
	return cache.IsExist(mood,"id = ?",mood.ID)
}

func Query(userId uint)[]MoodInfo  {
	var moods [] MoodInfo
	err := cache.Query(&moods,"user_id = ?",userId)
	if err != nil {
		fmt.Println("Query()called  with userId[",userId,"]error:\n",err)
		return nil
	}
	return moods
}
