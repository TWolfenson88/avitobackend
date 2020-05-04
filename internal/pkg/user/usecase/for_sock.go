package usecase

import (
	"fmt"
)

func (uc *userUseCase) SetOnline(username string) error {
	err := uc.rep.UpdateStatus(username, true)
	if err != nil {
		fmt.Println("error with setting online: db")
		return err
	}
	return nil
}

func (uc *userUseCase) SetOffline(username string) error {
	err := uc.rep.UpdateStatus(username, false)
	if err != nil {
		fmt.Println("error with setting offline: db")
		return err
	}
	return nil
}
