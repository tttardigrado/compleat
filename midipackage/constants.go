package midipackage

const numberOfNotes uint8 = 12

const (
	C3 MidiNote = iota + 48
	Cx3
	D3
	Dx3
	E3
	F3
	Fx3
	G3
	Gx3
	A3
	Ax3
	B3
	C4
	Cx4
	D4
	Dx4
	E4
	F4
	Fx4
	G4
	Gx4
	A4
	Ax4
	B4
)

var (
	// Major
	CM Scale = Scale{C3, D3, E3, F3, G3, A3, B3, C4}
	DM Scale = Scale{D3, E3, Fx3, G3, A3, B3, Cx4, D4}
	EM Scale = Scale{E3, Fx3, Gx3, A3, B3, Cx4, Dx4, E4}
	FM Scale = Scale{F3, G3, A3, Ax3, C4, D4, E4, F4}
	GM Scale = Scale{G3, A3, B3, C4, D4, E4, Fx4, G4}
	AM Scale = Scale{A3, B3, Cx4, D4, E4, Fx4, Gx4, A4}
	BM Scale = Scale{B3, Cx4, Dx4, E4, Fx4, Gx4, Ax4, B4}

	// Major # (sharp)
	CxM Scale = Scale{Cx3, Dx3, F3, Fx3, Gx3, Ax3, C4, Cx4}
	DxM Scale = Scale{Dx3, F3, G3, Gx3, Ax3, C4, D4, Dx4}
	FxM Scale = Scale{Fx3, Gx3, Ax3, B3, Cx4, Dx4, F4, Fx4}
	GxM Scale = Scale{Gx3, Ax3, C4, Cx4, Dx4, F4, G4, Gx4}
	AxM Scale = Scale{Ax3, C4, D4, Dx4, F4, G4, A4, Ax4}

	// Minor
	Cm Scale = Scale{C3, D3, Dx3, F3, G3, Gx3, Ax3, C4}
	Dm Scale = Scale{D3, E3, F3, G3, A3, Ax3, C4, D4}
	Em Scale = Scale{E3, Fx3, G3, A3, B3, C4, D4, E4}
	Fm Scale = Scale{F3, G3, Gx3, Ax3, C4, Cx4, Dx4, F4}
	Gm Scale = Scale{G3, A3, Ax3, C4, D4, Dx4, F4, G4}
	Am Scale = Scale{A3, B3, C4, D4, E4, F4, G4, A4}
	Bm Scale = Scale{B3, Cx4, D4, E4, Fx4, G4, A4, B4}

	// Minor # (sharp)
	Cxm Scale = Scale{Cx3, Dx3, E3, Fx3, Gx3, A3, B3, Cx4}
	Dxm Scale = Scale{Dx3, F3, Fx3, Gx3, Ax3, B3, Cx4, Dx4}
	Fxm Scale = Scale{Fx3, Gx3, A3, B3, Cx4, D4, E4, Fx4}
	Gxm Scale = Scale{Gx3, Ax3, B3, Cx4, Dx4, E4, Fx4, Gx4}
	Axm Scale = Scale{Ax3, C4, Cx4, Dx4, F4, Fx4, Gx4, Ax4}
)
