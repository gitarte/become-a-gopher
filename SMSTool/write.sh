#!/bin/bash
/bin/echo $1 | \
/usr/bin/gammu \
    -c ./gammu-config \
    sendsms \
    TEXT \
    $2
