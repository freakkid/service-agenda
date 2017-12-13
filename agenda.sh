#!/bin/sh
set -e

if [ "$0" = "agendad" ] || [ $# = 0 ]
then
    exec agendad
else
    exec agenda "$@"
fi