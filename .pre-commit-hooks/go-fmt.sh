#!/usr/bin/env bash

set -e

# Store and return last failure from fmt so this can validate every directory passed before exiting
FMT_ERROR=0

for file in "$@"; do
    # do not fix the files, just error out
    output=$(gofmt -l -e -d "$file")

    if [[ -n "${output}" ]]; then
        echo "${output}"
        FMT_ERROR=1
    fi
done

exit ${FMT_ERROR}
