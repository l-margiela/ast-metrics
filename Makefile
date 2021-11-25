run:
	@go run ./somepkg
run-mod: transform
	@go run ./somepkgmod
transform:
	@go run main.go