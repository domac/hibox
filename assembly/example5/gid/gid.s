#include "go_asm.h"
#include "textflag.h"

#define	get_tls(r)	MOVQ TLS, r
#define	g(r)	0(r)(TLS*1)

TEXT ·getg+0(SB), NOSPLIT, $0-8
    get_tls(CX)
    MOVQ    g(CX), AX
    MOVQ    AX, ret+0(FP)
    RET
