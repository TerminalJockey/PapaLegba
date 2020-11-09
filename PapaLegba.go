package PapaLegba

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
)

type Proc struct {
	Process *exec.Cmd
	stdin   io.WriteCloser
}

type RemoteProc struct {
	Process net.Conn
}

//creates connection object for further interaction
func ConnectRemProc(address string) (Proc RemoteProc) {
	outProc := RemoteProc{}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
	}
	outProc.Process = conn
	return outProc
}

//write argument to connection object
func SendRemArg(proc RemoteProc, payload string) {
	proc.Process.Write([]byte(payload))
}

//get response from connection object
func GetRemResp(proc RemoteProc) {
	resp, err := bufio.NewReader(proc.Process).ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(resp))
}

//converts string to hex
func StrToHex(in string) (out string) {
	hexout, err := hex.DecodeString(in)
	if err != nil {
		log.Println(err)
	}
	return string(hexout)
}

//start local process for further interaction
func StartProc(path string) (proc Proc) {
	outProc := Proc{}
	outProc.Process = exec.Command(path)
	var err error
	outProc.stdin, err = outProc.Process.StdinPipe()
	if err != nil {
		log.Println(err)
	}
	return outProc
}

//send argument to local process
func SendArg(proc Proc, payload string) {
	io.WriteString(proc.stdin, payload)
}

//get response from local process
func GetResp(proc Proc) {
	out, err := proc.Process.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
}

//strip newlines and \x from shellcode, useful in formatting msfvenom payloads
func FormatShellcode(in string) (out string) {
	out = strings.Replace(in, `\x`, "", -1)
	out = strings.Replace(out, "\n", "", -1)
	return out
}

//inverts byte order
func InvertEndian(in string) (out string) {
	runes := []byte(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
