TEXT	main(SB),512|7,$0
	ADD	R1, R2
	ADD	R1, R2, R2
	ADD R1, R2, R3
	AND	$42, R1
	AND	$42, R1, R1
	AND	$42, R1, R2
	SLLD	$16, R1
	SLLD	$16, R1, R1
	SLLD	$16, R1, R2
	FADDD	D2, D4
	FADDD	D2, D4, D4
	FADDD	D2, D4, D6
	CASD	(R1), R2, R3
	CMP	R1, R2
	RET
