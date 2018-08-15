package db

import (
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"
	"strings"
)

// GetGroup func
func GetGroup(id, name string) *utils.Result {
	result := &utils.Result{}
	group := Group{}
	if len(id) > 0 {
		if err := db.Model(&Group{}).Preload("Sounds").Where("id = ?", id).Find(&group).Error; err != nil {
			result = dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
		}

		result = dbSuccess(200, &group)

	} else {
		if err := db.Model(&Group{}).Preload("Sounds").Where("name = ?", name).Find(&group).Error; err != nil {
			result = dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
		}

		result = dbSuccess(200, &group)
	}

	return result
}

// GetGroups func
func GetGroups() *utils.Result {
	groups := []Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Find(&groups).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching groups from DB")
	}
	return dbSuccess(200, &groups)
}

// CreateGroup func
func CreateGroup(data map[string]*string) *utils.Result {
	var sounds []Sound

	if data["sound_ids"] != nil {
		jsonSoundIDs := strings.Split(*data["sound_ids"], ",")
		if err := db.Model(&Sound{}).Where("id in (?)", jsonSoundIDs).Find(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching sounds from DB")
		}
	}

	group := Group{
		Name:   data["name"],
		Sounds: sounds,
	}

	if err := db.Save(&group).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error saving group")
	}

	return dbSuccess(200, &group)
}

// DeleteGroup func
func DeleteGroup(id string) *utils.Result {
	group := Group{}
	if err := db.Model(&Group{}).Where("id = ?", id).Find(&group).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
	}
	if err := db.Delete(&group).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error fetching group from DB")
	}

	return dbSuccess(200, "Successfully deleted group from db")
}

// UpdateGroup func
func UpdateGroup(id string, data map[string]*string) *utils.Result {
	group := Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Where("id = ?", id).Find(&group).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
	}

	if err := db.Model(&group).Updates(&Group{
		Name: data["name"],
	}).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error updating group")
	}

	var sounds []Sound
	if data["sound_ids"] != nil {
		jsonSoundIDs := strings.Split(*data["sound_ids"], ",")
		if err := db.Model(Sound{}).Where("id in (?)", jsonSoundIDs).Find(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching sounds from DB")
		}

		if err := db.Model(&group).Association("Sounds").Replace(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusInternalServerError, "Error updating sounds")
		}
	}

	return dbSuccess(200, &group)
}
