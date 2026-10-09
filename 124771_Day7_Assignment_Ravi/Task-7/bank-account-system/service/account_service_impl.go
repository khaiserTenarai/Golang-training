package service

import (
	"errors"
	"strings"

	"bank-account-system/model"
	"bank-account-system/repository"
)

type AccountServiceImpl struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(
	accountRepository repository.AccountRepository,
) *AccountServiceImpl {
	return &AccountServiceImpl{
		accountRepository: accountRepository,
	}
}

func (s *AccountServiceImpl) CreateAccount(
	account model.Account,
) error {
	if strings.TrimSpace(account.AccountNumber) == "" {
		return errors.New("account number cannot be empty")
	}

	if strings.TrimSpace(account.CustomerName) == "" {
		return errors.New("customer name cannot be empty")
	}

	if account.Balance < 0 {
		return errors.New("initial balance cannot be negative")
	}

	return s.accountRepository.CreateAccount(account)
}

func (s *AccountServiceImpl) Deposit(
	accountNumber string,
	amount float64,
) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	account, err :=
		s.accountRepository.GetAccountByNumber(accountNumber)

	if err != nil {
		return err
	}

	newBalance := account.Balance + amount

	if err := s.accountRepository.UpdateBalance(
		account.ID,
		newBalance,
	); err != nil {
		return err
	}

	transaction := model.Transaction{
		AccountID:       account.ID,
		TransactionType: "DEPOSIT",
		Amount:          amount,
	}

	return s.accountRepository.AddTransaction(transaction)
}

func (s *AccountServiceImpl) Withdraw(
	accountNumber string,
	amount float64,
) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	account, err :=
		s.accountRepository.GetAccountByNumber(accountNumber)

	if err != nil {
		return err
	}

	if amount > account.Balance {
		return errors.New("insufficient balance")
	}

	newBalance := account.Balance - amount

	if err := s.accountRepository.UpdateBalance(
		account.ID,
		newBalance,
	); err != nil {
		return err
	}

	transaction := model.Transaction{
		AccountID:       account.ID,
		TransactionType: "WITHDRAW",
		Amount:          amount,
	}

	return s.accountRepository.AddTransaction(transaction)
}

func (s *AccountServiceImpl) GetBalance(
	accountNumber string,
) (float64, error) {
	account, err :=
		s.accountRepository.GetAccountByNumber(accountNumber)

	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (s *AccountServiceImpl) GetTransactionHistory(
	accountNumber string,
) ([]model.Transaction, error) {
	account, err :=
		s.accountRepository.GetAccountByNumber(accountNumber)

	if err != nil {
		return nil, err
	}

	return s.accountRepository.GetTransactions(account.ID)
}
