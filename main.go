package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/HarvestStars/PacketBundlingTest/payload"
)

var sep string = "]]>]]>"

type transportBasicIO struct {
	recvChan     chan []byte
	packetBuffer []byte
}

func (t *transportBasicIO) SendBundling(msg []string) {
	allMsg := ""
	previousMsg := ""
	for _, v := range msg {
		if previousMsg == "" {
			// first one
			previousMsg = v
			t.recvChan <- []byte(v)
		} else {
			// second one
			time.Sleep(2 * time.Second)
			t.recvChan <- []byte(sep + v[0:8])
			time.Sleep(2 * time.Second)
			t.recvChan <- []byte(v[8:])
		}
	}
	t.recvChan <- []byte(sep)

	fmt.Println("All msg ready to send:", allMsg)
}

func (t *transportBasicIO) Read(buf *[]byte, offset int) (int, error) {

	fmt.Println("buf inside the read is", string(*buf))

	if len(t.packetBuffer) == 0 {
		fmt.Println("init the packetBuffer")
		tmpPacket := <-t.recvChan
		t.packetBuffer = append(t.packetBuffer, tmpPacket...)
		fmt.Println("the packetBuffer read:", string(t.packetBuffer))
	}

	n := 0
	if len(t.packetBuffer) >= 4096 {
		fmt.Println("buffer is 4096")
		n = 4096
		copy((*buf)[offset:], t.packetBuffer[:4096])
		t.packetBuffer = t.packetBuffer[4096:]

	} else {
		fmt.Println("buffer less than 4096")
		n = len(t.packetBuffer)
		copy((*buf)[offset:], t.packetBuffer)
		fmt.Println("the buf is:", string(*buf))
		t.packetBuffer = []byte{}

	}
	fmt.Println("buf inside the read after copy", string(*buf), "copy start at the offset", offset)
	return n, nil
}

func (t *transportBasicIO) PacketBundling() ([][]byte, error) {
	out := &bytes.Buffer{}
	buf := make([]byte, 4096*2)
	allMsgGroup := [][]byte{}

	offset := 0
	for {
		fmt.Println("buf before read is", string(buf))
		n, err := t.Read(&buf, offset)
		if err != nil {
			break
		}
		fmt.Println("buf outside is", string(buf), "length is", len(buf), "n is", n)
		offset = t.RecursiveSplit(&buf, n, offset, []byte(sep), &allMsgGroup, out)
		if offset == 0 {
			return allMsgGroup, nil
		} else {
			continue
		}

	}

	return nil, fmt.Errorf("PacketBundling failed")
}

func (t *transportBasicIO) RecursiveSplit(buf *[]byte, n, offset int, sep []byte, allMsg *[][]byte, out *bytes.Buffer) int {
	loc, _ := t.Splitter((*buf)[0:n+offset], sep)
	if loc == -1 {
		if len((*buf)[0:n]) > len(sep) {
			out.Write((*buf)[0 : n+offset-len(sep)])
			context := (*buf)[n+offset-len(sep):]
			*buf = make([]byte, 4096*2)
			copy(*buf, context)
		}
		return len(sep)

	} else {

		out.Write((*buf)[0:loc])
		*allMsg = append(*allMsg, out.Bytes()) // recreate a new underlying array, not the slice
		out.Reset()

		if len((*buf)[loc:n+offset]) == len(sep) {
			// loc is the last one
			*buf = make([]byte, 4096*2)
			return 0

		} else {
			// another msg following the loc
			context := (*buf)[loc+len(sep) : n+offset]
			*buf = make([]byte, 4096*2)
			copy(*buf, context)
			offset := t.RecursiveSplit(buf, n+offset-loc-len(sep), 0, sep, allMsg, out)
			return offset
		}
	}
}

func (t *transportBasicIO) Splitter(buf, sep []byte) (int, error) {
	tailerNull := -1
	netconfTailer := bytes.Index(buf, sep)
	if netconfTailer > -1 {
		return netconfTailer, nil
	} else {
		return tailerNull, nil
	}
}

func main() {
	tio := &transportBasicIO{
		recvChan:     make(chan []byte),
		packetBuffer: []byte{},
	}
	go tio.SendBundling([]string{payload.HelloRsp, payload.A3TopoEventNtf, payload.OperTrxStateNtf})

	allMsgBytes, _ := tio.PacketBundling()
	for index, v := range allMsgBytes {
		msgStr := string(v)
		fmt.Printf("Msg %v is %v", index, msgStr)
	}
}
