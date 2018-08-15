package db

import (
	"go_apps/go_api_apps/mf_soundboard/utils"
	"net/http"
)

func dbWithError(err error, code int, text string) *utils.Result {
	result := utils.Result{}

	result.Error = &utils.Error{
		StatusCode: code,
		StatusText: http.StatusText(code) + " - " + text + " : " + err.Error(),
	}

	return &result
}

func dbSuccess(code int, data interface{}) *utils.Result {
	result := utils.Result{}

	result.Success = &utils.Success{
		StatusCode: code,
		Data:       data,
	}

	return &result
}
