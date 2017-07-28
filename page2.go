package sespage

// type StatusDiagnosticPage struct {
// 	PageCode Int8Element
// 	Unrecov Int8Element
// 	Crit Int8Element
// 	NonCrit Int8Element
// 	Info Int8Element
// 	Invop Int8Element
// 	PageLength Int16Element
// 	StatusDescriptorList []StatusDescriptor
// }

// func (p *StatusDiagnosticPage) HeadLength() int32 {
// 	return 4
// }

// func (p *StatusDiagnosticPage) DecodeHead(data []byte) error {
// 	p.PageCode.Decode(data[0:])
// 	p.Unrecov.Decode( []byte{data[1] & 0x01})
// 	p.Crit.Decode( []byte{(data[1] >> 1) & 0x01})
// 	p.NonCrit.Decode( []byte{(data[1] >> 2) & 0x01})
// 	p.Info.Decode( []byte{(data[1] >> 3) & 0x01})
// 	p.Invop.Decode( []byte{(data[1] >> 4) & 0x01})
// 	p.PageLength.Decode(data[2:4])

// 	return nil
// }

// func (p *StatusDiagnosticPage) Length() int32 {
// 	return int32(p.PageLength) + int32(4)
// }

// func (p *StatusDiagnosticPage) Decode(data []byte) error {
// 	p.DecodeHead(data)

// 	// offset := 8
// 	// for i, descriptor := range p.StatusDescriptorList {
// 	// 	p.StatusDescriptorList[i].Decode(data[offset:])
// 	// 	offset += p.StatusDescriptorList[i].Length()
// 	// }


// 	return nil
// }


// type StatusDescriptor struct {

// }