package main

import (
	"fmt"

	"github.com/google/gopacket"

	"github.com/google/gopacket/pcap"
)

func (box packetinfo) capture() {
	handle, _ := pcap.OpenLive(box.device, box.snaplen, box.promicious, box.timeout)
	defer handle.Close()

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		fmt.Println(packet)
	}
}
