.SILENT:

help:
	{ grep --extended-regexp '^[a-zA-Z_-]+:.*#[[:space:]].*$$' $(MAKEFILE_LIST) || true; } \
	| awk 'BEGIN { FS = ":.*#[[:space:]]*" } { printf "\033[1;32m%-22s\033[0m%s\n", $$1, $$2 }'

new: # create new directory for the new challange 
	sh create_challenge.sh $(path)

py: # run test of python script
	sh run_challange.sh "py"

go: # run test of golang script
	sh run_challange.sh "go"

rust: # run rust of golang script
	sh run_challange.sh "rust"

run: # run all the script
	sh run_challange.sh "all"