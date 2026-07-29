package psl

import "github.com/fxamacker/cbor/v2"

var strictDec, _ = cbor.DecOptions{
    MaxArrayElements: 1024,
    MaxMapPairs:      1024,
    MaxNestedLevels:  5, // set to actual
}.DecMode()