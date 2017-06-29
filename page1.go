package sespage

import (
	// "fmt"
)


type ConfigurationDiagnosticPage struct {
	PageCode Int8Element
	NumberOfSecondarySubenclosures Int8Element
	PageLength Int16Element
	GenerationCode Int32Element
	EnclosureDescriptorList []EnclosureDescriptor
	TypeDescriptorHeaderList []TypeDescriptorHeader
	TypeDescriptorTextList []StringElement
}

func (p *ConfigurationDiagnosticPage) HeadLength() int32 {
	return 4
}

func (p *ConfigurationDiagnosticPage) DecodeHead(data []byte) error {
	p.PageCode.Decode(data[0:])
	p.NumberOfSecondarySubenclosures.Decode(data[1:])
	p.PageLength.Decode(data[2:4])

	return nil
}

func (p *ConfigurationDiagnosticPage) Length() int32 {
	return int32(p.PageLength) + int32(4)
}


func (p *ConfigurationDiagnosticPage) Decode(data []byte) error {
	p.DecodeHead(data)

	p.GenerationCode.Decode(data[4:8])

	offset := int32(8)
	numSubenclosures := p.NumberOfSecondarySubenclosures+1
	p.EnclosureDescriptorList = make([]EnclosureDescriptor, numSubenclosures, numSubenclosures)
	for i, _ := range p.EnclosureDescriptorList {
		p.EnclosureDescriptorList[i].Decode(data[offset:])
		offset += p.EnclosureDescriptorList[i].Length()
	}

	numTypeDiscHeader := 0
	for _, descriptor := range p.EnclosureDescriptorList {
		numTypeDiscHeader += int(descriptor.NumberOfTypeDescriptorHeaders)
	}
	p.TypeDescriptorHeaderList = make([]TypeDescriptorHeader, numTypeDiscHeader, numTypeDiscHeader)
	for i, _ := range p.TypeDescriptorHeaderList {
		p.TypeDescriptorHeaderList[i].Decode(data[offset:])
		offset += p.TypeDescriptorHeaderList[i].Length()
	}

	p.TypeDescriptorTextList = make([]StringElement, 0, numTypeDiscHeader)
	for i, header := range p.TypeDescriptorHeaderList {
		if header.TypeDescriptorTextLength != 0 {
			var strelem StringElement
			strelem.Decode(data[offset:offset+int32(header.TypeDescriptorTextLength)])
			p.TypeDescriptorTextList = append(p.TypeDescriptorTextList, strelem)
			p.TypeDescriptorHeaderList[i].Link_TypeDescriptorText = strelem
			offset += int32(header.TypeDescriptorTextLength)
		}
	}

	// p.TypeDescriptorTextList = make([]StringElement, 0, numTypeDiscHeader)
	// for _, header := range p.TypeDescriptorHeaderList {
	// 	if header.TypeDescriptorTextLength != 0 {
	// 		var strelem StringElement
	// 		strelem.Decode(data[offset:offset+int32(header.TypeDescriptorTextLength)])
	// 		p.TypeDescriptorTextList = append(p.TypeDescriptorTextList, strelem)
	// 		offset += int32(header.TypeDescriptorTextLength)
	// 	}
	// }

	return nil
}

type EnclosureDescriptor struct {
	NumberOfEnclosureServicesProcesses Int8Element
	RelativeEnclosureServicesProcessIdentifier Int8Element
	SubenclosureIdentifier Int8Element
	NumberOfTypeDescriptorHeaders Int8Element
	EnclosureDescriptorLength Int8Element
	EnclosureLogicalIdentifier HexStringElement
	EnclosureVendorIdentification StringElement
	ProductIdentification StringElement
	ProductRevisionLevel StringElement
	VendorSpecificEnclosureInformation BytesElement
}

func (d *EnclosureDescriptor) HeadLength() int32 {
	return 4
}

func (d *EnclosureDescriptor) DecodeHead(data []byte) error {
	d.NumberOfEnclosureServicesProcesses.Decode([]byte{data[0] & 0x07})
	d.RelativeEnclosureServicesProcessIdentifier.Decode([]byte{(data[0] >> 4) & 0x07})
	d.SubenclosureIdentifier.Decode(data[1:])
	d.NumberOfTypeDescriptorHeaders.Decode(data[2:])
	d.EnclosureDescriptorLength.Decode(data[3:])
	return nil
}

func (d *EnclosureDescriptor) Length() int32 {
	return int32(d.EnclosureDescriptorLength) + 4
}

func (d *EnclosureDescriptor) Decode(data []byte) error {
	d.DecodeHead(data)
	d.EnclosureLogicalIdentifier.Decode(data[4:12])
	d.EnclosureVendorIdentification.Decode(data[12:20])
	d.ProductIdentification.Decode(data[20:36])
	d.ProductRevisionLevel.Decode(data[36:40])
	d.VendorSpecificEnclosureInformation.Decode(data[40:d.Length()])

	return nil
}

type TypeDescriptorHeader struct {
	ElementType ETypeElement
	NumberOfPossibleElements Int8Element
	SubenclosureIdentifier Int8Element
	TypeDescriptorTextLength Int8Element
	Link_TypeDescriptorText StringElement
}

func (d *TypeDescriptorHeader) Length() int32 {
	return 4
}

func (d *TypeDescriptorHeader) Decode(data []byte) error {
	d.ElementType.Decode(data[0:])
	d.NumberOfPossibleElements.Decode(data[1:])
	d.SubenclosureIdentifier.Decode(data[2:])
	d.TypeDescriptorTextLength.Decode(data[3:])
	return nil
}



