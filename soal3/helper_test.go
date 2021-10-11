package soal3

import "testing"
import "github.com/stretchr/testify/assert"

func TestFindFirstStringInBracket1(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket(""), "", "")
}

func TestFindFirstStringInBracket2(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my(name)is"), "name", "")
}

func TestFindFirstStringInBracket3(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my()name"), "", "")
}

func TestFindFirstStringInBracket4(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my name()"), "", "")
}

func TestFindFirstStringInBracket5(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("()my name is"), "", "")
}

func TestFindFirstStringInBracket6(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("(my name is)"), "my name is", "")
}

func TestFindFirstStringInBracket7(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("()"), "", "")
}

func TestFindFirstStringInBracket8(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my name is"), "", "")
}

func TestFindFirstStringInBracket9(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my name (is"), "", "")
}

func TestFindFirstStringInBracket10(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket("my name is)"), "", "")
}

func TestFindFirstStringInBracket11(t *testing.T) {
	assert.Equal(t, findFirstStringInBracket(")my name is("), "", "")
}
