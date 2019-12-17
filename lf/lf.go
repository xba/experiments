package lf

const (
	N = 16
	M = N - 1
	O = 60
)

type Map [N][3]uint64

func (m *Map) Set(key, value uint64) bool {
	data := &m[(key*11400714819323198485)>>O]
	if data[0] == 0 {
		data[0] = 1
		data[1] = key
		data[2] = value
		return true
	}
	return false
}

func (m *Map) Get(key uint64) (uint64, bool) {
	//hash := (key * 11400714819323198485)
	return 0, false
}
