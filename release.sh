#!/bin/bash

./build.sh
echo -e \\n

### hela
version=`cat VERSION`

mkdir -p hela-$version
#rm hela-$version.tar.gz || true
cp hela hela-$version/
cp config.json hela-$version/
cp Readme.md hela-$version/
cp VERSION hela-$version/
#cp -rf settings/ hela-$version/settings/

rm -rf ~/Desktop/release/hela || true
rm ~/Desktop/release/hela-$version.tar.gz || true
mkdir ~/Desktop/release/hela
cp -rf hela-$version ~/Desktop/release/hela
rm -rf hela-$version
cd ~/Desktop/release/hela
ln -s hela-$version latest
chmod 777 latest
cd ../

sleep 1s
tar -zcvf hela-$version.tar.gz hela
#rm -rf hela-$version
