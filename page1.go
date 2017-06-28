package sespage



// type ConfigurationDiagnosticPage struct {
// 	HeadLength int32
// 	Length int32
// 	HeadElements []Element
// 	PayloadElements []Element
// }

// func (p *ConfigurationDiagnosticPage) DecodeHead(buf []byte) (map[string]interface{}, error) {
// 	head := make(map[string]interface{})
// 	headbuf := buf[:p.HeadLength]
// 	offset := 0
// 	for _, e := range p.HeadElements {
// 		data, length := e.Decode(headbuf[offset:])
// 		offset += length
// 		head[e.Name()] = data
// 	}
// 	return head, nil
// }

// func (p *ConfigurationDiagnosticPage) Decode(buf []byte) (map[string]interface{}, error) {
// 	pagemap, _ := p.DecodeHead(buf)
// 	p.Length = pagemap["pageLength"].(int32) + 4
// 	offset := int(p.HeadLength)
// 	for _, e := range p.PayloadElements {
// 		data, length := e.Decode(buf[offset:])
// 		offset += length
// 		pagemap[e.Name()] = data
// 	}
// 	return pagemap, nil
// }




// func NewConfigurationDiagnosticPage() *ConfigurationDiagnosticPage {
// 	page := ConfigurationDiagnosticPage{
// 		HeadLength:8,
// 		HeadElements:[]Element{
// 			&IntByteElement{name:"pageCode"},
// 			&IntByteElement{name:"numberOfSecondarySubenclosures"},
// 			&Int2BytesElement{name:"pageLength"},
// 			&Int4BytesElement{name:"generationCode"}},
// 		PayloadElements:[]Element{
// 			&IntByteElement{name:"reserved"},
// 			&IntByteElement{name:"subenclosureIdentifier"},
// 			&IntByteElement{name:"numberOfTypeDescriptorHeaders"},
// 			&IntByteElement{name:"enclosureDescriptorLength"}}}
// 	return &page
// }
