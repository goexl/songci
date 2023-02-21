package songci

type codes []uint8

func (c codes) Add(code ...uint8) (new codes) {
	new = make(codes, 0, len(c)+len(code))
	new = append(new, c...)
	new = append(new, code...)

	return
}
