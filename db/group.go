package db

import (
	"fmt"
	"strings"
)

// GetGroup func
func GetGroup(name string) *Group {
	group := Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Where("name = ?", name).Find(&group).Error; err != nil {
		fmt.Println("Error fetching group ", err)
	}
	return &group
}

// GetGroups func
func GetGroups() *[]Group {
	groups := []Group{}
	if err := db.Model(&Group{}).Find(&groups).Error; err != nil {
		fmt.Println("Error fetching groups ", err)
	}
	return &groups
}

// CreateGroup func
func CreateGroup(data map[string]*string) *Group {
	var sounds []Sound

	if data["sound_ids"] != nil {
		jsonSoundIDs := strings.Split(*data["sound_ids"], ",")
		if err := db.Model(&Sound{}).Where("id in (?)", jsonSoundIDs).Find(&sounds).Error; err != nil {
			fmt.Println("Error fetching sounds from db ", err)
		}
	}

	group := Group{
		Name:   data["name"],
		Sounds: sounds,
	}

	if err := db.Save(&group).Error; err != nil {
		fmt.Println("Error saving group ", err)
	}

	return &group
}

// DeleteGroup func
func DeleteGroup(name string) string {
	group := Group{}
	if err := db.Model(&Group{}).Where("name = ?", name).Find(&group).Error; err != nil {
		fmt.Println("Error fetching group ", err)
	}
	if err := db.Delete(&group).Error; err != nil {
		fmt.Println("Error deleting group ", err)
	}

	return "Successfully deleted group from db"
}

// UpdateGroup func
func UpdateGroup(name string, data map[string]*string) *Group {
	group := Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Where("name = ?", name).Find(&group).Error; err != nil {
		fmt.Println("Error fetching group from db ", err)
	}

	var sounds []Sound
	if data["sound_ids"] != nil {
		jsonSoundIDs := strings.Split(*data["sound_ids"], ",")
		if err := db.Model(Sound{}).Where("id in (?)", jsonSoundIDs).Find(&sounds).Error; err != nil {
			fmt.Println("Error fetching sounds from db ", err)
		}
	}

	if err := db.Model(&group).Updates(&Group{
		Name:   data["name"],
		Sounds: sounds,
	}).Error; err != nil {
		fmt.Println("Error updating group ", err)
	}

	return &group
}
