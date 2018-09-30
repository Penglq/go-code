package dns

import (
	"testing"
	"fmt"
	"os"
	"net"
	"sync"
)

func TestAsk(t *testing.T) {
	//ips,err := Ask("114.114.114.114", "www.baidu.com")
	ips, err := Ask("114.114.114.114", "codespub.com")
	fmt.Println(ips, err)
}

func getDomain() chan string {
	domainList := []string{"www.baidu.com", "codespub.com"}
	ch := make(chan string, 10) // 创建一个通道，10个缓存空间，用于goroutine间的数据传输

	go func() { // 使用go关键字，创建一个新的goroutine来执行该匿名函数，即多线程
		for _, domain := range domainList {
			ch <- domain
		}
		close(ch) // 关闭通道
	}()

	return ch
}

func TestMoreAsk(t *testing.T) {
	var wg sync.WaitGroup // 官方sync包中对象，可用于阻塞主线程，等待所有goroutine执行结束
	ch := getDomain()     // 该方法中的逻辑是异步执行的

	for c := range ch { // 遍历直至通道关闭
		wg.Add(1)
		go func(domain string) { // 为每个域名单独创建一个goroutine处理
			defer wg.Done() // 相当于wg.Add(-1)

			if ips, err := Ask("8.8.8.8", domain); err == nil {
				fmt.Println(ips)
			}
		}(c)
	}

	wg.Wait()
}

func TestDns1(t *testing.T) {
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
	fmt.Println(n, buf[0:n])
	fmt.Println(out.Unpack(buf[0:n]))
	os.Exit(0)
}
