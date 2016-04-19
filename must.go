// Package must provides fatal wrappers for common stdlib functions.
package must

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func AbortIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// strconv.Atoi
func Atoi(buf []byte) int {
	i, err := strconv.Atoi(string(buf))
	AbortIf(err)
	return i
}

// os.File.Close, …
func Close(f io.Closer) {
	err := f.Close()
	AbortIf(err)
}

// os.Create
func Create(name string) *os.File {
	f, err := os.Create(name)
	AbortIf(err)
	return f
}

type Encoder interface {
	Encode(v interface{}) error
}

// json.Encoder.Encode, gob.Encoder.Encode, xml.Encoder.Encode, …
func Encode(e Encoder, v interface{}) {
	err := e.Encode(v)
	AbortIf(err)
}

// os.Open
func Open(name string) *os.File {
	f, err := os.Open(name)
	AbortIf(err)
	return f
}

// os.OpenFile
func OpenFile(name string, flag int, perm os.FileMode) *os.File {
	f, err := os.OpenFile(name, flag, perm)
	AbortIf(err)
	return f
}

// os.File.Read, bufio.Reader.Read, net.Conn.Read, rand.Read, …
func Read(r io.Reader, buf []byte) (n int) {
	n, err := r.Read(buf)
	AbortIf(err)
	return n
}

// ioutil.ReadAll
func ReadAll(f *os.File) []byte {
	buf, err := ioutil.ReadAll(f)
	AbortIf(err)
	return buf
}

// io.ReadFull
func ReadFull(f *os.File, buf []byte) int {
	n, err := io.ReadFull(f, buf)
	AbortIf(err)
	return n
}

// ioutil.ReadFile
func ReadFile(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	AbortIf(err)
	return buf
}

// os.Remove
func Remove(name string) {
	err := os.Remove(name)
	AbortIf(err)
}

// os.Rename
func Rename(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	AbortIf(err)
}

// os.File.Stat
func Stat(f *os.File) os.FileInfo {
	fi, err := f.Stat()
	AbortIf(err)
	return fi
}

// os.Stat
func StatPath(name string) os.FileInfo {
	fi, err := os.Stat(name)
	AbortIf(err)
	return fi
}

// os.File.Sync
func Sync(f *os.File) {
	err := f.Sync()
	AbortIf(err)
}

// os.File.Truncate
func Truncate(f *os.File, size int64) {
	err := f.Truncate(size)
	AbortIf(err)
}

// os.Truncate
func TruncatePath(name string, size int64) {
	err := os.Truncate(name, size)
	AbortIf(err)
}

// os.File.Write, bufio.Writer.Write, net.Conn.Write, …
func Write(f io.Writer, b []byte) int {
	n, err := f.Write(b)
	AbortIf(err)
	return n
}

// os.File.WriteAt
func WriteAt(f io.WriterAt, b []byte, off int64) int {
	n, err := f.WriteAt(b, off)
	AbortIf(err)
	return n
}

// ioutil.WriteFile
func WriteFile(name string, buf []byte, perm os.FileMode) {
	err := ioutil.WriteFile(name, []byte("0"), perm)
	AbortIf(err)
}
