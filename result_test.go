package result_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/bubunyo/go-result"
	"github.com/stretchr/testify/assert"
)

func ReadFile(input string) result.Result[[]string] {
	file, err := ioutil.ReadFile(input)
	if err != nil {
		return result.Error[[]string](err)
	}
	fileContent := string(file)
	return result.Ok(strings.Split(fileContent, "\n"))
}

func joinString(s1, s2 result.Result[[]string]) result.Result[[]string] {
	return result.Ok(append(s1.Result(), s2.Result()...))
}

func TestReadFileOk(t *testing.T) {
	count := 0
	for range ReadFile("./rick-roll-1").Result() {
		count++
	}
	assert.Equal(t, 4, count)
}

func TestReadWithError(t *testing.T) {
	assert.PanicsWithValue(t, "open ./rick-roll-3: no such file or directory, rick roll p3 not present", func() {
		for i, v := range ReadFile("./rick-roll-3").Expect("rick roll p3 not present") {
			fmt.Println(i, v)
		}
	})
}

func TestThen(t *testing.T) {
	song := ReadFile("./rick-roll-1").
		Then(func(p1 []string) result.Resolver[[]string] {
			return result.Ok(append(p1, ReadFile("./rick-roll-2").Result()...))
		}).
		Result()

	assert.Equal(t, strings.Join(song, "\n"),
		`Never gonna give you up
Never gonna let you down
Never gonna run around and desert you

Never gonna make you cry
Never gonna say goodbye
Never gonna tell a lie and hurt you
`)
}
