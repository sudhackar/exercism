package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type stats struct {
	mp, w, d, l, p int
	name           string
}

var scoreAdd = map[string][]int{
	"win": {3, 0}, "loss": {0, 3}, "draw": {1, 1},
}

var teamMap = make(map[string]*stats)

// Tally reads the results and writes table
func Tally(reader io.Reader, writer io.Writer) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	// split lines and calculate stats for teams
	lines := strings.Split(string(b), "\n")
	m, err := calculate(lines)
	if err != nil {
		return err
	}

	// convert to array to sort
	statArr := make([]*stats, 0, len(m))
	for _, value := range m {
		statArr = append(statArr, value)
	}

	sort.Slice(statArr, func(i, j int) bool {
		if statArr[i].p == statArr[j].p {
			return statArr[i].name < statArr[j].name
		}
		return statArr[i].p > statArr[j].p
	})

	//write
	_, err = fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")

	if err != nil {
		return err
	}

	for _, team := range statArr {
		_, err = fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.mp, team.w, team.d, team.l, team.p)
		if err != nil {
			return err
		}
	}
	return nil
}

func calculate(lines []string) (map[string]*stats, error) {
	mp := make(map[string]*stats)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if len(line) > 0 && (line[0] == '#' || line[0] == '\n') {
			continue
		}
		words := strings.Split(strings.TrimSpace(string(line)), ";")
		if len(words) != 3 {
			return nil, errors.New("incorrect non empty line")
		}
		teamName1 := words[0]
		teamName2 := words[1]
		matchResult := words[2]

		if _, ok := scoreAdd[matchResult]; !ok {
			return nil, errors.New("result must be one of win/loss/draw")
		}

		if _, ok := mp[teamName1]; !ok {
			mp[teamName1] = &stats{name: teamName1}
		}
		if _, ok := mp[teamName2]; !ok {
			mp[teamName2] = &stats{name: teamName2}
		}
		// Increment played/win/loss/draw
		mp[teamName1].mp++
		mp[teamName2].mp++
		if matchResult == "win" {
			mp[teamName1].w++
			mp[teamName2].l++
		} else if matchResult == "draw" {
			mp[teamName1].d++
			mp[teamName2].d++
		} else if matchResult == "loss" {
			mp[teamName1].l++
			mp[teamName2].w++
		}
		// Increment scores
		sc := scoreAdd[matchResult]
		mp[teamName1].p += sc[0]
		mp[teamName2].p += sc[1]

	}
	return mp, nil
}
