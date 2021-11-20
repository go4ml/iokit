package iokit

import (
	"fmt"
	"gotest.tools/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func Test_Version(t *testing.T) {
	assert.Assert(t, Version.Major() == 1)
	assert.Assert(t, Version.Minor() == 0)
	assert.Assert(t, Version.Patch() == 0)
	assert.Assert(t, Version.String() == "1.0.0")
}

func Test_Open(t *testing.T) {
	S := `test`
	w, err := os.Create(CacheFile("go-iokit/tests/create.txt"))
	assert.NilError(t, err)
	_, err = w.WriteString(S)
	assert.NilError(t, err)
	err = w.Close()
	assert.NilError(t, err)

	r := File(CacheFile("go-iokit/tests/create.txt")).MustOpen()
	defer r.Close()
	x := r.MustReadAll()
	assert.Assert(t, string(x) == S)

	r2 := Cache("go-iokit/tests/create.txt").MustOpen()
	defer r2.Close()
	x = r2.MustReadAll()
	assert.Assert(t, string(x) == S)
}

func Test_CreateOpen(t *testing.T) {
	S := `test`
	file := File(CacheFile("go-iokit/tests/createopen.txt"))
	w, err := file.Create()
	assert.NilError(t, err)
	defer w.End()
	_, err = w.Write([]byte(S))
	assert.NilError(t, err)
	err = w.Commit()
	assert.NilError(t, err)
	r, err := file.Open()
	assert.NilError(t, err)
	defer r.Close()
	x, err := ioutil.ReadAll(r)
	assert.NilError(t, err)
	assert.Assert(t, string(x) == S)
}

func Test_CacheOpen(t *testing.T) {
	S := `test`
	file := Cache("go-iokit/tests/createopen.txt").File()
	w, err := file.Create()
	assert.NilError(t, err)
	defer w.End()
	_, err = w.Write([]byte(S))
	assert.NilError(t, err)
	err = w.Commit()
	assert.NilError(t, err)
	r, err := file.Open()
	assert.NilError(t, err)
	defer r.Close()
	x, err := ioutil.ReadAll(r)
	assert.NilError(t, err)
	assert.Assert(t, string(x) == S)
}

func findSkills(s string) []string {
	j := strings.Index(s, "SKILLS")
	j = strings.Index(s[j:], "<li>") + j
	k := strings.Index(s[j:], "</li>") + j
	return strings.Split(s[j+4:k], ", ")
}

func Test_PathHttp(t *testing.T) {
	cache := Cache("go-iokit/test_httppath.txt")
	cache.Remove()
	file := Url("http://sudachen.github.io/cv", cache)
	r, err := file.Open()
	assert.NilError(t, err)
	defer r.Close()
	x, err := ioutil.ReadAll(r)
	assert.NilError(t, err)
	u := findSkills(string(x))
	assert.Assert(t, u[0] == "Go")
	r, err = Url("file://" + cache.Path()).Open()
	assert.NilError(t, err)
	defer r.Close()
	x, err = ioutil.ReadAll(r)
	assert.NilError(t, err)
	u = findSkills(string(x))
	assert.Assert(t, u[0] == "Go")
}

func Test_StringIO(t *testing.T) {
	S := fmt.Sprintf(`test text %v`, rand.Int())
	r := StringIO(S).MustOpen()
	assert.Assert(t, S == string(r.MustReadAll()))
}
