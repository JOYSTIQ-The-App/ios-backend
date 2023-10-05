package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseSocialType(socialStr string) (models.SocialType, error) {
	switch socialStr {
	case string(models.Discord):
		return models.Discord, nil

	case string(models.Kick):
		return models.Kick, nil

	case string(models.Twitch):
		return models.Twitch, nil

	case string(models.Youtube):
		return models.Youtube, nil

	case string(models.Xbox):
		return models.Xbox, nil

	case string(models.Playstation):
		return models.Playstation, nil

	default:
		return "", fmt.Errorf("invalid social type: %s", socialStr)
	}
}
