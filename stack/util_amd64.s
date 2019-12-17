#include "textflag.h"

TEXT ·swapInterface(SB),NOSPLIT,$0
	MOVQ addr+0(FP), BP
	XORQ AX, AX
	XORQ DX, DX
	MOVQ new+8(FP), BX
	MOVQ new+16(FP), CX
loop:
	LOCK
	CMPXCHG16B (BP)
	JNE loop
	MOVQ AX, ret+24(FP)
	MOVQ DX, ret+32(FP)
	RET

TEXT ·casInterface(SB),NOSPLIT,$0
	MOVQ addr+0(FP), BP
	MOVQ old+8(FP), AX
	MOVQ old+16(FP), DX
	MOVQ new+24(FP), BX
	MOVQ new+32(FP), CX
	LOCK
	CMPXCHG16B (BP)
	SETEQ ret+40(FP)
	RET

TEXT ·loadInterface(SB),NOSPLIT,$0
	MOVQ addr+0(FP), BP
	XORQ AX, AX
	XORQ DX, DX
	XORQ BX, BX
	XORQ CX, CX
	LOCK
	CMPXCHG16B (BP)
	MOVQ AX, ret+8(FP)
	MOVQ DX, ret+16(FP)
	RET

TEXT ·storeInterface(SB),NOSPLIT,$0
	MOVQ addr+0(FP), BP
	XORQ AX, AX
	XORQ DX, DX
	MOVQ new+8(FP), BX
	MOVQ new+16(FP), CX
loop:
	LOCK
	CMPXCHG16B (BP)
	JNE loop
	RET
