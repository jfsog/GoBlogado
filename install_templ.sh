#!/bin/bash
mkdir -p bin &&\
    wget -q https://github.com/a-h/templ/releases/latest/download/templ_Linux_x86_64.tar.gz -O - | tar -xzv -C bin
