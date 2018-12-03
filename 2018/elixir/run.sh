#!/bin/bash

if [ $# -lt 1 ]; then
    echo -e "Please tell me what daily quiz you want to run:"
    for i in `ls -d */`; do
        echo -e " => ${i}"
    done
    exit 1
fi

./$1/main.exs
