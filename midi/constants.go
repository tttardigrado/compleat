package midipackage

const numberOfNotes int = 12

const baseOctave int = 3

var (
	C3  Note = Note{0, 3, 100, 1}
	Cx3 Note = Note{1, 3, 100, 1}
	D3  Note = Note{2, 3, 100, 1}
	Dx3 Note = Note{3, 3, 100, 1}
	E3  Note = Note{4, 3, 100, 1}
	F3  Note = Note{5, 3, 100, 1}
	Fx3 Note = Note{6, 3, 100, 1}
	G3  Note = Note{7, 3, 100, 1}
	Gx3 Note = Note{8, 3, 100, 1}
	A3  Note = Note{9, 3, 100, 1}
	Ax3 Note = Note{10, 3, 100, 1}
	B3  Note = Note{11, 3, 100, 1}
	C4  Note = Note{0, 4, 100, 1}
	Cx4 Note = Note{1, 4, 100, 1}
	D4  Note = Note{2, 4, 100, 1}
	Dx4 Note = Note{3, 4, 100, 1}
	E4  Note = Note{4, 4, 100, 1}
	F4  Note = Note{5, 4, 100, 1}
	Fx4 Note = Note{6, 4, 100, 1}
	G4  Note = Note{7, 4, 100, 1}
	Gx4 Note = Note{8, 4, 100, 1}
	A4  Note = Note{9, 4, 100, 1}
	Ax4 Note = Note{10, 4, 100, 1}
	B4  Note = Note{11, 4, 100, 1}
)

var (
	// Major
	CM Scale = Scale{C3, D3, E3, F3, G3, A3, B3}
	DM Scale = Scale{D3, E3, Fx3, G3, A3, B3, Cx4}
	EM Scale = Scale{E3, Fx3, Gx3, A3, B3, Cx4, Dx4}
	FM Scale = Scale{F3, G3, A3, Ax3, C4, D4, E4}
	GM Scale = Scale{G3, A3, B3, C4, D4, E4, Fx4}
	AM Scale = Scale{A3, B3, Cx4, D4, E4, Fx4, Gx4}
	BM Scale = Scale{B3, Cx4, Dx4, E4, Fx4, Gx4, Ax4}

	// Major # (sharp)
	CxM Scale = Scale{Cx3, Dx3, F3, Fx3, Gx3, Ax3, C4}
	DxM Scale = Scale{Dx3, F3, G3, Gx3, Ax3, C4, D4}
	FxM Scale = Scale{Fx3, Gx3, Ax3, B3, Cx4, Dx4, F4}
	GxM Scale = Scale{Gx3, Ax3, C4, Cx4, Dx4, F4, G4}
	AxM Scale = Scale{Ax3, C4, D4, Dx4, F4, G4, A4}

	// Minor
	Cm Scale = Scale{C3, D3, Dx3, F3, G3, Gx3, Ax3}
	Dm Scale = Scale{D3, E3, F3, G3, A3, Ax3, C4}
	Em Scale = Scale{E3, Fx3, G3, A3, B3, C4, D4}
	Fm Scale = Scale{F3, G3, Gx3, Ax3, C4, Cx4, Dx4}
	Gm Scale = Scale{G3, A3, Ax3, C4, D4, Dx4, F4}
	Am Scale = Scale{A3, B3, C4, D4, E4, F4, G4}
	Bm Scale = Scale{B3, Cx4, D4, E4, Fx4, G4, A4}

	// Minor # (sharp)
	Cxm Scale = Scale{Cx3, Dx3, E3, Fx3, Gx3, A3, B3}
	Dxm Scale = Scale{Dx3, F3, Fx3, Gx3, Ax3, B3, Cx4}
	Fxm Scale = Scale{Fx3, Gx3, A3, B3, Cx4, D4, E4}
	Gxm Scale = Scale{Gx3, Ax3, B3, Cx4, Dx4, E4, Fx4}
	Axm Scale = Scale{Ax3, C4, Cx4, Dx4, F4, Fx4, Gx4}
)
