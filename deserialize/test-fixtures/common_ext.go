package protos

import (
	"errors"
	"fmt"
	"strings"
)

func (ct *SubCredType) UnmarshalJSON(val []byte) error {
	typ := strings.ToUpper(string(val))
	fmt.Println(typ)
	sct, ok := SubCredType_value[typ]
	if !ok {
		return errors.New("gasp! what is this supposed to be? " + typ)
	}
	//*f = append(*f, flag)
	*ct = SubCredType(sct)
	return nil
}