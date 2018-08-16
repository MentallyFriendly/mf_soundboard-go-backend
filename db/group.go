package db

import (
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"
)

// GroupPayload type
type GroupPayload struct {
	Name     *string `json:"name"`
	SoundIDs *[]int  `json:"sound_ids"`
}

// GetGroup func
func GetGroup(id string) *utils.Result {
	group := Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Where("id = ?", id).Find(&group).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
	}

	return dbSuccess(200, &group)
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
func CreateGroup(data *GroupPayload) *utils.Result {
	var sounds []Sound

	if len(*data.SoundIDs) > 0 {
		if err := db.Model(&Sound{}).Where("id in (?)", *data.SoundIDs).Find(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching sounds from DB")
		}
	}

	group := Group{
		Name:   data.Name,
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
func UpdateGroup(id string, data *GroupPayload) *utils.Result {
	group := Group{}
	if err := db.Model(&Group{}).Preload("Sounds").Where("id = ?", id).Find(&group).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching group from DB")
	}

	if err := db.Model(&group).Updates(&Group{
		Name: data.Name,
	}).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error updating group")
	}

	var sounds []Sound
	if len(*data.SoundIDs) > 0 {
		if err := db.Model(Sound{}).Where("id in (?)", *data.SoundIDs).Find(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching sounds from DB")
		}

		if err := db.Model(&group).Association("Sounds").Replace(&sounds).Error; err != nil {
			return dbWithError(err, http.StatusInternalServerError, "Error updating sounds")
		}
	}

	return dbSuccess(200, &group)
}
