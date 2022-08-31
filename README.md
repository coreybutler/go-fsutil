# fsutil

This cross-platform go module provides a lightweight abstraction of common file system methods:

- `Touch(path string)`: Like the Unix [touch command](https://en.wikipedia.org/wiki/Touch_(command)). Returns a string with the absolute path of the file/directory.
- `Mkdirp(path string)`: Like the Unix [mkdir -p](https://en.wikipedia.org/wiki/Mkdir) command. Returns a string with the absolute path of the directory.
- `Exists(path string)`: Returns a boolean indicating `true` if the path exists and `false` if it does not.
- `Abs(path string)`: Returns the absolute path as a string. Unlike the native [filepath.Abs](https://golang.org/pkg/path/filepath/#Abs), this method always returns a string (and only a string, no error). This method does not depend on the existence of the directory. Relative paths are always resolved from the current working directory.
- `Clean(path string)`: This method ensures an empty directory exists at the specified path.
- `IsFile(path string)`: Returns a boolean value indicating `true` if the path resolves to a file and `false` if it does not.
- `IsDirectory(path string)`: Returns a boolean value indicating `true` if the path resolves to a directory and `false` if it does not.
- `IsSymlink(path string) bool`: Determines if a path is a symbolic link.
- `ReadTextFile(path string)`: Reads a text file and returns a _string_.
- `WriteTextFile(path string, content string, <os.FileMode>)`: Writes a text file from a _string_. Optionally accepts a file mode.
- `IsReadable(path string) bool`: Determines whether the path is readable.
- `IsWritable(path string) bool`: Determines whether the path is writable.
- `IsExecutable(path string) bool`: Determines whether the path has execute permissions.
- `ByteSize(path string)`: Determines the size (in bytes) of a file or directory.
- `Size(path string, decimalPlaces int)`: A "pretty" label for the size of a file or directory. For example, `3.14MB`.
- `FormatSize(size int64, decimalPlaces int)`: Pretty-print the byte size, i.e. `3.14MB`.
- `Copy(source string, target string, ignoreErrors ...bool) error`: Copy a file/directory contents. Ignores symlinks. Optionally specify `true` as the last argument to ignore errors.
- `Move(source string, target string, ignoreErrors ...bool) error`: Move a file/directory contents. Ignores symlinks. Optionally specify `true` as the last argument to ignore errors.
- `Unzip(source string, target string) error`: Unzip a file into the target directory.
- `Zip(source string, target string) error`: Zip a file/directory into the target directory/filename.

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

## Notice

This API is stable, but subject to additions. I work on it whenever a "common file system need" comes up in other projects. As a result, the `1.0.X` release cycle will continue to receive new feature additions until I consider the API "well-defined". This deviates a tiny bit from traditional semantic versioning, because new "features" are being added in patch releases.

Upon the release of a `1.1.0` version, the API will be considered "well-defined" and will adhere more strictly to semantic versioning practices.
