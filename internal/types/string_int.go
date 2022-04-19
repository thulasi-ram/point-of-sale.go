package types

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"strconv"
)

type StringInt int

func (i *StringInt) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeString(strconv.Itoa(i.Int()))
}

func (i *StringInt) DecodeMsgpack(dec *msgpack.Decoder) error {
	s, err := dec.DecodeString()
	if err != nil {
		return err
	}
	return i.UnmarshalJSON([]byte(s))
}

func (i *StringInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(i.Int()))
}

func (i *StringInt) UnmarshalJSON(data []byte) error {
	val, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*i = StringInt(val)
	return nil
}

func (i StringInt) Int() int {
	return int(i)
}

func (i StringInt) Int64() int64 {
	return int64(i)
}
