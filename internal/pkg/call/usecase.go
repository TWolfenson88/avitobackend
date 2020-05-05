package call

import (
	"avitocalls/internal/pkg/forms"
	"avitocalls/internal/pkg/models"
)

type UseCase interface {
	SaveCallStarting(call forms.CallStartForm) (int, error)
	SaveCallEnding(call forms.CallEndForm) (int, error)
	GetUserCallHistory(form forms.CallHistReqForm) ([]models.Call, error)
}
