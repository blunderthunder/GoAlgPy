#!/bin/bash

val=$(cat .activechallange);

if [ $1 = "go" ]
then
    go run "${val}/main.${1}";
else
    python "${val}/main.${1}";
fi