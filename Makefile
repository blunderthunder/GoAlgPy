.SILENT:

help:
	{ grep --extended-regexp '^[a-zA-Z_-]+:.*#[[:space:]].*$$' $(MAKEFILE_LIST) || true; } \
	| awk 'BEGIN { FS = ":.*#[[:space:]]*" } { printf "\033[1;32m%-22s\033[0m%s\n", $$1, $$2 }'

new: # create new directory for the new challange 
	go run main.go create_challenge

run-py: # run test of python script
	go run main.go run_challenge "py"

run-go: # run test of golang script
	go run main.go run_challenge "go"

run-rust: # run rust of golang script
	go run main.go run_challenge "rust"

run: # run all the script
	go run main.go run_challenge "all"

select: #  select challenge as activechallenge
	go run main.go select_challenge