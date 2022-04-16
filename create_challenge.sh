#!/bin/bash

create_proj_dir(){
    touch "${1}/readme.md"
    touch "${1}/__init__.py"
    touch "${1}/main.py"
    touch "${1}/main.go"
    touch "${1}/solution.md"
}

create_rust_proj_dir(){
    mkdir "${1}/tests"
    touch "${1}/tests/${1}_test.rs"
}

projdir=$1

echo $projdir > .activechallange

# create the rust project
cargo new --lib $projdir

if [ $? -eq 0 ]; then
    create_proj_dir $projdir
    create_rust_proj_dir $projdir
else
    create_proj_dir $projdir
fi

echo "Successfully created project ${projdir}"