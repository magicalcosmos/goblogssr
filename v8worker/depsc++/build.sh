#!/bin/bash

g++ -c -O2 -Wall -std=c++20 v8binding.cpp -I/opt/homebrew/Cellar/v8/12.7.224.16/libexec/include
rm -f libdepsc++.a
ar -r libdepsc++.a v8binding.o
