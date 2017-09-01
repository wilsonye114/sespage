package sespage

import (
	"bytes"
)

type VpdStatusDescriptorElement struct {
	VpdType Uint8Field
	VpdId Uint8Field
	VpdData Element
}

func (e *VpdStatusDescriptorElement) Decode(data []byte) error {
	err := e.VpdType.Decode(data[0:1])
	if err != nil {
		return err
	}
	err = e.VpdId.Decode(data[1:2])
	if err != nil {
		return err
	}

	ef := GetEF("ses")
	switch e.VpdType.(*Uint8CodeElement).Code {
	// case 0x02:
	case 0x07:
		e.VpdData = ef.CreateElement("CanisterVpdDataElement")
	case 0x0e:
		e.VpdData = ef.CreateElement("MidplaneVpdDataElement")
	// case 0xff:
	default:
		e.VpdData = ef.CreateElement("BytesElement")
	}
	err = e.VpdData.Decode(data[2:])
	if err != nil {
		return err
	}
	return nil
}

func (e *VpdStatusDescriptorElement) Encode() ([]byte, error) {
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
	err = encodeField(e.VpdType)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.VpdId)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.VpdData)
	if err != nil {
		return data.Bytes(), err
	}
	return data.Bytes(), nil
}

func (e *VpdStatusDescriptorElement) Length() ElementLength {
	return e.VpdType.Length() + e.VpdId.Length() + e.VpdData.Length()
}

func NewVpdStatusDescriptorElement() *VpdStatusDescriptorElement {
	ef := GetEF("ses")
	elem := &VpdStatusDescriptorElement{
	VpdType: ef.CreateElement("TypeCodeElement").(Uint8Field),
	VpdId: ef.CreateElement("Uint8Element").(Uint8Field)}
	// VpdData type is unknown before decode.
	return elem
}

func CreateVpdStatusDescriptorElement() Element {
	return NewVpdStatusDescriptorElement()
}

type SesPage12 struct {
	PageCode Uint8Field
	Reserved_1_0 Uint8Field
	PageLength Uint16Field
	CompletionCode Uint8Field
	VpdStatusDescriptor Element
}

func (e *SesPage12) Decode(data []byte) error {
	err := e.PageCode.Decode(data[0:1])
	if err != nil {
		return err
	}
	err = e.Reserved_1_0.Decode(data[1:2])
	if err != nil {
		return err
	}
	err = e.PageLength.Decode(data[2:4])
	if err != nil {
		return err
	}
	err = e.CompletionCode.Decode(data[4:5])
	if err != nil {
		return err
	}
	err = e.VpdStatusDescriptor.Decode(data[5:e.PageLength.Uint16()+4])
	if err != nil {
		return err
	}
	return nil
}

func (e *SesPage12) Encode() ([]byte, error) {
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
	err = encodeField(e.PageCode)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.Reserved_1_0)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.PageLength)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.CompletionCode)
	if err != nil {
		return data.Bytes(), err
	}
	err = encodeField(e.VpdStatusDescriptor)
	if err != nil {
		return data.Bytes(), err
	}
	return data.Bytes(), nil
}

func (e *SesPage12) Length() ElementLength {
	return e.PageCode.Length() + e.Reserved_1_0.Length() + e.PageLength.Length() + e.CompletionCode.Length() + e.VpdStatusDescriptor.Length()
}

func NewSesPage12() *SesPage12 {
	ef := GetEF("ses")
	elem := &SesPage12{
		PageCode: ef.CreateElement("Uint8Element").(Uint8Field),
		Reserved_1_0: ef.CreateElement("Uint8Element").(Uint8Field),
		PageLength: ef.CreateElement("Uint16Element").(Uint16Field),
		CompletionCode: ef.CreateElement("Uint8Element").(Uint8Field),
		VpdStatusDescriptor: ef.CreateElement("VpdStatusDescriptorElement")}
	return elem
}

func CreateSesPage12() Element {
	return NewSesPage12()
}