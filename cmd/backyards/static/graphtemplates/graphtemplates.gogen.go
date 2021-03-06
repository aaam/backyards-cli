// Code generated by vfsgen; DO NOT EDIT.

package graphtemplates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// GraphTemplates statically implements the virtual filesystem provided to vfsgen.
var GraphTemplates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/base.json": &vfsgen۰CompressedFileInfo{
			name:             "base.json",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 4139,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\x51\x6f\xe2\x46\x10\xc7\xdf\xf9\x14\xab\x6d\x1f\x48\x85\x1c\x03\x76\x82\x11\x41\x4a\xef\x12\xa9\x52\xa5\x3b\x25\x55\x5f\xc2\x15\x2d\xf6\x00\xab\x33\xbb\xce\xec\x9a\x42\x23\xf7\xb3\x57\x6b\x4c\xf0\xda\x1c\x24\x97\x34\xb9\x97\x28\x8c\xf7\x3f\xb3\x33\xf3\x9b\xb1\x1f\x1a\x84\xd0\x25\xa0\xe2\x52\xd0\x3e\xa1\xcb\x36\x6d\x19\x53\xc4\x34\x53\x32\xc5\x10\x14\xed\x13\x73\xca\x18\x77\xff\x13\x42\x13\x94\x0b\xd0\x73\x48\xcb\x56\x42\x28\x8b\x22\x04\x65\x8c\x74\xae\x75\xd2\x3f\x3d\x6d\x77\xce\x1d\xd7\x71\x9d\x76\x3f\x70\x03\xf7\xb4\x24\x2c\x54\x59\x63\xfb\x37\x2b\xa2\xab\xf9\x44\x32\x8c\x76\xb1\x97\x0c\x39\x9b\xc4\x60\x5d\x41\x73\x1d\xc3\x6d\x3a\x9d\xf2\x95\x7d\x87\x50\x0a\xa5\x99\xd0\xc6\x6a\xb4\x71\x0a\xe6\x3e\x84\x16\xa1\x8a\x40\xf9\xd9\x29\x8f\x35\xe0\x93\xf4\x08\x89\x44\x0d\x78\x31\xa2\x9b\xe2\x8c\x68\x2b\x02\xa5\xb9\x60\x9a\x4b\x31\x56\x80\x4b\x1e\xc2\x58\xb0\x05\x5c\x8c\xe8\x14\xa5\xd0\x09\x9b\xc1\x88\xee\x0b\xcc\x85\x06\x5c\xb2\xd8\x0e\x5d\xb6\x12\xaa\x34\x24\x26\x65\xdf\xdd\x39\x68\x94\xdc\xd0\xbf\x79\x34\x03\x6d\x8e\xdc\x15\xcf\x4b\xbe\xf2\xfa\x98\x8b\xdf\x7c\xbe\x1d\x0c\x9c\x52\xb9\x86\x43\xda\xda\x9d\x9b\x21\x8f\x12\x69\xf7\xd1\xf8\xa6\x7d\xd2\xed\x3e\x5a\xb2\x92\x42\x71\x31\x8b\x41\x69\xa6\xab\xa2\xfb\x14\x70\x5d\x31\x12\x42\x61\x95\x98\x1a\x53\x95\x2e\x9a\xc8\x34\x34\xb9\xd2\x5c\x8e\x11\xee\x53\x50\x5a\x8d\xb5\xd4\x2c\x7e\x18\x0c\x9c\x4d\x3b\x86\xc3\xec\x6e\x30\x70\xb6\xc5\x18\x0e\xbf\x9c\x9c\x94\x6e\xbc\xc1\xf1\x91\xd1\xdf\x3e\x1a\xd7\xd1\x23\x4f\x95\xdb\x12\x42\x53\xc1\xcd\x4d\xa9\x90\x02\x2c\x3f\x9b\xe6\xfe\x01\xab\xfc\xf1\x60\xe0\xe4\xbf\xad\xf2\x98\x50\x10\xf2\x05\x8b\x4d\x85\x3a\xd6\x03\x3d\x47\x50\x73\x19\xe7\xa3\x71\xf7\x60\xc8\x89\x65\x9e\xe8\x4f\xe7\x67\xac\x1d\x9d\x51\x92\x7d\xd9\x55\xb0\xca\xc0\xbe\x5e\x5d\x21\x4a\x54\xa4\xe9\xaf\x56\x27\xef\xd4\xa4\xe6\x93\xbb\xd4\x42\x50\x89\x14\x0a\xc6\xa1\x8c\xe0\xe2\xdf\x11\xf5\x1d\x67\x44\xeb\xcd\x23\xa7\xe4\x05\xad\x27\x27\xe4\x17\xd2\x76\x5d\xf2\xe9\x86\x2c\x21\xd4\x12\x9b\xee\xab\xf0\x90\x00\x86\x20\xf4\xcb\x90\x68\x1f\x40\xc2\xba\xa2\x5d\x6e\x52\xa6\xa5\xd7\xf9\x70\x75\xdd\xa1\xd6\x81\xac\xf5\x54\xf5\xf5\xf5\x79\xcf\xbd\xac\x14\xc4\x10\xa0\x19\xea\x3f\x8b\xf5\xe5\x3a\x6e\xfb\x3b\xfd\x47\x9e\xc7\xba\xec\x88\xff\x8e\xed\xbc\xf4\xeb\x99\x13\xf0\x3b\xd3\x20\xc2\xf5\x37\xe0\x2f\x90\xf7\x5e\x07\xf5\x39\x57\x5a\xce\x90\x2d\xc6\xf7\x29\x13\x9a\xc7\xd0\x74\x9d\x20\x68\x55\x69\x9d\xb0\xf0\xeb\x9a\x61\xa4\xb6\xdc\x8e\xa3\x14\xb7\x5b\x3f\x94\x22\x52\xe3\x49\x1a\x7e\x05\x7d\x04\xe4\xc9\x9a\x34\x63\x78\x9d\x65\x56\x04\xfe\x0e\x78\x0f\x31\x5a\xee\x7b\x37\xec\x41\x18\xd1\x1a\x29\xfb\xd8\xab\xd1\xd6\xf5\x0f\xea\xb6\x4c\xd5\x74\x67\x2f\x61\xe7\xe6\xf3\xed\x41\x6e\xcc\x12\xc9\xac\x03\x2c\x99\x57\x99\x59\x72\x95\xb2\x98\xff\x93\xf7\xb7\xce\x8e\x02\xe4\xa0\x3e\x2d\x01\x91\x47\x50\xab\x60\x91\x29\xc2\x0c\xcc\x67\x09\xfd\xab\x63\x92\xdc\x25\xee\x85\x6c\xea\xf9\xb5\xa2\xd6\x64\x9e\x2d\x9b\x86\xd3\x73\x1f\x8e\xcb\x7c\x5b\x06\x7e\x9b\x79\x21\xb5\x6a\x9a\xd7\xb5\x36\xcc\xeb\xcb\x15\xaf\xbe\x5e\x48\x99\x37\x84\xfb\x44\xd5\x97\x40\xf9\xe5\xf8\xcd\x15\x60\x13\x6c\x86\x92\xc3\x13\x16\xe4\x9e\xb1\xa8\xc5\x7f\x85\x4f\x8b\x7c\x2c\xad\x57\x59\x75\x40\x09\xa1\x31\xcc\x40\x44\x9b\xb1\x22\x8e\x75\x9a\x0c\x87\xf4\x2d\xb7\xdf\x9b\x50\x6c\x51\x95\xf8\xee\xcf\x36\x58\x57\x97\xbf\xf6\xba\xc1\x1e\x1e\x6d\x5d\xe0\x57\x75\xd7\x3d\xaf\xfb\xe1\xb8\x2e\xa8\xea\x3a\xde\x47\xaf\x53\x05\xb9\x82\xf1\x71\x88\xb7\x4b\xf3\x3d\x49\xfd\x41\x5e\x3a\x16\xd3\x49\x10\x3c\xef\xeb\xe3\xe5\x09\xfb\xff\x7b\xc2\xe4\x70\xc6\xfe\x1b\x67\xec\xbb\xef\x9c\xb1\xef\x3e\x6b\x4d\x35\xb6\xf6\xac\x91\x35\xfe\x0b\x00\x00\xff\xff\xf3\x93\xc1\x45\x2b\x10\x00\x00"),
		},
		"/cb.json": &vfsgen۰CompressedFileInfo{
			name:             "cb.json",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 2108,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\x5d\x6f\xea\x46\x10\x7d\xe7\x57\x8c\xa6\x2f\x10\x59\xc6\x10\xb8\xb9\x17\x39\xae\x6e\x6f\x1b\xa9\x4f\x91\xfa\xf5\x72\x49\xd1\x62\x8f\xcd\xaa\xc6\xeb\xcc\xae\x5d\x68\x44\x7f\x7b\xb5\x98\x0f\xaf\x43\x14\xd4\x97\xfb\x92\xe0\xe3\x3d\x73\xd6\x67\xe6\xcc\x4b\x0f\x00\x6b\x62\x2d\x55\x81\x33\xc0\x7a\x84\x9e\x85\x12\x61\x84\x56\x15\xc7\xa4\x71\x06\xf6\x94\x05\xcf\xbf\x01\xb0\x64\xb5\x26\xb3\xa2\xaa\x8d\x02\xa0\x48\x12\x26\x6d\x41\x5c\x19\x53\xce\x86\xc3\xd1\xf8\xce\x0f\xfc\xc0\x1f\xcd\x3e\x05\x9f\x82\x61\x8b\x78\x60\xed\x7a\xc7\xbf\xbb\x83\xba\x5e\x2d\x95\xe0\xe4\xac\x5d\x0b\x96\x62\x99\x93\x73\x05\x23\x4d\x4e\xbf\x56\x69\x2a\x37\xee\x1d\x62\x55\x68\x23\x0a\x63\x51\xcb\xcd\x2b\xb2\xf7\x01\x3c\x48\x1d\x84\xf6\x67\x53\x99\x1b\xe2\xab\xf8\x4c\xa5\x62\x43\x7c\x3f\xc7\xc6\x9c\x39\x7a\x09\x69\x23\x0b\x61\xa4\x2a\x16\x9a\xb8\x96\x31\x2d\x0a\xb1\xa6\xfb\x39\xa6\xac\x0a\x53\x8a\x8c\xe6\x78\x49\x58\x16\x86\xb8\x16\xb9\x2b\xdd\x46\x01\xb5\xa1\xd2\x7e\xf2\x34\x38\x17\xe8\xb5\xca\xe0\xdf\x32\xc9\xc8\xd8\x23\x5f\x0f\xef\x5b\xb5\xf6\xfe\xd8\x8b\x7f\x91\x1c\x57\xd2\xc0\x0f\x4c\xe2\x2f\x62\x30\x2c\x4b\x0d\x2c\x0c\x85\xa1\xdf\x72\x31\x8a\xd0\x3b\xd3\x33\x96\x49\xa9\xdc\xf6\x5a\x49\x9c\xc1\xed\xed\x09\xd9\xb5\x18\x5a\x16\x59\x4e\xda\x08\xd3\x25\x3d\x57\xc4\xdb\x0e\x08\x80\xb4\x29\xad\xf5\xd8\xd7\xd5\xba\x2f\xed\x85\xfa\x52\x1b\xa9\x16\x4c\xcf\x15\x69\xa3\x17\x46\x19\x91\xbf\x84\xa1\xdf\xf4\x29\x8a\x3c\x26\x5d\xaa\x42\xd3\x22\x56\x09\xdd\xff\x3b\xc7\xa9\x7f\x33\xc7\x33\x9c\xe6\x22\xd3\x16\xef\xfb\x37\xdf\x0f\x7e\x7f\xdc\xff\x9b\xe3\xee\x6b\x18\xfa\x47\x77\xa3\xe8\x69\x30\x80\x21\x5c\x2f\x7b\x81\x3e\x80\x1b\x18\x05\x01\x3c\xfe\x02\x35\xc5\x46\x71\x3f\x18\xb4\xfc\x6b\x32\x73\x0a\xd2\xcf\x3f\xda\x0f\x4d\x4e\x43\xdf\xf1\x0e\x00\xab\x42\x5a\xdf\xb0\x24\x8e\xa9\x30\x4e\xa9\x66\x08\x7f\xa3\xcd\xfe\x44\x18\xfa\xfb\x67\xa7\x5f\x56\x8d\x62\xb9\x16\xb9\x6d\xd9\xd8\x79\x61\x56\x4c\x7a\xa5\xf2\xa4\x3d\x29\xdd\x79\x39\x0d\x7f\xae\xf6\x4d\xf9\xee\xe3\xf8\xcb\x4f\x0f\x63\x74\x0e\xec\xbc\x6b\xd9\x0f\x0f\x77\x1f\x83\xcf\x1d\x43\xec\x90\x18\xc1\xe6\x8f\x43\xa6\x02\x3f\x18\xfd\xcf\xfa\xc9\x64\x22\x6e\xc5\x3b\xf5\xc7\x6e\xf1\xd6\xd3\x53\xaf\x8b\x9e\xa4\xaf\x8f\xd0\x72\x0b\x4d\x77\xdf\xc8\x4d\x93\x96\x0f\x1f\x9c\x94\x64\x2c\xca\x55\x37\x20\xb5\xd4\x95\xc8\xe5\x3f\xfb\x4d\xf2\x3a\x28\x9a\x58\x92\x7e\xac\x89\x59\x26\xf4\xaa\x89\x60\xa5\x98\x32\xb2\x8b\x10\xff\x1c\xa3\xd7\x76\x6a\x12\x8b\x74\x32\xc5\xae\xb7\x1d\xd2\xc4\x25\xa5\x71\x7a\x37\xa5\xf7\x48\x53\x97\x44\xd3\x91\x98\xc4\xe8\x18\x0d\xf0\xd4\x09\xc5\xf6\xf3\x46\x76\xd7\x0a\xb4\x13\xc0\xf4\x5c\x6a\x7c\xb3\x75\x6e\x6c\xec\x6a\x91\x74\xc5\x60\x5f\xc8\xe2\xab\xe1\x39\x2e\x24\x56\x55\x91\x7c\x9b\xad\xb4\xdc\x42\xdf\xa9\xe6\x1d\x66\x6c\x21\xca\x72\xe0\xd9\xc8\x04\xa3\xee\xa2\x01\xc0\x9c\x32\x2a\x92\x66\x3d\x80\x7f\xa6\x40\x14\xcd\x2c\xe2\xd4\x84\x28\x6a\x46\xf8\x6d\x97\x2f\x04\xa4\x77\xc4\x77\xbd\x5d\xef\xbf\x00\x00\x00\xff\xff\x0d\xa1\xdf\xcb\x3c\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/base.json"].(os.FileInfo),
		fs["/cb.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
