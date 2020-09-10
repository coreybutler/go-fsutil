package fsutil

import (
	"io/ioutil"
	"log"
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

func TestWriteTextFile(t *testing.T) {
	clear()
	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir)
	path := filepath.Join(abs, "test.txt")
	content := "test content"

	err := WriteTextFile(path, content)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if string(data) != content {
		t.Log("Failed to read the file content.")
		t.Fail()
	}

	clear()
}

func TestReadTextFile(t *testing.T) {
	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir)
	os.MkdirAll(abs, os.ModePerm)
	path := filepath.Join(abs, "test.txt")
	content := "test content"

	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	data, err := ReadTextFile(path)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if data != content {
		t.Log("Failed to read the file content.")
		t.Fail()
	}

	clear()
}

func TestPermissions(t *testing.T) {
	clear()
	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir)
	os.MkdirAll(abs, os.ModePerm)
	path := filepath.Join(abs, "test.txt")
	content := "test content"

	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if !IsReadable(path) {
		t.Log("File is readable, but method suggests it is not.")
		t.Fail()
	}

	if !IsWritable(path) {
		t.Log("File is writable, but method suggests it is not.")
		t.Fail()
	}

	if IsReadable(path + "junk") {
		t.Log("File is not readable, but method suggests it is.")
		t.Fail()
	}

	if IsWritable(path + "junk") {
		t.Log("File is not writable, but method suggests it is.")
		t.Fail()
	}

	clear()
}

func TestList(t *testing.T) {
	clear()
	abs, _ := filepath.Abs("./")
	abs = filepath.Join(abs, testDir)
	os.MkdirAll(abs, os.ModePerm)
	path := filepath.Join(abs, "test.txt")
	content := "test content"

	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	testPath := filepath.Dir(abs)

	// Basic Recursive Directory Listing
	list, err := List(testPath, true)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 3 {
		t.Logf("Expected 3 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// Non-recursive directory listing
	list, err = List(testPath, false)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 1 {
		t.Logf("Expected 1 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// Basic recursive directory listing with blacklist.
	list, err = List(testPath, true, filepath.Join(testPath, "/**/test.txt"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 2 {
		t.Logf("Expected 2 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// Basic non-recursive directory listing with blacklist.
	list, err = List(testDir, true, filepath.Join(testPath, "/**/test.txt"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 1 {
		t.Logf("Expected 1 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// List directories only (recursively)
	list, err = ListDirectories(testPath, true)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 2 {
		t.Logf("Expected 2 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// List directories only (recursively) w/ blacklist
	list, err = ListDirectories(testDir, true, filepath.Join(testPath, "/a"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 1 {
		t.Logf("Expected 1 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	list, err = ListFiles(testPath, true)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 1 {
		t.Logf("Expected 1 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	// Non-existant directory
	list, err = List("./dne", true)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if len(list) != 0 {
		t.Logf("Expected 0 results, received %v", len(list))
		t.Log(list)
		t.Fail()
	}

	clear()
}

func TestByteSize(t *testing.T) {
	clear()
	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data/test.txt"

	_, err := os.Create(fp)
	if err != nil {
		t.Log("Problem creating test file.")
		t.Fail()
	}

	size, err := ByteSize(fp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Logf("ByteSize: %v", size)
}

func TestSize(t *testing.T) {
	clear()
	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data/test.txt"

	_, err := os.Create(fp)
	if err != nil {
		t.Log("Problem creating test file.")
		t.Fail()
	}

	size, err := Size(fp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Logf("Size: %v", size)
}

func TestIsSymlink(t *testing.T) {
	clear()

	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data"

	ok := IsSymlink(fp)
	if ok {
		t.Log("Non-existant symlink detected.")
		t.Fail()
	}

	e := os.Symlink(".data", ".test")

	// Users may not have permission to create symlink.
	if e != nil {
		log.Print(e)
		clear()
		return
	}

	ok = IsSymlink(".test")
	os.Remove(".test")
	if !ok {
		t.Log("Symlink not detected.")
		t.Fail()
	}
}

func TestLastModified(t *testing.T) {
	clear()

	os.MkdirAll("./.data", os.ModePerm)

	fp := "./.data/test.txt"

	_, err := os.Create(fp)
	if err != nil {
		t.Log("Problem creating test file.")
		t.Fail()
	}

	_, err = LastModified(fp)
	clear()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	clear()
}

func clear() {
	os.RemoveAll("./.data")
}
