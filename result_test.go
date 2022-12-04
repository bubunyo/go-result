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

func TestReadFileOk(t *testing.T) {
	for i, v := range ReadFile("./README.mb").Result() {
		fmt.Println(i, v)
	}
}

func TestReadWithError(t *testing.T) {
	assert.Panic(func() {
		for i, v := range ReadFile("./random").Expect("file not found") {
			fmt.Println(i, v)
		}
	})
}
