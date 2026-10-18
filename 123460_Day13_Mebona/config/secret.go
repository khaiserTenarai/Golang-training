package config

import "log/slog"

// Secret is a string that refuses to print itself.
// If a Secret ends up in a log line, fmt.Println or a JSON response by accident,
// only "[REDACTED]" is shown. The real value is available only through Reveal().
type Secret string

func (s Secret) String() string {
	if s == "" {
		return "(not set)"
	}
	return "[REDACTED]"
}

func (s Secret) GoString() string                { return s.String() }
func (s Secret) MarshalJSON() ([]byte, error)    { return []byte(`"` + s.String() + `"`), nil }
func (s Secret) LogValue() slog.Value            { return slog.StringValue(s.String()) }
func (s Secret) Reveal() string                  { return string(s) }
