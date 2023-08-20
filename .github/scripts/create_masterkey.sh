#!/bin/bash

cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-30} | head -n 1 > ./master.key
echo "Masterkey: $(cat ./master.key)"
