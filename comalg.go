package sespage

import (
	"bytes"
)

func EncodeElements(elems ...Element) ([]byte, error) {
	data := bytes.NewBuffer([]byte{})

	for _, elem := range elems { 
		buf, err := elem.Encode()
		if err != nil {
			return data.Bytes(), err
		}
		_, err = data.Write(buf)
		if err != nil {
			return data.Bytes(), err
		}
	}
	return data.Bytes(), nil
}