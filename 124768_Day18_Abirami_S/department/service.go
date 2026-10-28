package department

import "context"

type Service struct { repo *Repository }
func NewService(repo *Repository) *Service { return &Service{repo:repo} }
func (s *Service) Get(ctx context.Context, id int64) (Department, error) { return s.repo.Get(ctx, id) }
func (s *Service) List(ctx context.Context) ([]Department, error) { return s.repo.List(ctx) }
