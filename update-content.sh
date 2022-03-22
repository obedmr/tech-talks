#!/bin/bash

github_prefix=https://github.com/CodersSquad
projects=(hello-gophers hello-c-world)

for proj in "${projects[@]}"; do
    git clone $github_prefix/$proj
    pushd $proj
    ./generate-content.sh
    popd
    rm -rf content/$proj
    mv -f $proj/public content/$proj
    rm -rf $proj
done

git add .
git commit -s -m "Update on $(date)"
