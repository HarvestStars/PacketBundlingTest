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
	if len(t.packetBuffer) == 0 {
		tmpPacket := <-t.recvChan
		t.packetBuffer = append(t.packetBuffer, tmpPacket...)
		fmt.Println("Read, the packetBuffer read:", string(t.packetBuffer))
	}

	n := 0
	if len(t.packetBuffer) >= 4096 {
		fmt.Println("Read, buffer is 4096")
		n = 4096
		copy((*buf)[offset:], t.packetBuffer[:4096])
		t.packetBuffer = t.packetBuffer[4096:]

	} else {
		fmt.Println("Read, buffer less than 4096")
		n = len(t.packetBuffer)
		copy((*buf)[offset:], t.packetBuffer)
		t.packetBuffer = []byte{}

	}
	return n, nil
}

func (t *transportBasicIO) PacketBundling() ([][]byte, error) {
	out := &bytes.Buffer{}
	buf := make([]byte, 4096*2)
	allMsgGroup := [][]byte{}

	offset := 0
	for {
		fmt.Println("PacketBundling, buf before read is", string(buf))

		n, err := t.Read(&buf, offset)
		if err != nil {
			break
		}
		fmt.Println("buf outside is", string(buf), "length is", len(buf), "n is", n)
		offset = t.RecursiveSplit(&buf, n, offset, []byte(sep), &allMsgGroup, out)
		if offset == 0 {
			return allMsgGroup, nil
		} else {
			fmt.Println("PacketBundling, continue read again, offset is", offset)
			continue
		}

	}

	return nil, fmt.Errorf("PacketBundling failed")
}

func (t *transportBasicIO) RecursiveSplit(buf *[]byte, n, offset int, sep []byte, allMsg *[][]byte, out *bytes.Buffer) int {
	loc, _ := t.Splitter((*buf)[0:n+offset], sep)
	if loc == -1 {
		if len((*buf)[0:n+offset]) > len(sep) {
			out.Write((*buf)[0 : n+offset-len(sep)])
			contextLeft := (*buf)[n+offset-len(sep) : n+offset]
			*buf = make([]byte, 4096*2)
			copy(*buf, contextLeft)
			return len(contextLeft)
		} else {
			contextLeft := (*buf)[0 : n+offset]
			return len(contextLeft)
		}
	} else {

		out.Write((*buf)[0:loc])
		*allMsg = append(*allMsg, append([]byte(nil), out.Bytes()...)) // recreate a new underlying array, not the slice
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

func (t *transportBasicIO) WaitForFunc() ([][]byte, error) {
	var out bytes.Buffer
	buf := make([]byte, 4096*2)
	allMsgGroup := [][]byte{}

	pos := 0
	index := 0
	for {
		index++
		fmt.Println("buf before read is", string(buf))
		n, err := t.Read(&buf, pos)
		if err != nil {
			break
		}

		if n > 0 {
			// loop to spilt msg bundling
			for {
				netconfEnd, err := t.Splitter(buf[0:pos+n], []byte(payload.MsgSeperator))
				if err != nil {
					return nil, err
				}

				if netconfEnd > -1 {
					// there is 1 msg at least
					out.Write(buf[0:netconfEnd])

					fmt.Printf("Index %v, buf[0:netconfEnd] is: %v", index, string(buf))
					fmt.Println("")
					fmt.Printf("Index %v, out bytes, out is: %v", index, string(out.Bytes()))
					fmt.Println("")
					allMsgGroup = append(allMsgGroup, append([]byte(nil), out.Bytes()...))

					for _, v := range allMsgGroup {
						fmt.Println("")
						fmt.Printf("Index %v, WaitForFunc, one msg is: %v", index, string(v))
					}

					out.Reset()

					if len(buf[netconfEnd:pos+n]) == len([]byte(payload.MsgSeperator)) {
						// this is the only msg from read
						return allMsgGroup, nil
					} else {
						// there must be other msg following behind
						msgFollowing := buf[netconfEnd+len([]byte(payload.MsgSeperator)) : pos+n]
						n = len(msgFollowing)
						pos = 0
						buf = make([]byte, 4096*2)
						copy(buf, msgFollowing)
						continue
					}
				} else {
					// not found end
					// but maybe the netconf ned is accross two msg
					// so we should leave the last 6 bytes
					contextLeft := []byte{}
					if pos+n > len([]byte(payload.MsgSeperator)) {
						out.Write(buf[0 : pos+n-len([]byte(payload.MsgSeperator))])
						fmt.Println("")
						fmt.Printf("Index %v, netconfEnd is -1 out bytes, out is: %v", index, string(out.Bytes()))
						fmt.Println("")
						// reminder is a part of the netconf end
						contextLeft = buf[pos+n-len([]byte(payload.MsgSeperator)) : pos+n]
					} else {
						contextLeft = buf[0 : pos+n]
					}
					pos = len(contextLeft)
					buf = make([]byte, 4096*2)
					copy(buf, contextLeft)
					break
				}
			}
		}
	}

	return nil, fmt.Errorf("WaitForFunc failed")
}

func main() {
	tio := &transportBasicIO{
		recvChan:     make(chan []byte),
		packetBuffer: []byte{},
	}
	go tio.SendBundling([]string{payload.MeasurementResultStats, payload.MeasurementResultStats, payload.MeasurementResultStats})

	allMsgBytes, _ := tio.PacketBundling() // version recursive
	//allMsgBytes, _ := tio.WaitForFunc() // version iteration
	for index, v := range allMsgBytes {
		msgStr := string(v)
		fmt.Println("")
		fmt.Printf("Msg %v is %v", index, msgStr)
		fmt.Println("")
	}
}
