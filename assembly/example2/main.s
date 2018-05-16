"".show STEXT size=120 args=0x20 locals=0x40
	0x0000 00000 (main.go:3)	TEXT	"".show(SB), $64-32
	0x0000 00000 (main.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:3)	JLS	113
	0x000f 00015 (main.go:3)	SUBQ	$64, SP
	0x0013 00019 (main.go:3)	MOVQ	BP, 56(SP)
	0x0018 00024 (main.go:3)	LEAQ	56(SP), BP
	0x001d 00029 (main.go:3)	FUNCDATA	$0, gclocals·4032f753396f2012ad1784f398b170f4(SB)
	0x001d 00029 (main.go:3)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (main.go:4)	MOVQ	$0, (SP)
	0x0025 00037 (main.go:4)	LEAQ	go.string."hello "(SB), AX
	0x002c 00044 (main.go:4)	MOVQ	AX, 8(SP)
	0x0031 00049 (main.go:4)	MOVQ	$6, 16(SP)
	0x003a 00058 (main.go:3)	MOVQ	"".a+72(SP), AX
	0x003f 00063 (main.go:4)	MOVQ	AX, 24(SP)
	0x0044 00068 (main.go:3)	MOVQ	"".a+80(SP), AX
	0x0049 00073 (main.go:4)	MOVQ	AX, 32(SP)
	0x004e 00078 (main.go:4)	PCDATA	$0, $1
	0x004e 00078 (main.go:4)	CALL	runtime.concatstring2(SB)
	0x0053 00083 (main.go:4)	MOVQ	40(SP), AX
	0x0058 00088 (main.go:4)	MOVQ	48(SP), CX
	0x005d 00093 (main.go:4)	MOVQ	AX, "".~r1+88(SP)
	0x0062 00098 (main.go:4)	MOVQ	CX, "".~r1+96(SP)
	0x0067 00103 (main.go:4)	MOVQ	56(SP), BP
	0x006c 00108 (main.go:4)	ADDQ	$64, SP
	0x0070 00112 (main.go:4)	RET
	0x0071 00113 (main.go:4)	NOP
	0x0071 00113 (main.go:3)	PCDATA	$0, $-1
	0x0071 00113 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0076 00118 (main.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 62 48  eH..%....H;a.vbH
	0x0010 83 ec 40 48 89 6c 24 38 48 8d 6c 24 38 48 c7 04  ..@H.l$8H.l$8H..
	0x0020 24 00 00 00 00 48 8d 05 00 00 00 00 48 89 44 24  $....H......H.D$
	0x0030 08 48 c7 44 24 10 06 00 00 00 48 8b 44 24 48 48  .H.D$.....H.D$HH
	0x0040 89 44 24 18 48 8b 44 24 50 48 89 44 24 20 e8 00  .D$.H.D$PH.D$ ..
	0x0050 00 00 00 48 8b 44 24 28 48 8b 4c 24 30 48 89 44  ...H.D$(H.L$0H.D
	0x0060 24 58 48 89 4c 24 60 48 8b 6c 24 38 48 83 c4 40  $XH.L$`H.l$8H..@
	0x0070 c3 e8 00 00 00 00 eb 88                          ........
	rel 5+4 t=16 TLS+0
	rel 40+4 t=15 go.string."hello "+0
	rel 79+4 t=8 runtime.concatstring2+0
	rel 114+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=168 args=0x0 locals=0x70
	0x0000 00000 (main.go:7)	TEXT	"".main(SB), $112-0
	0x0000 00000 (main.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:7)	JLS	158
	0x0013 00019 (main.go:7)	SUBQ	$112, SP
	0x0017 00023 (main.go:7)	MOVQ	BP, 104(SP)
	0x001c 00028 (main.go:7)	LEAQ	104(SP), BP
	0x0021 00033 (main.go:7)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (main.go:7)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (main.go:7)	LEAQ	""..autotmp_3+64(SP), AX
	0x0026 00038 (main.go:9)	MOVQ	AX, (SP)
	0x002a 00042 (main.go:9)	LEAQ	go.string."hello "(SB), AX
	0x0031 00049 (main.go:9)	MOVQ	AX, 8(SP)
	0x0036 00054 (main.go:9)	MOVQ	$6, 16(SP)
	0x003f 00063 (main.go:9)	LEAQ	go.string."domac"(SB), AX
	0x0046 00070 (main.go:9)	MOVQ	AX, 24(SP)
	0x004b 00075 (main.go:9)	MOVQ	$5, 32(SP)
	0x0054 00084 (main.go:9)	PCDATA	$0, $0
	0x0054 00084 (main.go:9)	CALL	runtime.concatstring2(SB)
	0x0059 00089 (main.go:9)	MOVQ	40(SP), AX
	0x005e 00094 (main.go:9)	MOVQ	AX, "".~r1.ptr+96(SP)
	0x0063 00099 (main.go:9)	MOVQ	48(SP), CX
	0x0068 00104 (main.go:9)	MOVQ	CX, "".~r1.len+56(SP)
	0x006d 00109 (main.go:9)	PCDATA	$0, $1
	0x006d 00109 (main.go:9)	CALL	runtime.printlock(SB)
	0x0072 00114 (main.go:9)	MOVQ	"".~r1.ptr+96(SP), AX
	0x0077 00119 (main.go:9)	MOVQ	AX, (SP)
	0x007b 00123 (main.go:9)	MOVQ	"".~r1.len+56(SP), AX
	0x0080 00128 (main.go:9)	MOVQ	AX, 8(SP)
	0x0085 00133 (main.go:9)	PCDATA	$0, $0
	0x0085 00133 (main.go:9)	CALL	runtime.printstring(SB)
	0x008a 00138 (main.go:9)	PCDATA	$0, $0
	0x008a 00138 (main.go:9)	CALL	runtime.printnl(SB)
	0x008f 00143 (main.go:9)	PCDATA	$0, $0
	0x008f 00143 (main.go:9)	CALL	runtime.printunlock(SB)
	0x0094 00148 (main.go:10)	MOVQ	104(SP), BP
	0x0099 00153 (main.go:10)	ADDQ	$112, SP
	0x009d 00157 (main.go:10)	RET
	0x009e 00158 (main.go:10)	NOP
	0x009e 00158 (main.go:7)	PCDATA	$0, $-1
	0x009e 00158 (main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x00a3 00163 (main.go:7)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 8b  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
	0x0020 68 48 8d 44 24 40 48 89 04 24 48 8d 05 00 00 00  hH.D$@H..$H.....
	0x0030 00 48 89 44 24 08 48 c7 44 24 10 06 00 00 00 48  .H.D$.H.D$.....H
	0x0040 8d 05 00 00 00 00 48 89 44 24 18 48 c7 44 24 20  ......H.D$.H.D$ 
	0x0050 05 00 00 00 e8 00 00 00 00 48 8b 44 24 28 48 89  .........H.D$(H.
	0x0060 44 24 60 48 8b 4c 24 30 48 89 4c 24 38 e8 00 00  D$`H.L$0H.L$8...
	0x0070 00 00 48 8b 44 24 60 48 89 04 24 48 8b 44 24 38  ..H.D$`H..$H.D$8
	0x0080 48 89 44 24 08 e8 00 00 00 00 e8 00 00 00 00 e8  H.D$............
	0x0090 00 00 00 00 48 8b 6c 24 68 48 83 c4 70 c3 e8 00  ....H.l$hH..p...
	0x00a0 00 00 00 e9 58 ff ff ff                          ....X...
	rel 5+4 t=16 TLS+0
	rel 45+4 t=15 go.string."hello "+0
	rel 66+4 t=15 go.string."domac"+0
	rel 85+4 t=8 runtime.concatstring2+0
	rel 110+4 t=8 runtime.printlock+0
	rel 134+4 t=8 runtime.printstring+0
	rel 139+4 t=8 runtime.printnl+0
	rel 144+4 t=8 runtime.printunlock+0
	rel 159+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=79 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	72
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
	0x0038 00056 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0043 00067 (<autogenerated>:1)	ADDQ	$8, SP
	0x0047 00071 (<autogenerated>:1)	RET
	0x0048 00072 (<autogenerated>:1)	NOP
	0x0048 00072 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0048 00072 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004d 00077 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 39 48  eH..%....H;a.v9H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 02 48  ...............H
	0x0040 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb b1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 73+4 t=8 runtime.morestack_noctxt+0
go.string."hello " SRODATA dupok size=6
	0x0000 68 65 6c 6c 6f 20                                hello 
go.info."".show SDWARFINFO size=52
	0x0000 02 22 22 2e 73 68 6f 77 00 00 00 00 00 00 00 00  ."".show........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 05 61 00 01  .............a..
	0x0020 9c 00 00 00 00 05 7e 72 31 00 04 9c 11 10 22 00  ......~r1.....".
	0x0030 00 00 00 00                                      ....
	rel 9+8 t=1 "".show+0
	rel 17+8 t=1 "".show+120
	rel 33+4 t=28 go.info.string+0
	rel 47+4 t=28 go.info.string+0
go.range."".show SDWARFRANGE size=0
go.string."domac" SRODATA dupok size=5
	0x0000 64 6f 6d 61 63                                   domac
go.info."".main SDWARFINFO size=65
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 04 7e 72 31  .............~r1
	0x0020 2e 6c 65 6e 00 04 9c 11 40 22 00 00 00 00 04 7e  .len....@".....~
	0x0030 72 31 2e 70 74 72 00 04 9c 11 68 22 00 00 00 00  r1.ptr....h"....
	0x0040 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+168
	rel 42+4 t=28 go.info.int+0
	rel 60+4 t=28 go.info.*uint8+0
go.range."".main SDWARFRANGE size=0
go.info."".init SDWARFINFO size=29
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+79
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..hashfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..alg32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc32+0
	rel 8+8 t=1 type..eqfunc32+0
type..namedata.*[32]uint8- SRODATA dupok size=13
	0x0000 00 00 0a 2a 5b 33 32 5d 75 69 6e 74 38           ...*[32]uint8
type.*[32]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f4 c7 79 15 00 08 08 36 00 00 00 00 00 00 00 00  ..y....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[32]uint8-+0
	rel 48+8 t=1 type.[32]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[32]uint8 SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x0010 9c 59 ff a8 02 01 01 91 00 00 00 00 00 00 00 00  .Y..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 20 00 00 00 00 00 00 00                           .......
	rel 24+8 t=1 type..alg32+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[32]uint8-+0
	rel 44+4 t=6 type.*[32]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
gclocals·4032f753396f2012ad1784f398b170f4 SRODATA dupok size=10
	0x0000 02 00 00 00 04 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
