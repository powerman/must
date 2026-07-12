// Package must provides fatal wrappers for common stdlib functions.
package must

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

// AbortIf should point to FatalIf or PanicIf or similar user-provided
// function which will interrupt execution in case its param is not nil.
//
//nolint:gochecknoglobals // Package's whole purpose is a swappable abort strategy.
var AbortIf = FatalIf

// NoErr is just a synonym for AbortIf.
func NoErr(err error) {
	AbortIf(err)
}

// FatalIf will call [log.Fatal](err) in case given err is not nil.
func FatalIf(err error) {
	if err != nil {
		//nolint:revive // Calling log.Fatal is the documented purpose of FatalIf.
		log.Fatal(err)
	}
}

// PanicIf will call panic(err) in case given err is not nil.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// Atoi is a wrapper for [strconv.Atoi].
func Atoi(buf []byte) int {
	i, err := strconv.Atoi(string(buf))
	AbortIf(err)
	return i
}

// Close is a wrapper for [os.File.Close], ….
func Close(f io.Closer) {
	err := f.Close()
	AbortIf(err)
}

// Create is a wrapper for [os.Create].
func Create(name string) *os.File {
	//nolint:gosec // Package's whole purpose is opening caller-provided paths.
	f, err := os.Create(name)
	AbortIf(err)
	return f
}

// Decoder is an interface compatible with [json.Decoder], gob.Decoder,
// xml.Decoder, ….
type Decoder interface {
	Decode(v any) error
}

// Decode is a wrapper for [json.Decoder.Decode], gob.Decoder.Decode,
// xml.Decoder.Decode, ….
func Decode(e Decoder, v any) {
	err := e.Decode(v)
	AbortIf(err)
}

// Encoder is an interface compatible with [json.Encoder], gob.Encoder,
// xml.Encoder, ….
type Encoder interface {
	Encode(v any) error
}

// Encode is a wrapper for [json.Encoder.Encode], gob.Encoder.Encode,
// xml.Encoder.Encode, ….
func Encode(e Encoder, v any) {
	err := e.Encode(v)
	AbortIf(err)
}

// MarshalJSON is a wrapper for [json.Marshal].
func MarshalJSON(v any) []byte {
	data, err := json.Marshal(v)
	AbortIf(err)
	return data
}

// Open is a wrapper for [os.Open].
func Open(name string) *os.File {
	//nolint:gosec // Package's whole purpose is opening caller-provided paths.
	f, err := os.Open(name)
	AbortIf(err)
	return f
}

// OpenFile is a wrapper for [os.OpenFile].
func OpenFile(name string, flag int, perm os.FileMode) *os.File {
	//nolint:gosec // Package's whole purpose is opening caller-provided paths.
	f, err := os.OpenFile(name, flag, perm)
	AbortIf(err)
	return f
}

// Read is a wrapper for [os.File.Read], [bufio.Reader.Read], net.Conn.Read,
// rand.Read, ….
func Read(r io.Reader, buf []byte) (n int) {
	n, err := r.Read(buf)
	AbortIf(err)
	return n
}

// ReadAll is a wrapper for [io.ReadAll].
func ReadAll(f io.Reader) []byte {
	buf, err := io.ReadAll(f)
	AbortIf(err)
	return buf
}

// ReadDir is a wrapper for [os.ReadDir].
func ReadDir(dirname string) []os.DirEntry {
	entries, err := os.ReadDir(dirname)
	AbortIf(err)
	return entries
}

// ReadFull is a wrapper for [io.ReadFull].
func ReadFull(f io.Reader, buf []byte) int {
	n, err := io.ReadFull(f, buf)
	AbortIf(err)
	return n
}

// ReadFile is a wrapper for [os.ReadFile].
func ReadFile(name string) []byte {
	//nolint:gosec // Package's whole purpose is opening caller-provided paths.
	buf, err := os.ReadFile(name)
	AbortIf(err)
	return buf
}

// Remove is a wrapper for [os.Remove].
func Remove(name string) {
	err := os.Remove(name)
	AbortIf(err)
}

// Rename is a wrapper for [os.Rename].
func Rename(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	AbortIf(err)
}

// Seek is a wrapper for [os.File.Seek], bytes.Buffer.Seek, ….
func Seek(f io.Seeker, offset int64, whence int) int64 {
	n, err := f.Seek(offset, whence)
	AbortIf(err)
	return n
}

// Stat is a wrapper for [os.File.Stat].
func Stat(f *os.File) os.FileInfo {
	fi, err := f.Stat()
	AbortIf(err)
	return fi
}

// StatPath is a wrapper for [os.Stat].
func StatPath(name string) os.FileInfo {
	fi, err := os.Stat(name)
	AbortIf(err)
	return fi
}

// Sync is a wrapper for [os.File.Sync].
func Sync(f *os.File) {
	err := f.Sync()
	AbortIf(err)
}

// TempDir is a wrapper for [os.MkdirTemp].
func TempDir(dir, prefix string) string {
	name, err := os.MkdirTemp(dir, prefix)
	AbortIf(err)
	return name
}

// TempFile is a wrapper for [os.CreateTemp].
func TempFile(dir, prefix string) *os.File {
	f, err := os.CreateTemp(dir, prefix)
	AbortIf(err)
	return f
}

// Truncate is a wrapper for [os.File.Truncate].
func Truncate(f *os.File, size int64) {
	err := f.Truncate(size)
	AbortIf(err)
}

// TruncatePath is a wrapper for [os.Truncate].
func TruncatePath(name string, size int64) {
	err := os.Truncate(name, size)
	AbortIf(err)
}

// UnmarshalJSON is a wrapper for [json.Unmarshal].
func UnmarshalJSON(data []byte, v any) {
	err := json.Unmarshal(data, v)
	AbortIf(err)
}

// Write is a wrapper for [os.File.Write], [bufio.Writer.Write],
// net.Conn.Write, ….
func Write(f io.Writer, b []byte) int {
	n, err := f.Write(b)
	AbortIf(err)
	return n
}

// WriteAt is a wrapper for [os.File.WriteAt].
func WriteAt(f io.WriterAt, b []byte, off int64) int {
	n, err := f.WriteAt(b, off)
	AbortIf(err)
	return n
}

// WriteFile is a wrapper for [os.WriteFile].
func WriteFile(name string, buf []byte, perm os.FileMode) {
	err := os.WriteFile(name, buf, perm)
	AbortIf(err)
}
