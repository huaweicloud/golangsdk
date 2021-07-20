package utils

import (
	"bytes"
	"encoding/gob"
)

/**
deep-copies src to desc
using go.gob
**/
func CopyProperties(src interface{}, desc interface{}) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	dec := gob.NewDecoder(&buff)
	enc.Encode(src)
	dec.Decode(desc)
}
