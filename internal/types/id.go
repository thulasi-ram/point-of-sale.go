package types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"strconv"
)

type ID struct {
	val    string
	intVal int
	valid  bool
}

func (i *ID) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeString(i.val)
}

func (i *ID) DecodeMsgpack(dec *msgpack.Decoder) error {
	s, err := dec.DecodeString()
	if err != nil {
		return err
	}
	i.val = s
	return nil
}

func (i *ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *ID) UnmarshalJSON(data []byte) error {
	sVal := string(data)

	// verify if its a proper int
	ival, err := strconv.Atoi(sVal)
	if err != nil {
		return err
	}

	i.val = sVal
	i.intVal = ival

	return nil
}

func (i ID) String() string {
	return i.val
}

func (i ID) Int() int {
	return i.intVal
}

func (i ID) Int64() int64 {
	return int64(i.intVal)
}
