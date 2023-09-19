package parsers

import (
	"fmt"

	"github.com/JOYSTIQ-The-App/Backend/api/models"
)

func ParseMediaType(mediaStr string) (models.MediaType, error) {
	switch mediaStr {
	case string(models.None):
		return models.None, nil

	case string(models.Photo):
		return models.Photo, nil

	case string(models.Video):
		return models.Video, nil

	default:
		return "", fmt.Errorf("invalid media type: %s", mediaStr)
	}
}
