package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// readDotEnv parses a simple KEY=VALUE file. A missing file is not an error.
// Supported: comment lines (#), blank lines, "export KEY=VALUE", and "quoted" values.
// NOT supported: inline comments after a value (write comments on their own line).
// Error messages contain the file and line number but never the value (it might be a secret).
func readDotEnv(path string) (map[string]string, error) {
	out := map[string]string{}

	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return out, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	line := 0
	for sc.Scan() {
		line++
		text := strings.TrimSpace(strings.TrimPrefix(sc.Text(), "\ufeff")) // tolerate a Windows BOM
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}
		text = strings.TrimPrefix(text, "export ")
		key, val, ok := strings.Cut(text, "=")
		key = strings.TrimSpace(key)
		if !ok || key == "" {
			return nil, fmt.Errorf("%s:%d: expected KEY=VALUE", path, line)
		}
		val = strings.TrimSpace(val)
		if len(val) >= 2 && (val[0] == '"' || val[0] == '\'') && val[len(val)-1] == val[0] {
			val = val[1 : len(val)-1]
		}
		out[key] = val
	}
	return out, sc.Err()
}
