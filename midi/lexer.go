package midipackage

type itemType int

type item struct {
	itype itemType
	value string
}

const (
	itemError itemType = iota

	itemModifierStart
	itemModifierEnd
	itemModifierPlus
	itemModifierMinus
	itemModifierVelocity

	itemChrodStart
	itemChrodEnd

	itemGroupStart
	itemGroupEnd

	itemComment

	itemBpmModifier
	itemBpmValue

	itemKeyModifier
	itemKeyValue

	itemNote
)
