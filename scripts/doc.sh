#!/usr/bin/env bash

set -euo pipefail

godoc2md github.com/andy2046/goji \
    > $GOPATH/src/github.com/andy2046/goji/docs/goji.md

godoc2md github.com/andy2046/goji/pkg/strutil \
    > $GOPATH/src/github.com/andy2046/goji/docs/strutil.md

godoc2md github.com/andy2046/goji/pkg/intutil \
    > $GOPATH/src/github.com/andy2046/goji/docs/intutil.md

godoc2md github.com/andy2046/goji/pkg/float32util \
    > $GOPATH/src/github.com/andy2046/goji/docs/float32util.md
