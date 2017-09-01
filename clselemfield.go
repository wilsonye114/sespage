package sespage

import (
	"fmt"
	"bytes"
)

/****************************************************************************
* VPD type
*
* +-------------------+----------------------+
* | VPD Type Code | Name                     |
* +===================+======================+
* |      02h      | Power Supply             |
* +-------------------+----------------------+
* |      07h      | Canister                 |
* +-------------------+----------------------+
* |      0eh      | Midplane                 |
* +-------------------+----------------------+
* |      ffh      | All VPD in the Enclosure |
* +-------------------+----------------------+
*
* Issue: can't decode when 'VPD Type Cdoe' is ffh.
****************************************************************************/
var (
	defaultVpdTypeCodes Uint8Codes
	defaultVpdTypeCodesInited bool = false
)

func NewVpdTypeCodes() Uint8Codes {
	if defaultVpdTypeCodesInited {
		return defaultVpdTypeCodes
	}

	codes := make(Uint8Codes)
	for i := 0; i <= 0xff; i++ {
		codes[uint8(i)] = "Reserved"
	}
	codes[0x02] = "Power Supply"
	codes[0x07] = "Canister"
	codes[0x0e] = "Midplane"
	codes[0xff] = "All VPD in the Enclosure"
	defaultVpdTypeCodes = codes
	defaultVpdTypeCodesInited = true

	return defaultVpdTypeCodes
}

func NewVpdTypeCodeElement() *Uint8CodeElement {
	elem := &Uint8CodeElement{codes: NewVpdTypeCodes()}
	return elem
}

func CreateTypeCodeElement() Element {
	return NewVpdTypeCodeElement()
}

/****************************************************************************
*  Midplane VPD Data Format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                              Board Product Name                                               |
* +----------+                                                                                               |
* |   17     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   18     |                              Board Part Number                                                |
* +----------+                                                                                               |
* |   37     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   38     |                              Board Serial Number                                              |
* +----------+                                                                                               |
* |   57     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   58     |                              Board Hardware EC Level                                          |
* +----------+                                                                                               |
* |   65     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   66     |                              Product Name                                                     |
* +----------+                                                                                               |
* |   83     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   84     |                              Porduct Part Number                                              |
* +----------+                                                                                               |
* |   99     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   100    |                              Product Serial Number                                            |
* +----------+                                                                                               |
* |   115    |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   116    |                              Product Version                                                  |
* +----------+                                                                                               |
* |   119    |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/
type MidplaneVpdDataElement struct {
	BoardProductName StringField
	BoardPartNumber StringField
	BoardSerialNumber StringField
	BoardHardwareECLevel StringField
	ProductName StringField
	PorductPartNumber StringField
	ProductSerialNumber StringField
	ProductVersion StringField
}

func (e *MidplaneVpdDataElement) Decode(data []byte) error {
	var err error

	if len(data) < 120 {
		return fmt.Errorf("Midplane VPD should have 120 bytes, get %d bytes", len(data))
	}

	err = e.BoardProductName.Decode(data[0:18])
	if err != nil {
		return err
	}	
	err = e.BoardPartNumber.Decode(data[18:38])
	if err != nil {
		return err
	}
	err = e.BoardSerialNumber.Decode(data[38:58])
	if err != nil {
		return err
	}	
	err = e.BoardHardwareECLevel.Decode(data[58:66])
	if err != nil {
		return err
	}
	err = e.ProductName.Decode(data[66:84])
	if err != nil {
		return err
	}	
	err = e.PorductPartNumber.Decode(data[84:100])
	if err != nil {
		return err
	}
	err = e.ProductSerialNumber.Decode(data[100:116])
	if err != nil {
		return err
	}	
	err = e.ProductVersion.Decode(data[116:120])
	if err != nil {
		return err
	}
	return nil
}

func (e *MidplaneVpdDataElement) Encode() ([]byte, error) {
	var err error

	data := bytes.NewBuffer([]byte{})

	encodeField := func(field Element) error {
		buf, err := field.Encode()
		if err != nil {
			return err
		}
		_, err = data.Write(buf)
		if err != nil {
			return err
		}
		return nil
	}
	err = encodeField(e.BoardProductName)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardPartNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardSerialNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardHardwareECLevel)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.ProductName)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.PorductPartNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.ProductSerialNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.ProductVersion)
	if err != nil {
		return data.Bytes(), err
	}
	return data.Bytes(), nil
}

func (e *MidplaneVpdDataElement) Length() ElementLength {
	length := e.BoardProductName.Length() +
		e.BoardPartNumber.Length() +
		e.BoardSerialNumber.Length() +
		e.BoardHardwareECLevel.Length() +
		e.ProductName.Length() +
		e.PorductPartNumber.Length() +
		e.ProductSerialNumber.Length() +
		e.ProductVersion.Length()
	return length
}

func NewMidplaneVpdDataElement() *MidplaneVpdDataElement {
	ef := GetEF("ses")
	elem := &MidplaneVpdDataElement{
		BoardProductName: ef.CreateElement("StringElement").(StringField),
		BoardPartNumber: ef.CreateElement("StringElement").(StringField),
		BoardSerialNumber: ef.CreateElement("StringElement").(StringField),
		BoardHardwareECLevel: ef.CreateElement("StringElement").(StringField),
		ProductName: ef.CreateElement("StringElement").(StringField),
		PorductPartNumber: ef.CreateElement("StringElement").(StringField),
		ProductSerialNumber: ef.CreateElement("StringElement").(StringField),
		ProductVersion: ef.CreateElement("StringElement").(StringField)}
	return elem
}

func CreateMidplaneVpdDataElement() Element {
	return NewMidplaneVpdDataElement()
}

/****************************************************************************
*  Canister VPD Data Format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                              Board Product Name                                               |
* +----------+                                                                                               |
* |   15     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   16     |                              Board Part Number                                                |
* +----------+                                                                                               |
* |   35     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   36     |                              Board Serial Number                                              |
* +----------+                                                                                               |
* |   55     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   56     |                              Board Hardware EC Level                                          |
* +----------+                                                                                               |
* |   59     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/

type CanisterVpdDataElement struct {
	BoardProductName StringField
	BoardPartNumber StringField
	BoardSerialNumber StringField
	BoardHardwareECLevel StringField
}

func (e *CanisterVpdDataElement) Decode(data []byte) error {
	var err error

	if len(data) < 60 {
		return fmt.Errorf("Canister VPD should have 60 bytes, get %d bytes", len(data))
	}

	err = e.BoardProductName.Decode(data[0:16])
	if err != nil {
		return err
	}	
	err = e.BoardPartNumber.Decode(data[16:36])
	if err != nil {
		return err
	}
	err = e.BoardSerialNumber.Decode(data[36:56])
	if err != nil {
		return err
	}	
	err = e.BoardHardwareECLevel.Decode(data[56:60])
	if err != nil {
		return err
	}
	return nil
}

func (e *CanisterVpdDataElement) Encode() ([]byte, error) {
	var err error

	data := bytes.NewBuffer([]byte{})

	encodeField := func(field Element) error {
		buf, err := field.Encode()
		if err != nil {
			return err
		}
		_, err = data.Write(buf)
		if err != nil {
			return err
		}
		return nil
	}
	err = encodeField(e.BoardProductName)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardPartNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardSerialNumber)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.BoardHardwareECLevel)
	if err != nil {
		return data.Bytes(), err
	}
	return data.Bytes(), nil
}

func (e *CanisterVpdDataElement) Length() ElementLength {
	length := e.BoardProductName.Length() +
		e.BoardPartNumber.Length() +
		e.BoardSerialNumber.Length() +
		e.BoardHardwareECLevel.Length()
	return length
}

func NewCanisterVpdDataElement() *CanisterVpdDataElement {
	ef := GetEF("ses")
	elem := &CanisterVpdDataElement{
		BoardProductName: ef.CreateElement("StringElement").(StringField),
		BoardPartNumber: ef.CreateElement("StringElement").(StringField),
		BoardSerialNumber: ef.CreateElement("StringElement").(StringField),
		BoardHardwareECLevel: ef.CreateElement("StringElement").(StringField)}
	return elem
}

func CreateCanisterVpdDataElement() Element {
	return NewCanisterVpdDataElement()
}