package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"employee-api/model"
	"employee-api/repository"
)

// Options is the part of the configuration the service needs. The service does not know about
// environment variables - main.go copies the values from config.Config into this struct.
type Options struct {
	RequestTimeout  time.Duration
	DefaultPageSize int
	MaxPageSize     int
	SearchEnabled   bool // feature flag
}

type EmployeeService struct {
	repo   repository.EmployeeRepository
	opts   Options
	logger *slog.Logger
}

func NewEmployeeService(repo repository.EmployeeRepository, opts Options, logger *slog.Logger) *EmployeeService {
	return &EmployeeService{repo: repo, opts: opts, logger: logger}
}

type ListQuery struct {
	Search   string
	Page     int // 0 = first page
	PageSize int // 0 = default page size
}

type ListResult struct {
	Items      []model.Employee
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}

// withTimeout derives a context with the per-request deadline from the caller's context.
// If the HTTP client disconnects or the server force-closes, the parent is cancelled and
// the repository call stops too.
func (s *EmployeeService) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, s.opts.RequestTimeout)
}

func (s *EmployeeService) Create(ctx context.Context, in model.EmployeeInput) (model.Employee, error) {
	if err := in.Validate(); err != nil {
		return model.Employee{}, err
	}
	ctx, cancel := s.withTimeout(ctx)
	defer cancel()

	e, err := s.repo.Create(ctx, model.Employee{Name: in.Name, Email: in.Email, Department: in.Department, Salary: in.Salary})
	if err != nil {
		return model.Employee{}, err
	}
	s.logger.InfoContext(ctx, "employee created", "id", e.ID)
	return e, nil
}

func (s *EmployeeService) Get(ctx context.Context, id int) (model.Employee, error) {
	ctx, cancel := s.withTimeout(ctx)
	defer cancel()
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) List(ctx context.Context, q ListQuery) (ListResult, error) {
	// FEATURE FLAG: search can be switched off with FEATURE_SEARCH=false
	if q.Search != "" && !s.opts.SearchEnabled {
		return ListResult{}, fmt.Errorf("search: %w", model.ErrFeatureDisabled)
	}

	if q.Page == 0 {
		q.Page = 1
	}
	if q.PageSize == 0 {
		q.PageSize = s.opts.DefaultPageSize
	}
	fields := map[string]string{}
	if q.Page < 1 {
		fields["page"] = "must be 1 or greater"
	}
	if q.PageSize < 1 || q.PageSize > s.opts.MaxPageSize {
		fields["pageSize"] = fmt.Sprintf("must be between 1 and %d", s.opts.MaxPageSize)
	}
	if len(fields) > 0 {
		return ListResult{}, &model.ValidationError{Fields: fields}
	}

	ctx, cancel := s.withTimeout(ctx)
	defer cancel()
	items, total, err := s.repo.List(ctx, model.ListFilter{
		Search: q.Search,
		Limit:  q.PageSize,
		Offset: (q.Page - 1) * q.PageSize,
	})
	if err != nil {
		return ListResult{}, err
	}
	pages := (total + q.PageSize - 1) / q.PageSize
	return ListResult{Items: items, Total: total, Page: q.Page, PageSize: q.PageSize, TotalPages: pages}, nil
}

func (s *EmployeeService) Update(ctx context.Context, id int, in model.EmployeeInput) (model.Employee, error) {
	if err := in.Validate(); err != nil {
		return model.Employee{}, err
	}
	ctx, cancel := s.withTimeout(ctx)
	defer cancel()

	e, err := s.repo.Update(ctx, model.Employee{ID: id, Name: in.Name, Email: in.Email, Department: in.Department, Salary: in.Salary})
	if err != nil {
		return model.Employee{}, err
	}
	s.logger.InfoContext(ctx, "employee updated", "id", e.ID)
	return e, nil
}

func (s *EmployeeService) Delete(ctx context.Context, id int) error {
	ctx, cancel := s.withTimeout(ctx)
	defer cancel()
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	s.logger.InfoContext(ctx, "employee deleted", "id", id)
	return nil
}

// Ping is used by the readiness endpoint.
func (s *EmployeeService) Ping(ctx context.Context) error {
	ctx, cancel := s.withTimeout(ctx)
	defer cancel()
	return s.repo.Ping(ctx)
}
