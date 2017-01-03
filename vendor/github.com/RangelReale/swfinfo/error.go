package swfinfo

import (
	"fmt"
)

type BadHeader struct {
	Code uint8
	Err  error
}

func (b BadHeader) Error() string {
	switch b.Code {
	case 0:
		fmt.Sprintf("error while reading header: %q", b.Err)
	case 1:
		fmt.Sprintf("invalid signature: %q", b.Err)
	case 2:
		fmt.Sprintf("invalid compression type")
	}
	return "unknown error"
}
