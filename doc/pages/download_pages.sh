#! /bin/bash

#  Supported Diagnostic Pages [sdp] [0x0]
#  Configuration (SES) [cf] [0x1]
#  Enclosure Status/Control (SES) [ec,es] [0x2]
#  String In/Out (SES) [str] [0x4]
#  Threshold In/Out (SES) [th] [0x5]
#  Element Descriptor (SES) [ed] [0x7]
#  Additional Element Status (SES-2) [aes] [0xa]
#  Download Microcode (SES-2) [dm] [0xe]
#  <unknown> [0x10]
#  <unknown> [0x11]
#  <unknown> [0x12]
#  <unknown> [0x13]
#  <unknown> [0x14]
#  <unknown> [0x15]
#  <unknown> [0x16]

#declare -a pages
pages=(0x0 0x01 0x02 0x04 0x05 0x07 0x0a 0x0e 0x10 0x11 0x12 0x13 0x14 0x15 0x16)
dev=/dev/sg123

for i in ${pages[@]}
do
	echo "sg_ses --page=$i $dev > Page-$i.txt"
	sg_ses --page=$i $dev > Page-$i.txt
	echo "sg_ses --page=$i -rr $dev | od -t x1 -v > Page-$i-raw.txt"
	sg_ses --page=$i -rr $dev | od -t x1 -v > Page-$i-raw.txt
done

