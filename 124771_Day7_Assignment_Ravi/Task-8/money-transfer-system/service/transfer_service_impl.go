package service

import (
	"errors"
	"strings"

	"money-transfer-system/model"
	"money-transfer-system/repository"
)

type TransferServiceImpl struct {
	transferRepository repository.TransferRepository
}

func NewTransferService(
	transferRepository repository.TransferRepository,
) *TransferServiceImpl {
	return &TransferServiceImpl{
		transferRepository: transferRepository,
	}
}

func (s *TransferServiceImpl) TransferMoney(
	transfer model.Transfer,
) error {
	transfer.FromAccount = strings.TrimSpace(transfer.FromAccount)
	transfer.ToAccount = strings.TrimSpace(transfer.ToAccount)

	if transfer.FromAccount == "" {
		return errors.New("source account cannot be empty")
	}

	if transfer.ToAccount == "" {
		return errors.New("destination account cannot be empty")
	}

	if transfer.FromAccount == transfer.ToAccount {
		return errors.New("source and destination accounts must be different")
	}

	if transfer.Amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	return s.transferRepository.TransferMoney(transfer)
}

func (s *TransferServiceImpl) GetAccount(
	accountNumber string,
) (model.Account, error) {
	accountNumber = strings.TrimSpace(accountNumber)

	if accountNumber == "" {
		return model.Account{}, errors.New("account number cannot be empty")
	}

	return s.transferRepository.GetAccount(accountNumber)
}
