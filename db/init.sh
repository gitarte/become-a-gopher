#!/bin/bash
ls \
&& docker build -t psql-poc . \
&& docker run \
    -p 5432:5432 \
    psql-poc \
&& echo "ok"
#-v /home/artgaw/Dokumenty/projekty/vodeno:/key \
