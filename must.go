// Package must provides fatal wrappers for common stdlib functions.
package must

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// AbortIf should point to FatalIf or PanicIf or similar user-provided
// function which will interrupt execution in case it's param is not nil.
var AbortIf = FatalIf

// Alternative names to consider:
//  must.OK()
//  must.BeNil()
//  must.OrPanic()
//  must.OrAbort()
//  must.OrDie()

// NoErr is just a synonym for AbortIf.
func NoErr(err error) {
	AbortIf(err)
}

// FatalIf will call log.Fatal(err) in case given err is not nil.
func FatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PanicIf will call panic(err) in case given err is not nil.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// Atoi is a wrapper for strconv.Atoi.
func Atoi(buf []byte) int {
	i, err := strconv.Atoi(string(buf))
	AbortIf(err)
	return i
}

// Close is a wrapper for os.File.Close, …
func Close(f io.Closer) {
	err := f.Close()
	AbortIf(err)
}

// Create is a wrapper for os.Create.
func Create(name string) *os.File {
	f, err := os.Create(name)
	AbortIf(err)
	return f
}

// Decoder is an interface compatible with json.Decoder, gob.Decoder,
// xml.Decoder, …
type Decoder interface {
	Decode(v interface{}) error
}

// Decode is a wrapper for json.Decoder.Decode, gob.Decoder.Decode,
// xml.Decoder.Decode, …
func Decode(e Decoder, v interface{}) {
	err := e.Decode(v)
	AbortIf(err)
}

// Encoder is an interface compatible with json.Encoder, gob.Encoder,
// xml.Encoder, …
type Encoder interface {
	Encode(v interface{}) error
}

// Encode is a wrapper for json.Encoder.Encode, gob.Encoder.Encode,
// xml.Encoder.Encode, …
func Encode(e Encoder, v interface{}) {
	err := e.Encode(v)
	AbortIf(err)
}

// MarshalJSON is a wrapper for json.Marshal.
func MarshalJSON(v interface{}) []byte {
	data, err := json.Marshal(v)
	AbortIf(err)
	return data
}

// Open is a wrapper for os.Open.
func Open(name string) *os.File {
	f, err := os.Open(name)
	AbortIf(err)
	return f
}

// OpenFile is a wrapper for os.OpenFile.
func OpenFile(name string, flag int, perm os.FileMode) *os.File {
	f, err := os.OpenFile(name, flag, perm)
	AbortIf(err)
	return f
}

// Read is a wrapper for os.File.Read, bufio.Reader.Read, net.Conn.Read,
// rand.Read, …
func Read(r io.Reader, buf []byte) (n int) {
	n, err := r.Read(buf)
	AbortIf(err)
	return n
}

// ReadAll is a wrapper for ioutil.ReadAll.
func ReadAll(f io.Reader) []byte {
	buf, err := ioutil.ReadAll(f)
	AbortIf(err)
	return buf
}

// ReadDir is a wrapper for  ioutil.ReadDir.
func ReadDir(dirname string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dirname)
	AbortIf(err)
	return entries
}

// ReadFull is a wrapper for io.ReadFull.
func ReadFull(f io.Reader, buf []byte) int {
	n, err := io.ReadFull(f, buf)
	AbortIf(err)
	return n
}

// ReadFile is a wrapper for ioutil.ReadFile.
func ReadFile(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	AbortIf(err)
	return buf
}

// Remove is a wrapper for os.Remove.
func Remove(name string) {
	err := os.Remove(name)
	AbortIf(err)
}

// Rename is a wrapper for os.Rename.
func Rename(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	AbortIf(err)
}

// Seek is a wrapper for os.File.Seek, bytes.Buffer.Seek, …
func Seek(f io.Seeker, offset int64, whence int) int64 {
	n, err := f.Seek(offset, whence)
	AbortIf(err)
	return n
}

// Stat is a wrapper for os.File.Stat.
func Stat(f *os.File) os.FileInfo {
	fi, err := f.Stat()
	AbortIf(err)
	return fi
}

// StatPath is a wrapper for os.Stat.
func StatPath(name string) os.FileInfo {
	fi, err := os.Stat(name)
	AbortIf(err)
	return fi
}

// Sync is a wrapper for os.File.Sync.
func Sync(f *os.File) {
	err := f.Sync()
	AbortIf(err)
}

// TempDir is a wrapper for ioutil.TempDir.
func TempDir(dir, prefix string) string {
	name, err := ioutil.TempDir(dir, prefix)
	AbortIf(err)
	return name
}

// TempFile is a wrapper for ioutil.TempFile.
func TempFile(dir, prefix string) *os.File {
	f, err := ioutil.TempFile(dir, prefix)
	AbortIf(err)
	return f
}

// Truncate is a wrapper for os.File.Truncate.
func Truncate(f *os.File, size int64) {
	err := f.Truncate(size)
	AbortIf(err)
}

// TruncatePath is a wrapper for os.Truncate.
func TruncatePath(name string, size int64) {
	err := os.Truncate(name, size)
	AbortIf(err)
}

// UnmarshalJSON is a wrapper for json.Unmarshal.
func UnmarshalJSON(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	AbortIf(err)
}

// Write is a wrapper for os.File.Write, bufio.Writer.Write,
// net.Conn.Write, …
func Write(f io.Writer, b []byte) int {
	n, err := f.Write(b)
	AbortIf(err)
	return n
}

// WriteAt is a wrapper for os.File.WriteAt.
func WriteAt(f io.WriterAt, b []byte, off int64) int {
	n, err := f.WriteAt(b, off)
	AbortIf(err)
	return n
}

// WriteFile is a wrapper for ioutil.WriteFile.
func WriteFile(name string, buf []byte, perm os.FileMode) {
	err := ioutil.WriteFile(name, []byte("0"), perm)
	AbortIf(err)
}
