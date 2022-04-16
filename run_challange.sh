#!/bin/bash

val=$(cat .activechallange);

if [ "$1" = "go" ]; then
    go run "${val}/main.go";
elif [ "$1" = "rust" ]; then
    cd $val; cargo run; cd -;
elif [ "$1" = "python" ] || [ $1 = "py" ]; then
    python3 "${val}/main.py";
elif [ "$1" = "all" ]; then
    go run "${val}/main.go";
    python3 "${val}/main.py";
    cd $val; cargo run; cd -;
else
    echo "parameters are [ all , go, python/py, rust ]"
fi