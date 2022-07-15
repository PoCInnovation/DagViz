#!/bin/bash
CUEDIR=$1

if [ -z "$CUEDIR" ]; then
    echo "Usage: dagviz.sh <cue_dir>"
    exit 1
fi