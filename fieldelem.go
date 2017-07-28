package sespage

// import (
// 	"fmt"
// )

// type ETypeElement string

// func (e *ETypeElement) Decode(data []byte) error {
// 	if len(data) < 1 {
// 		return fmt.Errorf("Expect 1 bytes data, got %d", len(data))
// 	}
// 	t := uint8(data[0])
// 	switch {
// 		case t == 0x00: 
// 			*e = "Unspecified"
// 		case t == 0x01: 
// 			*e = "Device Slot"
// 		case t == 0x02: 
// 			*e = "Power Supply"
// 		case t == 0x03: 
// 			*e = "Cooling"
// 		case t == 0x04: 
// 			*e = "Temperature Sensor"
// 		case t == 0x05: 
// 			*e = "Door Lock"
// 		case t == 0x06: 
// 			*e = "Audible Alarm"
// 		case t == 0x07: 
// 			*e = "Enclosure Services Controller Electronics"
// 		case t == 0x08: 
// 			*e = "SCC Controller Electronics"
// 		case t == 0x09: 
// 			*e = "Nonvolatile Cache"
// 		case t == 0x0A: 
// 			*e = "Invalid Operation Reason c"
// 		case t == 0x0B: 
// 			*e = "Uninterruptible Power Supply"
// 		case t == 0x0C: 
// 			*e = "Display"
// 		case t == 0x0D: 
// 			*e = "Key Pad Entry"
// 		case t == 0x0E: 
// 			*e = "Enclosure"
// 		case t == 0x0F: 
// 			*e = "SCSI Port/Transceiver"
// 		case t == 0x10: 
// 			*e = "Language"
// 		case t == 0x11: 
// 			*e = "Communication Port"
// 		case t == 0x12: 
// 			*e = "Voltage Sensor"
// 		case t == 0x13: 
// 			*e = "Current Sensor"
// 		case t == 0x14: 
// 			*e = "SCSI Target Port"
// 		case t == 0x15: 
// 			*e = "SCSI Initiator Port"
// 		case t == 0x16: 
// 			*e = "Simple Subenclosure"
// 		case t == 0x17: 
// 			*e = "Array Device Slot"
// 		case t == 0x18: 
// 			*e = "SAS Expander"
// 		case t == 0x19: 
// 			*e = "SAS Connector"
// 		case t >= 0x1A && t <= 0x7F:
// 			*e = "Reserved"
// 		case t >= 0x80 && t <= 0xFF:
// 			*e = "Vendor specific"
// 	}
// 	return nil
// }