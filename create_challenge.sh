#!/bin/bash
projdir=$1

mkdir $projdir
touch "${projdir}/readme.md"
touch "${projdir}/__init__.py"
touch "${projdir}/main.py"
touch "${projdir}/main.go"
touch "${projdir}/solution.md"

