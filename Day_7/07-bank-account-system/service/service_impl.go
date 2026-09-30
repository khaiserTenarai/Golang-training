package service
import (
	"bank_account/model"
	"bank_account/repository"
	"bank_account/utility"
)

type AccountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(
	repository repository.AccountRepository,
) AccountService {

	return &AccountServiceImpl{
		repository: repository,
	}
}

func (s *AccountServiceImpl) Create(
	account model.Account,
) error {

	err := utility.ValidateAccount(account)

	if err != nil {
		return err
	}

	return s.repository.Create(account)
}

func (s *AccountServiceImpl) FindByID(
	id int,
) (model.Account, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Account{}, err
	}

	return s.repository.FindByID(id)
}

func (s *AccountServiceImpl) Deposit(
	id int,
	amount float64,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	err = utility.ValidateAmount(amount)

	if err != nil {
		return err
	}

	return s.repository.Deposit(id, amount)
}

func (s *AccountServiceImpl) Withdraw(
	id int,
	amount float64,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	err = utility.ValidateAmount(amount)

	if err != nil {
		return err
	}

	return s.repository.Withdraw(id, amount)
}

func (s *AccountServiceImpl) GetTransactions(
	id int,
) ([]model.Transaction, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return nil, err
	}

	return s.repository.GetTransactions(id)
}
