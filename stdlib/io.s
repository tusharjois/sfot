; Prevents double including
.define IO_HEADER

; Format specifiers
.define io_ascii_c $63
.define io_ascii_d $64
.define io_ascii_s $73
.define io_ascii_u $75
.define io_ascii_x $78

; Text lines
.define io_ascii_newline $0a
.define io_max_col_num $27 ; 40 total columns, position 39 zero indexed is last
.define io_max_row_num $17 ; 24 total rows, position 23 zero indexed is last
.define io_start_lo $00 ; Top left corner of screen, $0400
.define io_start_hi $04
.define io_end_lo $f7 ; Botton right corner of screen, $07f7
.define io_end_hi $07

; putc Print a char to the current cursor position
; [input] $01 the ASCII code of the char to print
; [zero page use] $01, $02, $03 TODO push to stack
; TODO outside cursor area

main:
	brk ; Useful only for subroutines

putc:
	ldx $01
	; Load address of cursor position
	lda $07f9
	sta $02
	lda $07fa
	sta $03
	; Store character at position
	ldy #00
	stx ($02),Y
	rts

newline:
	lda $07fb
	clc
	sbc #io_max_col_num

puts:

getc:

gets:

