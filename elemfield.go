package sespage

import (
	// "fmt"
)

/****************************************************************************
* ElemTypeCodeElement
****************************************************************************/
var (
	defaultElemTypeCodes Codes
	defaultElemTypeCodesInited bool = false
)

func NewElemTypeCodes() Codes {
	if defaultElemTypeCodesInited {
		return defaultElemTypeCodes
	}

	codes := make(Codes)
	codes[0x00] = "Unspecified"
	codes[0x01] = "Device Slot"
	codes[0x02] = "Power Supply"
	codes[0x03] = "Cooling"
	codes[0x04] = "Temperature Sensor"
	codes[0x05] = "Door Lock"
	codes[0x06] = "Audible Alarm"
	codes[0x07] = "Enclosure Services Controller Electronics"
	codes[0x08] = "SCC Controller Electronics"
	codes[0x09] = "Nonvolatile Cache"
	codes[0x0A] = "Invalid Operation Reason c"
	codes[0x0B] = "Uninterruptible Power Supply"
	codes[0x0C] = "Display"
	codes[0x0D] = "Key Pad Entry"
	codes[0x0E] = "Enclosure"
	codes[0x0F] = "SCSI Port/Transceiver"
	codes[0x10] = "Language"
	codes[0x11] = "Communication Port"
	codes[0x12] = "Voltage Sensor"
	codes[0x13] = "Current Sensor"
	codes[0x14] = "SCSI Target Port"
	codes[0x15] = "SCSI Initiator Port"
	codes[0x16] = "Simple Subenclosure"
	codes[0x17] = "Array Device Slot"
	codes[0x18] = "SAS Expander"
	codes[0x19] = "SAS Connector"
	for i := uint8(0x1A); i >= 0x1A && i <= 0x7F; i++ {
		codes[i] = "Reserved"
	}
	for i := uint8(0x80); i >= 0x80 && i <= 0xFF; i++ {
		codes[i] = "Vendor specific"
	}

	defaultElemTypeCodes = codes
	defaultElemTypeCodesInited = true
	return defaultElemTypeCodes
}

func NewElemTypeCodeElement() *CodeElement {
	elem := &CodeElement{codes: NewElemTypeCodes()}
	return elem
}

func CreateElemTypeCodeElement() Element {
	return NewElemTypeCodeElement()
}

/****************************************************************************
 Table
+----------------------------------------------------------------------------------------------------------+
| Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
+==========+===============================================================================================+
|   0      |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+
|   1      |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+
|   2      |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+
|   3      |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+

****************************************************************************/

/****************************************************************************
Control element format
+----------------------------------------------------------------------------------------------------------+
| Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
+==========+===============================================================================================+
|    0     |                                   COMMON CONTROL                                              |
|          |-----------------------------------------------------------------------------------------------|
|          |  SELECT   |  PRDFAIL  |  DISABLE  |  RST SWAP |                   Reserved                    |
+----------+-----------------------------------------------+-----------------------------------------------+
|    1     |                           Element type specific control information                           |
+----------+-----------                                                                        ------------+
|    3     |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+

****************************************************************************/
type CommonControlElement struct {
	Reserved_0_0 Uint8Field
	RstSwap Uint8Field
	Disable Uint8Field
	Prdfail Uint8Field
	Select Uint8Field
}

func (e *CommonControlElement) Decode(data []byte) error {
	b := data[0]
	e.Reserved_0_0.SetUint8(b & 0x0f)
	e.RstSwap.SetUint8((b >> 4) & 0x01)
	e.Disable.SetUint8((b >> 5) & 0x01)
	e.Prdfail.SetUint8((b >> 6) & 0x01)
	e.Select.SetUint8((b >> 7) & 0x01)
	return nil
}

func (e *CommonControlElement) Encode() ([]byte, error) {
	data := make([]byte, 1)
	data[0] = (e.Select.Uint8() << 7) | (e.Prdfail.Uint8() << 6) | (e.Disable.Uint8() << 5) | (e.RstSwap.Uint8() << 4) | e.Reserved_0_0.Uint8()
	return data, nil
}

func (e *CommonControlElement) Length() int32 {
	return 1
}

func NewCommonControlElement() *CommonControlElement {
	ef := GetEF("ses")
	elem := &CommonControlElement{
		Reserved_0_0: ef.CreateElement("Uint8Element").(Uint8Field),
		RstSwap: ef.CreateElement("Uint8Element").(Uint8Field),
		Disable: ef.CreateElement("Uint8Element").(Uint8Field),
		Prdfail: ef.CreateElement("Uint8Element").(Uint8Field),
		Select: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateCommonControlElement() Element {
	return NewCommonControlElement()
}

/****************************************************************************
Status element format
+----------------------------------------------------------------------------------------------------------+
| Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
+==========+===============================================================================================+
|    0     |                                    COMMON STATUS                                              |
|          |-----------------------------------------------------------------------------------------------|
|          |  Reserved |  PRDFAIL  |  DISABLE  |    SWAP   |            ELEMENT STATUS CODE                |
+----------+-----------------------------------------------+-----------------------------------------------+
|    1     |                           Element type specific status information                            |
+----------+-----------                                                                        ------------+
|    3     |                                                                                               |
+----------+-----------------------------------------------------------------------------------------------+

****************************************************************************/
var (
	gElemStatusCodes Codes
	gElemStatusCodesInited bool = false
)

func NewElemStatusCodes() Codes {
	if gElemStatusCodesInited {
		return gElemStatusCodes
	}

	codes := make(Codes)
	codes[0x00] = "Unsupported"
	codes[0x01] = "OK"
	codes[0x02] = "Critical"
	codes[0x03] = "Noncritical"
	codes[0x04] = "Unrecoverable"
	codes[0x05] = "Not Installed"
	codes[0x06] = "Unknown"
	codes[0x07] = "Not Available"
	codes[0x08] = "No Access Allowed"
	for i := uint8(0x09); (i >= 0x09) && (i <= 0x0f); i++ {
		codes[i] = "Reserved"
	}
	gElemStatusCodes = codes
	gElemStatusCodesInited = true
	return gElemStatusCodes
}

func NewElemStatusCodeElement() *CodeElement {
	elem := &CodeElement{codes: NewElemStatusCodes()}
	return elem
}

func CreateElemStatusCodeElement() Element {
	return NewElemStatusCodeElement()
}

type CommonStatusElement struct {
	ElementStatusCode Uint8Field
	Swap Uint8Field
	Disable Uint8Field
	Prdfail Uint8Field
	Reserved_0_7 Uint8Field
}

func (e *CommonStatusElement) Decode(data []byte) error {
	b := data[0]
	e.ElementStatusCode.SetUint8(b & 0x0f)
	e.Swap.SetUint8((b >> 4) & 0x01)
	e.Disable.SetUint8((b >> 5) & 0x01)
	e.Prdfail.SetUint8((b >> 6) & 0x01)
	e.Reserved_0_7.SetUint8((b >> 7) & 0x01)
	return nil
}

func (e *CommonStatusElement) Encode() ([]byte, error) {
	data := make([]byte, 1)
	data[0] = (e.Reserved_0_7.Uint8() << 7) | (e.Prdfail.Uint8() << 6) | (e.Disable.Uint8() << 5) | (e.Swap.Uint8() << 4) | e.ElementStatusCode.Uint8()
	return data, nil
}

func (e *CommonStatusElement) Length() int32 {
	return 1
}

func NewCommonStatusElement() *CommonStatusElement {
	ef := GetEF("ses")
	elem := &CommonStatusElement{
		ElementStatusCode: ef.CreateElement("ElemStatusCodeElement").(Uint8Field),
		Swap: ef.CreateElement("Uint8Element").(Uint8Field),
		Disable: ef.CreateElement("Uint8Element").(Uint8Field),
		Prdfail: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_0_7: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateCommonStatusElement() Element {
	return NewCommonStatusElement()
}
