package call

import "avitocalls/internal/pkg/models"

type UseCase interface {
	GetReceiverObject(call models.Call) (models.Call, int, error)
}
