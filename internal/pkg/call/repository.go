package call

import "avitocalls/internal/pkg/forms"

type Repository interface {
	SaveCallEndingInfo(form forms.CallEndForm) error
	SaveCallStartingInfo(form forms.CallStartForm) (int, error)
}
