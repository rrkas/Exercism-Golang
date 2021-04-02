package robotname

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var names []string

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = getRandomName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = getRandomName()
}

func getRandomName() string {
	var name string
	sort.Strings(names)
	for {
		var sb strings.Builder
		sb.WriteString(getRandomLetter())
		sb.WriteString(getRandomLetter())
		sb.WriteString(fmt.Sprintf("%03d", rand.Intn(1000)))
		name = sb.String()
		index := sort.SearchStrings(names, name)
		if index == len(names) {
			names = append(names, name)
			break
		}
	}
	return name
}

func getRandomLetter() string {
	return string(rune(65 + rand.Intn(25)))
}
