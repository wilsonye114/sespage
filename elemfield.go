package sespage

import (
	"fmt"
)

/****************************************************************************
*  Table
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   2      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   3      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/

/****************************************************************************
* ElemTypeCodeElement
*
* +-------------------+-------------------------------------------+
* | Element Type Code | Name                                      |
* +===================+===========================================+
* |        00h        | Unspecified                               |
* +-------------------+-------------------------------------------+
* |        01h        | Device Slot                               |
* +-------------------+-------------------------------------------+
* |        02h        | Power Supply                              |
* +-------------------+-------------------------------------------+
* |        03h        | Cooling                                   |
* +-------------------+-------------------------------------------+
* |        04h        | Temperature Sensor                        |
* +-------------------+-------------------------------------------+
* |        05h        | Door Lock                                 |
* +-------------------+-------------------------------------------+
* |        06h        | Audible Alarm                             |
* +-------------------+-------------------------------------------+
* |        07h        | Enclosure Services Controller Electronics |
* +-------------------+-------------------------------------------+
* |        08h        | SCC Controller Electronics                |
* +-------------------+-------------------------------------------+
* |        09h        | Nonvolatile Cache                         |
* +-------------------+-------------------------------------------+
* |        0Ah        | Invalid Operation Reason c                |
* +-------------------+-------------------------------------------+
* |        0Bh        | Uninterruptible Power Supply              |
* +-------------------+-------------------------------------------+
* |        0Ch        | Display                                   |
* +-------------------+-------------------------------------------+
* |        0Dh        | Key Pad Entry                             |
* +-------------------+-------------------------------------------+
* |        0Eh        | Enclosure                                 |
* +-------------------+-------------------------------------------+
* |        0Fh        | SCSI Port/Transceiver                     |
* +-------------------+-------------------------------------------+
* |        10h        | Language                                  |
* +-------------------+-------------------------------------------+
* |        11h        | Communication Port                        |
* +-------------------+-------------------------------------------+
* |        12h        | Voltage Sensor                            |
* +-------------------+-------------------------------------------+
* |        13h        | Current Sensor                            |
* +-------------------+-------------------------------------------+
* |        14h        | SCSI Target Port                          |
* +-------------------+-------------------------------------------+
* |        15h        | SCSI Initiator Port                       |
* +-------------------+-------------------------------------------+
* |        16h        | Simple Subenclosure                       |
* +-------------------+-------------------------------------------+
* |        17h        | Array Device Slot                         |
* +-------------------+-------------------------------------------+
* |        18h        | SAS Expander                              |
* +-------------------+-------------------------------------------+
* |        19h        | SAS Connector                             |
* +-------------------+-------------------------------------------+
* |     1Ah ~ 7Fh     | Reserved                                  |
* +-------------------+-------------------------------------------+
* |     80h ~ FFh     | Vendor specific                           |
* +-------------------+-------------------------------------------+
*
****************************************************************************/
var (
	defaultElemTypeCodes Uint8Codes
	defaultElemTypeCodesInited bool = false
)

func NewElemTypeCodes() Uint8Codes {
	if defaultElemTypeCodesInited {
		return defaultElemTypeCodes
	}

	codes := make(Uint8Codes)
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

func NewElemTypeCodeElement() *Uint8CodeElement {
	elem := &Uint8CodeElement{codes: NewElemTypeCodes()}
	return elem
}

func CreateElemTypeCodeElement() Element {
	return NewElemTypeCodeElement()
}

/****************************************************************************
* Control element format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |    0     |                                   COMMON CONTROL                                              |
* |          |-----------+-----------+-----------+-----------+-----------------------------------------------|
* |          |  SELECT   |  PRDFAIL  |  DISABLE  |  RST SWAP |                   Reserved                    |
* +----------+-----------+-----------+-----------+-----------+-----------------------------------------------+
* |    1     |                           Element type specific control information                           |
* +----------+                                                                                               |
* |    3     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/
type CommonControlElement struct {
	Reserved_0_0 Uint4Field
	RstSwap Uint1Field
	Disable Uint1Field
	Prdfail Uint1Field
	Select Uint1Field
}

func (e *CommonControlElement) Decode(data []byte) error {
	b := data[0]
	e.Reserved_0_0.SetUint4(b & 0x0f)
	e.RstSwap.SetUint1((b >> 4) & 0x01)
	e.Disable.SetUint1((b >> 5) & 0x01)
	e.Prdfail.SetUint1((b >> 6) & 0x01)
	e.Select.SetUint1((b >> 7) & 0x01)
	return nil
}

func (e *CommonControlElement) Encode() ([]byte, error) {
	data := make([]byte, 1)
	data[0] = (e.Select.Uint1() << 7) | (e.Prdfail.Uint1() << 6) | (e.Disable.Uint1() << 5) | (e.RstSwap.Uint1() << 4) | e.Reserved_0_0.Uint4()
	return data, nil
}

func (e *CommonControlElement) Length() ElementLength {
	return ElementLength(8)
}

func NewCommonControlElement() *CommonControlElement {
	ef := GetEF("ses")
	elem := &CommonControlElement{
		Reserved_0_0: ef.CreateElement("Uint4Element").(Uint4Field),
		RstSwap: ef.CreateElement("Uint1Element").(Uint1Field),
		Disable: ef.CreateElement("Uint1Element").(Uint1Field),
		Prdfail: ef.CreateElement("Uint1Element").(Uint1Field),
		Select: ef.CreateElement("Uint1Element").(Uint1Field)}
	return elem
}

func CreateCommonControlElement() Element {
	return NewCommonControlElement()
}

/****************************************************************************
* Status element format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |    0     |                                    COMMON STATUS                                              |
* |          |-----------+-----------+-----------+-----------+-----------------------------------------------|
* |          |  Reserved |  PRDFAIL  |  DISABLE  |    SWAP   |            ELEMENT STATUS CODE                |
* +----------+-----------+-----------+-----------+-----------+-----------------------------------------------+
* |    1     |                           Element type specific status information                            |
* +----------+                                                                                               |
* |    3     |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/
var (
	gElemStatusCodes Uint8Codes
	gElemStatusCodesInited bool = false
)

func NewElemStatusCodes() Uint8Codes {
	if gElemStatusCodesInited {
		return gElemStatusCodes
	}

	codes := make(Uint8Codes)
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

func NewElemStatusCodeElement() *Uint8CodeElement {
	elem := &Uint8CodeElement{codes: NewElemStatusCodes()}
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

func (e *CommonStatusElement) Length() ElementLength {
	return ElementLength(8)
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


/****************************************************************************
* Threshold control element format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                         REQUESTED HIGH CRITICAL THRESHOLD                                     |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                         REQUESTED HIGH WARNING THRESHOLD                                      |
* +----------+-----------------------------------------------------------------------------------------------+
* |   2      |                         REQUESTED LOW WARNING THRESHOLD                                       |
* +----------+-----------------------------------------------------------------------------------------------+
* |   3      |                         REQUESTED LOW CRITICAL THRESHOLD                                      |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/

type ThresholdControlElement struct {
	RequestedHighCriticalThreshold Uint8Field
	RequestedHighWarningThreshold Uint8Field
	RequestedLowWarningThreshold Uint8Field
	RequestedLowCriticalThreshold Uint8Field
}

func (e *ThresholdControlElement) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	e.RequestedHighCriticalThreshold.SetUint8(data[0])
	e.RequestedHighWarningThreshold.SetUint8(data[1])
	e.RequestedLowWarningThreshold.SetUint8(data[2])
	e.RequestedLowCriticalThreshold.SetUint8(data[3])
	return nil
}

func (e *ThresholdControlElement) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data[0] = e.RequestedHighCriticalThreshold.Uint8()
	data[1] = e.RequestedHighWarningThreshold.Uint8()
	data[2] = e.RequestedLowWarningThreshold.Uint8()
	data[3] = e.RequestedLowCriticalThreshold.Uint8()
	return data, nil
}

func (e *ThresholdControlElement) Length() ElementLength {
	return ElementLength(32)
}

func NewThresholdControlElement() *ThresholdControlElement {
	ef := GetEF("ses")
	elem := &ThresholdControlElement{
		RequestedHighCriticalThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		RequestedHighWarningThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		RequestedLowWarningThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		RequestedLowCriticalThreshold: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateThresholdControlElement() Element {
	return NewThresholdControlElement()
}

/****************************************************************************
*  Threshold status element format
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                         HIGH CRITICAL THRESHOLD                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                         HIGH WARNING THRESHOLD                                                |
* +----------+-----------------------------------------------------------------------------------------------+
* |   2      |                         LOW WARNING THRESHOLD                                                 |
* +----------+-----------------------------------------------------------------------------------------------+
* |   3      |                         LOW CRITICAL THRESHOLD                                                |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/

type ThresholdStatusElement struct {
	HighCriticalThreshold Uint8Field
	HighWarningThreshold Uint8Field
	LowWarningThreshold Uint8Field
	LowCriticalThreshold Uint8Field
}

func (e *ThresholdStatusElement) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	e.HighCriticalThreshold.SetUint8(data[0])
	e.HighWarningThreshold.SetUint8(data[1])
	e.LowWarningThreshold.SetUint8(data[2])
	e.LowCriticalThreshold.SetUint8(data[3])
	return nil
}

func (e *ThresholdStatusElement) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data[0] = e.HighCriticalThreshold.Uint8()
	data[1] = e.HighWarningThreshold.Uint8()
	data[2] = e.LowWarningThreshold.Uint8()
	data[3] = e.LowCriticalThreshold.Uint8()
	return data, nil
}

func (e *ThresholdStatusElement) Length() ElementLength {
	return ElementLength(32)
}

func NewThresholdStatusElement() *ThresholdStatusElement {
	ef := GetEF("ses")
	elem := &ThresholdStatusElement{
		HighCriticalThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		HighWarningThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		LowWarningThreshold: ef.CreateElement("Uint8Element").(Uint8Field),
		LowCriticalThreshold: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateThresholdStatusElement() Element {
	return NewThresholdStatusElement()
}


/****************************************************************************
*  Unspecified control element
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                                    COMMON CONTROL                                             |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                                       Reserved                                                |
* +----------+                                                                                               |
* |   3      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/
type UnspecifiedControlElement struct {
	CommonControl Element
	Reserved_1_0 Uint8Field
	Reserved_2_0 Uint8Field
	Reserved_3_0 Uint8Field
}

func (e *UnspecifiedControlElement) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	e.CommonControl.Decode(data)
	e.Reserved_1_0.SetUint8(data[1])
	e.Reserved_2_0.SetUint8(data[2])
	e.Reserved_3_0.SetUint8(data[3])
	return nil
}

func (e *UnspecifiedControlElement) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data0, err := e.CommonControl.Encode()
	if err != nil {
		return data, err
	}
	data[0] = data0[0]
	data[1] = e.Reserved_1_0.Uint8()
	data[2] = e.Reserved_2_0.Uint8()
	data[3] = e.Reserved_3_0.Uint8()
	return data, nil
}

func (e *UnspecifiedControlElement) Length() ElementLength {
	return ElementLength(32)
}

func NewUnspecifiedControlElement() *UnspecifiedControlElement {
	ef := GetEF("ses")
	elem := &UnspecifiedControlElement{
		CommonControl: ef.CreateElement("CommonControlElement").(Element),
		Reserved_1_0: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_2_0: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_3_0: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateUnspecifiedControlElement() Element {
	return NewUnspecifiedControlElement()
}

/****************************************************************************
*  Unspecified status element
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                                    COMMON STATUS                                              |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                                       Reserved                                                |
* +----------+                                                                                               |
* |   3      |                                                                                               |
* +----------+-----------------------------------------------------------------------------------------------+
* 
****************************************************************************/
type UnspecifiedStatusElement struct {
	CommonStatus Element
	Reserved_1_0 Uint8Field
	Reserved_2_0 Uint8Field
	Reserved_3_0 Uint8Field
}

func (e *UnspecifiedStatusElement) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	err := e.CommonStatus.Decode(data)
	if err != nil {
		return err
	}
	e.Reserved_1_0.SetUint8(data[1])
	e.Reserved_2_0.SetUint8(data[2])
	e.Reserved_3_0.SetUint8(data[3])
	return nil
}

func (e *UnspecifiedStatusElement) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data0, err := e.CommonStatus.Encode()
	if err != nil {
		return data, err
	}
	data[0] = data0[0]
	data[1] = e.Reserved_1_0.Uint8()
	data[2] = e.Reserved_2_0.Uint8()
	data[3] = e.Reserved_3_0.Uint8()
	return data, nil
}

func (e *UnspecifiedStatusElement) Length() ElementLength {
	return ElementLength(32)
}

func NewUnspecifiedStatusElement() *UnspecifiedStatusElement {
	ef := GetEF("ses")
	elem := &UnspecifiedStatusElement{
		CommonStatus: ef.CreateElement("CommonStatusElement").(Element),
		Reserved_1_0: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_2_0: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_3_0: ef.CreateElement("Uint8Element").(Uint8Field)}
	return elem
}

func CreateUnspecifiedStatusElement() Element {
	return NewUnspecifiedStatusElement()
}

/****************************************************************************
*  Device Slot control element
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* | Byte\Bit |     7     |     6     |     5     |     4     |     3     |     2     |     1     |     0     |
* +==========+===========+===========+===========+===========+===========+===========+===========+===========+
* |   0      |                                   COMMON CONTROL                                              |
* +----------+-----------------------------------------------------------------------------------------------+
* |   1      |                                    Reserved                                                   |
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* |   2      | RQST      | DO NOT    | Reserved  | RQST      | RQST      | RQST      | RQST      | Reserved  |
* |          | ACTIVE    | REMOVE    |           | MISSING   | INSERT    | REMOVE    | IDENT     |           |
* +----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
* |   3      | Reserved              | RQST      | DEVICE    | ENABLE    | ENABLE    | Reserved              |
* |          |                       | FAULT     | OFF       | BYP A     | BYP B     |                       |
* +----------+-----------------------+-----------+-----------+-----------+-----------+-----------------------+
* 
****************************************************************************/
type DeviceSlotControlElement struct {
	CommonStatus Element
	Reserved_1_0 Uint8Field
	Reserved_2_0 Uint8Field
	RqstIdent Uint8Field
	RqstRemove Uint8Field
	RqstInsert Uint8Field
	RqstMissing Uint8Field
	Reserved_2_5 Uint8Field
	DoNotRemove Uint8Field
	RqstActive Uint8Field
	Reserved_3_0 Uint8Field
	EnableBypB Uint8Field
	EnableBypA Uint8Field
	DeviceOff Uint8Field
	RqstFault Uint8Field
	Reserved_3_6 Uint8Field
}

func (e *DeviceSlotControlElement) Decode(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Expect 4 bytes data, got %d", len(data))
	}
	err := e.CommonStatus.Decode(data)
	if err != nil {
		return err
	}
	e.Reserved_1_0.SetUint8(data[1])
	e.Reserved_2_0.SetUint8(data[2]&0x01)
	e.RqstIdent.SetUint8((data[2]>>1)&0x01)
	e.RqstRemove.SetUint8((data[2]>>2)&0x01)
	e.RqstInsert.SetUint8((data[2]>>3)&0x01)
	e.RqstMissing.SetUint8((data[2]>>4)&0x01)
	e.Reserved_2_5.SetUint8((data[2]>>5)&0x01)
	e.DoNotRemove.SetUint8((data[2]>>6)&0x01)
	e.RqstActive.SetUint8((data[2]>>7)&0x01)
	e.Reserved_3_0.SetUint8(data[3]&0x3)
	e.EnableBypB.SetUint8((data[3]>>2)&0x1)
	e.EnableBypA.SetUint8((data[3]>>3)&0x1)
	e.DeviceOff.SetUint8((data[3]>>4)&0x1)
	e.RqstFault.SetUint8((data[3]>>5)&0x1)
	e.Reserved_3_6.SetUint8((data[3]>>6)&0x3)

	return nil
}

func (e *DeviceSlotControlElement) Encode() ([]byte, error) {
	data := make([]byte, 4, 4)
	data0, err := e.CommonStatus.Encode()
	if err != nil {
		return data, err
	}
	data[0] = data0[0]
	data[1] = e.Reserved_1_0.Uint8()
	data[2] = (e.RqstActive.Uint8() << 7 |
			   e.DoNotRemove.Uint8() << 6 |
			   e.Reserved_2_5.Uint8() << 5 |
			   e.RqstMissing.Uint8() << 4 |
			   e.RqstInsert.Uint8() << 3 |
			   e.RqstRemove.Uint8() << 2 |
			   e.RqstIdent.Uint8() << 1 |
			   e.Reserved_2_0.Uint8())
	data[3] = (e.Reserved_3_6.Uint8() << 6 |
			   e.RqstFault.Uint8() << 5 |
			   e.DeviceOff.Uint8() << 4 |
			   e.EnableBypA.Uint8() << 3 |
			   e.EnableBypB.Uint8() << 2 |
			   e.Reserved_3_0.Uint8())
	return data, nil
}

func (e *DeviceSlotControlElement) Length() ElementLength {
	return ElementLength(32)
}

func NewDeviceSlotControlElement() *DeviceSlotControlElement {
	ef := GetEF("ses")
	elem := &DeviceSlotControlElement{
		CommonStatus: ef.CreateElement("CommonStatusElement").(Element),
		Reserved_1_0: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_2_0: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstIdent: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstRemove: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstInsert: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstMissing: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_2_5: ef.CreateElement("Uint8Element").(Uint8Field),
		DoNotRemove: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstActive: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_3_0: ef.CreateElement("Uint8Element").(Uint8Field),
		EnableBypB: ef.CreateElement("Uint8Element").(Uint8Field),
		EnableBypA: ef.CreateElement("Uint8Element").(Uint8Field),
		DeviceOff: ef.CreateElement("Uint8Element").(Uint8Field),
		RqstFault: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_3_6: ef.CreateElement("Uint8Element").(Uint8Field)}

	return elem
}

func CreateDeviceSlotControlElement() Element {
	return NewDeviceSlotControlElement()
}


