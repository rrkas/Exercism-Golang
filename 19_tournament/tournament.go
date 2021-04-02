package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type score struct {
	name  string
	mp    int
	win   int
	draw  int
	loss  int
	point int
}

func (s score) String() string {
	name := s.name
	for len(name) < 31 {
		name += " "
	}
	return fmt.Sprintf("%s|  %d |  %d |  %d |  %d |  %d\n", name, s.mp, s.win, s.draw, s.loss, s.point)
}

func Tally(r io.Reader, w io.Writer) error {
	scores := map[string]score{}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	s := string(buf)
	for _, row := range strings.Split(s, "\n") {
		if len(row) == 0 || strings.HasPrefix(row, "#") {
			continue
		}
		items := strings.Split(row, ";")

		if len(items) != 3 {
			return errors.New("invalid row")
		}

		switch items[2] {
		case "win":
			calcForWin(scores, items[0])
			calcForLoss(scores, items[1])
		case "loss":
			calcForWin(scores, items[1])
			calcForLoss(scores, items[0])
		case "draw":
			calcForDraw(scores, items[0])
			calcForDraw(scores, items[1])
		default:
			return errors.New("invalid op")
		}
	}

	var result []score
	for _, v := range scores {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].point > result[j].point {
			return true
		} else if result[i].point == result[j].point {
			if result[i].name < result[j].name {
				return true
			}
		}
		return false
	})

	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, v := range result {
		w.Write([]byte(v.String()))
	}

	return nil
}

func calcForWin(scores map[string]score, team string) {
	v := scores[team]
	v.name = team
	v.mp++
	v.win++
	v.point += 3
	scores[team] = v
}

func calcForLoss(scores map[string]score, team string) {
	v := scores[team]
	v.name = team
	v.mp++
	v.loss++
	scores[team] = v
}

func calcForDraw(scores map[string]score, team string) {
	v := scores[team]
	v.name = team
	v.mp++
	v.draw++
	v.point++
	scores[team] = v
}
