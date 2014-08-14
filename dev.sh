#!/bin/bash

function add_path()
{
  # $1 path variable
  # $2 path to add
  if [ -d "$2" ] && [[ ":$1:" != *":$2:"* ]]; then
    echo "$1:$2"
  else
    echo "$1"
  fi
}

TOKUDB_DIR=/usr/local/tokuft

export CGO_CFLAGS="$CGO_CFLAGS -I$TOKUDB_DIR/include"
export CGO_CXXFLAGS="$CGO_CXXFLAGS -I$TOKUDB_DIR/include"
export CGO_LDFLAGS="$CGO_LDFLAGS -L$TOKUDB_DIR/lib -ltokufractaltree -ltokuportability"
export LD_LIBRARY_PATH=$(add_path $LD_LIBRARY_PATH $TOKUDB_DIR/lib)
export DYLD_LIBRARY_PATH=$(add_path $DYLD_LIBRARY_PATH $TOKUDB_DIR/lib)
