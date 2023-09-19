package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseAccountType(accountStr string) (models.AccountType, error) {
	switch accountStr {
	case string(models.Public):
		return models.Public, nil

	case string(models.Private):
		return models.Private, nil

	case string(models.Hidden):
		return models.Hidden, nil

	default:
		return "", fmt.Errorf("invalid account type: %s", accountStr)
	}
}
