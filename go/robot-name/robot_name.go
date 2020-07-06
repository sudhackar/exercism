package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot has a name
type Robot struct {
	name string
}

var mp = make(map[string]map[int]bool)
var count uint32
var source = rand.NewSource(time.Now().UnixNano())
var randSource = rand.New(source)

// Name returns a name if possible else error if namespace is exhausted
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	// lets target bonus, not the benchmark
	if count == 26*26*10*10*10 {
		count = 0
		return "", errors.New("namespace exhausted")
	}
	p, s := name()
	for checkExisting(p, s) {
		p, s = name()
	}
	if _, ok := mp[p]; !ok {
		mp[p] = make(map[int]bool)
	}
	mp[p][s] = true
	count++
	r.name = p + fmt.Sprintf("%03d", s)
	return r.name, nil
}

// Reset will reset the name
func (r *Robot) Reset() {
	r.name = ""
}

func checkExisting(p string, s int) bool {
	if _, ok := mp[p]; ok {
		if b, ok := mp[p][s]; ok {
			return b
		}
	}
	return false
}

func name() (string, int) {
	return prefix(), suffix()
}

func prefix() string {
	b := make([]rune, 2)
	b[0] = rune(0x41 + randSource.Intn(26))
	b[1] = rune(0x41 + randSource.Intn(26))
	return string(b)
}

func suffix() int {
	return randSource.Intn(1000)
}
