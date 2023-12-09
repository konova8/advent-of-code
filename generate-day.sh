#!/bin/bash

set -x

D=0
Y=0

if [[ $# -gt 1 ]]; then
    D=$2
    Y=$1
    D=$(expr $D + 0)
    Y=$(expr $Y + 0)
else
    Y=$(date +"%Y")
    D=$(date +"%d")
    D=$(expr $D + 0)
    Y=$(expr $Y + 0)
fi

printf -v YYYY "%04d" $Y
printf -v DD "%02d" $D

echo "Generating dir and files for Year" $YYYY "and day" $D

dayDD='day'$DD

mkdir -p $YYYY/${dayDD}
cp -i templ/main.go $YYYY/${dayDD}/main.go
cp -i templ/main_test.go $YYYY/${dayDD}/main_test.go

if [ $(TZ=America/New_York date +"%Y") -ge $YYYY -a $(TZ=America/New_York date +"%d") -lt $DD ]; then
    echo "We have to wait before problem goes live"
    ~/go/bin/aocdl -year $YYYY -day $DD -wait
else
    ~/go/bin/aocdl -year $YYYY -day $DD
fi

firefox --new-window "https://adventofcode.com/${YYYY}/day/${D}"
mv input.txt $YYYY/${dayDD}/input.txt
touch $YYYY/${dayDD}/example.txt
touch $YYYY/${dayDD}/example2.txt
$EDITOR $YYYY/${dayDD}/main.go
