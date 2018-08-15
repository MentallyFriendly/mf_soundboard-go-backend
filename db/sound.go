package db

import (
	"fmt"
	"strings"
)

// GetSounds func
func GetSounds() *[]Sound {
	sounds := []Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Find(&sounds).Error; err != nil {
		fmt.Println("Error fetching sounds ", err)
	}
	return &sounds
}

// GetSound func
func GetSound(id string) *Sound {
	sound := Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Where("id = ?", id).Find(&sound).Error; err != nil {
		fmt.Println("Error fetching sound ", err)
	}

	return &sound
}

// CreateSound func
func CreateSound(data map[string]*string) *Sound {
	var groups []Group

	if data["group_ids"] != nil {
		jsonGroupIds := strings.Split(*data["group_ids"], ",")

		if err := db.Model(&Group{}).Where("id in (?)", jsonGroupIds).Find(&groups).Error; err != nil {
			fmt.Println("Error fetching groups")
		}
	}

	sound := Sound{
		Name:         data["name"],
		Path:         data["path"],
		Letter:       data["letter"],
		EmojiUnicode: data["emoji_unicode"],
		Groups:       groups,
	}

	if err := db.Save(&sound).Error; err != nil {
		fmt.Println("Error saving new sound ", err)
	}

	return &sound
}

// BulkCreateSounds func
func BulkCreateSounds(data []map[string]*string) string {
	for _, sound := range data {
		go func(sound map[string]*string) {
			groups := []Group{}

			if sound["groups"] != nil {
				jsonGroupNames := strings.Split(*sound["groups"], ",")
				if err := db.Model(&Group{}).Where("name in (?)", jsonGroupNames).Find(&groups).Error; err != nil {
					fmt.Println("Error fetching group by name ", err)
				}
			}

			if err := db.Save(&Sound{
				Name:         sound["name"],
				Path:         sound["path"],
				Letter:       sound["letter"],
				EmojiUnicode: sound["emoji_unicode"],
				Groups:       groups,
			}).Error; err != nil {
				fmt.Println("Error saving sound ", err)
			}
		}(sound)
	}

	return "Successfully saved new sounds"
}

// DeleteSound func
func DeleteSound(id string) string {
	sound := Sound{}
	if err := db.Model(&Sound{}).Where("id = ?", id).Find(&sound).Error; err != nil {
		fmt.Println("Error fetching sound for delete ", err)
	}

	if err := db.Delete(&sound).Error; err != nil {
		fmt.Println("Error deleting sound ", err)
	}

	return "Successfully deleted sound from DB"
}

// UpdateSound func
func UpdateSound(id string, data map[string]*string) *Sound {
	sound := Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Where("id = ?", id).Find(&sound).Error; err != nil {
		fmt.Println("Error fetching sound from db ", err)
	}

	var groups []Group
	if data["group_ids"] != nil {
		jsonGroupIDs := strings.Split(*data["group_ids"], ",")
		if err := db.Model(&Group{}).Where("id in (?)", jsonGroupIDs).Find(&groups).Error; err != nil {
			fmt.Println("Error fetching groups from db ", err)
		}
	}

	if err := db.Model(&sound).Updates(&Sound{
		Name:   data["name"],
		Path:   data["path"],
		Letter: data["letter"],
		Groups: groups,
	}).Error; err != nil {
		fmt.Println("Error updating sound ", err)
	}

	return &sound
}
