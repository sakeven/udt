#!/bin/sh

echo 'Preparing to build UDT C library.'

os=''
arch=''

case "$(uname -s)" in
  'Darwin')
  os='OSX'
  ;;
  'Linux')
  os='LINUX'
  ;;
  *)
  echo 'Unknown OS.'
  exit 1
  ;;
esac

case "$(uname -m)" in
  'x86_64')
  arch='AMD64'
  ;;
  *)
  echo 'Unknown architecture.'
  exit 1
  ;;
esac

echo "Building for ${os} ${arch}"

# get script abs dir
script_dir="$( cd "$( dirname $0 )" && pwd )"
udt4_project_dir="${script_dir}/udt4"

cd "${udt4_project_dir}"

make clean
make -e os=$os arch=$arch

if [ $? -ne 0 ]; then
  echo "Build failed with error ${?}"
  exit 1
else
  echo "UDT4 build succeeded."
fi

case $os in
  'OSX')
  export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:${udt4_project_dir}/src
  ;;
  'LINUX')
  export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:${udt4_project_dir}/src
  ;;
esac

# app/test
# if [ $? -ne 0 ]; then
#   echo "UDT4 tests failed."
#   exit 1
# else
#   echo "UDT4 tests succeeded."
# fi

# echo "UDT4 build was successful."