#!/bin/sh
protoc -I=. --go_out=. ./*.proto