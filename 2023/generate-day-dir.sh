#!/bin/bash

if [[ $# -gt 0 ]]; then
    set -x

    printf -v NN "%02d" $1

    daynn='day'$NN

    cp -ir dayNN/ ${daynn}

    sed -i 's/dayNN/'${daynn}'/' ${daynn}/go.mod

    ~/go/bin/aocdl -year 2023 -day $1

    mv input.txt ${daynn}/
else
    echo 'You need to pass a number for the day'
fi
