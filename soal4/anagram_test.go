package soal4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkCheckAndListAnagram1(b *testing.B) {
	var words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	checkAnagram(words)
}

func TestCountGroup(t *testing.T) {
	var words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	var result = checkAnagram(words)
	assert.Equal(t, len(result), 4, "expect there are 4 groups")
	var expect = []string{"kita", "atik", "tika"}
	assert.Contains(t, result, expect, "expect contains : kita, atik, tika")
	var expect2 = []string{"makan"}
	assert.Contains(t, result, expect2, "expect contain : makan")
	var expect3 = []string{"kia"}
	assert.Contains(t, result, expect3, "expect contain : kia")
	var expect4 = []string{"aku", "kua"}
	assert.Contains(t, result, expect4, "expect contains : aku, kua")
}

func TestCountGroup2(t *testing.T) {
	var words = []string{""}
	var result = checkAnagram(words)
	assert.Equal(t, len(result), 1, "")
	var expect = []string{""}
	assert.Contains(t, result, expect, "expect empty string")
}

func TestCountGroup3(t *testing.T) {
	var words = []string{"helicopter", "jet", "plane"}
	var result = checkAnagram(words)
	assert.Equal(t, len(result), 3, "")
	var expect = []string{"helicopter"}
	assert.Contains(t, result, expect, "helicopter")
	var expect2 = []string{"jet"}
	assert.Contains(t, result, expect2, "jet")
	var expect3 = []string{"plane"}
	assert.Contains(t, result, expect3, "plane")
}
