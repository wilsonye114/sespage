package sespage

import (
	"testing"
	"log"
	"encoding/json"
	"os/exec"
	"bytes"
)

func TestPage12MidplaneVirtual(t *testing.T) {
	data := []byte{
	0x12, 0x00, 0x00, 0x7b, 0x00, 0x0e, 0x00, 0x45, 0x72, 0x6f, 0x73, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x00, 0x00, 0x52, 0x30, 0x39, 0x33, 0x34, 0x2d, 0x47,
	0x30, 0x30, 0x30, 0x34, 0x2d, 0x30, 0x33, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x46, 0x46, 0x46,
	0x46, 0x46, 0x48, 0x48, 0x48, 0x48, 0x48, 0x48, 0x59, 0x59, 0x57, 0x57, 0x53, 0x53, 0x53, 0x53,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x45, 0x72, 0x6f, 0x73, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x52, 0x30, 0x39, 0x33, 0x34,
	0x2d, 0x58, 0x30, 0x30, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x20, 0x20, 0x30, 0x39, 0x38, 0x37, 0x36,
	0x35, 0x34, 0x33, 0x32, 0x31, 0x30, 0x39, 0x38, 0x37, 0x36, 0x35, 0x30, 0x30, 0x30, 0x31}

	log.Printf("Data Length: %d\n", len(data))
	elem := NewSesPage12()
	err := elem.Decode(data)
	if err != nil {
		panic(err)
	}
	js, _ := json.MarshalIndent(elem, "", "    ")
	log.Printf("%s\n", js)
	FieldCommonTest(data, elem)
}

func TestPage12MidplaneSg(t *testing.T) {
	var out bytes.Buffer

	log.Printf("sg_senddiag")
	cmdsend := exec.Command("sg_senddiag", "-p", "-r", "12,00,00,03,00,0e,00", "/dev/sg63")
	err := cmdsend.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sg_ses")
	cmd := exec.Command("sg_ses", "--page=0x12", "-rr", "/dev/sg63")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sg123 - SES Page1 Raw Data:\n%q\n", out.String())
	raw := out.Bytes()
	page := NewSesPage12()
	page.Decode(raw)
	js, _ := json.MarshalIndent(page, "", "    ")
	log.Printf("Json:\n%s\n", js)
}

func TestPage12CanisterVirtual(t *testing.T) {
	data := []byte{
		0x12, 0x00, 0x00, 0x3f, 0x00, 0x07, 0x00, 0x45, 0x72, 0x6f, 0x73, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x52, 0x30, 0x39, 0x33, 0x34, 0x47, 0x30, 0x30, 0x30,
		0x32, 0x30, 0x33, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x31, 0x48, 0x38, 0x56, 0x30,
		0x31, 0x34, 0x30, 0x33, 0x30, 0x30, 0x30, 0x30, 0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x30, 0x31, 0x20}
    
	log.Printf("Data Length: %d\n", len(data))
	elem := NewSesPage12()
	err := elem.Decode(data)
	if err != nil {
		panic(err)
	}
	js, _ := json.MarshalIndent(elem, "", "    ")
	log.Printf("%s\n", js)
	FieldCommonTest(data, elem)
}

func TestPage12CanisterSg(t *testing.T) {
	var out bytes.Buffer

	log.Printf("sg_senddiag")
	cmdsend := exec.Command("sg_senddiag", "-p", "-r", "12,00,00,03,00,07,00", "/dev/sg63")
	err := cmdsend.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sg_ses")
	cmd := exec.Command("sg_ses", "--page=0x12", "-rr", "/dev/sg63")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sg123 - SES Page1 Raw Data:\n%q\n", out.String())
	raw := out.Bytes()
	page := NewSesPage12()
	page.Decode(raw)
	js, _ := json.MarshalIndent(page, "", "    ")
	log.Printf("Json:\n%s\n", js)
}