package middleware

import (
	"encoding/json"
	"net/http"
)

func jsonEncode(w http.ResponseWriter, value any) error {
	return json.NewEncoder(w).Encode(value)
}
