package call

import "avitocalls/internal/pkg/forms"

type UseCase interface {
	SaveCallStarting(call forms.CallStartForm) (int, error)
	SaveCallEnding(call forms.CallEndForm) (int, error)
}
