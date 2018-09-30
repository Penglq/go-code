package dns

import (
	"testing"
	"bytes"
	"fmt"
	"os"
	"net"
)

func TestPack(t *testing.T) {
	b := Pack(TypeA, "codespub.com")
	fmt.Println("pack->", b)
	Unpack(bytes.NewReader(b))
}

func TestAsk(t *testing.T) {
	ips,err := Ask("114.114.114.114", "www.baidu.com")
	//ips,err := Ask("114.114.114.114", "codespub.com")
	fmt.Println(ips,err)
}

func TestDns1(t *testing.T)  {
	service := "114.114.114.114:53"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	question := dnsQuestion{"www.baidu.com", dnsTypeA, dnsClassINET}
	//question := dnsQuestion{"www.cyeam.com", dnsTypeA, dnsClassINET}
	// question1 := dnsQuestion{"blog.cyeam.com", dnsTypeALL, dnsClassINET}
	out := DnsMsg{}
	out.Id = 2015
	out.Bits |= _RD
	out.Questions = append(out.Questions, question)
	// out.Questions = append(out.Questions, question1)
	// fmt.Println(append(out.Pack(), []byte{1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 5, 99, 121, 101, 97, 109, 3, 99, 111, 109, 0, 0, 1, 0, 1}...))
	fmt.Println(out.Pack())
	_, err = conn.Write(out.Pack())
	// _, err = conn.Write([]byte{11, 117, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 119, 119, 119, 5, 99, 121, 101, 97, 109, 3, 99, 111, 109, 0, 0, 1, 0, 1})
	checkError(err)

	buf := []byte{}
	buf = make([]byte, 512)
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(n,buf[0:n])
	fmt.Println(out.Unpack(buf[0:n]))
	os.Exit(0)
}