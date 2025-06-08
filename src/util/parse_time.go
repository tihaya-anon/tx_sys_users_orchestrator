package util

import (
	"time"
	config_model "users_orchestrator/config/model"
)

func GetTime(configTime config_model.Time) time.Duration {
	return time.Hour*time.Duration(configTime.Hour) + time.Minute*time.Duration(configTime.Minute) + time.Second*time.Duration(configTime.Second)
}
