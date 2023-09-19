package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseRoleType(roleStr string) (models.RoleType, error) {
	switch roleStr {
	case string(models.Admin):
		return models.Admin, nil

	case string(models.Org):
		return models.Org, nil

	case string(models.Gamer):
		return models.Gamer, nil

	default:
		return "", fmt.Errorf("invalid role type: %s", roleStr)
	}
}
