// Task 14: Configuration-Based Feature Flags
// Features are toggled via environment variables or config.

package features

import "configurable_api/config"

type FeatureFlags struct {
	Pagination bool
	Filtering  bool
	Sorting    bool
	Auth       bool
}

func LoadFlags(cfg *config.Config) *FeatureFlags {
	return &FeatureFlags{
		Pagination: cfg.EnablePagination,
		Filtering:  cfg.EnableFiltering,
		Sorting:    cfg.EnableSorting,
		Auth:       cfg.EnableAuth,
	}
}

func (f *FeatureFlags) IsEnabled(name string) bool {
	switch name {
	case "pagination":
		return f.Pagination
	case "filtering":
		return f.Filtering
	case "sorting":
		return f.Sorting
	case "auth":
		return f.Auth
	default:
		return false
	}
}
