package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseStatusType(statusStr string) (models.StatusType, error) {
	switch statusStr {
	case string(models.Live):
		return models.Live, nil

	case string(models.Draft):
		return models.Draft, nil

	case string(models.Archived):
		return models.Archived, nil

	case string(models.Reported):
		return models.Reported, nil

	default:
		return "", fmt.Errorf("invalid status type: %s", statusStr)
	}
}
