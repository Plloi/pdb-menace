package menace

type Board struct {
	Layout [3][3]State
}

type State int

const (
	X State = 1 - iota
	EMPTY
	O
	MOVE
)

func NewBoard(a, b, c, d, e, f, h, i, j State) Board {
	return Board{Layout: [3][3]State{
		{a, b, c},
		{d, e, f},
		{h, i, j},
	}}
}

func BoardMatch(a, b Board) bool {
	var matches = 0

	if a.Layout[1][1] == b.Layout[1][1] {
		var aFlat = BoardFlatten(a)
		var bFlat = BoardFlatten(b)

		for rot := 0; rot < 4 && matches != 8; rot++ {
			matches = 0
			aFlat = BoardRotate(aFlat)
			for i := 0; i < 16 && matches != 8; i++ {
				if i == 8 {
					aFlat = BoardRevers(aFlat)
					matches = 0
				}
				if aFlat[i%8] == bFlat[i%8] {
					matches++
				}
			}
			if matches != 8 {
				aFlat = BoardRevers(aFlat)
			}
		}
	}

	return matches == 8
}

func BoardFlatten(b Board) [8]State {
	return [8]State{
		b.Layout[0][0], b.Layout[0][1], b.Layout[0][2], b.Layout[1][2], b.Layout[2][2], b.Layout[2][1], b.Layout[2][0], b.Layout[1][0],
	}
}
func BoardRevers(b [8]State) [8]State {
	return [8]State{b[0], b[7], b[6], b[5], b[4], b[3], b[2], b[1]}
}

func BoardRotate(b [8]State) [8]State {
	return [8]State{b[2], b[3], b[4], b[5], b[6], b[7], b[0], b[1]}
}
