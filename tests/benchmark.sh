#!/bin/bash

# This script run all unit tests and store coverage output into coverage.out
# Require go >= 1.13

endpath=$(basename $(pwd))
echo $endpath
if [ $endpath == "tests" ]
then
    cd ..
fi

go test ./tests/benchmarks -bench=. > tests/benchmarks.out