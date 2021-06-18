package menace

type Board struct {
	Layout [3][3]State
}

type State int

const (
	X State = 1 - iota
	EMPTY
	O
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

		for rot := 0; rot < 8 && matches != 8; rot += 2 {
			matches = 0
			for i := 0; i < 16; i++ {
				if i == 8 {
					aFlat = BoardRevers(aFlat)
				}
				if aFlat[(i+rot)%8] == bFlat[i%8] {
					matches++
				}
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
	return [8]State{b[7], b[6], b[5], b[4], b[3], b[2], b[1], b[0]}
}
