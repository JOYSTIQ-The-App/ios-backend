package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseEnvironmentType(environmentStr string) (models.EnvironmentType, error) {
	switch environmentStr {
	case string(models.Classic):
		return models.Classic, nil

	case string(models.GameRoom):
		return models.GameRoom, nil

	default:
		return "", fmt.Errorf("invalid account type: %s", environmentStr)
	}
}
