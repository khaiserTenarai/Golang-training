# 5. Run gofmt

`before_unformatted.go` has inconsistent indentation (mixing tabs/spaces),
missing spaces around operators, and messy struct field alignment - all
things that still compile fine, but are painful to read and would look
different depending on whose editor you open it in.

## Command

```
gofmt -l .
```

`-l` just lists which files need formatting, without changing them. Running
it against `before_unformatted.go` reports the file as needing formatting.

To actually rewrite the file in place:

```
gofmt -w before_unformatted.go
```

`after_formatted.go` is what that file looks like once `gofmt -w` has been
run on it - consistent tabs for indentation, single spaces around `:=` and
`>`, and the struct fields lined up in a column. This is exactly what
`gofmt -w` would produce automatically; I don't have a Go toolchain in this
sandbox to actually run it, so `after_formatted.go` shows the expected
result by hand.

In everyday use you'd almost never call `gofmt` directly - most editors
(VS Code, GoLand) run it automatically every time you save a `.go` file.
