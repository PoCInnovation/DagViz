#!/bin/bash

sudo rm -f "/usr/local/bin/dagviz"

cd electron || exit; npm install; cd - || exit;
cd go || exit; go build -o dagviz; cd - || exit;

echo "#!/bin/bash" > dagviz.sh
echo "PATHTOREPO=\"$PWD/electron\"" >> dagviz.sh
echo "ARG=\$1" >> dagviz.sh
echo "if [ -z \"\$ARG\" ]; then" >> dagviz.sh
echo "  ARG=\$PWD" >> dagviz.sh
echo "elif [[ \"\$ARG::1\" != \"/\" ]]; then" >> dagviz.sh
echo "  NEWARG=\"\$PWD/\$ARG\"" >> dagviz.sh
echo "  ARG=\$NEWARG" >> dagviz.sh
echo "fi" >> dagviz.sh
echo "cd \$PATHTOREPO || exit; CUEDIR=\"-cueDir=\\\"\$ARG\\\"\" npm run dev; cd - || exit;" >> dagviz.sh

sudo ln -s "$PWD/dagviz.sh" "/usr/local/bin/dagviz"
