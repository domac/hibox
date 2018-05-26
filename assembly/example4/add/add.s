//TEXT 指令定义符号 runtime·profileloop, RET 表示结尾
//frame_size 是 $0, 表示了需要 stack 的空间大小，这里是0， 表示不需要stack，只使用 寄存器。函数的参数和返回值的大小为 3 * 8 = 24 bytes
//ret 符号是编译器默认的返回值符号,也可以定义其它名称

TEXT ·AddNum+0(SB),$0-24
MOVQ x+0(FP),AX
MOVQ y+8(FP),BX
ADDQ BX,AX
MOVQ AX, ret+16(FP)
RET 
