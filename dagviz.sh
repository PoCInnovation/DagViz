#!/bin/bash
PATHTOREPO="/home/isma/Developement/Go/DagViz/electron"
ARG=$1

if [ -z "$ARG" ]; then
  ARG=$PWD
fi

# check the if the directory exists

cd $PATHTOREPO || exit; CUEDIR="-cueDir=\"$ARG\"" npm run dev; cd - || exit;
