package fsutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testDir string = "./.data/a/b"

func TestExists(t *testing.T) {
	clear()

	abs, _ := filepath.Abs("./")

	if !Exists(abs) {
		t.Log("An existing directory is unrecognized.")
		t.Fail()
	}

	if !Exists(filepath.Join(abs, "go.mod")) {
		t.Log("An existing file is unrecognized.")
		t.Fail()
	}

	abs = filepath.Join(abs, testDir)

	if Exists(abs) {
		t.Log("A non-existant directory is recognized.")
		t.Fail()
	}

	if Exists(filepath.Join(abs, "noexist.ext")) {
		t.Log("A non-existant file is recognized.")
		t.Fail()
	}

	clear()
}

func TestAbs(t *testing.T) {
	clear()

	dir := testDir + "/c/d"
	absdir := Abs(dir)
	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir+"/c/d")

	if absdir != abs {
		t.Logf("Identifying absolute path failed. \"%v\" should be \"%v\".", absdir, abs)
		t.Fail()
	}

	clear()
}

func TestMkdirp(t *testing.T) {
	clear()

	dir := testDir + "/c/d"

	Mkdirp(dir)

	abs, err := filepath.Abs(dir)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(abs) == 0 {
		t.Log("Failed to created nested directories.")
		t.Fail()
	}

	stat, err := os.Stat(abs)
	if err != nil {
		if os.IsNotExist(err) {
			t.Logf("Mkdirp did not produce the appropriate directory structure: %v", abs)
			t.Fail()
		}
	}

	if !stat.IsDir() {
		t.Logf("Created a directory instead of a file at \"%v\"", abs)
	}

	clear()
}

func TestTouch(t *testing.T) {
	clear()

	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir)

	Touch(testDir + "/test.txt")

	if _, err := os.Stat(filepath.Join(abs, "test.txt")); err != nil {
		if os.IsNotExist(err) {
			t.Log("Failed to create empty file.")
			t.Fail()
		}
	}

	stat, err := os.Stat(filepath.Join(abs, "test.txt"))
	if err != nil {
		if os.IsNotExist(err) {
			t.Logf("Mkdirp did not produce the appropriate directory structure: %v", abs)
			t.Fail()
		}
	}

	if stat.IsDir() {
		t.Logf("Created a directory instead of a file at \"%v\"", abs)
		t.Fail()
	}

	clear()

	// Touch a directory
	abs = filepath.Join(abs, "dummydir.old")
	Touch(abs, false, true)

	stat2, err2 := os.Stat(abs)
	if err2 != nil {
		if os.IsNotExist(err) {
			t.Logf("Mkdirp did not produce the appropriate directory structure: %v", abs)
			t.Fail()
		}
	}

	if !stat2.IsDir() {
		t.Logf("Created a file instead of a directory at \"%v\"", abs)
		t.Fail()
	}

	clear()

	// Test forced file
	abs = filepath.Join(abs, "dummyshellscript")
	Touch(abs, true)

	stat3, err3 := os.Stat(abs)
	if err3 != nil {
		if os.IsNotExist(err) {
			t.Logf("Mkdirp did not produce the appropriate directory structure: %v", abs)
			t.Fail()
		}
	}

	if stat3.IsDir() {
		t.Logf("Created a directory instead of a file at \"%v\"", abs)
		t.Fail()
	}

	clear()
}

func TestIsFile(t *testing.T) {
	clear()

	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data/test.txt"

	_, err := os.Create(fp)
	if err != nil {
		t.Log("Problem creating test file.")
		t.Fail()
	}

	if !IsFile(fp) {
		t.Log("A file is not recognized as a file.")
		t.Fail()
	}

	if IsFile("./.data") {
		t.Log("A directory is recognized as a file.")
		t.Fail()
	}

	clear()
}

func TestIsDirectory(t *testing.T) {
	clear()

	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data/test.txt"

	_, err := os.Create(fp)
	if err != nil {
		t.Log("Problem creating test file.")
		t.Fail()
	}

	if IsDirectory(fp) {
		t.Log("A file is recognized as a directory.")
		t.Fail()
	}

	if !IsDirectory("./.data") {
		t.Log("A directory is not recognized as a directory.")
		t.Fail()
	}

	clear()
}

func TestClean(t *testing.T) {
	clear()

	Clean(testDir)

	stat, err := os.Stat(testDir)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if !stat.IsDir() {
		t.Log("Clean does not guarantee the existence of a directory.")
		t.Fail()
	}

	files, _ := ioutil.ReadDir(testDir)
	if len(files) > 0 {
		t.Logf("Content still exists after cleaning the directory. (%v files/directories)", len(files))
		t.Fail()
	}

	clear()
}

func clear() {
	os.RemoveAll("./.data")
}
