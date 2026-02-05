#include "textflag.h"

TEXT ·_ElementProviderFromPoint_ASM(SB), NOSPLIT, $80-32
    MOVQ CX, 0(SP)
    MOVQ R9, 8(SP)
    MOVSD X2, 16(SP)
    MOVSD X1, 24(SP)

    MOVQ ·_ElementProviderFromPoint(SB), AX
    MOVQ AX, 48(SP)
    
    LEAQ 0(SP), BX
    MOVQ BX, 56(SP)
    
    MOVQ $32, 64(SP)
    MOVQ $0, 72(SP)
    
    CALL runtime·cgocallback(SB)
    RET
