gen:
	PATH="/opt/homebrew/opt/llvm/bin:${PATH}" CGO_LDFLAGS="-L`llvm-config --libdir`" go run ./internal/generator
