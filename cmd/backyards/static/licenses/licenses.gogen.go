// Code generated by vfsgen; DO NOT EDIT.

package licenses

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

// Licenses statically implements the virtual filesystem provided to vfsgen.
var Licenses = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/LICENSE-v1.txt": &vfsgen۰CompressedFileInfo{
			name:             "LICENSE-v1.txt",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 18529,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x7c\xcd\x72\xdb\x38\xf6\xef\x3a\x7a\x8a\x53\xd9\x8c\x5d\xc5\x68\x3a\x33\xf7\xde\x9a\x8f\x15\x2d\xd1\x31\xab\x65\xc9\x97\xa2\x92\x71\xef\x20\x12\x92\xd0\xa1\x00\x0d\x00\xda\xd1\xac\x72\x9f\xe1\xae\xe6\x5f\x3d\x4f\x33\x6f\x92\x27\xf9\xd7\x39\x00\x48\x50\x96\xdc\x5d\x93\xff\x2a\x96\x44\x02\x07\xe7\xf3\x77\x3e\x90\x74\xab\x39\xdf\x73\x69\x61\xa3\x34\x2c\xdb\xc3\xa1\x39\x0a\xb9\x85\xec\x89\x35\x2d\xb3\x42\x49\x78\xd0\xaa\x6e\x2b\x6b\x46\xa3\x6c\xb3\xe1\x95\x15\x4f\x1c\x6a\x66\xf9\x5f\xe0\x3d\xcc\xd5\x13\xdf\xaf\xb9\x86\x3f\xfc\xf0\xfe\xcf\xa3\x51\x79\x97\x2f\x21\xfb\x98\xce\x56\x69\x99\x2f\xe6\x90\x7e\x28\xb2\xec\x3e\x9b\x97\x70\xf5\xed\xeb\x2f\xdd\xa7\x6f\x5f\xff\x75\x0d\xf9\x12\x52\x98\x65\x1f\xd2\x19\xcc\xf2\x49\x36\x5f\x66\xd1\xe3\x37\x59\xf9\x29\xcb\xe6\xf0\xb8\x58\x41\x3e\x9f\xe6\x1f\xf3\xe9\x2a\x9d\xcd\x1e\x21\xbf\xa5\xef\xd2\xc2\x3f\x9d\xcf\x3f\x40\xb9\x80\xbc\x84\x9c\x9e\x2e\x60\xf1\x69\x0e\x93\xf4\x21\x9d\xe4\xe5\x63\x02\x8b\x62\xf0\xca\xaa\xbc\x5b\x14\xf9\x4f\xd9\x14\x5f\x4a\x27\xff\x77\x95\x17\x19\x94\x77\x59\x47\xc2\x62\x0e\x37\xd9\x5d\x3a\xbb\x85\xc5\xad\x5b\x6f\xb2\xb8\x7f\x48\xe7\x8f\xb8\xd2\xa2\xf8\x90\xce\xf3\x9f\xe8\x6c\x49\x47\x23\xbe\x9e\xcd\xcb\xbc\x7c\x84\xdb\x45\x01\x9f\xee\x16\xcb\x0c\x6e\xb2\x79\x76\x9b\x97\x6e\xe7\x49\x49\xbf\x5c\xbd\x7d\x5c\xac\xde\xe2\x42\xdf\xbe\xfe\x32\x59\x2d\xcb\xc5\x7d\x56\x10\x2f\xd2\xf9\x14\x6e\xd2\xf9\x4f\x69\x0e\x93\xd9\x62\x35\x85\x9f\x8a\x72\x0c\x57\x6f\xe3\xef\xde\x5e\x27\xf0\x29\x2f\xef\x20\x2f\x97\xf0\x50\xe4\xf3\x49\xfe\x90\xce\xe0\x61\x96\x4e\x32\x24\xf6\x66\xb5\xcc\xe7\xd9\x72\x09\x69\x09\xc5\xbf\xff\xdf\x8f\x8b\xc9\x4f\x39\xfc\xfb\xbf\x4a\x78\xff\xee\x8f\x09\xbc\xff\xe1\x4f\x7f\x82\x9b\xd5\x34\x7d\xc8\x96\x65\x02\x77\xab\xf9\x87\xb4\x78\xf4\x2b\x16\xd9\xf2\x21\x9b\x94\xc8\x13\x3c\xcc\x43\xb1\x98\xae\x26\xe5\x12\xc9\xfa\xfd\xa2\x80\x65\x56\x7c\xcc\x27\xd9\x12\x26\x8b\x8f\x59\x91\x4d\xe1\xe6\x11\x48\xd2\x9d\xbc\x88\xcf\xa7\x4b\x21\xd3\x4a\xda\x6d\x9a\xdd\x2f\xf0\x89\x48\x33\xfc\x1e\xd0\x2f\x0f\x9f\xee\xf2\xc9\x1d\xdc\xa7\x3f\x66\x4b\x22\x23\x9d\x4c\xb2\x87\x32\x9d\xbb\xe3\x0d\x37\x84\x14\x26\x8b\xf9\x34\xa7\xb5\x16\xb7\xb0\x5a\x66\x70\xf5\xf6\xe5\xfa\xcb\xb7\xd7\xe3\xd1\xe8\x26\xbb\x5d\x14\x19\xe4\xf3\x65\x99\xce\x66\xa8\x32\x57\x28\xcd\xf2\x2e\x2b\x3e\xe5\xcb\x0c\x45\x7c\x9f\xcd\x27\xf8\xc3\x6a\x99\x7e\xc0\xfd\xae\x9d\x58\x5f\x2c\x98\xc0\xc3\x2c\x4b\xf1\x9d\xb4\xc8\x6e\x57\xa8\x92\x45\x96\x4e\xe9\xe9\xdb\xc5\x6c\xb6\xf8\x84\xab\x74\x74\xe2\xe6\x8f\xf1\xc6\xe7\x57\x85\x01\x39\xa8\xf0\x93\xb4\xc4\xc7\x49\x03\x23\x46\x7c\xba\xcb\xf0\x31\xc8\xfe\xf6\x50\xa0\xac\x51\xbd\xef\x1f\x66\x79\x36\x85\x2b\x94\xff\x87\x04\xb5\x2d\xff\x48\x6f\x27\x30\x59\x3c\x3c\xd2\x1f\xd3\x7c\x59\x16\xf9\xcd\x0a\xbf\xbe\x4e\xce\x9a\xd0\x4d\x06\x37\x8b\xd5\xdc\x4b\x37\x83\x32\x2b\xee\x97\x2f\x39\x3f\x0e\xf6\x34\x5d\xc0\x7c\x51\xba\x1f\x82\xe6\xb8\x77\x50\x9b\x3b\xe9\x2c\xe1\x2e\x2b\xb2\x7c\x9e\xe0\x03\xce\xa0\xef\xd2\x8f\x19\xcc\x17\x50\xe4\x1f\xee\x48\x53\x50\x7a\xf8\xfa\x7d\x5a\x66\x45\x9e\xce\x50\xbf\x17\x1f\xf3\x69\x36\x85\xd5\x7c\x9a\x15\x70\xd1\xaf\x8c\xd1\xec\xe9\xd7\x49\xba\xcc\x92\x8e\xa8\xb2\xcc\xee\x1f\x68\xed\xe9\xe2\xd3\x7c\xb6\x48\xa7\x49\x90\x42\x82\xdb\x39\xd6\x20\xfb\x3a\xce\x64\x97\x64\x73\xf3\x48\xaa\x7c\x9f\xa5\x73\x77\x36\x74\x36\x4b\x27\x1a\x54\xf7\xc5\x7c\x99\xdf\xe4\x33\xb4\x7f\xfc\x75\x96\xa7\xfe\x53\xb9\x70\x47\xcd\x4b\x98\x66\xb3\xac\x74\x96\x93\xce\x66\xf0\x90\x16\x65\x9e\x2d\xd1\x57\x40\x91\x4d\xb2\xfc\x63\x46\xcb\xde\x16\x8b\x7b\xc7\xdc\xbc\xc8\x26\xe5\x8c\x28\x44\x75\x70\x9f\xc6\xa3\x91\xf3\x27\x3f\xce\x17\x9f\x66\xd9\xf4\x43\x46\x3b\x7a\x11\xdc\xa5\xa5\x63\x45\xf0\x65\x28\x36\x52\x1d\xb7\x31\xf9\x52\x54\x96\x49\x8e\x2b\xa3\xb9\x2e\x56\x25\x2c\xf3\x0f\xf3\xb4\x5c\x15\x64\x65\x78\x50\x24\xee\x11\x5f\x38\x63\x16\x27\x62\xea\xa4\xd9\x1d\x22\x92\xd7\xd0\xf9\x87\x9d\x49\x33\x90\xad\x83\x68\x30\xa6\x9d\x83\x56\xbb\xa8\xd0\xbf\x8f\x24\x66\xd3\xa1\x91\x4c\x66\x59\x5a\xcc\x1c\xcb\xef\x57\xa5\x8b\x0e\xf4\xca\x74\x10\x3d\xd0\xd2\x23\x0f\x1d\xfc\x71\xec\xe7\xcf\xf8\xdf\x4f\x39\x4a\xa9\xc8\x3e\xa6\xf9\x0c\xd0\xeb\x9d\x1e\x29\x9f\x93\xc6\xe1\xd2\x85\xe3\x34\x1e\xec\x76\x96\x4f\xca\x28\x30\x0c\x5d\xd6\x7c\x0a\xcb\xd5\xe4\xee\xd2\x31\xc7\xa3\xd1\xc8\x8b\x6e\x0c\xf0\xa8\x5a\x60\x55\xc5\x0f\x16\x98\xac\x81\x61\x9c\x06\xbb\x63\xf6\x5c\x68\x86\x2b\x7c\xa6\x52\x52\xf2\xca\xf2\x1a\x6a\x55\xb5\x18\xd4\xe9\xa9\x6b\x30\x18\xd6\x05\xaf\x61\x7d\x84\x1b\x26\xff\xc1\x04\x4c\x1a\xd5\xd6\xc0\x34\x87\x46\x54\x5c\x1a\x5e\xd3\x36\x68\x3c\xcb\xc5\x6c\x3a\x1e\x3e\xb7\xd5\x4c\x5a\x03\x56\x39\xb2\x64\x4d\xff\xee\xb8\xe6\xeb\xa3\xa7\x32\x01\x06\x52\x49\xfe\xa5\x6a\x5a\x23\x9e\x78\x82\x9f\xac\x66\xd2\x6c\xb8\x66\xeb\x86\x27\x20\xb4\xe6\x4f\xaa\xc2\x0f\x61\x57\x5c\x52\x48\x63\x59\xd3\xd0\xb2\xad\xe1\xc0\x40\x73\x66\x94\xa4\xe7\xae\x0e\x5a\x1d\x94\xc6\x73\x30\x8b\x8f\xbf\x03\x7f\xd4\x8d\xd0\x7b\x44\x2a\xcf\xc2\xee\xde\x81\xdd\x71\x38\xb4\xfa\xa0\x0c\xbf\x06\xd9\x12\x26\x51\x1b\xa8\xd4\x41\x70\x83\x7f\xe1\x03\x46\x6d\xec\x33\x9e\xf9\x1c\x0b\x71\x59\xcd\x1b\xf6\x82\x7f\x20\x24\x9e\x51\xe9\x9a\xc9\x8a\xd3\x7e\xb4\x9a\xe5\x7a\x6f\x02\x35\xb5\xc0\x47\x0d\x18\x4e\x48\xca\xee\x88\x3b\x42\x8e\x61\x65\x38\xee\x7f\x6e\xcb\x35\x3f\x2a\x59\x77\x8b\xd1\x5a\xf8\xe1\xc1\x1d\x04\x9e\x45\xd3\xc0\x9a\xc3\x9a\xa1\x7c\x94\x04\x06\x86\x1f\x98\x46\x4e\xb0\x00\xdc\xc6\x64\x5e\x33\xc7\x4f\xb4\xa5\xd1\xe8\x61\x55\x3c\x2c\x50\x8b\x26\xad\xb1\x6a\xcf\x35\x08\x03\xac\x69\xd4\x33\xaf\x91\xe3\xc8\x65\xdc\xa7\x13\xfd\x39\xe2\x10\x0f\x0a\x69\xb9\x96\xac\x01\xfe\x85\xed\x85\xa4\x27\x12\xb0\xdc\x58\xe4\x3c\x92\xcb\xfb\x37\xd5\x06\x4c\x5b\xed\x2e\x32\xd7\xee\xb8\xd0\x50\xa9\xfd\x81\x59\xb1\x16\x8d\xb0\x47\xfa\x9e\x36\x51\x07\xd4\x12\xf7\x25\x71\xf8\xa8\x5a\x0d\x07\xff\x7e\x02\x86\xeb\x27\x51\x71\xb7\xd2\x41\xab\x8a\x1b\xc3\xcd\x18\x4a\x64\x57\xbc\xcb\x74\x20\xba\x3d\x3b\x82\x54\x16\x99\xd8\xe2\x41\xf1\x50\x4c\x1e\x41\xd9\x1d\xd7\x41\x61\x12\x10\xb2\x6a\xda\x1a\xcf\xb4\x6e\x2d\xbd\xd0\x88\xbd\xb0\x8e\x5b\x9e\x08\x84\xbd\x8a\xe8\xdf\x73\x5d\x09\xd6\xe0\x8a\x8e\xf9\x5e\x5e\xe4\xc8\x46\xa3\x69\x36\xcb\x3f\x66\x45\x7a\x33\xcb\x96\x63\x6f\xc7\x9f\xa5\x7a\x6e\x78\xbd\xf5\x26\x2c\x24\xf1\xbf\x52\xad\x76\xda\x71\xe0\x7a\xa3\x9c\x3e\x23\x79\x4b\x7f\xda\x64\x68\x86\x78\x9a\x4a\x73\x14\x7f\xa7\xca\x4a\xfb\xc3\x3c\x2b\xfd\x99\x34\x9d\xb5\x76\xa7\xb4\xd9\x89\x03\xd4\xbc\x11\x4f\x5c\xbb\x73\x20\x25\x87\x56\x9b\x96\x49\x8b\x9f\x49\xbe\xc1\x69\x20\xb3\x3a\xcd\xf6\xd4\x90\xba\xab\x4d\x47\x0d\x5c\x31\x44\xae\x53\xb7\x28\x9a\xa7\x73\xdd\xcb\x76\xfd\x33\xaf\xac\xdf\x43\x83\x16\xdb\x9d\x35\xe1\x90\x9d\x0e\x4e\xd0\x66\x6b\x2e\x2d\xf2\x2e\x97\xb4\x03\xed\x7b\x55\xf3\x8d\x90\xe8\x9e\x78\xa3\x9e\xaf\x4f\x0e\x6d\x76\xe8\x1c\xd4\xb3\x44\x0d\x76\x6b\x27\x60\x85\x6d\x78\xaf\x3d\xdc\x10\x4f\x49\xfd\x22\xf2\x4c\x2c\x59\x7c\x1b\x1f\x6e\x1a\x5e\xd9\x96\x35\x28\xd7\x03\xd7\xf6\x18\xe8\xb5\xce\x62\x83\xae\x6a\x6e\xd5\x89\x1b\xf4\x0e\xaf\xf7\x86\xdd\xd9\x62\xcf\xe7\xd6\xa3\x65\x22\x2f\x87\x36\xf7\x82\x3e\xe8\x5e\x6a\x8e\x5e\x18\x3f\xb7\xf2\x44\x1a\x03\xf5\x8e\xf4\x92\x3c\x5a\xa5\x0e\x3c\xb8\x37\xaf\x86\xe3\xd1\x88\xc2\x64\x1f\xa7\xbd\x16\x52\x08\xa1\x60\x72\xaa\x8f\xf4\xb3\xe6\xa4\xf9\x6a\x6d\x99\x90\x41\x13\x5f\xe1\x18\x12\xac\x34\x52\x82\x0f\xee\x99\xe5\x5a\xb0\xc6\xe0\x53\x4f\xa2\x3e\x13\x6d\xbc\x0e\x5e\x52\x3a\x7c\xcd\x78\x47\xd2\xab\x5c\x27\xc0\xe4\x9c\x6d\xe2\xce\x9d\x29\xe0\xd1\x4e\x22\x1f\x5a\xe7\x7d\xa0\x0c\xb5\x35\xf1\xd6\x62\x77\xcc\xa9\xa7\x17\xbe\xda\x38\x09\x1d\x78\x25\x36\xa2\x62\x4d\xe3\xe5\xcc\x6b\xa7\xc8\xc2\x40\xda\xbb\x5c\x3c\x47\xf0\xce\xa8\xd1\xb6\x71\xf4\x7c\xe6\xfc\xd0\x07\xb2\xa6\x81\xfb\xdf\xca\x16\x77\x3c\xe4\x86\x6a\x6d\xe4\xa3\x22\x15\xe2\x5f\xd0\xf1\xe0\xc2\xf8\xab\x93\xf9\x6f\xd1\xe8\x84\x28\x51\x1b\x78\xde\x89\x6a\x47\x82\xf6\x9a\x6c\xac\x16\x95\x6d\x8e\xa0\x39\x3a\xd8\x97\xc4\x8d\x21\x97\x70\x60\xda\x8a\xaa\x6d\x98\xa6\xa3\x05\x1a\x49\x0c\x3e\x24\xf4\xc7\x8c\x3c\xee\x5e\xd5\x62\x23\x78\x9d\x80\xe6\xef\x6a\x81\x9b\xad\x5b\x8b\x9f\x6b\x61\xaa\x46\x99\xa0\xc7\x42\xd7\x6e\x13\xb4\xd8\x86\x4b\x9b\xc0\x4e\xa0\xd7\x52\xad\x4d\x60\xcf\x6a\x0e\xec\x89\x89\x86\x10\x81\xd7\xfc\x43\xbb\x6e\x44\x95\x80\x51\x4d\x9d\x80\xda\x6c\xc8\xcd\xa1\x7f\x37\x0c\xc1\x86\xd9\x31\x8d\x5b\xa1\x86\x7a\x14\xa2\x9d\x2c\x7b\xde\x3e\xb3\xe3\x78\x34\x4a\x9b\xe6\x44\x20\x9a\xd5\x7c\xcf\xf4\x67\x93\xb8\xbf\x41\xb2\x3d\x91\xa6\xb6\xca\x99\xa1\x54\x96\xd4\xf3\x80\x8c\x93\x16\x63\x33\x12\xd5\xb3\x21\x28\xc7\xa1\xe3\xac\x7f\x0d\x6a\xbe\x61\x15\x52\x16\xd8\x43\x2e\x7c\xdd\x08\xcb\x31\xb0\xd7\xe3\xd1\x68\xb2\x58\x3c\x64\x85\x2f\x33\xdc\x66\xd9\xf4\x26\x9d\xfc\x18\x1b\xb1\x8b\x4a\xa8\x4e\x43\xca\xc9\x98\xc8\xcb\x54\x8a\x22\xaa\x93\x4e\x67\x9b\xe8\x0f\x3b\xbf\x9b\x10\x72\x33\xee\x40\x08\x11\x95\xb6\x5e\x45\x06\x8b\xd6\x9c\xef\xc9\xdb\xad\x79\x8f\xcc\x50\x65\xfe\xde\x0a\x1f\x5a\x08\x5b\xbc\xd0\x6a\xd3\xd2\xfa\x9b\x16\x8d\x29\xd0\x8b\x6c\xba\x84\x10\x82\xd5\xbb\xa3\xb6\xb2\xe6\xda\x58\x72\x58\x43\x04\x1c\x6f\xf4\xed\xeb\x3f\x0d\xb1\x6f\xcb\x1c\x02\x43\xd5\xa6\x57\x49\xd1\xf9\x17\x94\x80\x69\x8e\x3d\x48\xe3\x35\xb4\x07\x25\x5d\x3c\x24\xaa\xd0\xd9\x7d\x0f\xd7\xd0\xe3\x1e\x84\xc4\xed\x13\xe8\xcf\xa6\xb9\x69\x1b\x04\x2f\x08\x98\xfa\x4f\x1b\xce\x6c\xab\x39\x88\x9a\x33\x03\x1e\x90\x44\x3b\x80\xe6\x5b\xa6\x89\xa8\x0b\xcc\x4a\x22\x80\x30\x08\x27\x4a\x47\x2a\xa8\x36\x27\x81\x34\xca\x02\xf0\xec\xe8\x1b\x6f\x39\xaf\xd7\xac\xfa\x4c\x55\xa8\xa0\xb3\x84\x94\x4e\xbd\xd4\x46\x73\x0c\x52\x2e\xe4\xd6\xfc\xc0\x25\x86\xf2\x86\x1c\x11\xae\xe6\xcf\xd1\x1b\xf3\x31\x0e\xbe\xad\x15\x8d\xf8\x47\xc7\xe7\xb0\x2d\xd9\x2b\xb9\xb1\x3d\x4a\xc2\x15\x22\x4f\xe8\x26\x01\x1f\x62\x35\x19\x60\x41\xb2\x7c\xcf\xab\xfd\xd0\x9d\x1c\xa1\xf2\xf1\xd9\x60\xe6\x7c\x31\xa1\x7a\x89\xb1\x14\xb9\x54\xa8\x99\x65\x48\x0e\x61\xd1\xd6\xb0\x2d\x27\xa6\x31\xc9\x9a\xa3\x15\x95\xa1\x07\x88\x75\xde\x1b\xbf\x94\x55\x17\x0c\x84\xed\xc0\xe7\x19\xd8\xc9\xe5\x0e\xd1\x16\x9e\x1f\xc1\x2e\xba\x55\x43\x3b\xe0\xeb\xeb\xd6\x08\x89\x2a\xe7\x5f\x34\x08\xb9\xd0\x91\x0f\xe8\x70\x18\xff\x99\x1d\xf1\x25\x25\x8f\x7b\xf1\x0f\xef\x75\x3c\xa4\x43\xc3\x6d\xc0\x21\x30\x7b\x44\x81\x1d\xb4\xda\x88\x86\xf0\x03\x73\x11\xa4\x35\x5c\x9b\xcb\xa7\xf1\xee\x9d\x3f\x71\x8d\x8a\xe2\xd6\x72\x2e\xde\x6a\xf4\x6a\xb8\xaa\xe6\x95\x92\xc6\xea\xb6\x72\xde\x6c\x96\xdf\xe7\x25\x39\xb3\x31\x10\x5e\xf7\xb9\x0a\x12\x7c\x02\x67\x42\xe2\x13\x32\xcc\x80\x49\x0e\x5c\xef\x85\xb5\xc3\xec\xe5\x3f\x36\x0c\x1f\x03\xf6\x4c\x4a\x82\x6e\x36\xf2\x12\x0e\x3c\x13\xeb\xd6\xc7\xb3\x61\x3f\x06\x4f\x27\x8a\xa4\x9e\x25\x77\xc0\xfb\x7c\xc2\xf7\xdb\xc8\xa3\x45\x9b\x06\x4c\xbb\x36\xfc\xef\x2d\x9a\x84\x4f\x63\x29\x9e\xab\x8d\x77\x10\x0d\xaa\x84\x17\x15\x3a\x0f\x5c\x65\xcf\x6b\xc1\x7c\x84\x6f\xf0\x08\xef\x7c\x5e\xbb\x67\xc2\x63\xad\x41\x78\x1f\x8d\xca\xac\xb8\xff\x2b\x06\x19\x97\x4b\xf5\xfd\x00\x61\x80\x77\xf5\xfe\x8d\x56\x7b\x07\xee\x9d\x09\xa1\xb2\x26\xf0\xbc\xe3\x14\x47\xf9\x97\x43\x23\x2a\x61\x29\xa5\xd8\xfb\xbf\xaf\xfa\x27\x91\x95\x94\x37\x79\x15\xf7\x68\xeb\x7a\x5c\x86\xcc\xd7\x1f\x23\x68\x46\x87\xbc\x0c\xb4\xd2\xeb\x48\x02\xcc\x40\xa3\x10\x97\x1a\x47\x4a\x6b\xf8\x6b\xaa\x7a\x51\xbd\x06\x91\xf4\xac\x03\x40\x9a\x84\x2b\x38\x0c\x34\x00\x55\x47\x58\x83\xa0\x83\x13\x88\xd1\xbc\x07\xb1\x62\x4f\xdc\xb7\xdc\xf3\x6d\x00\x96\xb6\xe2\x09\xed\xfc\xa8\x5a\x1f\x43\x07\x50\xaf\x11\x21\xf7\xf5\x18\xf9\xa5\x37\x1d\x8f\x46\xb7\x0a\xc3\x2c\xf9\x3a\x4f\x1f\xa1\x65\xe2\xbf\xd0\x5d\x12\x6e\x63\x31\x26\xb4\x9c\xcb\xa1\x2a\xce\x3c\x2a\xfd\x15\xce\xb9\x90\xdb\xa7\xa1\xf4\x45\x0f\xf0\x7b\x4c\x3e\x08\x5b\x5b\x2e\x1d\x82\x41\x71\xa3\xcb\x7a\x25\x7e\x35\xe2\x33\x47\x38\x65\x12\xa8\x94\xe6\x50\xb7\xfb\xc3\x49\xa0\xbc\x76\xbe\x1e\x29\x4f\x80\x59\x5a\x2b\x62\xf9\x49\x94\x48\x80\x0b\xd2\x45\xcd\x6d\xab\x89\x2b\x35\x37\x56\xab\xe3\x6b\x8e\x79\x3f\x30\xba\x4b\x09\xea\x78\x34\x5a\xba\x6c\xc5\xc0\x1f\x12\xf8\x63\x02\xff\x2b\x81\xff\x93\xc0\x9f\x12\xf8\x73\x02\xef\xdf\xd3\xcb\xef\xff\xe0\xb9\x6c\x5a\xfd\x84\x26\x73\x4e\x46\x1e\xb6\x0f\x9c\xca\x68\x54\xf6\xa5\x98\x4e\xf3\x3d\xa0\x5c\x33\x23\xcc\xcb\x97\x50\xb7\x31\x20\xe3\x2f\x5a\x1d\x59\x63\x85\xf3\x22\x0e\xd7\x6e\x78\x17\x23\xf0\x4f\x52\xe9\x35\xef\x6a\x47\xe8\xe6\x70\x25\x4c\xb8\xed\x33\xe7\xb2\xab\xea\x0d\xa1\x3f\x7d\xdb\x18\x15\x81\x4e\x76\x24\xfd\xb1\xec\x0b\x37\x5e\x50\x28\x67\x72\x0f\x27\x24\x7a\x15\x66\xe8\x40\x04\xd1\xd7\x21\xb3\x10\x8d\xc3\xb2\x6b\xce\xdc\xc3\x95\x32\x54\x38\xa8\x5a\x02\xeb\x1e\xa8\x5c\x48\x18\x7d\xdd\xd0\xa3\xb2\x9d\xa2\xa2\x54\x02\xba\x95\x92\xfe\x60\x15\x02\x54\xfa\xb3\x52\x07\xea\x6a\x0e\xaa\x3f\xf6\xe8\x52\xe7\x5f\x31\x06\xf4\xaa\x9d\x9a\x54\xca\x95\x53\x5e\xc6\x86\xd1\x68\x54\x64\xcb\xb2\xc8\x27\xd4\xf5\x18\x8d\x56\xd2\x39\x68\xdc\xeb\x19\xb3\xb6\x28\xc8\x78\xee\x9f\x40\xac\xde\x52\x31\x26\x5d\x89\x38\xf3\xba\xfe\xcb\x68\x74\xc5\xae\xdf\x84\x5c\x26\xa1\x94\x0d\x4d\xc3\x50\x22\x84\xca\x83\xb9\x13\xc3\x7f\xb4\xe7\x7f\xc3\xa5\x8b\xcc\xf4\x0a\x01\xf5\x01\x3c\x0a\x0a\x43\xe4\xf5\x09\xda\x77\x45\xd7\x18\xd5\x58\xb5\x75\x01\xe2\x65\x45\x0f\xdf\x09\x38\xee\xaf\xa3\xd1\xd5\xfa\xfa\x8d\x46\x58\x81\x7c\x92\x5b\x21\x39\x1e\xb1\xe6\x95\xda\x1f\x04\xa6\x73\xad\x6c\x54\xf5\x19\x8f\xe6\x1e\xea\x7e\x40\x43\x17\x86\x19\xc3\xf7\x54\x64\x66\x8d\xc5\x57\x29\xb9\x3a\xfa\xc3\x37\xcc\xe2\x2f\x7b\x2e\x6b\xf0\x55\xcd\x2d\x15\x51\x89\xae\xb8\x7c\x76\xd0\xca\xba\x12\xfa\xfa\xd5\xda\x87\xf9\x4e\x68\x8e\xb2\xb1\xc3\xcc\x17\xc5\x53\x2b\x30\x8a\x9a\xa9\xfc\x89\xb2\x5e\x79\x0c\x24\xe1\x26\xcc\x03\x78\xef\x1c\xbe\x67\x7b\xab\x3a\x16\x55\x42\x57\xed\xfe\xc9\xeb\x8c\x5a\x3f\x51\x0c\x23\x9c\x1e\xed\x6d\x77\x5a\xb5\xdb\x1d\x58\x5e\xed\xa4\x6a\xd4\x56\x54\xac\xc1\x17\xd6\xc7\xc8\xa6\xf6\x9c\x49\x92\x67\x75\xfd\xe6\xbb\x71\xda\x09\x5e\x3e\xad\xdc\xc4\x60\x11\x05\xe9\xeb\x8d\x83\x92\x98\x07\x14\x48\x51\x7d\xfd\xa6\x6e\x11\x9b\xb0\xef\xd4\x70\xb8\x42\xe6\xb1\xcf\x1c\x18\x39\x96\x00\xcb\xae\x83\x7b\x41\x81\x22\x01\xc2\x26\xc0\xbf\x50\xc6\xe1\x01\x08\xff\x62\x91\xcd\xa1\x14\xff\xc2\x8b\xd0\x0a\x07\x22\x92\x9a\x24\xec\x79\x0c\xb7\x3e\x45\x0a\xe0\xbf\x8b\x06\x7d\xc1\xec\x2d\x92\xf1\x36\xf2\x1d\xce\x04\x79\xef\xf7\x36\xc3\x4a\x99\xef\xa8\x38\xf4\x23\x31\x1d\xa5\x82\x97\xaa\xaa\x56\x1b\xa8\x5b\x4a\xa8\x24\xc6\x3f\x52\xfc\xad\x66\x7b\xe0\x5f\x78\xd5\x3a\x0f\xbc\xe3\xd2\x25\x89\x67\x3a\x23\x28\x8d\x90\xd5\x05\x09\xe1\xa9\x4c\xc7\x16\xcd\xb7\x54\x4b\xc2\x04\xb0\x3d\x90\xb6\x70\xef\xd8\x9a\xef\x97\x4d\xe4\x7d\xfc\x8a\x54\x62\x76\xfe\xc2\x57\x05\xad\x0a\x0d\x00\x77\x36\xaa\xbb\x37\x4c\x6e\x5b\xb6\xe5\xe6\x7a\xe8\x17\x83\x99\x90\xec\x3a\x77\x8a\x78\xfb\x7b\x09\x85\xe7\x1d\xe2\x48\x57\x87\x47\xee\x5c\x27\x54\x36\x09\x2d\x1e\x81\x82\xa4\xa8\xab\x2d\x01\xa5\xfd\x5a\x48\xfa\x3a\xf2\x5b\x7d\x8b\x24\xf4\xd1\x84\xab\xcc\x9e\xfc\xea\x32\x70\x7e\xf6\xd7\xa0\xa5\xcc\x80\xe4\x15\x37\x86\xe9\xe3\xa9\x10\x51\x4c\x1b\x74\xd2\x7b\xf5\xc4\x83\x7f\x75\x0e\xc3\x54\xad\xe6\x3e\x7e\x1f\x8e\xa1\x42\x1f\x8a\x68\xbe\x56\x96\xb8\xd4\xd0\x55\x1c\x11\x1f\x35\x7c\xcb\xa9\x98\x40\xaa\x5a\x87\x76\x41\x1f\x6b\xa9\xc5\xc8\x9e\x38\x26\x94\x5c\x3c\xb9\xe4\xb2\xb3\x2f\x6f\x74\x48\xd6\xee\xfa\x8d\xaf\xc9\x0c\x04\xd7\x9a\x28\x45\x59\x1f\x9d\x6f\xf2\xe5\x0a\xdd\xdb\x82\x51\x4d\xeb\xb0\x5d\xfc\x76\xe2\xc0\x8b\xb3\x45\x12\x6f\xac\x2b\xd4\x74\xd8\x30\x2a\x50\x7f\xfb\xfa\x4b\xfa\x90\x53\x15\x80\xcc\x29\xe4\xac\xa1\x12\x7d\x26\xd0\x0b\x79\xd2\x61\xd4\x9c\x22\x77\xdd\x83\x9b\xff\x54\xaf\x90\x1f\xa2\xf7\xbd\x83\x6d\xff\x63\x5d\xc5\xc8\xc4\x9f\x78\xa3\x0e\xde\x8d\x38\x8e\xb9\xf3\x06\xfb\xa8\x79\xc3\x31\xcc\xf6\x4e\x96\x04\x7e\x68\x58\xe5\x92\x0e\x2f\xbd\x20\xe0\xce\x1f\x50\xfe\xea\xbb\x1e\x0c\x93\x20\x52\x15\x5f\x25\x33\x09\x39\xaf\xfe\xf5\xe0\xba\x2e\x23\xb6\xf1\x68\x34\x5f\xc0\x4d\x36\x9f\xdc\xdd\xa7\xc5\x8f\xf9\xfc\x83\x83\xb2\x27\xf8\xcb\x51\x7f\x0c\x85\x7c\x14\x59\x48\xc9\x9c\xab\xa6\xa7\xe1\xa0\x85\xd2\xf0\xac\x85\xb5\x9c\x70\xa8\x39\x53\xa3\xf2\xce\x91\x8e\xcc\xe5\x96\x6d\x11\x9d\xc2\x9a\xcb\x6a\x87\x26\xf0\x02\x7a\x7a\x50\x7c\x74\xdd\x50\xf2\xb6\xd5\x8e\x69\x56\x21\x6b\x5c\xe1\xe7\x72\x45\x29\x32\x7e\x23\xf6\x02\x5d\x69\xaf\xc3\xa8\x8e\xf1\xbe\xa4\x97\xae\x51\xcb\x25\x88\xcd\x90\xa8\x67\x66\xe2\x60\x34\x4c\xa8\x90\x35\xa1\x9a\xdf\x25\xab\xbd\x77\x0a\x0c\x0a\xac\x61\x07\x0c\x48\x28\x57\x04\xa4\x66\x97\x44\x88\x12\x9d\x97\x7c\xe2\x2f\x11\x27\x75\x01\x42\xc6\x17\xc2\xd4\x80\xc4\x73\xa0\x95\xbe\x73\x0d\x00\x60\x6d\x2d\xb8\xac\xf8\x78\x34\x5a\x3d\x4c\xd3\x32\x5b\x9e\x34\xed\xa8\x86\x50\x53\xa1\x27\xaa\x13\xc7\x25\x74\xd7\xcc\xa1\xd8\x75\x84\x75\xbb\x85\x8d\xf8\xc2\xfb\x9e\x33\x1c\x58\xf5\xd9\x60\x86\x61\xc3\x2f\xed\xa1\x66\xd6\xfd\xb1\x45\x3f\x67\x92\x93\x0a\x9e\xab\xed\x3b\xcf\x61\x12\x57\xfb\xe7\xcf\xc1\xce\xc9\xbc\x10\xc7\x3a\xb7\xb3\x39\x71\x7c\x5d\xa7\xc8\xd1\x7d\x02\x11\x50\xc2\x2b\xda\xdf\x75\x5f\xef\xd4\x33\x62\xe2\x04\x85\x1b\x6a\x66\xdd\x0a\xdd\xb1\x56\x81\x62\xc2\x75\xfe\x93\x47\x0c\x98\x1a\xf6\x2d\xdc\x8b\xa3\x0d\xe7\x52\xd7\x53\x4e\xa3\xae\x68\x6e\x0e\x4a\x1a\xe1\x66\x3e\x9c\x48\xab\x1d\x93\x5b\x87\x71\xd3\x87\x9c\xce\x1f\x68\xb0\xea\x9c\x92\x9f\x38\x0b\x03\x6b\xcd\xd9\x67\xd3\x25\x82\x64\x3d\x6a\xf3\xce\x39\x19\xea\x14\x9c\xf8\xf5\x2e\x1b\xf6\x51\xae\x97\xa7\xab\x7c\x52\xc0\x68\x65\x6b\x08\x6a\x51\x10\x0e\xb3\x09\x0d\xa7\xbe\x97\x6b\x09\xa0\xbc\x12\x0a\x9b\x7e\x87\xee\x7c\x5d\xc5\x26\xb8\x40\xdd\xe7\x2b\xae\xea\x16\xe7\x54\x0a\x95\x27\xb4\x19\xe2\x19\x88\x83\x56\xeb\x86\x7b\x7e\x9f\x78\x95\xe0\xb1\xc8\x54\x98\xd8\x93\xb9\x46\xe5\x22\xea\x77\xe1\x92\x68\xde\xe4\xfa\x26\x8b\xf9\x6d\x3e\xcd\xe6\x65\x9e\xce\xf2\xf2\x11\xbf\x09\x4d\x6a\xf2\x73\xad\xe1\xf1\x7c\x4d\x85\xac\x22\x5b\xa0\x15\x82\x4d\x7a\xcb\x23\x06\x1a\xdf\x04\x33\xbc\xd2\xdc\xbe\xec\x2b\x74\x51\xaf\x2f\x07\xa3\xf7\x0b\x7e\xc1\x85\x7f\xe7\x83\xf6\x6b\x55\x8b\x3e\xd6\x5f\x48\xb7\xcf\x4e\xe1\x8c\x01\xde\x16\x43\xba\x03\xe2\xc5\xc0\xee\xa8\xa0\xd3\xb8\x16\x56\xd4\x9a\x47\x18\xdb\x1c\xf1\xe4\xc6\x9b\x3d\x26\x35\x54\xc8\x53\xcf\xb2\xcf\x68\xa8\x3a\xc3\x3a\xa7\x4a\xb8\x98\x8f\xa1\x5b\xea\x77\x71\xa3\x29\x36\xce\x07\xa6\xd9\x56\xb3\xc3\xae\x0f\x25\x88\x1c\x48\x37\x7c\xbc\x7a\xb5\xa8\x7e\x79\xf4\xc8\x9d\xe5\x99\x3b\x7d\xf6\x0d\x9a\x4a\xed\xf9\x20\x7a\x7a\x57\x58\xab\x3d\x13\x32\x89\xda\xbc\x28\x55\xa1\x5a\xd4\xa0\xcf\x92\xce\xda\xcf\x2c\x50\x94\xed\x1f\x1d\x36\x75\x7c\xb8\x77\x61\x21\xbc\x31\xa6\x3c\x1d\xde\x0c\x15\x8a\xea\xf3\xbf\xa1\xb9\x77\xfe\x80\xbe\x87\xef\xda\x13\x83\xe6\x7c\xb7\x0b\xdf\x1f\x1a\x75\xe4\x68\x85\xdd\x77\x6c\xeb\xfc\xac\xd2\xd1\x77\x9b\x8d\x68\x5c\xfd\xb5\x7f\x23\x04\xae\x58\x63\x7f\x77\x21\xaa\x77\xd9\x9a\x4f\x80\x7c\xdf\x44\x28\x57\x0d\x75\x19\x31\x11\x4a\x41\xc5\x84\xb0\x4e\xd8\xf3\x95\x72\xa7\xf3\xc2\x91\x16\x51\x3b\xc3\x50\xea\x86\x06\xdc\x65\x76\x83\xe9\x83\xc0\xaf\x48\x64\xbf\x33\xbf\x52\xac\x1a\x53\xe6\x0d\x91\x88\x08\xea\x38\xbf\x6c\x15\x38\xf4\x4e\xa7\x09\x1d\x6c\xb5\x81\x03\xb3\x74\xfa\x0e\xbd\xfb\xa6\xb9\xc7\xef\x7d\x57\xc3\x97\xf9\x86\xd3\x47\x68\x46\x03\x99\x26\xa1\x49\x18\x90\xb3\xda\x40\xc5\xb5\x8d\x03\xe2\x9e\xc9\x76\xc3\x2a\x34\x30\xfd\xc2\xef\x51\xc7\x8f\xfa\xef\x6e\xc7\xc3\xee\x68\xa8\xda\x70\x88\x8e\x59\x5f\x9f\xd1\x43\x66\x8c\xd8\x52\xb9\x95\xea\xa8\x51\x6d\xcc\xf8\xe2\x98\xd3\x19\xa7\x66\x61\xf4\x43\x0f\xfa\xc7\x61\xb2\xe6\xb5\x23\x46\xbe\x31\xc6\x8c\x91\x5c\x4f\x9c\xe4\xd8\x25\xb9\x2f\x29\x0e\xf5\xac\xbe\x80\x15\x15\xbd\xc2\x8f\x51\x11\xcb\x25\x8b\x28\xc0\x67\x86\x66\xaa\xbb\xee\xcc\x65\xdb\x43\xc4\xde\xea\x8a\x43\xa5\x6a\xb7\x04\xff\x42\xcd\xfe\xd3\x73\xba\x36\x1b\xc6\x3a\xd7\xed\x77\x8f\xf5\x93\x36\x95\x6a\xa5\xd5\xc7\xf1\x68\x94\x3f\x40\x3e\x9f\x66\xf7\x73\x0c\x31\x70\x32\x84\xc0\xb4\x1f\x90\x42\x97\xec\xd9\xb9\xe6\xc6\xd7\x44\x0c\xc4\xc3\x47\x7c\x50\x9e\xec\xd5\x56\xd6\xbf\x5d\xeb\x5f\xf3\x9f\x67\xea\x42\x83\xe9\x4a\x57\x2d\xd9\xa0\xbd\x77\xe5\xb6\x83\x16\xdc\x62\x0a\xec\x15\x24\xaa\x29\x38\x53\x31\x91\xad\x98\x28\xd5\xf5\x73\x8e\x71\xb4\xbc\x0e\xa0\xd6\xa9\x9c\xeb\xbd\x12\xba\xa0\x24\x83\x62\x2e\x03\x8a\xed\xc4\x1e\x4f\x4a\x28\xf2\xa3\x5a\xe8\xd0\xfc\xa2\xae\x93\xe6\x7b\x5e\x87\x3e\xfc\xd9\x66\xd4\x95\xb8\x7e\xd9\xd1\xe2\x8d\x87\x77\x21\x23\x53\x3a\xe4\xf0\x97\xb8\x1a\xaa\x5c\xc2\x82\x54\xf2\x9d\x27\x8d\xea\xe9\x1d\xbe\x24\x11\x9f\x3a\x1b\xfa\x92\xbc\xa5\xdf\x2c\x94\xb6\x06\xa8\x18\x58\xb5\x13\xfc\xc9\xb5\x35\xa9\xd9\x69\x99\xa4\xca\x40\x94\x0e\xe0\x4f\x4a\x8b\xad\x90\xac\x09\x88\xf9\x15\x4d\x20\xe5\xbe\x12\xe7\x18\x80\x88\x8f\xca\x14\x2e\xf0\xf3\x2f\x07\xea\xb7\xb0\xae\xf3\xb2\x89\x63\x49\xd4\x66\xf6\x4a\xc0\x9a\x86\x6f\x79\x4d\x73\x7a\x4e\x46\x75\xd8\xed\xdc\x76\x17\x3b\x88\x71\xdb\x6f\xd3\x6a\x3f\xe3\x15\xb5\xff\xa2\x68\xdb\x77\xce\x61\x71\xdb\x5f\x2c\x18\x43\xa9\x7c\xd6\xdc\x34\x68\x57\x2f\xcb\x8a\x0d\x3b\x19\xc7\x89\x00\xa4\x3c\x76\x2d\xe6\x21\x90\xc4\x5f\x6a\xb6\x67\x5b\x97\xa0\x50\xd2\xf5\xce\x0f\x52\xd0\xbb\x49\x6c\x0a\xb5\xd0\xa1\xc9\x89\xbe\x1e\x3f\x24\x2e\x82\xb1\x86\x9a\x00\x04\x07\xf1\x6f\x0a\xaf\xd4\xc8\xa6\x9f\x28\xe2\x49\xe1\xef\x91\xd1\x76\x09\x34\xca\xd8\x6e\xd0\xc1\x7f\xac\x99\x65\xd7\x49\x97\xb9\x0e\x47\x31\x99\x89\xd1\xcd\x9a\x63\x06\x5a\x3f\x09\x1a\x94\xde\x0c\xf3\x4c\x7c\xd6\x37\x5d\xad\x22\xef\xd3\x01\x27\x65\x3a\x1c\x1f\xc6\x97\x3d\x49\x63\xf8\xb4\x43\xd8\x15\x7f\xd7\xd7\x76\x3c\x37\x09\x8a\xd6\x2d\x77\x19\xe9\x9e\xc9\x9a\x59\xa5\x8f\xd0\xf0\x2d\x6b\xa8\x20\xa7\x5b\xd7\x85\x7d\x11\xdf\x86\xac\x8f\xd8\x8e\x40\x6c\x2f\x89\xa2\x28\x2a\xf5\xb5\x5d\xc4\x28\xdc\x0d\x75\xa0\xb8\x69\xcc\x1d\x4d\x9e\x7c\xa9\xb2\xac\x01\xb6\x47\x3f\x8d\xdf\x53\xf7\xaf\x2b\x9f\x39\xc4\xea\xdb\x0d\xdf\xbe\xfe\x73\x58\x77\x26\xb3\xf9\xdf\x3f\x24\xef\x20\x5b\x15\x8b\x25\x5c\x6d\xc4\xc6\x1e\x81\xb7\x5a\x99\x6b\xd7\xa5\x34\x3c\x1a\xfa\x33\xa0\xe4\x39\xba\x05\x8f\xa9\xed\x70\x30\x92\xe8\x5b\xb4\x11\x05\x98\xd5\x55\xbb\xe0\x23\xab\xa8\x01\x7b\xc2\x80\xe4\xa5\x59\x13\x1b\xfb\x05\x04\x75\x0e\xcf\xf1\x6e\xe3\x42\x9b\x9f\x7f\x91\x3f\xb7\x3a\x4c\xbf\xb8\x8e\x8e\x3b\x44\xe2\x4d\x99\x08\x71\xa6\x42\x8e\x2c\x78\x23\x0a\xb2\x6b\x56\xc3\x86\x09\xbb\x4b\x60\xa3\x59\x5b\x27\x54\x0a\x94\xae\x74\x05\x7b\x61\x30\x5f\xf6\xf9\xeb\x93\x50\x4d\x07\x80\xd0\x20\xa3\xca\x3e\x95\x79\x7d\x73\x3d\x1e\x0c\x8d\x38\x84\x61\x76\xf1\x69\x9e\x15\xcb\xbb\xfc\x61\x0c\xf9\xc6\x4d\xe7\x9e\x69\x27\x7a\xa8\xc5\x6b\x7f\xe7\x00\xae\xf8\x78\x3b\xc6\xf0\xe0\xdb\xd6\xf0\xfe\x07\x5c\xd2\xe7\x28\xdf\xbe\xfe\xb2\x38\x70\x09\x4b\x07\x0a\x82\xef\xa4\x4a\x51\xef\xca\xfb\x29\x17\x73\x0a\x9e\x35\xf7\x63\x8f\xac\xaf\x02\x0e\x8c\xd3\x43\xa6\xb8\x6a\xe3\x87\xcf\x34\xa5\x5c\x17\xbb\x67\x27\x4d\xf9\xd7\x0b\x21\xae\x3c\x11\xc2\xb0\x6b\xcd\x76\x13\x66\xbf\xbe\x78\x0c\x31\x7a\x2c\xe0\xfa\x3a\x41\x05\x50\x23\xad\x46\x28\x60\xd5\x29\x17\x06\xc3\x3e\x97\xf0\xf8\x4b\xa7\xe5\xba\x55\x4d\x03\x9a\x5b\x26\xa2\xa1\x76\x73\x69\xaa\x9d\x30\x85\xba\x1c\xf2\x86\x37\x1e\x68\x26\xce\x15\x07\xc3\xbc\xdf\xb9\xc9\x8d\xe4\xec\x18\x2f\x12\xf3\x22\x1b\x88\x00\x4e\x32\x00\x37\xd1\x5e\xaf\xb5\x4a\xaf\xbb\x70\xe3\xc6\x98\x92\x53\x1c\x20\x6b\x07\x64\x19\x85\x04\x6a\xc4\x8e\x2f\x68\xa2\xbb\x3d\x41\xd0\xa8\x9f\x53\x57\xb2\x39\x52\x20\x77\xc3\x3e\x21\xa0\xbb\x54\xe6\x7f\x86\x73\x89\x1f\xa4\xc3\xdd\xb5\xe6\x18\xe6\xdc\x85\x10\xe7\xb8\x22\x7c\xfa\xac\x74\x6d\xd0\xc6\x0e\xad\xae\x76\xcc\xf0\xe4\xdb\xd7\x7f\xe1\x67\x1a\x4b\xc6\xbf\x95\x76\xe3\x2f\xae\xc8\xd6\x8f\xab\x4a\x15\xe9\x54\x40\x1d\x1a\xa1\x33\xd5\x87\xa9\x80\x3a\xcc\x07\xcf\xdb\x45\x1c\xfb\xc6\xa3\xd1\xa7\xb4\x28\xd2\x79\x99\x67\xcb\x31\x8d\x10\x00\xbc\xf9\x5e\x35\x1a\x94\x18\xd3\xe5\xbb\x7c\xe9\xd4\x89\xee\x36\xde\xa6\xab\x59\xb9\x1c\xc3\x5c\x85\x4c\xc0\x85\x38\xae\xb5\xd2\x7e\x6a\x25\x0c\xd9\xba\x0b\x35\xb2\xbb\x6c\xe2\xe6\x97\x43\x16\x1c\xa9\xc7\xf0\x3e\x96\x81\xad\x40\x50\x70\xee\xf8\x91\xe6\x74\xd5\xcd\xa8\x22\xe9\xae\x35\xf9\xca\x8f\x9b\xef\xd4\x60\x8e\xc6\xf2\xbd\xe9\x7a\x9a\xe7\xa6\x4d\x2e\x64\xdb\x6b\xe4\xe7\x2b\xe9\x0f\x65\x5f\xc1\xd3\xbb\x42\xe5\xf9\x86\xd0\xd9\xab\x1f\xe7\x8b\x09\x7e\xb2\xdc\x5d\xd2\xf1\x5d\x58\x9a\xe5\xe4\x94\x7b\x62\x20\x72\xa3\x8d\xeb\x23\xfc\xbd\x65\x8d\x2b\xc9\xb9\x08\x28\x79\xe3\x8a\x04\x6f\x4e\xdc\x99\x27\x5b\x78\x94\x33\x98\x9c\xc3\x6c\xa9\xbf\xd8\xf6\xe2\x3a\x84\x17\xd1\xa0\xf3\x1c\xa5\x3c\xbe\x4f\x10\x66\x03\x79\x1d\xed\x15\x4f\x8b\x9c\xfc\x4a\x88\x6c\xcf\xd1\x88\xa4\x65\xd1\x7d\xb1\x8d\xb0\x34\x0d\x4b\x90\x35\xbe\x9c\x70\xe8\xdb\xcd\x98\x42\x37\x22\xee\x4a\x13\x26\x4e\x7c\xf3\xd9\x79\x3b\x9a\x2d\x67\xda\x0f\xad\x86\xef\x11\xea\xf8\x0b\x6e\x18\x58\x9f\x44\xdd\x86\x70\xe1\x38\x81\x59\xb1\xbf\xba\x70\x76\x42\xb6\x35\x16\xb3\x0c\x9a\x93\x12\x7b\x6e\x50\x58\x8c\x52\x6c\x97\x10\xef\xb8\x1b\x08\x96\xfe\x32\xcb\x13\xef\x9d\x4b\x3c\x3e\xf0\xda\x08\xee\x9e\x35\xa1\x65\xe6\xa1\x0c\xd5\xdf\x84\xac\x94\x26\x58\xde\xe5\x50\xe8\x24\x9e\xdd\x4c\xd9\x9a\xef\xd8\x93\x50\x7a\xb8\x1d\xd2\xb2\x3e\x0e\xce\x8a\x89\xa7\xf5\xf7\xf5\xd0\x4e\xb4\x30\x9f\xc7\xa3\x51\x4a\x55\x87\x4e\x4d\x2e\x5e\x68\x89\xf5\x07\x95\x45\xec\x79\xdc\x58\x78\x31\xa3\x88\x1c\xbf\x70\xd9\x1e\xbf\x0f\xff\xfd\x40\x42\xee\x65\xb9\x7a\x70\xf7\xdd\xfb\x5b\xc8\xfe\xbe\x79\x57\xf6\x86\x7c\x7e\xbb\x28\xee\xfd\x55\xf1\x22\xeb\x6f\x93\x7f\xfb\xfa\x4b\xba\x84\x7c\x89\x6e\x18\xdf\x0a\x77\xa0\xd3\xf9\x23\xf4\x7e\x32\x81\xec\x6f\x93\xec\xa1\x84\x74\x19\xee\xea\xce\x1e\x61\x99\xd1\xff\xe2\x50\xde\xf9\x0b\xed\xe3\xe1\x8d\xe1\x69\xbe\x9c\xcc\xd2\x9c\xae\xbf\x3f\x76\x97\xf2\x17\x05\x2c\xcb\xb4\x5c\x95\x8b\x62\xb8\x43\x3e\x9f\xcc\x56\xd3\x7c\xfe\xa1\xa3\x21\x4a\xf6\x86\xe4\x60\xee\x77\x9f\x15\x93\xbb\x74\x5e\xfa\x0c\x30\x81\xdb\xbc\xa4\xff\xec\xe1\x76\x51\x40\xea\xee\x93\x4f\x56\xb3\xb4\x00\x7f\xf7\x13\x77\x9e\x2f\xe6\xef\xf2\xf9\x6d\x91\xcf\x3f\x84\x0b\xc6\x8b\x87\x6c\x0e\xcb\xc5\xaa\x98\xf4\xff\xcd\x83\x1b\xfc\x3d\x1f\x13\x4e\x5a\xc2\x67\x62\x01\x6a\x63\x30\xfa\x5e\x87\xd0\xc5\xcb\xbe\xec\xb4\x3f\x28\xe9\xea\x26\x01\xad\xbb\xc2\x33\x02\xa2\xe0\x8c\xfb\xea\x23\x01\x20\xff\x6d\x3c\x0d\xe3\xe3\x7a\xdf\x92\x72\x94\xc7\x5b\xf5\xf7\x04\xe2\x9b\xcd\xe7\x42\xe6\x6b\x4d\x2e\x1e\x6f\x1b\x2f\xff\x92\x04\x7c\xff\xf7\xe4\x70\x02\xf5\xa1\xb0\xea\x1b\x12\x12\xcd\x7f\xa3\x5a\xd9\x75\x3d\x42\x45\x12\x36\x54\xe4\xc3\xaf\x86\xb1\xb7\x2b\x50\x46\x11\xb7\xa2\x9e\x91\x3c\x76\x00\x37\xaa\x85\xa2\xd6\x17\x59\xfc\x9f\x22\x40\xf9\x9b\x9b\x78\x94\x9d\x5b\x61\xc3\xdc\x1f\x26\x33\x3a\x9e\x03\x0f\x33\xa2\x83\xda\xdc\xcb\x5b\x47\x1e\x14\x05\x43\x0f\x9c\x8a\xef\x16\xfd\x4a\x05\xef\xa5\x72\x11\x88\x70\x63\xc5\x0d\x04\x3c\xe5\xcf\xa5\x36\x3d\x45\xdf\xbe\xfe\xff\xae\x4a\xa0\xb9\xbf\xe9\x65\x95\xbb\x31\x1d\xb0\x43\xd2\x0d\xa2\x08\xf9\xa4\x42\x1f\x30\x2c\xaa\x74\xed\x17\x42\xe7\xd7\x4b\x7f\x50\xde\xa5\xa6\x41\x37\xec\x7e\x66\x26\x80\xca\x55\x83\x0c\x6e\x96\x7e\x22\x4f\xf3\x31\x9b\xaf\x32\x12\xcc\x80\xf9\x1d\x44\xd9\x62\xf6\x22\xc3\x4c\x18\x8d\x7e\xd1\x19\xef\x5a\xb9\x65\xfa\x38\x06\x98\x09\x1b\x5a\x4f\x4a\xfa\xc4\xc4\xe1\xed\x8a\x51\x6b\x8b\x69\x61\x2e\x0d\xe3\x76\xfb\xc4\x17\x4e\xd7\x34\xd1\xe7\x87\x8c\xdc\x7d\xe4\xfd\x81\x5b\xa7\x14\xad\xb6\x31\x01\x80\xa1\x43\x6e\xe1\xe7\x56\x0b\x53\x0b\x17\x79\xfa\x7b\xe9\x3d\xc3\xdc\xf3\x82\x49\x3a\xc2\x78\xf4\xdf\x01\x00\x00\xff\xff\xd1\x06\x58\x63\x61\x48\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/LICENSE-v1.txt"].(os.FileInfo),
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
