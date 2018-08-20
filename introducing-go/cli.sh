#!/bin/bash

go build cli.go

for (( i=0; i<20; ++i)); do
	./cli -max=1000 non_flag -test
done

rm cli
