// Code generated by vfsgen; DO NOT EDIT.

package istio_assets

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

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/istio.yaml": &vfsgen۰CompressedFileInfo{
			name:             "istio.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 2069,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x0c\x7c\xb7\xd6\x4a\x93\x62\xab\x9b\x90\x7a\x51\xb7\x71\x23\xc4\xda\x5c\x83\x31\x35\xeb\x30\xa1\x48\x82\x1c\x29\x76\xb6\xf9\xef\x05\x25\x59\xd6\xfa\x03\x88\x6f\xe2\xf0\xcd\xa3\xe6\xcd\xf0\x11\xad\x7c\x24\xe7\xa5\xd1\x09\x48\xcf\xd2\x44\x2b\xd4\xef\x28\x85\x32\x55\x11\x49\xf3\xa5\x8e\x57\xc4\x18\x8f\x5e\xa5\x2e\x12\x98\x07\xc8\xa8\x24\xc6\x02\x19\x93\x11\x80\xc6\x92\x12\x28\xc9\x3f\x8f\xbc\x25\x11\x42\xf5\x8e\x70\x1c\x47\x37\xd1\x74\x3c\x82\x66\x3f\x33\x4a\x8a\x6d\x00\x00\x94\xac\xfc\xc2\x14\x94\x40\x36\x7b\x58\xcc\x97\xcb\xf9\xe3\x6c\x04\x80\x15\x9b\x05\x2b\x9f\x00\xbb\x8a\xba\xbc\xd9\xc6\xa2\x6e\x09\xbb\xa8\x30\x9a\x9d\x51\x99\x42\x4d\x4b\x12\x95\x93\xbc\x9d\x69\x5c\x29\x2a\x7a\x8c\x97\x05\x09\x74\x73\xfd\x42\x82\x8d\x6b\x8f\xa5\x1d\xe8\x07\x2a\x4f\x4d\xc8\xd1\x9b\x93\x4c\xa9\xb5\x7f\xe5\x79\x96\x39\xb3\xa2\x9e\x03\x40\x36\xe9\x54\xdc\x1a\xcd\x28\x35\xb9\xb4\x28\x24\x4b\xa3\x51\xcd\x74\xfd\x88\xce\xb7\xc4\x93\x4e\x87\xf9\x32\x9f\xdf\x3f\x2d\x66\x79\xfa\xf7\xf2\xfe\xdf\xa7\xec\x2e\xcd\xbf\xdd\x3f\x2c\x9a\xc8\x9f\x69\x9e\x36\x60\x80\x1a\x55\x45\x09\xfc\xd7\x2d\x01\x7e\x8e\x8f\xa0\xe3\xe4\xe7\x58\xa8\xca\x33\xb9\x27\x59\x8c\x93\x71\x89\xe1\x7b\xfc\xf1\x31\x02\x90\x25\xae\x29\xab\x94\xea\x44\x85\x54\xbd\xe1\xd6\x07\x6d\x24\x63\x41\xea\x4c\xbd\x6b\x54\x8a\xb6\x67\x37\x99\x02\x4b\xd2\x95\xbe\x76\xe4\xbb\x05\x40\x89\x9b\x07\xb2\x4a\x0a\xf4\x09\xc4\x5d\xd0\x1a\xc7\x3d\x62\xa7\x81\x67\xe4\xca\x4f\xc2\x5e\x5f\x5f\x58\x24\x10\xdf\x4c\xaf\xa6\xfb\x98\x33\x6c\x84\x51\x09\xe4\xb7\x59\x1f\x65\x74\x6b\xe2\xec\x08\xbf\x63\x7f\x66\xb6\x57\x07\xbc\x5f\x2f\x21\xfd\x7a\x8a\xd1\x1f\x30\x5e\x5f\xff\x76\x01\xe5\x1e\xbd\xe3\x64\x75\xc8\x18\xdf\x5c\xc6\x39\xc4\xf7\xac\xc2\x4e\x50\xf9\xc9\x31\xfb\xcd\xf4\xf7\xe9\x25\x22\x0c\xf1\x43\xf6\x77\x69\x5f\xa5\x3e\x75\xc0\x1f\xd7\x71\x7c\xc9\x01\x3d\x9e\x7e\x19\xa3\xa3\xa1\xb3\x52\x19\x3e\x7b\x3b\x4f\x4c\x5d\x33\xfa\x09\x0c\x6c\xea\x4b\x63\x5c\x93\x96\xa9\xb1\x9c\xc9\xea\x5d\x44\x61\x4a\xac\x33\x9b\x6e\xdc\xcf\xe7\x05\x4c\x7d\x75\x90\x09\x50\x79\x5a\x74\x46\x37\xdb\x88\x67\xd4\x6b\xfa\x26\x15\x93\x1b\xf8\x03\xe9\xda\x6c\x53\x21\xc8\xfb\x3b\xb3\x5e\x92\xab\xa5\xa0\xa3\x5a\x7b\x38\xc0\xb3\xf1\x1c\x7e\x42\xbc\x6e\xd1\x15\x3e\xb4\x33\xda\xaf\xfc\xd6\x33\x95\x91\xaf\x45\xd4\x5d\xfd\x48\x19\x81\x6a\x74\xae\xd7\x2c\xec\x3f\x44\x16\x95\xac\xfb\x63\xc3\xd5\x65\x72\x35\xaa\x04\xe2\xa9\x1f\x36\x6d\x45\x3e\x81\xfd\x14\xb2\x0c\x7d\xdf\x63\x58\xf9\x25\x31\x4b\xbd\xf6\x7b\xb2\xb2\xb1\xe9\xce\xd8\xbe\xe7\xdf\xd3\xbb\xe0\xcb\x72\x43\x9d\xa7\x5e\xd2\xa2\x36\xed\x50\xe8\xb2\x52\x2c\x6f\xdb\x7a\x97\x95\x6d\xab\xec\x34\x63\x52\x54\x12\xbb\x73\x9e\x75\xe2\x7c\x3b\x78\x67\x3e\x83\x67\x87\x42\xea\xf5\x41\x42\xdf\xb3\xb0\x1d\x7a\xde\x5e\x8d\x26\xd4\x7e\xee\x24\xc2\xa2\x68\x26\x7c\xd0\xd5\x16\xf0\x89\xc6\x26\xc3\x7b\xf5\x79\xf9\x1b\x35\x8b\xd3\x7f\xfc\xf2\xc6\xbb\x47\xe1\x87\x74\x9e\x27\x16\x1d\x6f\x27\x2f\x6f\xbc\xeb\x9b\x22\xef\xf3\x33\xba\x36\x24\xff\x07\x00\x00\xff\xff\x62\x21\x57\x0a\x15\x08\x00\x00"),
		},
		"/without-istiod.yaml": &vfsgen۰CompressedFileInfo{
			name:             "without-istiod.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 1867,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x51\x6f\xab\x3a\x0c\x7e\xe7\x57\x44\x7b\x87\x95\xdd\xf5\x6a\x97\xb7\x6a\xea\xd5\xad\xee\xaa\x83\x46\xb7\xd7\x23\x37\x78\x34\x5b\x48\xa2\xc4\xd0\xb2\x5f\x7f\x14\xa0\x94\x75\x45\x5a\xf7\x34\xec\xcf\x9f\x1d\x7f\xb6\x0b\x46\xbc\xa2\x75\x42\xab\x84\x09\x47\x42\x47\x5b\x50\x9f\x20\xb8\xd4\x55\x1e\x09\x7d\x5b\xc7\x5b\x24\x88\x83\x0f\xa1\xf2\x84\xad\x3c\x24\x28\x91\x20\x07\x82\x24\x60\x4c\x41\x89\x09\x2b\xd1\xed\x02\x67\x90\x7b\x53\x7d\x24\xbc\x89\xa3\x79\x34\xbb\x09\x58\xeb\x4f\xb5\x14\xbc\xf1\x00\xc6\x4a\x92\x6e\xad\x73\x4c\x58\xba\x7c\x5e\xaf\xb2\x6c\xf5\xba\x0c\x18\x83\x8a\xf4\x9a\xa4\x4b\x18\xd9\x0a\xfb\xb8\xe5\xc1\x80\xea\x08\x7b\x2b\xd7\x8a\xac\x96\xa9\x04\x85\x19\xf2\xca\x0a\x6a\x96\x0a\xb6\x12\xf3\x01\xe3\x44\x8e\x1c\xec\x4a\xbd\x23\x27\x6d\xbb\xb4\x78\x06\x62\xcc\xe2\xde\x0a\xc2\x85\x31\xff\x6d\x36\x69\x6a\xf5\x16\x07\xaf\x28\xa1\xc0\xb4\x92\xb2\xaf\x9c\x2d\xe4\x1e\x1a\xe7\x0b\x10\x04\x39\xca\xcb\xa4\x05\x48\x89\xcd\x94\x8f\xd0\x73\x74\x5e\xa1\x0a\x8b\xae\xff\x60\xac\x84\xc3\x33\x1a\x29\x38\xb8\x84\xc5\xbd\xd1\x68\x4b\x03\x22\xec\xdb\xed\x08\xa8\x72\xa1\xf7\xf5\x9e\x0e\x98\xb0\x78\x3e\xbb\x9b\x9d\x6c\x56\x93\xe6\x5a\x26\x6c\xf3\x98\x0e\x56\x02\x5b\x20\xa5\xdf\xf0\x47\xf6\x1d\x91\xb9\x3b\xe3\x7d\xb8\x86\xf4\xe1\x12\xa3\x3b\x63\xbc\xbf\xff\xeb\x0a\xca\x13\xfa\xc8\x49\xf2\x9c\x31\x9e\x5f\xc7\x39\xc6\x0f\xac\xdc\x84\x20\x5d\xf8\x9d\x7d\x3e\xfb\x7b\x76\x4d\x13\xc6\xf8\x31\xfb\xa7\x30\x1f\x42\x5d\x4a\xf0\xcf\x7d\x1c\x5f\x93\x60\xc0\xe3\x97\x31\x1a\x66\xee\x0d\xa4\xf3\x43\x67\x84\xd4\x34\xb5\x00\x17\x86\xae\x9d\xfb\x84\x8d\x0e\xc1\x6d\x7b\x1a\xc2\x8e\xa8\x5d\xea\x70\xfb\xc9\x83\xb6\xca\x43\x3f\xea\xd3\x51\x1e\x53\xdf\x7d\x89\x63\xac\x72\xb8\xee\xcf\xc8\xf2\xc0\x77\xa0\x0a\xfc\x57\x48\x42\x3b\xaa\x0d\x55\xad\x9b\x05\xe7\xe8\xdc\x93\x2e\x32\xb4\xb5\xe0\x38\xf9\x4c\xff\xb7\xd3\x8e\x7c\x0d\xfc\xa3\x01\x9b\x3b\x2f\x65\x74\xfa\x72\x8d\x23\x2c\x23\x57\xf3\x88\xcb\xca\x11\xda\x48\x6a\x0e\x32\x98\xd2\x99\xb8\xf9\x1f\xd1\x80\x14\xf5\x90\xd7\xaf\x2d\xa1\xad\x41\x26\x2c\x9e\xb9\xb1\x60\x5b\x74\x09\x3b\x4d\x20\x09\xaf\xf9\x09\x43\xd2\x65\x48\x24\x54\xe1\x4e\x64\x65\x7b\x05\x57\xd9\x66\xf5\xeb\xf7\xfa\x65\xf3\xb2\x78\xf2\x67\x4f\x1c\xb0\x3f\x59\xd7\xe8\xd3\x85\x7d\xed\x73\x59\x49\x12\x8f\xdd\x6b\xb3\xca\x74\x6f\xec\x3b\x4c\x28\xb1\x44\xb2\x13\xc7\xea\x42\x72\x33\xba\xe1\x17\x04\xf8\x86\x27\x0b\x5c\xa8\x62\x32\xc0\xfb\xbd\xe4\xdd\x52\xb4\xa6\xee\xdf\x63\x83\x20\xcf\xdb\xd9\x1e\x69\xda\x01\x7e\x20\x6b\x32\xde\xa8\x9f\x37\xbf\xed\x65\x3e\x51\xf2\xfb\x9e\x8e\xbf\x06\x6f\xc2\x3a\x0a\x0d\x58\x6a\xc2\xf7\x3d\x05\x7f\x02\x00\x00\xff\xff\x8d\xa5\xcc\xfc\x4b\x07\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/istio.yaml"].(os.FileInfo),
		fs["/without-istiod.yaml"].(os.FileInfo),
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
