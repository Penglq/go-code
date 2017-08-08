# command-line-arguments
"".main t=1 size=158 args=0x0 locals=0x50
	0x0000 00000 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	TEXT	"".main(SB), $80-0
	0x0000 00000 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	JLS	148
	0x0013 00019 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	SUBQ	$80, SP
	0x0017 00023 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	MOVQ	BP, 72(SP)
	0x001c 00028 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	LEAQ	72(SP), BP
	0x0021 00033 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	FUNCDATA	$1, gclocals·e226d4ae4a7cad8835311c6a4683c14f(SB)
	0x0021 00033 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	$16, ""..autotmp_1+48(SP)
	0x002a 00042 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	$0, ""..autotmp_0+56(SP)
	0x0033 00051 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	$0, ""..autotmp_0+64(SP)
	0x003c 00060 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	LEAQ	type.int(SB), AX
	0x0043 00067 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	AX, (SP)
	0x0047 00071 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	LEAQ	""..autotmp_1+48(SP), AX
	0x004c 00076 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	AX, 8(SP)
	0x0051 00081 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	PCDATA	$0, $1
	0x0051 00081 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	CALL	runtime.convT2E(SB)
	0x0056 00086 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	24(SP), AX
	0x005b 00091 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	16(SP), CX
	0x0060 00096 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	CX, ""..autotmp_0+56(SP)
	0x0065 00101 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	AX, ""..autotmp_0+64(SP)
	0x006a 00106 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	LEAQ	""..autotmp_0+56(SP), AX
	0x006f 00111 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	AX, (SP)
	0x0073 00115 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	$1, 8(SP)
	0x007c 00124 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	MOVQ	$1, 16(SP)
	0x0085 00133 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	PCDATA	$0, $1
	0x0085 00133 (/Users/luqiangpeng/Go/src/go-code/rune.go:9)	CALL	fmt.Println(SB)
	0x008a 00138 (/Users/luqiangpeng/Go/src/go-code/rune.go:10)	MOVQ	72(SP), BP
	0x008f 00143 (/Users/luqiangpeng/Go/src/go-code/rune.go:10)	ADDQ	$80, SP
	0x0093 00147 (/Users/luqiangpeng/Go/src/go-code/rune.go:10)	RET
	0x0094 00148 (/Users/luqiangpeng/Go/src/go-code/rune.go:10)	NOP
	0x0094 00148 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	PCDATA	$0, $-1
	0x0094 00148 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0099 00153 (/Users/luqiangpeng/Go/src/go-code/rune.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 81  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 50 48 89 6c 24 48 48 8d 6c 24  ...H..PH.l$HH.l$
	0x0020 48 48 c7 44 24 30 10 00 00 00 48 c7 44 24 38 00  HH.D$0....H.D$8.
	0x0030 00 00 00 48 c7 44 24 40 00 00 00 00 48 8d 05 00  ...H.D$@....H...
	0x0040 00 00 00 48 89 04 24 48 8d 44 24 30 48 89 44 24  ...H..$H.D$0H.D$
	0x0050 08 e8 00 00 00 00 48 8b 44 24 18 48 8b 4c 24 10  ......H.D$.H.L$.
	0x0060 48 89 4c 24 38 48 89 44 24 40 48 8d 44 24 38 48  H.L$8H.D$@H.D$8H
	0x0070 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7 44 24  ..$H.D$.....H.D$
	0x0080 10 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 48 48  ..........H.l$HH
	0x0090 83 c4 50 c3 e8 00 00 00 00 e9 62 ff ff ff        ..P.......b...
	rel 5+4 t=16 TLS+0
	rel 63+4 t=15 type.int+0
	rel 82+4 t=8 runtime.convT2E+0
	rel 134+4 t=8 fmt.Println+0
	rel 149+4 t=8 runtime.morestack_noctxt+0
"".init t=1 size=91 args=0x0 locals=0x8
	0x0000 00000 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	TEXT	"".init(SB), $8-0
	0x0000 00000 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	JLS	84
	0x000f 00015 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	SUBQ	$8, SP
	0x0013 00019 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVQ	BP, (SP)
	0x0017 00023 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	LEAQ	(SP), BP
	0x001b 00027 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	CMPB	AL, $1
	0x0024 00036 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	JLS	$0, 47
	0x0026 00038 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVQ	(SP), BP
	0x002a 00042 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	ADDQ	$8, SP
	0x002e 00046 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	RET
	0x002f 00047 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	JNE	$0, 56
	0x0031 00049 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	PCDATA	$0, $0
	0x0031 00049 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	CALL	runtime.throwinit(SB)
	0x0036 00054 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	UNDEF
	0x0038 00056 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVB	$1, "".initdone·(SB)
	0x003f 00063 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	PCDATA	$0, $0
	0x003f 00063 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	CALL	fmt.init(SB)
	0x0044 00068 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVB	$2, "".initdone·(SB)
	0x004b 00075 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	MOVQ	(SP), BP
	0x004f 00079 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	ADDQ	$8, SP
	0x0053 00083 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	RET
	0x0054 00084 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	NOP
	0x0054 00084 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	PCDATA	$0, $-1
	0x0054 00084 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x0059 00089 (/Users/luqiangpeng/Go/src/go-code/rune.go:11)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 45 48  eH..%....H;a.vEH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8  ................
	0x0040 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24 48  ...........H.,$H
	0x0050 83 c4 08 c3 e8 00 00 00 00 eb a5                 ...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 64+4 t=8 fmt.init+0
	rel 70+4 t=15 "".initdone·+-1
	rel 85+4 t=8 runtime.morestack_noctxt+0
go.string."我是一颗abcd" t=8 dupok size=16
	0x0000 e6 88 91 e6 98 af e4 b8 80 e9 a2 97 61 62 63 64  ............abcd
gclocals·e226d4ae4a7cad8835311c6a4683c14f t=8 dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 03                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 t=8 dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
go.info."".main t=45 size=27
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 00                 ...........
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+158
gclocals·33cdeccccebe80329f1fdbee7f5874cb t=8 dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
go.info."".init t=45 size=27
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 00                 ...........
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+91
"".initdone· t=32 size=1
runtime.gcbits.01 t=8 dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}. t=8 dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}.+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 t=8 dupok size=1
	0x0000 03                                               .
type.interface {} t=8 dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}.+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}. t=8 dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}.+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}.+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}. t=8 dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}.+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} t=8 dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}.+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..importpath.fmt. t=8 dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
