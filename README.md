# fsutil

[![Version](https://img.shields.io/github/tag/coreybutler/go-fsutil.svg)](https://github.com/coreybutler/go-fsutil)
[![GoDoc](https://godoc.org/github.com/coreybutler/go-fsutil?status.svg)](https://godoc.org/github.com/coreybutler/go-fsutil)
[![Build Status](https://travis-ci.org/coreybutler/go-fsutil.svg?branch=master)](https://travis-ci.org/coreybutler/go-fsutil)

This cross-platform library provides a lightweight abstraction of common file system methods:

- `Touch(path string)`: Like the Unix [touch command](https://en.wikipedia.org/wiki/Touch_(command)). Returns a string with the absolute path of the file/directory.
- `Mkdirp(path string)`: Like the Unix [mkdir -p](https://en.wikipedia.org/wiki/Mkdir) command. Returns a string with the absolute path of the directory.
- `Exists(path string)`: Returns a boolean indicating `true` if the path exists and `false` if it does not.
- `Abs(path string)`: Returns the absolute path as a string. Unlike the native [filepath.Abs](https://golang.org/pkg/path/filepath/#Abs), this method always returns a string (and only a string, no error). This method does not depend on the existance of the directory. Relative paths are always resolved from the current working directory.
- `Clean(path string)`: This method ensures an empty directory exists at the specified path.
- `IsFile(path string)`: Returns a boolean value indicating `true` if the path resolves to a file and `false` if it does not.
- `IsDirectory(path string)`: Returns a boolean value indicating `true` if the path resolves to a directory and `false` if it does not.

For complete details, see the [Godoc](https://godoc.org/github.com/coreybutler/go-fsutil). Examples are available in the [test files](https://github.com/coreybutler/go-fsutil/blob/master/fsutil_test.go).