package db

import (
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"
)

// SoundPayload type
type SoundPayload struct {
	Name     *string `json:"name"`
	Path     *string `json:"path"`
	Letter   *string `json:"letter"`
	Emoji    *string `json:"emoji"`
	GroupIDs *[]int  `json:"group_ids"`
}

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
func CreateSound(data *SoundPayload) *utils.Result {
	var groups []Group

	if len(*data.GroupIDs) > 0 {
		if err := db.Model(&Group{}).Where("id in (?)", *data.GroupIDs).Find(&groups).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching groups from DB")
		}
	}

	sound := Sound{
		Name:   data.Name,
		Path:   data.Path,
		Letter: data.Letter,
		Emoji:  data.Emoji,
		Groups: groups,
	}

	if err := db.Save(&sound).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error saving new sound to DB")
	}

	return dbSuccess(200, &sound)
}

// BulkCreateSounds func
func BulkCreateSounds(data []*SoundPayload) *utils.Result {
	var result *utils.Result
	for _, sound := range data {
		go func(sound *SoundPayload) {
			groups := []Group{}
			if len(*sound.GroupIDs) > 0 {
				if err := db.Model(&Group{}).Where("id in (?)", *sound.GroupIDs).Find(&groups).Error; err != nil {
					result = dbWithError(err, http.StatusNotFound, "Error fetching group by name")
					return
				}
			}

			if err := db.Save(&Sound{
				Name:   sound.Name,
				Path:   sound.Path,
				Letter: sound.Letter,
				Emoji:  sound.Emoji,
				Groups: groups,
			}).Error; err != nil {
				result = dbWithError(err, http.StatusInternalServerError, "Error saving sound to DB")
				return
			}
		}(sound)
		if result != nil {
			return result
		}
	}

	return dbSuccess(200, "Successfully added sounds")
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
func UpdateSound(id string, data *SoundPayload) *utils.Result {
	sound := Sound{}
	if err := db.Model(&Sound{}).Preload("Groups").Where("id = ?", id).Find(&sound).Error; err != nil {
		return dbWithError(err, http.StatusNotFound, "Error fetching sound from DB")
	}

	if err := db.Model(&sound).Updates(&Sound{
		Name:   data.Name,
		Path:   data.Path,
		Letter: data.Letter,
		Emoji:  data.Emoji,
	}).Error; err != nil {
		return dbWithError(err, http.StatusInternalServerError, "Error updating sound")
	}

	var groups []Group
	if len(*data.GroupIDs) > 0 {
		if err := db.Model(&Group{}).Where("id in (?)", *data.GroupIDs).Find(&groups).Error; err != nil {
			return dbWithError(err, http.StatusNotFound, "Error fetching groups from DB")
		}

		if err := db.Model(&sound).Association("Groups").Replace(&groups).Error; err != nil {
			return dbWithError(err, http.StatusInternalServerError, "Error updating new groups")
		}
	}

	return dbSuccess(200, &sound)
}
