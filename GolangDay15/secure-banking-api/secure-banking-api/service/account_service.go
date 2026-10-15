package service

import (
	"context"
	"strings"

	"secure-banking-api/model"
	"secure-banking-api/repository"
)

type AccountService interface {
	Create(ctx context.Context, userID int64, req model.CreateAccountRequest) (*model.Account, error)
	GetByID(ctx context.Context, id, userID int64, role string) (*model.Account, error)
	ListMine(ctx context.Context, userID int64, q model.PageQuery) (model.PageResult[model.Account], error)
	ListAll(ctx context.Context, q model.PageQuery) (model.PageResult[model.Account], error)
	AttachDocument(ctx context.Context, id int64, path string) error
}

type accountService struct{ repo repository.AccountRepository }

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) Create(ctx context.Context, userID int64, req model.CreateAccountRequest) (*model.Account, error) {
	a := &model.Account{
		UserID:     userID,
		HolderName: strings.TrimSpace(req.HolderName),
		Balance:    req.InitialDeposit,
	}
	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}


func (s *accountService) GetByID(ctx context.Context, id, userID int64, role string) (*model.Account, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role != model.RoleAdmin && a.UserID != userID {
		return nil, model.ErrForbidden
	}
	return a, nil
}

func (s *accountService) ListMine(ctx context.Context, userID int64, q model.PageQuery) (model.PageResult[model.Account], error) {
	return s.list(ctx, userID, q)
}

func (s *accountService) ListAll(ctx context.Context, q model.PageQuery) (model.PageResult[model.Account], error) {
	return s.list(ctx, 0, q)
}

func (s *accountService) list(ctx context.Context, userID int64, q model.PageQuery) (model.PageResult[model.Account], error) {
	items, total, err := s.repo.List(ctx, userID, q.Limit, q.Offset())
	if err != nil {
		return model.PageResult[model.Account]{}, err
	}
	return model.NewPageResult(items, q, total), nil
}

func (s *accountService) AttachDocument(ctx context.Context, id int64, path string) error {
	return s.repo.SetDocument(ctx, id, path)
}
