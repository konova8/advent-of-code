#!/bin/bash

if [[ $# -gt 1 ]]; then
    set -x

    printf -v YYYY $1
    printf -v NN "%02d" $2

    daynn='day'$NN

    mkdir -p $YYYY/${daynn}
    cp -i templ/main.go $YYYY/${daynn}/main.go

    ~/go/bin/aocdl -year $YYYY -day $NN

    mv input.txt $YYYY/${daynn}/input.txt
    touch $YYYY/${daynn}/example.txt
else
    echo 'You need to pass a number for year and one for the day'
fi
