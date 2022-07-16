#!/bin/bash

sudo rm -f "/usr/local/bin/dagviz"

echo "#!/bin/bash" > dagviz.sh
echo "PATHTOREPO=\"$PWD/electron\"" >> dagviz.sh
echo "cd \$PATHTOREPO || echo \"can't find the location\"; npm run dev; cd -" >> dagviz.sh

sudo ln -s "$PWD/dagviz.sh" "/usr/local/bin/dagviz"