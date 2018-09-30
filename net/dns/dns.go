package dns

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type Header struct {
	ID      uint16
	Flag    uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (h *Header) SetFlag(qr, opcode, aa, tc, rd, ra, rcode uint16) {
	h.Flag = qr<<15 + opcode<<11 + aa<<10 + tc<<9 + rd<<8 + ra<<7 + rcode
}

const (
	eof     = byte(0)
	pointer = 0xC0 // byte(196)
)

// Question
// TYPE和CLASS两种定义方法
// 形式一
const (
	TypeA     = 1
	TypeNS    = 2
	TypeCNAME = 5
	TypeMX    = 15
	TypeTXT   = 16

	ClassIN = 1
	ClassCS = 2
	ClassCH = 3
	ClassHS = 4
)

// 形式二
const QTypeAny = 255
const QClassAny = 255

type Question struct {
	QName  []byte // byte切片
	QType  uint16
	QClass uint16
}

func (q *Question) SetQName(qname string) {
	var buf bytes.Buffer

	for _, n := range strings.Split(qname, ".") {
		binary.Write(&buf, binary.BigEndian, byte(len(n))) // 标签长度
		binary.Write(&buf, binary.BigEndian, []byte(n))    // 标签内容
		fmt.Println("pack qname query->", []byte(n), n)
	}
	binary.Write(&buf, binary.BigEndian, eof) // 以0x00结束

	q.QName = buf.Bytes()
}

// Answer

type resourceData interface {
	value() string
}

type Answer struct {
	Name     []byte
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    resourceData
}

func (r *Answer) SetRData(rdata, data []byte) error {
	var rd resourceData // 接口类型
	switch r.Type {
	case TypeA:
		rd = new(rdataA) // 普通类型转接口是隐式的
		if len(rdata) != 4 {
			return errors.New("invalid resource record data")
		}
		for i, d := range rdata {
			// 接口转普通类型需要使用断言：rd.(*rdataA)，即断言接口rd为rdataA指针类型
			binary.Read(bytes.NewBuffer([]byte{d}), binary.BigEndian, &rd.(*rdataA).addr[i])

			// 断言和强制转换是不同的，Go中的强制转换用于普通类型之间的转换
			// 当然，得是互相之间可以转换的类型
			// var a float64 = 1
			// b := int(a)
		}
		// ...
	}
	r.RData = rd

	return nil
}

type rdataA struct {
	addr [4]uint8
}

func (r *rdataA) value() string {
	return fmt.Sprintf("%d.%d.%d.%d", r.addr[0], r.addr[1], r.addr[2], r.addr[3])
}

// 叶子对象（基类）
type rdataDomain struct {
	name []byte
}

func (r *rdataDomain) value() string {
	var labels []string
	for i := 0; i < len(r.name)-1; {
		l := int(r.name[i]) //3
		labels = append(labels, string(r.name[i+1:i+l+1]))
		i += l + 1
	}

	return strings.Join(labels, ".")
}

// 组合对象
type rdataNS struct {
	rdataDomain // 匿名叶子，可以看做继承，直接拥有叶子对象的全部属性和函数
	// 也可以使用常规属性来描述
	// domain rdataDomain

	// 当然还可以定义当前结构体独有的其他属性
}

type rdataCNAME struct {
	rdataDomain
}

func Ask(server, qname string) ([]net.IP, error) {
	var names []net.IP
	//
	//reqData := Pack(TypeA, qname) // 封包
	//// 使用官方net包进行UDP连接
	//conn, err := net.Dial("udp", server+":53")
	//if err != nil {
	//	return nil, err
	//}
	//defer conn.Close() // 延迟处理
	//conn.SetDeadline(time.Now().Add(time.Second * 3))
	//// 发包
	//if i, err := conn.Write(reqData); err != nil || i <= 0 {
	//	return nil, err
	//}
	//buf := make([]byte, 4024)
	//n, _ := conn.Read(buf)
	//fmt.Println("read buf", buf[:n])
	buf := []byte{0, 1, 129, 128, 0, 1, 0, 3, 0, 0, 0, 0, 3, 119, 119, 119, 5, 98, 97, 105, 100, 117, 3, 99, 111, 109, 0, 0, 1, 0, 1, 192, 12, 0, 5, 0, 1, 0, 0, 1, 253, 0, 15, 3, 119, 119, 119, 1, 97, 6, 115, 104, 105, 102, 101, 110, 192, 22, 192, 43, 0, 1, 0, 1, 0, 0, 1, 10, 0, 4, 180, 97, 33, 107, 192, 43, 0, 1, 0, 1, 0, 0, 1, 10, 0, 4, 180, 97, 33, 108}
	answers, err := Unpack(bytes.NewReader(buf[:])) // 拆包
	if err != nil {
		return nil, err
	}
	for _, a := range answers {
		if a.Type != TypeA {
			continue
		}
		if ip := net.ParseIP(a.RData.value()); ip != nil {
			names = append(names, ip)
		}
	}

	return names, nil
}

func Pack(qtype uint16, qname string) []byte {
	// 封Header
	// 结构体的实例化
	header := Header{
		ID:      0x0001,
		QDCount: 1,
	}
	// 也可以使用另一种形式
	// header := new(Header)
	// header.ID = 0x0001
	// header.QDCount = 1
	header.SetFlag(0, 0, 0, 0, 1, 0, 0)

	// 封Question
	question := Question{
		QType:  qtype,
		QClass: ClassIN,
	}
	question.SetQName(qname)

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, header)
	fmt.Println("pack header->", buf.Bytes())
	fmt.Println("pack qname->", question.QName)
	binary.Write(&buf, binary.BigEndian, question.QName)
	binary.Write(&buf, binary.BigEndian, []uint16{question.QType, question.QClass})
	fmt.Println("pack question->", buf.Bytes()[12:])

	return buf.Bytes()
}

func Unpack(rd io.Reader) ([]*Answer, error) {
	var (
		reader = bufio.NewReader(rd)
		data   []byte // 应答数据包缓存
		buf    []byte // 临时缓存
		err    error
		n      int
	)

	// 拆Header
	header := new(Header)
	buf = make([]byte, 12)
	if n, err = reader.Read(buf); err != nil || n != 12 {
		return nil, err
	}
	binary.Read(bytes.NewReader(buf[:2]), binary.BigEndian, &header.ID)
	binary.Read(bytes.NewReader(buf[2:4]), binary.BigEndian, &header.Flag)
	binary.Read(bytes.NewReader(buf[4:6]), binary.BigEndian, &header.QDCount)
	binary.Read(bytes.NewReader(buf[6:8]), binary.BigEndian, &header.ANCount)
	binary.Read(bytes.NewReader(buf[8:10]), binary.BigEndian, &header.NSCount)
	binary.Read(bytes.NewReader(buf[10:12]), binary.BigEndian, &header.ARCount)
	data = append(data, buf...)
	fmt.Println("unpack header->", header)

	// 拆Question
	question := new(Question)
	if buf, err = reader.ReadBytes(eof); err != nil { // 域名以0x00结尾
		return nil, err
	}
	fmt.Println("unpack QName->", buf)
	data = append(data, buf...)
	question.QName = buf
	buf = make([]byte, 4)
	if n, err = reader.Read(buf); err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, nil
	}

	data = append(data, buf...)
	binary.Read(bytes.NewBuffer(buf[0:2]), binary.BigEndian, &question.QType)
	binary.Read(bytes.NewBuffer(buf[2:]), binary.BigEndian, &question.QClass)

	fmt.Println("unpack question->", question)

	// 拆Answer(s)
	answers := make([]*Answer, header.ANCount)
	buf, _ = reader.Peek(59)
	for i := 0; i < int(header.ANCount); i++ { // 根据Header中的ANCOUNT标识判断有几个Answer
		answer := new(Answer)
		// NAME
		var b byte
		var p uint16
		for {
			if b, err = reader.ReadByte(); err != nil {
				return nil, err
			}
			data = append(data, b)
			//fmt.Println("test1----->", data, b, pointer)
			if b&pointer == pointer { // pointer是一个值为0xC0的byte类型常量
				buf = []byte{b ^ pointer, 0}
				if b, err = reader.ReadByte(); err != nil {
					return nil, err
				}
				data = append(data, b)
				buf[1] = b
				binary.Read(bytes.NewBuffer(buf), binary.BigEndian, &p)
				fmt.Println("test2----->", buf, data, p)
				if buf = getRefData(data, p); len(buf) == 0 {
					return nil, errors.New("invalid answer packet")
				}
				answer.Name = append(answer.Name, buf...)
				break
			} else {
				answer.Name = append(answer.Name, b)
				if b == eof {
					break
				}
			}
		}

		// TYPE、CLASS、TLL、RDLENGTH等其他数据
		buf = make([]byte, 10)
		if n, err = reader.Read(buf); err != nil || n != 10 {
			return nil, err
		}
		binary.Read(bytes.NewReader(buf[:2]), binary.BigEndian, &answer.Type)
		binary.Read(bytes.NewReader(buf[2:4]), binary.BigEndian, &answer.Class)
		binary.Read(bytes.NewReader(buf[4:8]), binary.BigEndian, &answer.TTL)
		binary.Read(bytes.NewReader(buf[8:10]), binary.BigEndian, &answer.RDLength)
		data = append(data, buf...)
		//fmt.Println("unpack rdlength", int(answer.RDLength))
		// RDATA
		buf = make([]byte, int(answer.RDLength))
		if n, err = reader.Read(buf); err != nil || n < int(answer.RDLength) {
			return nil, err
		}
		data = append(data, buf...)
		// 调用之前定义的SetRData()函数处理不同类型的RDATA
		if err = answer.SetRData(buf, data); err != nil {
			return nil, err
		}

		answers[i] = answer
		fmt.Println("unpack answer->", answer)
	}

	// 拆Authority和Additional，如果有的话

	return answers, nil
}

func getRefData(data []byte, p uint16) []byte {
	var refData []byte
	fmt.Println("refdata", data)
	// 从初始偏移量开始对应答数据包缓存进行遍历
	for i := int(p); i < len(data); i++ {
		// fmt.Print(i, p, " ")
		// 读到新指针
		if b := data[i]; b&pointer == pointer {
			if i+1 >= len(data) {
				return []byte{}
			}
			// 更新偏移量，继续遍历
			binary.Read(bytes.NewBuffer([]byte{b ^ pointer, data[i+1]}), binary.BigEndian, &p)
			i = int(p - 1)
		} else {
			refData = append(refData, b)
			// 读到0x00即可结束
			if b == eof {
				break
			}
		}
	}

	return refData
}
