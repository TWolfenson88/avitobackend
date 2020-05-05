package call

import (
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
)

type Repository interface {
	SaveCallEndingInfo(form forms.CallEndForm) error
	SaveCallStartingInfo(form forms.CallStartForm) (int, error)
	GetCallHistoryByID(userID int) ([]models.Call, error)
}
