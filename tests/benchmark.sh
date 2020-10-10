#!/bin/bash

# This script run benchmark for all algorithms through StringSimilarity function and store output into tests/benchmarks.out
# Require go >= 1.13

endpath=$(basename $(pwd))
if [ $endpath == "tests" ]
then
    cd ..
fi

go test ./tests/benchmarks -bench=. -benchmem | tee tests/outputs/benchmarks.txt