"".test STEXT size=179 args=0x10 locals=0x28
	0x0000 00000 (main.go:5)	TEXT	"".test(SB), $40-16
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:5)	JLS	169
	0x0013 00019 (main.go:5)	SUBQ	$40, SP
	0x0017 00023 (main.go:5)	MOVQ	BP, 32(SP)
	0x001c 00028 (main.go:5)	LEAQ	32(SP), BP
	0x0021 00033 (main.go:5)	FUNCDATA	$0, gclocals·d8b28f51bb91e05d264803f0f600a200(SB)
	0x0021 00033 (main.go:5)	FUNCDATA	$1, gclocals·9a26515dfaeddd28bcbc040f1199f48d(SB)
	0x0021 00033 (main.go:5)	LEAQ	type.int(SB), AX|
	0x0028 00040 (main.go:5)	MOVQ	AX, (SP)
	0x002c 00044 (main.go:5)	PCDATA	$0, $0
	0x002c 00044 (main.go:5)	CALL	runtime.newobject(SB)
	0x0031 00049 (main.go:5)	MOVQ	8(SP), AX
	0x0036 00054 (main.go:5)	MOVQ	AX, "".&a+24(SP)
	0x003b 00059 (main.go:5)	MOVQ	"".a+48(SP), CX
	0x0040 00064 (main.go:5)	MOVQ	CX, (AX)
	0x0043 00067 (main.go:5)	LEAQ	type.struct { F uintptr; "".a *int }(SB), CX
	0x004a 00074 (main.go:6)	MOVQ	CX, (SP)
	0x004e 00078 (main.go:6)	PCDATA	$0, $1
	0x004e 00078 (main.go:6)	CALL	runtime.newobject(SB)
	0x0053 00083 (main.go:6)	MOVQ	8(SP), AX
	0x0058 00088 (main.go:6)	LEAQ	"".test.func1(SB), CX
	0x005f 00095 (main.go:6)	MOVQ	CX, (AX)
	0x0062 00098 (main.go:6)	TESTB	AL, (AX)
	0x0064 00100 (main.go:6)	MOVL	runtime.writeBarrier(SB), CX
	0x006a 00106 (main.go:6)	LEAQ	8(AX), DX
	0x006e 00110 (main.go:6)	TESTL	CX, CX
	0x0070 00112 (main.go:6)	JNE	138
	0x0072 00114 (main.go:6)	MOVQ	"".&a+24(SP), CX
	0x0077 00119 (main.go:6)	MOVQ	CX, 8(AX)
	0x007b 00123 (main.go:6)	MOVQ	AX, "".~r1+56(SP)
	0x0080 00128 (main.go:6)	MOVQ	32(SP), BP
	0x0085 00133 (main.go:6)	ADDQ	$40, SP
	0x0089 00137 (main.go:6)	RET
	0x008a 00138 (main.go:6)	MOVQ	AX, ""..autotmp_6+16(SP)
	0x008f 00143 (main.go:6)	MOVQ	DX, (SP)
	0x0093 00147 (main.go:6)	MOVQ	"".&a+24(SP), AX
	0x0098 00152 (main.go:6)	MOVQ	AX, 8(SP)
	0x009d 00157 (main.go:6)	PCDATA	$0, $2
	0x009d 00157 (main.go:6)	CALL	runtime.writebarrierptr(SB)
	0x00a2 00162 (main.go:6)	MOVQ	""..autotmp_6+16(SP), AX
	0x00a7 00167 (main.go:6)	JMP	123
	0x00a9 00169 (main.go:6)	NOP
	0x00a9 00169 (main.go:5)	PCDATA	$0, $-1
	0x00a9 00169 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00ae 00174 (main.go:5)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 96  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
	0x0020 20 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00   H......H..$....
	0x0030 00 48 8b 44 24 08 48 89 44 24 18 48 8b 4c 24 30  .H.D$.H.D$.H.L$0
	0x0040 48 89 08 48 8d 0d 00 00 00 00 48 89 0c 24 e8 00  H..H......H..$..
	0x0050 00 00 00 48 8b 44 24 08 48 8d 0d 00 00 00 00 48  ...H.D$.H......H
	0x0060 89 08 84 00 8b 0d 00 00 00 00 48 8d 50 08 85 c9  ..........H.P...
	0x0070 75 18 48 8b 4c 24 18 48 89 48 08 48 89 44 24 38  u.H.L$.H.H.H.D$8
	0x0080 48 8b 6c 24 20 48 83 c4 28 c3 48 89 44 24 10 48  H.l$ H..(.H.D$.H
	0x0090 89 14 24 48 8b 44 24 18 48 89 44 24 08 e8 00 00  ..$H.D$.H.D$....
	0x00a0 00 00 48 8b 44 24 10 eb d2 e8 00 00 00 00 e9 4d  ..H.D$.........M
	0x00b0 ff ff ff                                         ...
	rel 5+4 t=16 TLS+0
	rel 36+4 t=15 type.int+0
	rel 45+4 t=8 runtime.newobject+0
	rel 70+4 t=15 type.struct { F uintptr; "".a *int }+0
	rel 79+4 t=8 runtime.newobject+0
	rel 91+4 t=15 "".test.func1+0
	rel 102+4 t=15 runtime.writeBarrier+0
	rel 158+4 t=8 runtime.writebarrierptr+0
	rel 170+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=319 args=0x0 locals=0x70
	0x0000 00000 (main.go:11)	TEXT	"".main(SB), $112-0
	0x0000 00000 (main.go:11)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:11)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:11)	JLS	309
	0x0013 00019 (main.go:11)	SUBQ	$112, SP
	0x0017 00023 (main.go:11)	MOVQ	BP, 104(SP)
	0x001c 00028 (main.go:11)	LEAQ	104(SP), BP
	0x0021 00033 (main.go:11)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x0021 00033 (main.go:11)	FUNCDATA	$1, gclocals·dd3278db0732f8fe4cd86d560ab06a9b(SB)
	0x0021 00033 (main.go:12)	MOVQ	$1, (SP)
	0x0029 00041 (main.go:12)	PCDATA	$0, $0
	0x0029 00041 (main.go:12)	CALL	"".test(SB)
	0x002e 00046 (main.go:12)	MOVQ	8(SP), DX
	0x0033 00051 (main.go:12)	MOVQ	DX, "".f+64(SP)
	0x0038 00056 (main.go:13)	MOVQ	$2, (SP)
	0x0040 00064 (main.go:13)	MOVQ	(DX), AX
	0x0043 00067 (main.go:13)	PCDATA	$0, $1
	0x0043 00067 (main.go:13)	CALL	AX
	0x0045 00069 (main.go:13)	MOVQ	8(SP), AX
	0x004a 00074 (main.go:14)	MOVQ	AX, ""..autotmp_4+56(SP)
	0x004f 00079 (main.go:14)	MOVQ	$0, ""..autotmp_3+88(SP)
	0x0058 00088 (main.go:14)	MOVQ	$0, ""..autotmp_3+96(SP)
	0x0061 00097 (main.go:14)	LEAQ	type.int(SB), AX
	0x0068 00104 (main.go:14)	MOVQ	AX, (SP)
	0x006c 00108 (main.go:14)	LEAQ	""..autotmp_4+56(SP), CX
	0x0071 00113 (main.go:14)	MOVQ	CX, 8(SP)
	0x0076 00118 (main.go:14)	PCDATA	$0, $2
	0x0076 00118 (main.go:14)	CALL	runtime.convT2E64(SB)
	0x007b 00123 (main.go:14)	MOVQ	16(SP), AX
	0x0080 00128 (main.go:14)	MOVQ	24(SP), CX
	0x0085 00133 (main.go:14)	MOVQ	AX, ""..autotmp_3+88(SP)
	0x008a 00138 (main.go:14)	MOVQ	CX, ""..autotmp_3+96(SP)
	0x008f 00143 (main.go:14)	LEAQ	""..autotmp_3+88(SP), AX
	0x0094 00148 (main.go:14)	MOVQ	AX, (SP)
	0x0098 00152 (main.go:14)	MOVQ	$1, 8(SP)
	0x00a1 00161 (main.go:14)	MOVQ	$1, 16(SP)
	0x00aa 00170 (main.go:14)	PCDATA	$0, $2
	0x00aa 00170 (main.go:14)	CALL	fmt.Println(SB)
	0x00af 00175 (main.go:15)	MOVQ	$3, (SP)
	0x00b7 00183 (main.go:15)	MOVQ	"".f+64(SP), DX
	0x00bc 00188 (main.go:15)	MOVQ	(DX), AX
	0x00bf 00191 (main.go:15)	PCDATA	$0, $0
	0x00bf 00191 (main.go:15)	CALL	AX
	0x00c1 00193 (main.go:15)	MOVQ	8(SP), AX
	0x00c6 00198 (main.go:16)	MOVQ	AX, ""..autotmp_6+48(SP)
	0x00cb 00203 (main.go:16)	MOVQ	$0, ""..autotmp_5+72(SP)
	0x00d4 00212 (main.go:16)	MOVQ	$0, ""..autotmp_5+80(SP)
	0x00dd 00221 (main.go:16)	LEAQ	type.int(SB), AX
	0x00e4 00228 (main.go:16)	MOVQ	AX, (SP)
	0x00e8 00232 (main.go:16)	LEAQ	""..autotmp_6+48(SP), AX
	0x00ed 00237 (main.go:16)	MOVQ	AX, 8(SP)
	0x00f2 00242 (main.go:16)	PCDATA	$0, $3
	0x00f2 00242 (main.go:16)	CALL	runtime.convT2E64(SB)
	0x00f7 00247 (main.go:16)	MOVQ	16(SP), AX
	0x00fc 00252 (main.go:16)	MOVQ	24(SP), CX
	0x0101 00257 (main.go:16)	MOVQ	AX, ""..autotmp_5+72(SP)
	0x0106 00262 (main.go:16)	MOVQ	CX, ""..autotmp_5+80(SP)
	0x010b 00267 (main.go:16)	LEAQ	""..autotmp_5+72(SP), AX
	0x0110 00272 (main.go:16)	MOVQ	AX, (SP)
	0x0114 00276 (main.go:16)	MOVQ	$1, 8(SP)
	0x011d 00285 (main.go:16)	MOVQ	$1, 16(SP)
	0x0126 00294 (main.go:16)	PCDATA	$0, $3
	0x0126 00294 (main.go:16)	CALL	fmt.Println(SB)
	0x012b 00299 (main.go:17)	MOVQ	104(SP), BP
	0x0130 00304 (main.go:17)	ADDQ	$112, SP
	0x0134 00308 (main.go:17)	RET
	0x0135 00309 (main.go:17)	NOP
	0x0135 00309 (main.go:11)	PCDATA	$0, $-1
	0x0135 00309 (main.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x013a 00314 (main.go:11)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 22  eH..%....H;a..."
	0x0010 01 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
	0x0020 68 48 c7 04 24 01 00 00 00 e8 00 00 00 00 48 8b  hH..$.........H.
	0x0030 54 24 08 48 89 54 24 40 48 c7 04 24 02 00 00 00  T$.H.T$@H..$....
	0x0040 48 8b 02 ff d0 48 8b 44 24 08 48 89 44 24 38 48  H....H.D$.H.D$8H
	0x0050 c7 44 24 58 00 00 00 00 48 c7 44 24 60 00 00 00  .D$X....H.D$`...
	0x0060 00 48 8d 05 00 00 00 00 48 89 04 24 48 8d 4c 24  .H......H..$H.L$
	0x0070 38 48 89 4c 24 08 e8 00 00 00 00 48 8b 44 24 10  8H.L$......H.D$.
	0x0080 48 8b 4c 24 18 48 89 44 24 58 48 89 4c 24 60 48  H.L$.H.D$XH.L$`H
	0x0090 8d 44 24 58 48 89 04 24 48 c7 44 24 08 01 00 00  .D$XH..$H.D$....
	0x00a0 00 48 c7 44 24 10 01 00 00 00 e8 00 00 00 00 48  .H.D$..........H
	0x00b0 c7 04 24 03 00 00 00 48 8b 54 24 40 48 8b 02 ff  ..$....H.T$@H...
	0x00c0 d0 48 8b 44 24 08 48 89 44 24 30 48 c7 44 24 48  .H.D$.H.D$0H.D$H
	0x00d0 00 00 00 00 48 c7 44 24 50 00 00 00 00 48 8d 05  ....H.D$P....H..
	0x00e0 00 00 00 00 48 89 04 24 48 8d 44 24 30 48 89 44  ....H..$H.D$0H.D
	0x00f0 24 08 e8 00 00 00 00 48 8b 44 24 10 48 8b 4c 24  $......H.D$.H.L$
	0x0100 18 48 89 44 24 48 48 89 4c 24 50 48 8d 44 24 48  .H.D$HH.L$PH.D$H
	0x0110 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7 44  H..$H.D$.....H.D
	0x0120 24 10 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 68  $..........H.l$h
	0x0130 48 83 c4 70 c3 e8 00 00 00 00 e9 c1 fe ff ff     H..p...........
	rel 5+4 t=16 TLS+0
	rel 42+4 t=8 "".test+0
	rel 67+0 t=11 +0
	rel 100+4 t=15 type.int+0
	rel 119+4 t=8 runtime.convT2E64+0
	rel 171+4 t=8 fmt.Println+0
	rel 191+0 t=11 +0
	rel 224+4 t=15 type.int+0
	rel 243+4 t=8 runtime.convT2E64+0
	rel 295+4 t=8 fmt.Println+0
	rel 310+4 t=8 runtime.morestack_noctxt+0
"".test.func1 STEXT nosplit size=21 args=0x10 locals=0x0
	0x0000 00000 (main.go:6)	TEXT	"".test.func1(SB), NOSPLIT|NEEDCTXT, $0-16
	0x0000 00000 (main.go:6)	FUNCDATA	$0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0000 00000 (main.go:6)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:6)	MOVQ	8(DX), AX
	0x0004 00004 (main.go:6)	MOVQ	"".i+8(SP), CX
	0x0009 00009 (main.go:7)	ADDQ	(AX), CX
	0x000c 00012 (main.go:7)	MOVQ	CX, (AX)
	0x000f 00015 (main.go:8)	MOVQ	CX, "".~r1+16(SP)
	0x0014 00020 (main.go:8)	RET
	0x0000 48 8b 42 08 48 8b 4c 24 08 48 03 08 48 89 08 48  H.B.H.L$.H..H..H
	0x0010 89 4c 24 10 c3                                   .L$..
"".init STEXT size=91 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	84
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JLS	47
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	JNE	56
	0x0031 00049 (<autogenerated>:1)	PCDATA	$0, $0
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0036 00054 (<autogenerated>:1)	UNDEF
	0x0038 00056 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	PCDATA	$0, $0
	0x003f 00063 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0044 00068 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	(SP), BP
	0x004f 00079 (<autogenerated>:1)	ADDQ	$8, SP
	0x0053 00083 (<autogenerated>:1)	RET
	0x0054 00084 (<autogenerated>:1)	NOP
	0x0054 00084 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0054 00084 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0059 00089 (<autogenerated>:1)	JMP	0
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
go.info."".test SDWARFINFO size=65
	0x0000 02 22 22 2e 74 65 73 74 00 00 00 00 00 00 00 00  ."".test........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 04 26 61 00  .............&a.
	0x0020 04 9c 11 68 22 00 00 00 00 05 61 00 01 9c 00 00  ...h".....a.....
	0x0030 00 00 05 7e 72 31 00 04 9c 11 08 22 00 00 00 00  ...~r1....."....
	0x0040 00                                               .
	rel 9+8 t=1 "".test+0
	rel 17+8 t=1 "".test+179
	rel 37+4 t=28 go.info.*int+0
	rel 46+4 t=28 go.info.int+0
	rel 60+4 t=28 go.info.func(int) int+0
go.range."".test SDWARFRANGE size=0
go.info."".main SDWARFINFO size=41
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 04 66 00 04  .............f..
	0x0020 9c 11 48 22 00 00 00 00 00                       ..H".....
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+319
	rel 36+4 t=28 go.info.func(int) int+0
go.range."".main SDWARFRANGE size=0
go.info."".test.func1 SDWARFINFO size=58
	0x0000 02 22 22 2e 74 65 73 74 2e 66 75 6e 63 31 00 00  ."".test.func1..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 01 05 69 00 01 9c 00 00 00 00 05 7e 72 31 00  ...i........~r1.
	0x0030 04 9c 11 08 22 00 00 00 00 00                    ....".....
	rel 15+8 t=1 "".test.func1+0
	rel 23+8 t=1 "".test.func1+21
	rel 39+4 t=28 go.info.int+0
	rel 53+4 t=28 go.info.int+0
go.range."".test.func1 SDWARFRANGE size=0
go.info."".init SDWARFINFO size=29
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+91
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata.*struct { F uintptr; a *int }- SRODATA dupok size=32
	0x0000 00 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ...*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 61 20 2a 69 6e 74 20 7d  intptr; a *int }
type..namedata..F- SRODATA dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type..namedata.a- SRODATA dupok size=4
	0x0000 00 00 01 61                                      ...a
type.struct { F uintptr; "".a *int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 bf fb c7 5a 02 08 08 19 00 00 00 00 00 00 00 00  ...Z............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+96
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int }-+0
	rel 44+4 t=6 type.*struct { F uintptr; "".a *int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.struct { F uintptr; "".a *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.a-+0
	rel 112+8 t=1 type.*int+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type.*struct { F uintptr; "".a *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a8 2a eb 0e 00 08 08 36 00 00 00 00 00 00 00 00  .*.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int }-+0
	rel 48+8 t=1 type.struct { F uintptr; "".a *int }+0
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*func(int) int- SRODATA dupok size=17
	0x0000 00 00 0e 2a 66 75 6e 63 28 69 6e 74 29 20 69 6e  ...*func(int) in
	0x0010 74                                               t
type.*func(int) int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2f eb d1 1f 00 08 08 36 00 00 00 00 00 00 00 00  /......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) int-+0
	rel 48+8 t=1 type.func(int) int+0
type.func(int) int SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 98 3c 32 87 02 08 08 33 00 00 00 00 00 00 00 00  .<2....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int) int-+0
	rel 44+4 t=6 type.*func(int) int+0
	rel 56+8 t=1 type.int+0
	rel 64+8 t=1 type.int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·d8b28f51bb91e05d264803f0f600a200 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 00 00                 ...........
gclocals·9a26515dfaeddd28bcbc040f1199f48d SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 02 01                 ...........
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·dd3278db0732f8fe4cd86d560ab06a9b SRODATA dupok size=12
	0x0000 04 00 00 00 05 00 00 00 00 01 19 06              ............
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
