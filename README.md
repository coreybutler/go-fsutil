# fsutil

[![Version](https://img.shields.io/github/tag/coreybutler/go-fsutil.svg)](https://github.com/coreybutler/go-fsutil)
[![GoDoc](https://godoc.org/github.com/coreybutler/go-fsutil?status.svg)](https://godoc.org/github.com/coreybutler/go-fsutil)
[![Build Status](https://travis-ci.org/coreybutler/go-fsutil.svg?branch=master)](https://travis-ci.org/coreybutler/go-fsutil)

This cross-platform go module provides a lightweight abstraction of common file system methods:

- `Touch(path string)`: Like the Unix [touch command](https://en.wikipedia.org/wiki/Touch_(command)). Returns a string with the absolute path of the file/directory.
- `Mkdirp(path string)`: Like the Unix [mkdir -p](https://en.wikipedia.org/wiki/Mkdir) command. Returns a string with the absolute path of the directory.
- `Exists(path string)`: Returns a boolean indicating `true` if the path exists and `false` if it does not.
- `Abs(path string)`: Returns the absolute path as a string. Unlike the native [filepath.Abs](https://golang.org/pkg/path/filepath/#Abs), this method always returns a string (and only a string, no error). This method does not depend on the existance of the directory. Relative paths are always resolved from the current working directory.
- `Clean(path string)`: This method ensures an empty directory exists at the specified path.
- `IsFile(path string)`: Returns a boolean value indicating `true` if the path resolves to a file and `false` if it does not.
- `IsDirectory(path string)`: Returns a boolean value indicating `true` if the path resolves to a directory and `false` if it does not.
- `ReadTextFile(path string)`: Reads a text file and returns a _string_.
- `WriteTextFile(path string, content string, <os.FileMode>)`: Writes a text file from a _string_. Optionally accepts a file mode.
- `IsReadable(path string)`: Determines whether the path is readable. Returns a boolean.
- `IsWritable(path string)`: Determines whether the path is writable. Returns a boolean.

## Example

```go
package main

import (
  "github.com/coreybutler/go-fsutil"
)

func main() {
  fsutil.Touch("./path/to/test.txt")
}
```

The code above would create an empty file called `test.txt` in `<<current working directory>>/path/to`. If the directory does not exist, it will be created.

---

For complete details, see the [Godoc](https://godoc.org/github.com/coreybutler/go-fsutil). Examples are available in the [test files](https://github.com/coreybutler/go-fsutil/blob/master/fsutil_test.go).
