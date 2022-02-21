package geohash

// quickNeighbors returns the neighbors of the given hash byte in order of:
// NW, N, NE, E, SE, S, SW, W.
type quickNeighbors [8]byte

const (
	_quickIndex_NW = iota
	_quickIndex_N
	_quickIndex_NE
	_quickIndex_E
	_quickIndex_SE
	_quickIndex_S
	_quickIndex_SW
	_quickIndex_W
)

const overflow byte = 0x80

func isOverflow(b byte) bool {
	return b & overflow != 0
}

// oddquick is index of the odd mapping table. in order of NW, N, NE, E, SE, S, SW, W.
// .eg. oddquick['d'] = { c (NorthWest) f g e 7 6 3 9 (West) }
//
// 'b', 'c', 'f', 'g', 'u', 'v', 'y', 'z'
// '8', '9', 'd', 'e', 's', 't', 'w', 'x'
// '2', '3', '6', '7', 'k', 'm', 'q', 'r'
// '0', '1', '4', '5', 'h', 'j', 'n', 'p'
//
var oddquick = map[byte]quickNeighbors{
	'b': {overflow|'p', overflow|'0', overflow|'1', 'c', '9', '8', overflow, overflow},
	'c': {overflow, overflow, overflow, 'f', 'd', '9', '8', 'b'},
	'f': {overflow, overflow, overflow, 'g', 'e', 'd', '9', 'c'},
	'g': {overflow, overflow, overflow, 'u', 's', 'e', 'd', 'f'},
	'u': {overflow, overflow, overflow, 'v', 't', 's', 'e', 'g'},
	'v': {overflow, overflow, overflow, 'y', 'w', 't', 's', 'u'},
	'y': {overflow, overflow, overflow, 'z', 'x', 'w', 't', 'v'},
	'z': {overflow, overflow, overflow, overflow, overflow, 'x', 'w', 'y'},
	'8': {overflow, 'b', 'c', '9', '3', '2', overflow, overflow},
	'9': {'b', 'c', 'f', 'd', '6', '3', '2', '8'},
	'd': {'c', 'f', 'g', 'e', '7', '6', '3', '9'},
	'e': {'f', 'g', 'u', 's', 'k', '7', '6', 'd'},
	's': {'g', 'u', 'v', 't', 'm', 'k', '7', 'e'},
	't': {'u', 'v', 'y', 'w', 'q', 'm', 'k', 's'},
	'w': {'v', 'y', 'z', 'x', 'r', 'q', 'm', 't'},
	'x': {'y', 'z', overflow, overflow, overflow, 'r', 'q', 'w'},
	'2': {overflow, '8', '9', '3', '1', '0', overflow, overflow},
	'3': {'8', '9', 'd', '6', '4', '1', '0', '2'},
	'6': {'9', 'd', 'e', '7', '5', '4', '1', '3'},
	'7': {'d', 'e', 's', 'k', 'h', '5', '4', '6'},
	'k': {'e', 's', 't', 'm', 'j', 'h', '5', '7'},
	'm': {'s', 't', 'w', 'q', 'n', 'j', 'h', 'k'},
	'q': {'t', 'w', 'x', 'r', 'p', 'n', 'j', 'm'},
	'r': {'w', 'x', overflow, overflow, overflow, 'p', 'n', 'q'},
	'0': {overflow, '2', '3', '1', overflow, overflow, overflow, overflow},
	'1': {'2', '3', '6', '4', overflow, overflow, overflow, '0'},
	'4': {'3', '6', '7', '5', overflow, overflow, overflow, '1'},
	'5': {'6', '7', 'k', 'h', overflow, overflow, overflow, '4'},
	'h': {'7', 'k', 'm', 'j', overflow, overflow, overflow, '5'},
	'j': {'k', 'm', 'q', 'n', overflow, overflow, overflow, 'h'},
	'n': {'m', 'q', 'r', 'p', overflow, overflow, overflow, 'j'},
	'p': {'q', 'r', overflow, overflow, overflow, overflow, overflow, 'n'},
}

// evenquick
// p,r,x,z
// n,q,w,y
// j,m,t,v
// h,k,s,u
// 5,7,e,g
// 4,6,d,f
// 1,3,9,c
// 0,2,8,b
var evenquick = map[byte]quickNeighbors{
	'p': {overflow, overflow, overflow, 'r', 'q', 'n', overflow, overflow},
	'r': {overflow, overflow, overflow, 'x', 'w', 'q', 'n', 'p'},
	'x': {overflow, overflow, overflow, 'z', 'y', 'w', 'q', 'r'},
	'z': {overflow, overflow, overflow, overflow, overflow, 'y', 'w', 'x'},
	'n': {overflow, 'p', 'r', 'q', 'm', 'j', overflow, overflow},
	'q': {'p', 'r', 'x', 'w', 't', 'm', 'j', 'n'},
	'w': {'r', 'x', 'z', 'y', 'v', 't', 'm', 'q'},
	'y': {'x', 'z', overflow, overflow, overflow, 'v', 't', 'w'},
	'j': {overflow, 'n', 'q', 'm', 'k', 'h', overflow, overflow},
	'm': {'n', 'q', 'w', 't', 's', 'k', 'h', 'j'},
	't': {'q', 'w', 'y', 'v', 'u', 's', 'k', 'm'},
	'v': {'w', 'y', overflow, overflow, overflow, 'u', 's', 't'},
	'h': {overflow, 'j', 'm', 'k', '7', '5', overflow, overflow},
	'k': {'j', 'm', 't', 's', 'e', '7', '5', 'h'},
	's': {'m', 't', 'v', 'u', 'g', 'e', '7', 'k'},
	'u': {'t', 'v', overflow, overflow, overflow, 'g', 'e', 's'},
	'5': {overflow, 'h', 'k', '7', '6', '4', overflow, overflow},
	'7': {'h', 'k', 's', 'e', 'd', '6', '4', '5'},
	'e': {'k', 's', 'u', 'g', 'f', 'd', '6', '7'},
	'g': {'s', 'u', overflow, overflow, overflow, 'f', 'd', 'e'},
	'4': {overflow, '5', '7', '6', '3', '1', overflow, overflow},
	'6': {'5', '7', 'e', 'd', '9', '3', '1', '4'},
	'd': {'7', 'e', 'g', 'f', 'c', '9', '3', '6'},
	'f': {'e', 'g', overflow, overflow, overflow, 'c', '9', 'd'},
	'1': {overflow, '4', '6', '3', '2', '0', overflow, overflow},
	'3': {'4', '6', 'd', '9', '8', '2', '0', '1'},
	'9': {'6', 'd', 'f', 'c', 'b', '8', '2', '3'},
	'c': {'d', 'f', overflow, overflow, overflow, 'b', '8', '9'},
	'0': {overflow, '1', '3', '2', overflow, overflow, overflow, overflow},
	'2': {'1', '3', '9', '8', overflow, overflow, overflow, '0'},
	'8': {'3', '9', 'c', 'b', overflow, overflow, overflow, '2'},
	'b': {'9', 'c', overflow, overflow, overflow, overflow, overflow, '8'},
}

func FastNeighbors(hash string) (result []string) {
	org := []byte(hash)
	result = make([]string, 8)
	findIndex := func(odd bool, c byte) quickNeighbors {
		return oddquick[c]
	}
	pos := len(org)
	overf := 8

	for overf != 0 || pos <= 0 {
		q := findIndex(pos%2 == 0, org[pos-1])
		for idx, v := range q {
			if v != overflow {
				result[idx] = 
			}
		}

		matchAll = (countOverflow(result) == 0)
		pos--
	}

	return result
}
