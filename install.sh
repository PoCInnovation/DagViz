#!/bin/bash
sudo mkdir /usr/local/lib/DagViz

echo "Copying files to /usr/local/lib/DagViz"

sudo cp -r ./electron/ /usr/local/lib/DagViz/

echo "building electron"

cd /usr/local/lib/DagViz/electron || exit; npm install; npm run build; cd - || exit

echo "building go"

go build -o /usr/local/lib/DagViz/dagviz
