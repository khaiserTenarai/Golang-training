
set -euo pipefail
APP=app; PKG=./cmd/app; BIN_DIR=bin
VERSION=$(git describe --tags --always 2>/dev/null || echo dev)

case "${1:-help}" in
  run)   go run "$PKG" ;;
  build) mkdir -p "$BIN_DIR"
         go build -ldflags "-s -w -X main.version=$VERSION" -o "$BIN_DIR/$APP" "$PKG"
         echo "Built $BIN_DIR/$APP ($VERSION)" ;;
  test)  go test -v -cover ./... ;;
  fmt)   go fmt ./... ;;
  vet)   go vet ./... ;;
  all)   "$0" fmt; "$0" vet; "$0" test; "$0" build ;;
  clean) rm -rf "$BIN_DIR" ;;
  *)     echo "Usage: $0 {run|build|test|fmt|vet|all|clean}"; exit 1 ;;
esac
