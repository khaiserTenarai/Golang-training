package service
import (
	"money_transfer/model"
	"money_transfer/repository"
	"money_transfer/utility"
)

type TransferServiceImpl struct {
	repository repository.TransferRepository
}

func NewTransferService(
	repository repository.TransferRepository,
) TransferService {

	return &TransferServiceImpl{
		repository: repository,
	}
}

func (s *TransferServiceImpl) TransferMoney(
	transfer model.Transfer,
) error {

	err := utility.ValidateTransfer(transfer)

	if err != nil {
		return err
	}

	return s.repository.TransferMoney(transfer)
}
