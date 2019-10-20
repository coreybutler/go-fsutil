package fsutil

// The fsutil module provides a lightweight cross-platform
// set of helper methods for interacting with the file system.
// Most of the methods are designed with the intention of
// "guaranteeing" a system resource does or does not exist.
// They are designed to abstract common functionality into
// more easily understood code.

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Similar to the touch command on *nix, where the file
// or directory will be created if it does not already exist.
// Returns the absolute path.
// The optional second boolean argument will force
// the method to treat the path as a file instead of a directory
// (useful when the filename has not extension).
// An optional 3rd boolean argument will force the method
// to treat the path as a directory even if a file extension is present.
//
// For example:
// `fsutil.Touch("./path/to/archive.old", false, true)`
//
// Normally, any file path with an extension is determined
// to be a file. However; the second argument (`false`)
// instructs the command to **not** force a file. The third
// argument (`true`) instructs the command to **treat the path
// like a directory**.
func Touch(path string, flags ...interface{}) string {
	abs := Abs(path)

	if !Exists(path) {
		forceFile := false
		forceDir := false

		if len(flags) > 0 {
			for i, flag := range flags {
				if i == 0 {
					forceFile = flag.(bool)
				} else if i == 1 {
					forceDir = flag.(bool)
				}
			}
		}

		ext := filepath.Ext(abs)

		if !forceDir && (forceFile || len(ext) > 0) {
			Mkdirp(filepath.Dir(abs))

			file, err := os.Create(abs)
			if err != nil {
				panic(err)
			}

			file.Close()
		} else {
			Mkdirp(abs)
		}
	}

	return abs
}

// Mkdirp is the equivalent of [mkdir -p](https://en.wikipedia.org/wiki/Mkdir)
// It will generate the full directory path if it does not already
// exist.
func Mkdirp(path string) string {
	path = Abs(path)
	os.MkdirAll(path, os.ModePerm)
	return path
}

// Exists is a helper method to quickly
// determine whether a directory or file exists.
func Exists(path string) bool {
	if len(Abs(path)) == 0 {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// Returns the fully resolved path, even if the
// path does not exist.
//
// ```
// fsutil.Abs("./does/not/exist")
// ```
// If the code above was run within `/home/user`, the
// result would be `/home/user/does/not/exist`.
func Abs(path string) string {
	abs, _ := filepath.Abs(path)
	return abs
}

// Clean will ensure the specified directory exists.
// If the directory already exists, all of contents
// are deleted. If the directory does not exist, it
// is automatically created.
func Clean(path string) {
	path = Abs(path)

	if IsFile(path) {
		path = filepath.Dir(path)
	}

	if Exists(path) {
		os.RemoveAll(path)
	}

	Mkdirp(path)
}

// IsFile determines whether the specified path
// represents a file.
func IsFile(path string) bool {
	if !Exists(path) {
		return false
	}

	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !stat.IsDir()
}

// IsDirectory determines whether the specified path
// represents a directory.
func IsDirectory(path string) bool {
	if !Exists(path) {
		return false
	}

	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// Writes text to a file (automatically converts string to
// a byte array). If the path does not exist, it will be
// created automatically. This is the equivalent of using
// the Touch() method first, then writing text content to
// the file.
//
// It is also possible to pass a third argument, a custom permission.
// By default, os.ModePerm is used.
func WriteTextFile(path string, content string, args ...interface{}) error {
	path = Touch(path, true)
	perm := os.ModePerm

	if len(args) > 0 {
		perm = args[0].(os.FileMode)
	}

	return ioutil.WriteFile(path, []byte(content), perm)
}

// Reads a text file and converts results from bytes
// to a string.
func ReadTextFile(path string) (string, error) {
	data, err := ioutil.ReadFile(Abs(path))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Determines whether the file/directory is Readable
// for the active system user.
func IsReadable(path string) bool {
	return allowFileAction(path, os.O_RDONLY, 0666)
}

// Determines whether the file/directory is Writable
// for the active system user.
func IsWritable(path string) bool {
	return allowFileAction(path, os.O_WRONLY, 0666)
}

func allowFileAction(path string, flag int, perm os.FileMode) bool {
	path = Abs(path)

	if !Exists(path) {
		return false
	}

	file, err := os.OpenFile(path, flag, perm)
	allowed := true
	if err != nil {
		if os.IsPermission(err) {
			allowed = false
		}
	}
	file.Close()

	return allowed
}

type listpath struct {
	Path string
	Stat os.FileInfo
}

func list(directory string, recursive bool, ignore ...string) ([]*listpath, error) {
	directory = Abs(directory)
	response := make([]*listpath, 0)
	var ignored error

	// Walk recursive lists
	if recursive {
		_ = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			ignored = isIgnoredPath(path, ignore...)
			if ignored != nil {
				return ignored
			}

			response = append(response, &listpath{
				Path: path,
				Stat: info,
			})

			return nil
		})
	} else {
		paths, matchErr := filepath.Glob(filepath.Clean(filepath.Join(directory, "/*")))

		if matchErr == nil {
			for _, path := range paths {
				ignored = isIgnoredPath(path, ignore...)
				if ignored == nil {
					stat, _ := os.Stat(path)
					response = append(response, &listpath{
						Path: path,
						Stat: stat,
					})
				}
			}
		} else {
			return make([]*listpath, 0), matchErr
		}
	}

	return response, nil
}

func isIgnoredPath(path string, ignore ...string) error {
	if len(ignore) > 0 {
		for _, pattern := range ignore {
			matched, matchErr := filepath.Match(pattern, path)

			if matchErr != nil {
				return matchErr
			}

			if matched {
				return errors.New("Ignored")
			}
		}
	}

	return nil
}

// Generate a list of path names for the given directory.
// Optionally provide a list of ignored paths, using
// [glob](https://en.wikipedia.org/wiki/Glob_%28programming%29) syntax.
func List(directory string, recursive bool, ignore ...string) ([]string, error) {
	response, err := list(directory, recursive, ignore...)
	if err != nil {
		return make([]string, 0), err
	}

	paths := make([]string, len(response))
	for i := range response {
		paths[i] = response[i].Path
	}

	return paths, nil
}

// ListDirectories provides absolute paths of directories only, ignoring files.
func ListDirectories(directory string, recursive bool, ignore ...string) ([]string, error) {
	paths := make([]string, 0)
	response, err := list(directory, recursive, ignore...)
	if err != nil {
		return paths, err
	}

	if len(response) == 0 {
		return paths, nil
	}

	for _, item := range response {
		if item.Stat.IsDir() {
			paths = append(paths, item.Path)
		}
	}

	return paths, nil
}

// ListFiles provides absolute paths of files only, ignoring directories.
func ListFiles(directory string, recursive bool, ignore ...string) ([]string, error) {
	paths := make([]string, 0)
	response, err := list(directory, recursive, ignore...)
	if err != nil {
		return paths, err
	}

	if len(response) == 0 {
		return paths, nil
	}

	for _, item := range response {
		if !item.Stat.IsDir() {
			paths = append(paths, item.Path)
		}
	}

	return paths, nil
}

// TODO List
// LastModified
// Created
// Size: bytes, kb, mb, gb
// IsSymlink
// Symlink
// Move
// Copy
// Rename
// Append
// Prepend
// IsExecutable?
