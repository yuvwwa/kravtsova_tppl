%macro pushd 0
    push rax
    push rbx
    push rcx
    push rdx
%endmacro

%macro popd 0
    pop rdx
    pop rcx
    pop rbx
    pop rax
%endmacro

%macro cout 2
    pushd
    mov rax, 1
    mov rdi, 1
    mov rsi, %1
    mov rdx, %2
    syscall
    popd
%endmacro

; если число положительное, то просто выводим
; если число отрицательное, то оно преобразовывается в положительное и дописывается впереди минус

%macro dprint 0
    cmp rax, 0
    jge _positive
    
    push rax
    cout minus_sign, 1
    pop rax
    
    neg rax
    
_positive:
    mov rcx, 10
    mov rdx, 0
    mov rbx, 0

_divide:
    mov rdx, 0
    div rcx
    push rdx
    inc rbx
    cmp rax, 0
    jnz _divide
    
_digit:
    pop rax
    add rax, '0'
    mov [buffer], rax
    cout buffer, 1
    dec rbx
    cmp rbx, 0
    jg _digit
%endmacro

section .text
global _start

; rax - сумма разностей между элементами двух массивов x и y
; rcx - счётчик цикла (i)
; r8 - количество итераций (количество элементов в массиве)

_start:
    xor rax, rax 
    xor rcx, rcx
    mov r8, 7

calculate_loop:
    ; умножение rcx на 4, чтобы получить смещение до нужного элемента (каждый элемент массива x является define doubleword, 4 байта)
    ; [x + rcx * 4] - это адрес i-го элемента массива x
    movsxd rbx, dword [x + rcx * 4]
    movsxd rdx, dword [y + rcx * 4]
    
    ; x[i] - y[i]
    sub rbx, rdx

    ; (x[0]-y[0]) + ... + (x[i]-y[i])
    add rax, rbx

    ; rcx +=1
    inc rcx
    
    ; если прошли цикл, то выходим, если нет, то продолжаем
    cmp rcx, r8
    jl calculate_loop

    ; вычисление среднего арифметического
    mov rdx, 0
    mov rbx, r8
    cqo
    idiv rbx
    
    cout msg_result, msg_result_len
    dprint

end:
    cout newline, nlen
    mov rax, 60
    xor rdi, rdi
    syscall

section .data
    x dd 5, 3, 2, 6, 1, 7, 4
    y dd 0, 10, 1, 9, 2, 8, 5
    
    msg_result db "average: "
    msg_result_len equ $ - msg_result
    
    newline db 0xA, 0xD
    nlen equ $ - newline
    
    minus_sign db "-"

section .bss
    buffer resb 1