#include "textflag.h"

TEXT ·Begin(SB),NOPTR|NOSPLIT,$0
	MOVL $0xffffffff, AX
	XBEGIN fallback
fallback:
	MOVL AX, ret+0(FP)
	RET

TEXT ·End(SB),NOPTR|NOSPLIT,$0
	XEND
	RET

TEXT ·Abort(SB),NOPTR|NOSPLIT,$0
	XABORT $0xf0
	RET

TEXT ·Test(SB),NOPTR|NOSPLIT,$0
	XTEST
	SETNE ret+0(FP)
	RET
