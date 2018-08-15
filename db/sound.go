package db

import (
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"
	"strings"
)

// GetSounds func
func GetSounds() *utils.Result {
	sounds := []Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Find(&sounds).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching sounds from DB")
	}

	return dbSuccess(200, &sounds)
}

// GetSound func
func GetSound(id string) *utils.Result {
	sound := Sound{}

	err := db.Model(&Sound{}).Preload("Groups").Where("id = ?", id).Find(&sound).Error
	if err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching sound from DB")
	}

	return dbSuccess(200, &sound)
}

// CreateSound func
func CreateSound(data map[string]*string) *utils.Result {
	var groups []Group

	if data["group_ids"] != nil {
		jsonGroupIds := strings.Split(*data["group_ids"], ",")

		if err := db.Model(&Group{}).Where("id in (?)", jsonGroupIds).Find(&groups).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching groups from DB")
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
		return dbWithError(err, http.StatusInternalServerError, "Error saving new sound to DB")
	}

	return dbSuccess(200, &sound)
}

// BulkCreateSounds func
func BulkCreateSounds(data []map[string]*string) *utils.Result {
	result := &utils.Result{}

	for _, sound := range data {
		go func(sound map[string]*string) {
			groups := []Group{}

			if sound["groups"] != nil {
				jsonGroupNames := strings.Split(*sound["groups"], ",")
				if err := db.Model(&Group{}).Where("name in (?)", jsonGroupNames).Find(&groups).Error; err != nil {
					result = dbWithError(err, http.StatusNotFound, "Error fetching group by name")
				}
			}

			if err := db.Save(&Sound{
				Name:         sound["name"],
				Path:         sound["path"],
				Letter:       sound["letter"],
				EmojiUnicode: sound["emoji_unicode"],
				Groups:       groups,
			}).Error; err != nil {
				result = dbWithError(err, http.StatusInternalServerError, "Error saving sound to DB")
			}
		}(sound)
	}

	result = dbSuccess(200, "Successfully added sounds")
	return result
}

// DeleteSound func
func DeleteSound(id string) *utils.Result {
	sound := Sound{}
	if err := db.Model(&Sound{}).Where("id = ?", id).Find(&sound).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching sound from DB")
	}

	if err := db.Delete(&sound).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error deleting sound from DB")
	}

	return dbSuccess(200, "Successfully deleted sound from DB")
}

// UpdateSound func
func UpdateSound(id string, data map[string]*string) *utils.Result {
	sound := Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Where("id = ?", id).Find(&sound).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching sound from DB")
	}

	if err := db.Model(&sound).Updates(&Sound{
		Name:   data["name"],
		Path:   data["path"],
		Letter: data["letter"],
	}).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error updating sound")
	}

	var groups []Group
	if data["group_ids"] != nil {
		jsonGroupIDs := strings.Split(*data["group_ids"], ",")
		if err := db.Model(&Group{}).Where("id in (?)", jsonGroupIDs).Find(&groups).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching groups from DB")
		}

		if err := db.Model(&sound).Association("Groups").Replace(&groups).Error; err != nil {
			return dbWithError(err, http.StatusInternalServerError, "Error updating groups")
		}
	}

	return dbSuccess(200, &sound)
}
