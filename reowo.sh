#!/bin/bash
gocc -a -o gocc owo.bnf &&
echo "gocc'd; building" &&
go build &&
echo "built; running" &&
./owo < test.owo
