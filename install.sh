#!/bin/bash

sudo rm -f "/usr/local/bin/dagviz"

echo "#!/bin/bash" > dagviz.sh
echo "PATHTOREPO=\"$PWD/electron\"" >> dagviz.sh
echo "ARG=\$1" >> dagviz.sh
echo "if [ -z \"\$ARG\" ]; then" >> dagviz.sh
echo "  ARG=\$PWD" >> dagviz.sh
echo "fi" >> dagviz.sh
echo "cd \$PATHTOREPO || exit; CUEDIR=\"-cueDir=\\\"\$ARG\\\"\" npm run dev; cd - || exit;" >> dagviz.sh

sudo ln -s "$PWD/dagviz.sh" "/usr/local/bin/dagviz"
