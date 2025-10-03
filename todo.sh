#! /usr/bin/env bash

grep -rni --exclude=todo.sh --exclude-dir=.git --exclude-dir=tmp --color=always "todo" .
