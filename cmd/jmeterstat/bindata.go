// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// web/compare-data.js
// web/compare-tables.js
// web/compare.html
// web/report-data.js
// web/report-tables.js
// web/report.html
// web/template/compare.html
// web/template/report.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _webCompareDataJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4b\x6f\xdb\x46\x10\x3e\x53\xbf\x62\xb0\x40\x00\xb2\x56\x24\xd9\x69\x0b\x44\x8a\x7b\xc8\x03\x68\x91\xba\x69\xad\xa0\x17\x41\x87\x35\x39\xb2\x08\x53\x5c\x79\x77\x19\x59\x08\xf4\xdf\x8b\x7d\xf1\x25\x92\xb2\x65\xb9\x88\x2f\x16\x77\x67\x76\xbe\xe1\xcc\xc7\xd9\x99\xe1\x10\x12\x7a\x83\xc9\x3f\x19\xf2\x18\xc5\x57\x7a\x93\xe0\x35\xdb\x08\x58\x73\x5c\x53\x8e\x20\xd5\x0a\x44\x54\x52\x58\x30\x6e\x84\x61\xc1\xd9\x0a\xe2\x74\x9d\x49\xb3\xd3\x1b\x0e\x81\xa3\xcc\x78\xaa\x7f\x2e\x62\x2e\xe4\x18\x16\x8c\x49\x31\x0b\x59\x92\xad\x52\x31\xef\xab\x1d\x81\x21\x4b\xa3\x31\x70\xb6\x11\xb3\x8b\x79\x75\x53\x2e\x63\x1e\x8d\x81\x72\x4e\xb7\xb3\x7b\x03\xa8\x10\xe9\x2d\xb2\x34\x94\x31\x4b\x0d\x86\x0f\xab\x75\x1d\xb3\xaf\x37\xfa\x1a\x52\x1f\xe8\xed\xed\xef\x48\x23\x11\xc0\xf7\x9e\x97\xa0\x34\x70\xe0\x12\x66\x73\xf3\xac\x30\xe8\xc7\x89\x79\xb6\x86\xce\x1b\xd6\x2e\xec\x5a\xcf\x73\x42\x83\x75\x26\x96\xfe\x77\x89\x0f\x72\xac\xcd\xee\x82\x49\xbe\x7b\x61\x76\xc9\x35\xde\x67\x28\x24\x09\x94\xe6\x70\x08\xd3\x6c\xb5\xa2\x7c\xdb\xf3\xbe\x51\x0e\x22\x5b\x4d\x25\x95\x70\xa9\xf1\xce\xc8\xd4\x3c\x13\x65\x7a\x38\x84\x4f\x0f\x18\x66\xca\x5d\x31\x86\x0f\x2c\x4b\x65\x1f\x3e\x71\xce\xb8\x80\x57\xcd\x28\x80\x14\x2a\xa4\x0f\x24\x64\x89\x58\xd3\x94\x8c\xe1\xbc\x09\x9b\x3b\x4c\x81\xf3\xf4\xab\x31\x1b\x57\x54\x2e\x07\x9c\x65\x69\xe4\x9f\x8f\x46\xf0\x93\xc3\x39\xb3\x1a\x64\x1e\xc0\x10\xce\x47\xa3\xc0\xe0\xbc\x46\xb1\x66\x69\x88\xf0\x35\x5e\xa1\x68\x83\x56\x95\x02\xff\x55\x50\x85\xe8\x82\x35\x48\x30\xbd\x95\xcb\x5d\x50\xc3\x3b\xa0\xeb\x75\xb2\xf5\xdd\x5a\x29\xba\x0a\x3c\x07\x5f\x05\x6b\x89\x34\x02\xb6\xa8\x46\xbe\xcb\xb7\xc2\xb5\x84\xae\x05\x46\x64\x3e\x53\x67\x94\x3d\xdc\x99\xc0\x61\x2a\xdb\x5c\x53\x7b\x3f\x9c\x43\x0a\xd4\xfb\xad\x44\xd1\xea\xd2\x35\x86\x18\x7f\xc3\xa8\x3d\x62\x66\xff\x87\x73\xad\xdd\x2d\x7b\x7c\xc6\x13\x88\x53\xc7\x2a\x4d\x29\x73\xba\x25\xbd\xe3\xb7\x61\x61\x85\x82\x5a\x78\x96\xf1\x44\xef\x57\x58\x58\x7f\x6e\x60\xa5\xc7\xd9\xc6\x80\xcf\x78\xd2\x87\x7d\x26\xb5\xd1\xa8\x89\x47\x87\x5e\x55\x61\x6c\xef\x4d\x1d\x48\x69\x95\x00\x45\x52\x3f\xd7\x4e\x57\xa6\x39\x4b\x45\xae\x3d\xd7\x5a\xa7\x25\xf5\x41\x37\xca\x9c\x6d\x5c\x4a\xa8\x20\xdb\x54\x54\x81\x77\xc9\xde\xcf\x3f\xed\x2a\xd4\xb6\x76\xe9\x14\x2c\xd5\xae\x7e\xcf\xf3\xca\x15\x4b\x3d\xdb\x22\x65\xb5\x7b\xde\x6e\xd2\xdb\xf5\x8a\xda\x74\x5f\xaa\x49\xc2\x57\x69\x95\xd7\x1f\xe7\xa9\x82\x41\xae\x90\xa6\x8a\x56\x57\xb1\xf9\x47\x1f\xd4\xbf\xbf\xdf\x8e\xcc\xbf\x5f\xcc\xbf\xb7\x64\x6e\x5d\xb8\x61\xd1\x56\xe5\xe9\x9b\x81\xc0\x04\x43\xe9\x13\xb5\xa2\x3f\xde\xea\x87\xe2\x1c\xa6\x91\x4f\x96\x6f\x48\x30\x50\x14\xf6\xc9\x7b\x4c\xc3\xe5\x8a\xf2\x3b\xe0\xb8\x66\x5c\x82\x4f\xe0\x0c\x52\xdc\xc0\x47\x2a\x51\x63\x1b\x4c\x25\xe5\x12\xa3\x60\x20\xd9\x1f\xd3\x2f\x53\xc9\xe3\xf4\xd6\x0f\xe0\x0c\x08\x10\x15\x8e\xfc\xaf\xae\xf8\x29\x8d\x1a\xd5\x02\x12\x58\xc4\x2a\x60\xc2\x52\x4b\xd9\x91\x13\xcb\xcf\x3b\xdc\x2a\x7e\xea\x7d\x13\x75\x25\x6e\xee\x19\x97\x50\xf1\x46\x2f\x6a\x27\x3d\xfd\xb3\xd8\x58\xe6\x6e\xfe\xa9\xaa\x6f\xb7\xcc\x1d\x6e\x83\x01\x95\x92\xfb\xf9\x67\xac\x0f\xe4\xe2\x67\x53\x98\x55\x86\x5e\xd1\x38\x05\x15\x1d\xe4\x16\x0f\x67\x1b\xf5\x0c\x97\x50\x3b\x95\x6b\x0f\x3d\xbb\xdf\x84\x88\x38\x63\x9b\x38\x92\x4b\x65\xea\x9c\x04\xdd\x2a\xa5\xd2\xdd\x84\xf4\x90\xfa\x11\x16\xeb\x15\x79\x25\x82\x26\xd3\xbf\xbe\x80\x69\x53\x31\x6f\x14\x95\xff\x37\x9b\x45\x39\x3b\xb9\xdd\x6a\xc2\x5c\x34\x67\x4c\x7e\xee\x45\x33\x3a\x7d\x51\x14\xe6\xb4\x7a\xdd\xe9\xd2\xd4\x45\xe8\xe0\xf9\xc5\x65\xef\x80\x60\xdb\x9b\x3d\xb2\x40\xb5\xd9\x51\x1a\x41\xf1\xd5\x3e\x06\xce\xe3\xab\xd7\x0b\x82\x78\x5a\x61\x7b\x19\x20\x36\x63\xf2\xde\xc2\x35\x17\xed\x5f\x2f\x23\x42\x93\x64\x6a\x3f\xd1\xfa\x53\x3c\xbb\xc3\xed\xbc\x50\xb7\xbd\x89\x93\x2a\xf5\x27\xaa\x28\x9a\xe3\x8b\x63\xa3\x82\xdc\x06\x87\x7b\x43\x95\x44\x6e\x57\xcb\x6f\x77\x26\x9f\xe7\x0e\xe4\xe7\x2f\x70\x59\xbb\x48\xe9\xab\x94\xbf\xa7\x01\xaf\x8b\x4e\x65\x9a\x85\x21\x0a\x7b\xc7\xda\x3b\xda\x5e\x1b\x3a\xf1\x7c\xfe\x12\x74\xbb\x79\x4a\xa2\xb4\x9b\x79\x6a\xc3\x62\x93\xe8\x38\xdc\x8f\x64\xd4\x11\x68\x5b\xef\x88\xcf\xc2\xfb\x04\xf2\x1d\x81\xf9\xa4\x78\x1d\x44\xdb\x9a\x94\x48\x55\x6a\x4f\xca\xfd\x48\x5d\x42\xf7\x24\x4e\x66\xd9\xc1\x6c\x6f\xd9\x0c\x2c\xe3\x49\x5e\xf1\x12\x2a\xf4\x84\x80\x1a\x8d\x3a\x4b\xdb\x8e\x10\x75\x82\x1e\x62\x68\x9d\x9e\x0d\xdc\x6c\x21\x66\x1b\x04\x43\xcb\xb6\xdd\xd6\x64\x69\x64\x65\x29\x69\x1a\x32\xa6\xcd\xc6\xc1\x66\xab\x9c\x2b\x3a\x59\x8e\x42\x6b\xb9\x78\x4a\x8c\x15\x12\x9e\x06\x65\x89\x81\xa7\x44\x7a\x42\x94\x3b\xd5\x04\x96\xfb\xb3\x90\xad\xd6\x94\xa3\xb9\x10\x7d\x44\x11\xfa\xb4\x0f\x37\x1a\x66\xbc\x00\x9f\x0e\xfe\xa5\x49\x86\xf0\x1b\xdc\x98\x5f\xfa\xc2\x64\xba\xc3\xd7\xe7\x93\x9e\x87\x89\x40\x28\x4b\xbe\x6b\x90\x74\x82\xc5\xca\x48\x77\x89\x6e\xd8\x6a\xac\x3f\x7f\xd6\xca\xd2\x64\x0b\x2b\xfa\xa0\xce\x13\x79\x4b\xfa\xf8\x01\xec\xb3\xa6\xaf\x35\x27\xaa\xc3\x57\x07\xea\xa9\xc3\xd7\xfa\x9c\xd5\xb4\xf3\xfa\xe8\xc0\xf5\x8f\x31\x5c\xc2\x68\x02\x31\xbc\xcb\xcd\x4c\x20\x3e\x3b\x33\xb9\x56\x51\xac\x4e\x39\x9b\xf6\x88\x9b\x13\x3c\x61\x3c\x6b\x10\xa3\x39\xda\xfa\xa0\x9f\x3e\xb0\x08\xf5\x7d\xaa\x32\x32\xd5\xab\x7a\xac\x5b\x6e\x7f\x0b\x05\x83\xdb\x1c\x67\x67\x70\xf0\x17\x5d\xe1\x18\xee\x70\xdb\x07\x9d\x5c\xe3\x92\xbc\xbe\xa8\xed\xec\x24\xcf\xaa\x09\xc6\xa5\xbf\x97\xdb\xba\xbd\x2d\xcd\xd4\xca\x97\xb3\x47\xbc\x4a\x95\xe6\x6a\xcb\xda\x30\xe3\x3e\xcb\xe8\xae\x89\xb1\x91\x9f\xc5\x73\x4b\x8c\xd2\x98\xa6\xac\x57\x88\x29\x6f\xcd\x14\x07\x34\xbf\xea\x16\x08\xd9\x53\xb6\x4b\xbb\x97\x98\xfb\x55\x63\xd9\x10\x48\x7b\x66\x35\x03\x0e\xab\x79\x9d\x09\x70\x54\x06\xe8\xaf\xe2\xe1\x1c\xa8\x4c\x26\x83\x1c\x49\x77\xf8\x3b\xe3\xdf\x3e\xa8\xeb\x0c\x7f\xa1\xd6\x10\xfd\x4a\xf8\x0b\x49\x17\xfd\xfd\x95\x5d\xd7\xe0\xef\xb9\x33\xbd\xff\x02\x00\x00\xff\xff\xfc\xc6\x2f\xe8\x1a\x1b\x00\x00")

func webCompareDataJsBytes() ([]byte, error) {
	return bindataRead(
		_webCompareDataJs,
		"web/compare-data.js",
	)
}

func webCompareDataJs() (*asset, error) {
	bytes, err := webCompareDataJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/compare-data.js", size: 6938, mode: os.FileMode(436), modTime: time.Unix(1591200812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webCompareTablesJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xdb\x38\x10\x3e\x4b\xbf\x62\x96\xe8\xc2\x52\x63\x48\x09\x16\x7b\xa8\x53\xf7\xe2\x04\xd8\x02\xdb\x7d\x25\x37\xc3\x07\xc6\xa2\x2d\xb6\xb2\xa8\x25\xe9\x38\x46\xe1\xff\xbe\x18\x8e\x68\xcb\x96\xfc\x48\xb0\xed\xa1\x54\xe6\xc5\x6f\x86\xf3\xcd\xf8\x99\x6b\xe0\xf3\xf9\x6f\x82\x67\x06\x86\x30\x66\x5f\x04\x2f\x59\x1f\xd8\x17\x49\x07\x7f\xc1\xe3\xaf\x0f\xd7\x74\xfc\x4a\xc7\x07\x36\x09\xc3\xd9\xb2\x9c\x5a\xa9\x4a\x18\x2d\xaa\x3b\x61\xa6\x5a\x56\xf8\xe7\x23\x7f\x2a\x84\x89\x32\x6e\x79\x0c\xdf\xc3\x00\xaf\xb0\x28\x83\x21\xbc\x8b\x7a\x1f\xdd\xf7\xa7\x8f\x29\x9d\xbd\x38\xe1\xd6\xea\xe8\x3b\x3c\x29\x9d\x09\x3d\x80\x1b\xd8\xc4\xb7\xb5\x9b\xf6\x3e\x1a\x1d\xf4\xa7\x1e\x6a\x9c\x20\x47\x41\x8e\xee\x56\xbc\xd8\x88\x3d\x0a\x63\x59\x9c\xf0\xaa\x12\x65\xf6\xa8\x22\xab\x8f\x9a\x3e\x58\xae\xad\xc8\x2e\xb4\xbe\x2f\xb3\x2e\x5b\x87\xbe\x16\xd6\xa2\x33\x88\x33\x14\x64\xdb\xc0\x58\x9f\xe4\x0f\xbe\x10\xdd\x30\x0e\xac\x4b\xb1\x82\x3b\x6e\x05\xb9\xd5\x19\xc4\x89\x55\x9f\x1f\xfe\x7c\xb0\x5a\x96\xf3\x28\x7e\x43\x20\x97\xdc\xb9\x30\xed\x54\x5f\xf1\x3a\xee\x71\xd9\x54\x15\xa6\xa2\xc6\xfa\x85\xf9\xca\x8e\xd4\xa2\xe2\x5a\xc0\x4a\xda\xfc\x92\x02\xbf\xa9\xc2\xa3\x45\xf5\xc6\x22\x8f\x16\xd5\xff\x56\xe7\xd1\xa2\x7a\x73\xa9\x9b\xb2\x47\x15\xb1\x27\x95\xad\x59\x7c\x1b\x6e\xc2\x30\x4d\x61\xa4\x05\xb7\x02\xfe\x5d\x0a\x2d\x85\x21\xa2\x19\x98\x69\xb5\x00\xbc\x79\x8f\xa4\x7f\x93\x51\x83\xa0\xfd\x2d\xf9\xfb\xb0\xe0\x2f\xf7\x5a\x2b\x6d\xb6\xac\x35\x96\x5b\x9c\x0a\xbe\xe7\xec\x6d\x18\xa4\x29\xfc\xae\x54\x05\xea\x59\x68\x28\xf8\x93\x28\xc8\xd6\x7d\xba\x11\x32\x09\x83\x99\xd2\x10\x7d\x13\x6b\x90\x25\x05\x71\x21\x03\xb2\x49\xaa\xa5\xc9\x51\x1b\x87\xc1\x26\xf4\x42\xa3\xb4\x8d\xe2\xb0\xf6\x95\xe8\x49\x1a\x72\x4d\x53\x7f\x6c\xe9\x3b\x80\x24\xe9\xd7\x32\x22\x69\x53\x82\x70\xd9\xc0\x3b\x05\xc0\x5c\xb4\x9b\x9d\x88\xfe\x25\x49\x18\x06\xbb\x04\x60\x58\xdf\x3a\x96\x13\xaf\x78\xe6\xc5\x52\x18\xaf\x39\x28\xe3\x3f\x6a\x65\x22\xa7\xe8\x53\xa6\x63\xf7\xc7\x64\x57\x58\x7c\xd6\x3a\xd2\x4c\x29\x57\x4f\x8a\x98\xcc\xa4\x36\xb6\x56\x69\xb5\x6a\x68\x8c\x98\xaa\x32\xab\x55\x53\x55\x2c\x17\x65\x43\x6b\x73\xa9\x51\x39\x75\x4f\x7f\x00\x28\x92\x7d\xa8\xf1\xd4\x8e\x7d\xba\xb7\xef\xee\x70\x1d\xd5\x9d\x16\x3d\xfe\x99\xac\x76\x4d\x72\xfb\x03\x92\x0a\x82\x66\x5a\x0d\x40\xe7\xb3\x0a\x36\x48\x88\x6d\xb7\x77\xd7\xc6\xb1\xe3\x73\x76\x3a\x96\x6b\x90\xce\xa5\x05\xd3\x82\x1b\x33\x64\x99\x34\x55\xc1\xd7\x0c\x8c\x5d\x17\x62\xc8\x56\x32\xb3\xf9\xe0\xe6\xfa\xfa\x67\xd6\xb1\xd7\x64\x36\x00\xe6\x84\x3b\x38\x0c\xae\x60\x0b\x66\x7f\xf1\xf9\xcb\x73\xc1\x33\x7f\x39\x7e\xd3\x9c\x6b\x8d\x40\x92\xee\x48\x53\xa7\x34\xbe\x9e\x50\x22\xce\x01\x91\xc0\x10\x18\x43\x41\x93\x9c\x3b\xeb\xb1\xf4\x0e\x81\x9c\x91\xfe\xa7\x21\xf4\x70\x96\xf5\xbc\xc2\x69\x28\x16\x06\xdb\x8a\x03\x27\xbb\x1a\x02\x03\x46\x92\x4d\xb8\x27\x77\xe1\xae\x80\x0d\x31\xef\xbd\x3b\xc7\xdf\xc4\x7a\xe2\x52\xa8\x9d\x36\x1e\xb3\xcd\xb7\xd9\xf7\xe0\x8a\x52\xb8\x82\x9e\x5f\x2c\x68\x66\x73\x9a\xb5\x7b\x11\x9d\xc8\xab\x0f\xc6\xab\x8b\xee\xaa\xd9\x98\xb1\x61\x10\xbc\xa2\xae\x37\xbe\x4c\x5d\x3f\x16\x76\x46\x58\xce\x8b\x6e\x77\x0f\xbe\x3f\xf6\xd1\xa2\xd1\x08\xd8\x9b\x1e\x17\x7e\x5f\xd8\x08\xae\xa5\x4f\x60\x75\xfa\xe3\x30\x51\xbd\xff\xf3\xe6\x00\x25\x1a\x34\x50\xe2\x56\xf2\x60\xf0\xbb\x85\x67\xc7\xad\x23\xd0\xc9\xf6\xab\xb7\x6d\x34\x64\xd7\x7e\xad\x4d\xc6\x5f\x3b\x12\xa0\x26\x72\x30\x0e\x52\xd8\x60\xb8\x56\x2a\x68\xd8\x7e\x88\xc6\xae\x0d\x50\xf7\x8e\xc6\x47\x9c\xdc\x71\xcb\x69\x9e\x10\xbe\x34\xc5\x05\x39\x70\x6b\xf2\x41\xd8\x3e\xc9\x8c\xe0\x7a\x9a\xcb\x72\x3e\x80\x19\x2f\x8c\x20\xf1\x7b\xfc\xbf\xe2\xf3\x03\xf9\xfb\x14\xc1\x01\x6d\xf5\xf6\x10\x6b\x4e\xc2\x8b\x67\xd8\x8f\x1a\x61\xf7\x5a\x9f\x1c\x5f\xc7\xa7\x17\x2e\xc2\xce\xb7\x6f\xd1\x8c\x5e\xfe\x14\xc5\x3a\x1a\x77\x13\x76\xb2\xeb\x08\xb9\x4e\x70\xeb\x0c\xc6\x06\xb3\x5e\x45\x2c\xc4\xd7\xa6\xd5\x11\x56\x9d\x20\x55\x17\xa7\x4e\x4c\x83\x0e\x46\xbd\x96\x50\x6e\x22\xb4\xe9\x54\xb3\xe9\x08\x99\x4e\x70\xe9\x18\x95\x3a\x98\xd4\x4d\x24\xc7\xa3\x16\x8d\x90\x45\x35\x89\xfe\x0b\x00\x00\xff\xff\x86\x95\xaf\x7f\xd8\x0e\x00\x00")

func webCompareTablesJsBytes() ([]byte, error) {
	return bindataRead(
		_webCompareTablesJs,
		"web/compare-tables.js",
	)
}

func webCompareTablesJs() (*asset, error) {
	bytes, err := webCompareTablesJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/compare-tables.js", size: 3800, mode: os.FileMode(436), modTime: time.Unix(1591200873, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webCompareHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xcb\x4e\xeb\x30\x10\x86\xd7\xf5\x53\xf8\x78\x75\x2a\xb5\x36\xa1\xaa\x84\x68\x92\x05\x05\x89\x2d\x12\x2f\xe0\xda\xd3\x26\x21\xb6\x83\x3d\xa9\x9a\xb7\x47\x89\x83\x1a\x55\xc0\x82\xcb\x2a\x92\xfd\x5f\xbe\xcc\xc8\xe9\x3f\xed\x14\x76\x0d\xd0\x02\x4d\x9d\x93\xf4\xfd\x03\x52\xe7\x24\xad\x4b\xfb\x42\x3d\xd4\x19\x0b\xd8\xd5\x10\x0a\x00\x64\xb4\xd7\x67\x0c\xe1\x84\x42\x85\xc0\x68\xe1\x61\x9f\xb1\x02\xb1\x09\xb7\x42\x28\x6d\xb9\x96\x28\x51\xee\x6a\x08\xdc\x02\x8a\x84\x27\x57\xfc\x3a\xe9\xd5\xa2\x7a\x6d\xc1\x77\x83\xe2\x39\x2a\xfa\x8c\x9c\xa4\x41\xf9\xb2\x41\x1a\xbc\x9a\x64\x39\x0d\x7c\x74\x28\x67\x46\xf3\x72\xc5\xd7\x3c\xe1\x55\x60\x79\x2a\xa2\xed\xec\x9f\xc0\x55\xf2\x28\xe3\x29\xa3\xaa\x90\x3e\x00\x66\xac\xc5\xfd\x0d\xbb\x68\xf9\x9c\xb8\xfa\x08\xf8\x2f\x8b\xf7\xe5\x09\xb4\x72\x75\x6b\x6c\x10\x2b\xbe\xe2\x03\xc3\xa4\x7c\x10\x6c\xa3\x80\x9b\xd2\x5e\xd2\x60\x89\x35\xe4\x77\x60\x55\x61\xa4\xef\xb7\xd7\x38\x8f\xa9\x88\xe7\x24\x15\xe3\x6a\x77\x4e\x77\xdf\x82\x57\xce\x34\xd2\xc3\xcf\xa7\x30\x06\x2d\xfb\x9f\xfb\xbd\x34\xfc\x62\x47\x39\x99\x1d\xa5\xa7\x46\x9e\x1e\xbc\x77\x3e\xd0\x8c\xae\x37\x84\xcc\xb6\xa6\xb9\x87\x28\x29\x9d\x8d\x83\xfe\xdf\x63\xcd\x37\xc3\xe5\x53\x0b\xbe\x84\x30\xb9\x58\x50\x79\x38\x3c\x82\xd4\x61\x71\x8e\x9b\x6f\xc8\xa4\x52\x8c\x13\x16\xf1\x49\xbd\x05\x00\x00\xff\xff\x14\x69\x67\x71\x6a\x03\x00\x00")

func webCompareHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webCompareHtml,
		"web/compare.html",
	)
}

func webCompareHtml() (*asset, error) {
	bytes, err := webCompareHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/compare.html", size: 874, mode: os.FileMode(436), modTime: time.Unix(1592906779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webReportDataJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4b\x6f\xdb\x46\x10\x3e\x53\xbf\x62\xb0\x40\x01\x32\x66\x24\x59\x69\x0b\x44\x8a\x7b\x48\x62\xa0\x45\xea\xa6\xb5\x82\x5e\x04\x1d\x56\xe4\xc8\x22\x4c\x91\xca\xee\x32\xb2\x10\xe8\xbf\x17\xfb\xe0\x53\x4b\xca\x96\xe5\x22\xbe\x58\xdc\x9d\xc7\x37\xe4\x37\x3b\x3b\x33\x18\x40\x4c\x17\x18\xff\x93\x21\x8b\x90\x7f\xa1\x8b\x18\x6f\xd3\x2d\x87\x0d\xc3\x0d\x65\x08\x42\xae\x40\x48\x05\x85\x65\xca\xb4\x30\x2c\x59\xba\x86\x28\xd9\x64\x42\xef\xf4\x06\x03\x60\x28\x32\x96\xa8\x9f\xcb\x88\x71\x31\x86\x65\x9a\x0a\x3e\x0b\xd2\x38\x5b\x27\x7c\xee\xcb\x1d\x8e\x41\x9a\x84\x63\x60\xe9\x96\xcf\x46\xf3\xfa\xa6\x58\x45\x2c\x1c\x03\x65\x8c\xee\x66\x5f\x35\xa0\x52\xa4\xb7\xcc\x92\x40\x44\x69\x62\x07\xec\xaa\x55\x5f\xe1\xf1\x81\xde\xdd\xfd\x8e\x34\xe4\x1e\x7c\xef\x39\x31\x0a\x8d\x05\xae\x60\x36\xd7\xcf\x12\x80\x7a\x9c\xe8\x67\xe3\xe5\xd2\xb2\x36\x32\x6b\x3d\x27\x17\xea\x6f\x32\xbe\x72\xbf\x0b\x7c\x10\x63\xe5\x76\xef\x4d\x8a\xdd\x91\xde\x25\xb7\xf8\x35\x43\x2e\x88\x27\x35\x07\x03\x98\x66\xeb\x35\x65\xbb\x9e\xf3\x8d\x32\xe0\xd9\x7a\x2a\xa8\x80\x2b\x85\x77\x46\xa6\xfa\x99\x48\xd7\x83\x01\x5c\x3f\x60\x90\xc9\x58\xf9\x18\x3e\xa4\x59\x22\x7c\xb8\x66\x2c\x65\x1c\x7e\xb2\xa3\x00\x52\xaa\x10\x1f\x48\x90\xc6\x7c\x43\x13\x32\x86\x91\x0d\xdb\x94\xae\x37\x31\x2a\xc9\xdc\xae\xc4\xe9\xa8\xb7\x94\xcb\x68\xbc\xc4\xcf\xc1\xce\x88\x82\x42\xe6\x3e\xdc\x50\xb1\xea\xb3\x34\x4b\x42\xf7\x72\x38\x1c\x0e\xe1\x15\xb8\x6e\x53\x0c\x5e\x97\x9a\xd3\x2c\x08\x90\x73\x32\xf7\x60\x70\x60\xcf\x93\x8b\x97\xc3\xa1\xa7\x83\xbf\x45\xbe\x49\x93\x00\xe1\x4b\xb4\x46\xde\x16\x6f\x5d\x0a\xdc\x35\xf7\xea\x81\xe7\x14\xe8\xc7\x98\xdc\x89\xd5\xde\x6b\xbc\x85\x3e\xdd\x6c\xe2\x9d\x9b\xaf\x55\x38\x23\xdf\x03\x03\x57\x52\x60\x85\x34\x84\x74\x59\xe7\x53\xf5\x35\xd5\xdf\x04\xbc\x2a\x62\xbb\x8e\xe9\x86\x63\x48\xe6\x33\x69\x63\x5e\x09\x71\xaf\xe9\x80\x89\x68\x8b\x4d\xee\x81\xbb\xd8\x09\xfc\xc1\x82\x92\xc0\xde\x4b\x58\xad\x61\xdd\x62\x80\xd1\x37\x0c\xdb\x3f\x9b\xde\xff\x21\xc3\x6b\x0f\xcd\x98\xcf\x58\x0c\x51\x92\xe7\xac\x4a\x58\x6d\xdd\x1c\x29\xf9\xe9\xa1\x73\xbc\x96\xe0\x4a\x78\x96\xb1\x58\xed\xd7\x72\xbc\xf9\x6c\xc9\x79\x87\xa5\x5b\x0d\x3e\x63\xb1\xaf\x4c\x1f\x4d\xc7\x66\x2e\x5a\x12\xb1\x2d\x0b\x6d\x69\x78\xec\x05\x97\x10\x0f\xde\xef\x91\x84\x90\xd4\x29\x53\xe2\xb9\x7e\xba\x38\x9a\x7b\x2a\x59\xfa\x5c\x6f\x9d\x9e\x64\x91\xd1\xca\x2c\xdd\xe6\x44\x92\xd4\x30\x04\x96\x74\xc9\xd3\xc4\x2f\xca\x8d\x24\x88\x29\xa6\x8a\xb8\x95\x62\xea\xf7\x1c\xa7\x5a\x42\xe5\xb3\xa9\x9a\x46\xbb\xe7\xec\x27\xbd\x7d\xaf\x2c\x96\x5f\x2b\x75\x92\xbb\x92\x8c\x45\x4d\xcc\x23\x95\x30\xc8\x0d\xd2\x44\x26\xe3\x4d\xa4\xff\xd1\x07\xf9\xef\xef\xb7\x43\xfd\xef\x17\xfd\xef\x2d\x99\x9b\x10\x16\x69\xb8\x93\xec\x7e\xd3\xe7\x18\x63\x20\x5c\x22\x57\x54\x15\x91\x3f\x64\xa6\x62\x12\xba\x64\xf5\x86\x78\x7d\x99\xfc\x2e\x79\x8f\x49\xb0\x5a\x53\x76\x0f\x0c\x37\x29\x13\xe0\x12\xb8\x80\x04\xb7\xf0\x91\x0a\x54\xd8\xfa\x53\x41\x99\xc0\xd0\xeb\x8b\xf4\x8f\xe9\xe7\xa9\x60\x51\x72\xe7\x7a\x70\x01\x04\x88\xfc\x1c\xc5\x5f\x53\xf1\x3a\x09\xad\x6a\x1e\xf1\x0c\x62\xf9\xc1\xb8\x49\x48\xe9\x47\x4c\x4c\x56\xdf\xe3\x4e\x66\xb5\xda\xd7\x5f\x5d\x8a\xeb\x8b\xcf\x15\xd4\xa2\x51\x8b\x2a\x48\x47\xfd\x2c\x37\x56\x45\x98\x7f\xca\x1b\x41\xb7\xcc\x3d\xee\xbc\x3e\x15\x82\xb9\xc5\xe1\xe7\x03\x19\xfd\xac\x2f\x0b\x92\xa1\x37\x34\x4a\x40\x7e\x1d\x64\x06\x0f\x4b\xb7\xf2\x19\xae\xa0\x61\x95\xa9\x08\x1d\xb3\x6f\x43\x44\x72\x67\xdb\x28\x14\x2b\xe9\xea\x92\x78\xdd\x2a\x95\xeb\x84\x0d\xe9\x31\xf5\x13\x3c\xda\x0a\xba\xc5\xf5\xaf\x2f\xe0\xba\x56\x6f\xff\x27\x9f\x07\x85\xf0\x7c\x7e\xeb\x84\x19\xd9\x19\x53\xd8\x1d\xd9\xd1\xa9\xcb\x2b\xd7\xd6\x9a\xd5\xaa\x4b\x53\x57\x93\x63\xf6\xcb\x5b\xe7\x11\xc1\xb6\x37\x7b\x62\x81\x6a\xf3\x23\x35\xbc\xf2\xd4\x3e\x05\xce\xe3\xab\xd7\x0b\x82\x78\x5a\x61\x7b\x19\x20\x86\x31\x45\xbf\x93\x37\x3c\xed\xa7\x97\x16\xa1\x71\x3c\x35\x47\xb4\x3a\x8a\x67\xf7\xb8\x9b\x97\xea\xa6\x5f\xca\xa5\x2a\x3d\x93\x2c\x8a\xda\x7c\x69\x36\x2c\x93\xdb\xf4\x31\x9e\x85\xc8\xed\x6a\x07\x3d\x8a\x01\xf2\xe9\x33\x5c\xbd\x50\xfb\xd3\x89\xe7\xd3\x67\xaf\x3b\xcc\x73\x26\x4a\xbb\x9b\xa7\xb6\x3b\x86\x44\xa7\xe1\x7e\x64\x46\x9d\x80\xb6\xf5\x8e\xf8\x2c\xbc\x4f\x48\xbe\x13\x30\x9f\x15\x6f\x0e\xd1\x34\x34\x95\xa4\xaa\x34\x35\xd5\x2e\xa6\x29\xa1\x3a\x99\x5c\x66\xd5\x91\xd9\xce\xca\x0e\x2c\x63\x71\x51\xf1\x62\xca\xd5\x2c\x82\x6a\x8d\x66\x96\xb6\x99\xe0\xcd\x04\x3d\x96\xa1\x27\x76\x44\x1d\x10\x74\x5a\xb6\xed\xb6\x92\xc5\x9a\x95\x15\xd2\x58\x18\xd3\xe6\xe3\x68\xb3\x55\xe5\x8a\x22\xcb\x49\x68\x4d\x2e\x9e\x13\x63\x2d\x09\xcf\x83\xb2\x92\x81\xe7\x44\x7a\x46\x94\x7b\xd9\x04\x56\xfb\xb3\x20\x5d\x6f\x28\x43\x7d\x21\xfa\x88\x3c\x70\xa9\x0f\x0b\x05\x33\x5a\x82\x4b\xfb\xff\xd2\x38\x43\xf8\x0d\x16\xfa\x97\xba\x30\xe9\xee\xf0\xf5\xe5\xa4\xe7\x60\xcc\x11\xaa\x92\xef\x2c\x92\xb9\x60\xb9\x32\x54\x5d\x62\x3e\xfd\xd5\xde\x9f\x3f\xfc\x4d\x93\x78\x07\x6b\xfa\x20\xed\xf1\xa2\x25\x7d\xfc\x44\xf8\xf4\x71\x70\x23\x82\xfa\x34\x38\x47\xf4\xd4\x69\x70\x73\xf0\xab\x7b\x79\x65\xba\x32\x54\x6d\xcc\x54\xbd\xbc\xad\x8c\xe0\x0a\x86\x13\x88\xe0\x5d\x01\x60\x02\xd1\xc5\x85\xa6\x60\x5d\x59\xc3\xd7\x4d\xa3\x65\x87\xe4\xc3\x83\x27\xcc\x91\x75\x24\xa8\x2f\xda\x26\x36\xf5\xf4\x21\x0d\x51\x5d\xb2\x8a\xda\x5d\xac\xaa\xf9\x73\xb5\x27\x2e\x15\x34\x6a\x6d\xce\x8c\xf4\xe0\x2f\xba\xc6\x31\xdc\xe3\xce\x07\xc5\xb8\x71\x45\x5e\xdd\xde\xf6\x66\x30\x68\xd4\x78\xca\x84\x7b\x40\x78\xd5\xf3\xda\x26\xcf\x5e\x6d\xd9\x72\x23\x7b\xc4\x6b\x96\x99\x21\xb7\x0c\x02\x3d\x57\x34\x87\x40\xc5\xb8\xde\x9e\x45\x73\x93\x3a\x93\xd6\x7d\x19\xb3\x1e\xf0\x80\x4a\xbd\xa6\x25\x42\x0e\x94\xcd\xd2\xfe\x25\x06\x89\xf5\x2f\x6a\xf9\x9c\xc6\x66\x9d\x07\xc7\xd5\x9c\x4e\x1a\x9c\xc4\x03\x75\x60\x1e\x67\x42\x6d\xd4\xa9\xf4\x8a\x85\x7a\x75\x2e\x40\x76\x33\xa0\x93\x02\xa5\x6d\x2b\x03\x6c\xdb\x05\x01\x6a\x0c\x28\x25\x73\x02\x1c\xae\xec\xbb\xc6\x82\xcf\x9d\xf8\xfd\x17\x00\x00\xff\xff\xaf\xce\x8e\x52\xc9\x1b\x00\x00")

func webReportDataJsBytes() ([]byte, error) {
	return bindataRead(
		_webReportDataJs,
		"web/report-data.js",
	)
}

func webReportDataJs() (*asset, error) {
	bytes, err := webReportDataJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/report-data.js", size: 7113, mode: os.FileMode(436), modTime: time.Unix(1590398236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webReportTablesJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xdb\x38\x10\x3e\x4b\xbf\x62\x96\xe8\xc2\x52\x6d\x48\xc9\x61\x0f\x75\xea\x5e\x36\x01\xb6\xc0\x76\x5f\xc9\xcd\xf0\x81\xb1\x68\x5b\xad\x22\x7a\x49\x3a\x89\x51\xf8\xbf\x2f\x66\x46\x94\x64\x4b\x7e\x74\x81\xe4\x10\x31\xdf\x3c\xf8\x0d\x39\xdf\x30\xcf\xd2\x80\x5c\x2e\x7f\x53\x32\xb3\x30\x81\xa9\xf8\xa2\x64\x29\x46\x20\xbe\xe4\xfc\x91\xaf\xf8\xf9\xeb\xc3\x15\x7f\x7e\xe1\xcf\x07\x31\x0b\xc3\xc5\xa6\x9c\xbb\x5c\x97\x70\xab\xec\xdc\xe4\x6b\x5c\x3f\xc8\xc7\x42\xd9\x28\x93\x4e\xc6\xf0\x3d\x0c\x30\xbf\x43\x0c\x26\xf0\x2e\x1a\x7c\xa4\xf5\xa7\x8f\x29\x7f\x07\x71\x22\x9d\x33\xd1\x77\x78\xd4\x26\x53\x66\x0c\xd7\xb0\x8b\x6f\xaa\x30\xe3\x63\x0c\x06\x98\x4f\x03\xb4\x10\xb0\x42\x60\x85\xe1\x4e\xbd\xba\x48\x3c\x28\xeb\x44\x9c\xc8\xf5\x5a\x95\xd9\x83\x8e\x9c\x39\xea\x7a\xef\xa4\x71\x2a\xbb\xd0\xfb\xae\xcc\xfa\x7c\x89\x7d\x05\x56\xd0\x19\xc6\x19\x02\x59\x9d\x18\xcf\x27\xf9\x43\x3e\xa9\x7e\x1a\x07\xde\xa5\x7a\x81\x5b\xe9\x14\x87\x55\x15\xc4\x89\xd3\x9f\xef\xff\xbc\x77\x26\x2f\x97\x51\xfc\x3f\x12\x51\x71\xe7\xd2\xf4\x94\xda\x86\x1e\x74\x24\x1e\x75\xb6\x15\xf1\x4d\xb8\x0b\xc3\x34\x85\x5f\x8d\x92\x4e\xc1\xbf\x1b\x65\x72\x65\xf9\xf2\x2d\x2c\x8c\x7e\x02\xdc\xb5\xe9\x9a\x56\xab\x8c\xea\x1e\x1c\xc1\x93\x7c\xbd\x33\x46\x1b\x5b\xf7\x8f\x75\xd2\x61\x73\xfa\xea\xdd\x4d\x18\xa4\x29\xfc\xae\xf5\x1a\xf4\xb3\x32\x50\xc8\x47\x55\xb0\x2f\x2d\xa9\x93\x67\x61\xb0\xd0\x06\xa2\x6f\x6a\x0b\x79\xc9\x49\x28\x65\xc0\x3e\xc9\x7a\x63\x57\x68\x8d\xc3\x60\x17\x7a\xd0\x6a\xe3\xa2\x38\xac\x62\x73\x8c\x64\x0b\x87\xa6\xa9\xff\xd4\x8d\x34\x86\x24\x19\x55\x18\xb7\x4b\x1b\x41\xba\x62\xec\x83\x02\x10\x94\xed\xba\x81\xf8\x27\x49\xc2\x30\x68\x0a\x80\x49\xb5\xeb\x34\x9f\x79\xc3\xb3\x2c\x36\xca\x7a\xcb\xdf\x7c\xbc\x74\x86\xff\xe8\x17\x1b\x11\x3a\xe2\x32\xa7\xf4\xc7\xac\x39\x55\xbc\x35\x4a\xb2\xd0\x9a\x8e\x92\x93\x25\x8b\xdc\x58\x57\x99\x8c\x7e\x69\x59\xac\x9a\xeb\x32\xab\x4c\x73\x5d\x6c\x9e\xca\x96\xd5\xad\x72\x83\xc6\x39\xdd\x75\x9b\x4b\x94\x8f\xa0\xa2\x52\x45\x8d\x78\xd3\x11\x6d\x10\xdf\xf4\x97\xc3\x37\x7e\xa6\x9a\xa6\x33\xde\xa2\x9c\x20\xa8\x0b\x6a\xb1\x39\x5f\x4f\xb0\xc3\xc6\xaf\xbb\xba\xe7\x48\x48\x02\x9f\xb3\xd3\x89\xa8\x1f\x7a\xa7\x25\xcc\x0b\x69\xed\x44\x64\xb9\x5d\x17\x72\x2b\xc0\xba\x6d\xa1\x26\xe2\x25\xcf\xdc\x6a\x7c\x7d\x75\xf5\xb3\xe8\x19\xa8\x79\x36\x06\x41\x20\x72\xd9\x0a\x18\x42\xcd\x63\x7f\xd8\xfa\x7d\x57\x4a\x66\x7e\x5f\x5c\xf3\xf4\xea\x0c\x36\x46\x1b\x79\x54\xd5\x4c\xaf\x66\x5c\x03\x05\x20\x09\x98\x80\x10\x08\xb4\x65\xd8\x78\x4f\x73\x1f\x10\xe4\x0b\xb6\xff\x34\x81\x01\xce\xa9\x81\x37\x90\x85\x73\x61\xb2\x1a\x0e\x08\x1b\x4e\x40\x80\x60\x64\x17\xee\xe1\x94\x6e\x08\x62\x82\x75\xef\xed\x39\xfd\xa6\xb6\x33\x2a\xa1\x0a\xda\x79\xce\x6e\x55\x57\x3f\x80\x21\x97\x30\x84\x81\x7f\x11\xd0\xcd\xad\x78\x8e\xee\x65\x24\xc8\x9b\x0f\xa6\x27\x65\xa7\xd3\x6c\x8d\xd0\x90\x5a\xed\xd2\x73\xbd\xf6\xc7\xd4\xf7\x40\x35\x4e\x78\x9c\x17\xed\x4e\x17\xbe\x3f\xd5\xd1\xa3\xd5\x08\xd8\x96\x9e\x17\xae\x2f\x6c\x04\xea\xe6\x13\x5c\xc9\x7e\x9c\x26\x9a\xf7\xdf\x99\x03\x96\xe8\xd0\x62\x89\xaf\x8e\x27\x83\xeb\x0e\x9f\x46\x56\x47\xa8\xb3\xef\x57\xef\xdb\x6a\xc8\xbe\xb7\xb3\x72\x99\x7e\xed\x29\x80\x9b\x88\x68\x1c\x94\xb0\xc3\x74\x9d\x52\xd0\xb1\x7b\x11\xad\xb7\x34\x40\xdb\x3b\x9e\x1c\x71\x72\x2b\x9d\xe4\x51\xc2\xfc\xd2\x14\x9f\xc2\x31\x3d\x88\xf7\xca\x8d\x18\xb3\x4a\x9a\xf9\x2a\x2f\x97\x63\x58\xc8\xc2\x2a\x86\xdf\xe3\xef\xb5\x5c\x1e\xe0\xef\x53\x24\x07\xfc\x6a\x1f\x0c\xaf\xf6\xf8\xbb\x78\x76\xbd\xc9\xe8\xba\x33\xe6\xe4\xe0\x3a\x3e\xb7\xc2\xa0\xfb\x2f\x19\xdd\x7a\x47\x60\x7c\xe7\xa7\xc4\xd5\xd3\xb2\xbb\xb0\x57\x57\x47\x64\x75\x42\x55\x67\x38\xb6\x34\xf5\x43\x92\x42\x7e\x5d\x41\x1d\xd1\xd3\x09\x39\xf5\xa9\xe9\xc4\x1c\xe8\xd1\xd2\x8f\x4a\x89\x66\x41\x57\x48\x95\x8e\x8e\xc8\xe8\x84\x8a\x8e\x89\xa8\x47\x43\xfd\x12\x22\x05\x75\x04\x84\xfa\xa9\xe4\xf3\x5f\x00\x00\x00\xff\xff\xe4\x0e\x2d\xca\x43\x0d\x00\x00")

func webReportTablesJsBytes() ([]byte, error) {
	return bindataRead(
		_webReportTablesJs,
		"web/report-tables.js",
	)
}

func webReportTablesJs() (*asset, error) {
	bytes, err := webReportTablesJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/report-tables.js", size: 3395, mode: os.FileMode(436), modTime: time.Unix(1590396848, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webReportHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x4b\x6e\xb3\x30\x10\xc7\xd7\xf1\x29\xfc\x79\xf5\x45\x4a\xec\xd2\x28\x52\xd5\x18\x16\x7d\x48\x3d\x40\x2f\xe0\xd8\x93\x00\x35\x36\xb5\x27\x51\xb8\x7d\x05\xa6\x0a\x8a\xda\x2e\xfa\x58\x21\xcc\xff\xf1\x63\x46\x96\xff\x8c\xd7\xd8\xb5\x40\x4b\x6c\x6c\x41\xe4\xfb\x03\x94\x29\x88\xb4\x95\x7b\xa1\x01\x6c\xce\x22\x76\x16\x62\x09\x80\x8c\xf6\xfa\x9c\x21\x9c\x50\xe8\x18\x19\x2d\x03\xec\x72\x56\x22\xb6\xf1\x56\x08\x6d\x1c\x37\x0a\x15\xaa\xad\x85\xc8\x1d\xa0\xc8\x78\x76\xc5\xaf\xb3\x5e\x2d\xea\xd7\x03\x84\x6e\x50\x3c\x27\x45\x9f\x51\x10\x19\x75\xa8\x5a\xa4\x31\xe8\x49\x96\x37\xc0\x47\x87\xf6\xcd\x68\x5e\xae\xf8\x9a\x67\xbc\x8e\xac\x90\x22\xd9\xce\xfe\x09\x5c\xad\x8e\x2a\x9d\x32\xaa\x4b\x15\x22\x60\xce\x0e\xb8\xbb\x61\x17\x2d\x9f\x13\xd7\x1f\x01\xff\x65\xf1\xae\x3a\x81\xd1\xde\x1e\x1a\x17\xc5\x8a\xaf\xf8\xc0\x30\x29\x1f\x04\xf7\x49\xc0\x9b\xca\x5d\xd2\x60\x85\x16\x8a\x3b\x70\xba\x6c\x54\xe8\xb7\xd7\xfa\x80\x52\xa4\x73\x22\xc5\xb8\xda\xad\x37\xdd\xb7\xe0\x53\xe0\xcf\x87\x90\x72\x96\xfd\xaf\xfd\x5a\x18\x7e\xb1\xa0\x82\xcc\x8e\x2a\xd0\x46\x9d\x1e\x43\xf0\x21\xd2\x9c\xae\x37\x84\xcc\x1e\x20\x7d\xaf\xbc\x4b\x23\xfe\xdf\x23\xcd\x37\x64\x36\x79\x5d\x50\xb5\xdf\x3f\x81\x32\x71\x71\x4e\x98\x6f\xc8\xa4\x45\x8c\x13\x15\xe9\x0a\xbd\x05\x00\x00\xff\xff\xef\x9e\xf3\x05\x5a\x03\x00\x00")

func webReportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webReportHtml,
		"web/report.html",
	)
}

func webReportHtml() (*asset, error) {
	bytes, err := webReportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/report.html", size: 858, mode: os.FileMode(436), modTime: time.Unix(1592906771, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplateCompareHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\x51\x8f\xe2\x36\x10\x7e\x26\xbf\xc2\x75\xf7\xb4\x20\xb5\xf6\x71\xe8\xaa\xde\xd6\xe4\xa1\x7b\x54\xaa\xd4\x4a\x15\xbb\x7f\xc0\xd8\x43\x63\x2e\x89\xb3\xf6\xc0\x81\xa2\xfc\xf7\x2a\x36\x09\x81\x65\x7b\x57\x75\xfb\xd2\xa7\x30\xe3\xef\x9b\xf9\x3c\x9e\xb1\x11\xdf\x68\xab\xf0\x50\x01\xc9\xb0\xc8\xd3\x44\x74\x1f\x90\x3a\x4d\x44\x6e\xca\x4f\xc4\x41\x3e\xa7\x1e\x0f\x39\xf8\x0c\x00\x29\x69\xf1\x73\x8a\xb0\x47\xae\xbc\xa7\x24\x73\xb0\x9e\xd3\x0c\xb1\xf2\x77\x9c\x2b\x5d\x32\x2d\x51\xa2\x5c\xe5\xe0\x59\x09\xc8\xa7\x6c\xfa\x96\xbd\x9b\xb6\x68\xbe\x79\xda\x82\x3b\x04\xc4\x63\x44\xb4\x31\xd2\x44\x78\xe5\x4c\x85\xc4\x3b\x35\x88\x65\x35\xb0\x23\x43\xd9\xe2\x48\xfe\x7e\xc6\xde\xb3\x29\xdb\x78\x9a\x0a\x1e\x69\x27\xfe\x40\xdc\x46\xee\x64\xf4\x52\xa2\x32\xe9\x3c\xe0\x9c\x6e\x71\xfd\x23\xbd\xc8\xf2\xb2\xe2\xcd\x35\xc1\xff\x65\xe2\xb5\xd9\x83\x56\x36\xdf\x16\xa5\xe7\x33\x36\x63\x41\xc3\x20\x79\x00\xdc\x47\x00\x2b\x4c\x79\xa9\x06\x0d\xe6\x90\xfe\x0c\xa5\xca\x0a\xe9\x3e\x11\x65\x8b\x4a\x3a\x20\x0e\x2a\xeb\x90\xd4\x35\x61\xcb\xf0\x93\x3d\xb6\xc8\xa6\x11\x3c\x52\x12\xc1\x8f\xa7\xbe\xb2\xfa\xd0\x46\x6a\x13\x92\x95\x75\x1a\xdc\x9c\x4e\x69\x9a\x8c\x04\xba\x54\x60\x96\x3e\x82\x47\xc1\x31\x0b\xc6\x03\x4a\x87\xa0\x7b\x7b\x51\xea\xce\xe2\xe8\x7a\x96\x4e\xaf\xe6\xd6\x97\x4b\xc7\x70\xd7\x17\x43\xec\x6e\x69\x18\x3d\x23\xca\xe6\xbe\x92\xe5\x9c\xce\x68\x7a\x7f\xdc\xf4\x67\x83\xd9\x4b\x4a\xee\x8b\xea\x6f\xc4\x9c\x56\xaf\xeb\x39\xad\x3f\x97\x24\x78\x28\x5d\x2a\x56\xad\x91\xcd\xd2\xdf\xe4\x0a\x72\x4f\x3c\x4a\x34\x1e\x8d\x12\x3c\x9b\xa5\x49\x5d\x3b\x59\xfe\x09\x84\xc5\x93\x6d\x9a\x44\x68\xb3\x23\x46\xcf\x69\x9b\xe1\xd7\x8f\x4d\x43\xaf\x1c\x43\x00\x2c\x07\x88\x91\xc0\x78\x70\xa3\xb0\xbf\x64\xd4\x7e\x33\xe2\xec\xe7\x58\x8f\x29\x3d\xd5\x66\x4a\x83\xfa\x20\x28\x68\xce\x7a\x7c\x8f\x79\x47\x87\xdc\x74\xb1\x07\xb5\x45\x63\x4b\x7f\x1d\xfe\xc3\x39\x7c\x09\xbe\xb2\xa5\x02\xf2\x68\x0a\xf0\x64\xfc\x66\xf2\x55\xb4\x07\x28\xf1\xab\xc1\x4b\x50\x60\x76\xa0\x5b\x42\xcf\x38\x1e\xf1\xb0\x06\xe9\x12\x9e\xb6\x5d\xab\x76\xbe\x85\x73\xd6\x79\xf2\xe6\xcc\xf9\x3b\xc8\xf2\xdc\x61\x2e\x6c\xb9\x3f\xb3\xff\xf8\xf0\xf6\xc2\x7e\x7f\x61\x7f\xe8\xec\xff\x4b\x8e\x6e\x86\x78\xd7\x6e\x02\xd7\xd6\xe2\xb0\xe6\x5d\x4b\xdf\x18\x84\x82\xdc\xcd\x09\xfb\xc5\x5a\x04\xb7\x84\xa7\xa6\x49\x08\x21\xe4\x18\xb6\xae\x03\xa4\x6b\xc1\xe0\xaf\x6b\x28\x75\xd3\x9c\x65\x8a\x09\xba\x81\x7a\x61\x18\x16\x5f\x1c\x86\x97\x9a\x3e\x7d\x90\x45\x95\xc3\xa9\xb3\x9f\x6f\x20\xb6\x4b\x90\x35\xec\x9f\x9e\x71\xf2\x0d\x82\x3c\xdf\xc8\x3f\x2e\xd9\xc2\xb9\xd7\x29\x19\xd7\x66\xd7\x5e\x36\x11\x29\xba\x87\x62\x74\x33\xd6\x56\x6d\x0b\x28\x71\xc2\x1c\x48\x7d\x18\xaf\xb7\xa5\x6a\x07\x7d\x3c\x21\x75\x72\x52\xd6\xdf\x4f\xb1\xcd\x6e\xc6\xb7\xdf\xf6\xd7\xcf\xed\x84\x7d\xec\xde\xa6\x71\x60\x8d\x46\x1e\xdc\x0e\xdc\x83\xd1\x70\x47\xd0\x6d\xe1\xbb\xe0\x95\x1b\xb9\xbf\x23\xb7\x5c\x9b\xf5\x3a\xbe\x45\xdc\xc5\xe9\xf4\xbc\x8f\xd6\x22\x1b\x32\xf9\x29\x19\x6c\x2d\xda\xff\x46\xed\xe2\x55\xd4\x42\x38\xf7\x2f\x6b\x1d\x25\x83\xd7\x98\x1f\xdf\x52\x1e\xff\x57\xfd\x15\x00\x00\xff\xff\x15\x1f\xd9\x6f\x6f\x09\x00\x00")

func webTemplateCompareHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplateCompareHtml,
		"web/template/compare.html",
	)
}

func webTemplateCompareHtml() (*asset, error) {
	bytes, err := webTemplateCompareHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/template/compare.html", size: 2415, mode: os.FileMode(436), modTime: time.Unix(1594043745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplateReportHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\xd1\x8e\xea\x36\x10\x7d\x4e\xbe\xc2\x4d\xb7\x5a\x90\x5a\xfb\x72\xd1\xad\x7a\xa9\xc9\x43\xbb\x54\xaa\xd4\x4a\x15\xf0\x03\xc6\x1e\x1a\xd3\xc4\x66\xed\x09\x05\x45\xf9\xf7\x2a\x0e\xc9\x06\x96\xd5\x6e\xb5\xed\xcb\x7d\x22\x33\x3e\x67\xe6\x8c\x3d\x63\xc3\xbf\x52\x56\xe2\x69\x0f\x24\xc3\x22\x4f\x63\xde\xfd\x80\x50\x69\xcc\x73\x6d\xfe\x22\x0e\xf2\x79\xe2\xf1\x94\x83\xcf\x00\x30\x21\x0d\x7e\x9e\x20\x1c\x91\x49\xef\x13\x92\x39\xd8\xce\x93\x0c\x71\xef\x67\x8c\x49\x65\xa8\x12\x28\x50\x6c\x72\xf0\xd4\x00\xb2\x09\x9d\x7c\xa0\x1f\x27\x0d\x9a\xed\x1e\x4b\x70\xa7\x80\x58\xb7\x88\x26\x46\x1a\x73\x2f\x9d\xde\x23\xf1\x4e\x0e\x62\x59\x05\xf4\xcc\x90\xb6\x38\x93\xbf\x9b\xd2\x4f\x74\x42\x77\x3e\x49\x39\x6b\x69\x4f\xfc\x81\xb8\x9d\x38\x88\xd6\x9b\x10\x99\x09\xe7\x01\xe7\x49\x89\xdb\x1f\x92\xab\x2c\x2f\x2b\xde\xdd\x12\xfc\x7f\x26\xde\xea\x23\x28\x69\xf3\xb2\x30\x9e\x4d\xe9\x94\x06\x0d\x83\xe4\x01\xf0\x73\x0b\xa0\x85\x36\xd7\x6a\x50\x63\x0e\xe9\x4f\x60\x64\x56\x08\xd7\x9c\xde\xde\x3a\x24\x55\x45\xe8\x32\x7c\xd2\x75\x83\xa8\x6b\xce\x5a\x68\xcc\xd9\xf9\xb4\x37\x56\x9d\x9a\x08\x4d\x22\xb2\xb1\x4e\x81\x9b\x27\x93\x24\x8d\x23\x8e\x2e\xe5\x98\xa5\x6b\xf0\xc8\x19\x66\xc1\x58\xa1\x70\x08\xaa\xb7\x17\x46\x75\x16\x43\xd7\xb3\x54\x7a\x33\xb7\xba\x5e\x3a\x87\xbb\xbd\x18\x62\x77\x4b\x21\x3a\x67\x41\x67\xca\x37\x8d\x91\x4d\xd3\xdf\xc4\x06\x72\x4f\x3c\x0a\xd4\x1e\xb5\xe4\x2c\x9b\xa6\x71\x55\x39\x61\xfe\x04\x42\xdb\xed\xab\xeb\x98\x2b\x7d\x20\x5a\xcd\x93\x26\xfc\xaf\x0f\x75\x9d\xdc\xa8\x39\x00\x96\x03\x44\xc4\xb1\xdd\xa5\x28\x94\x15\x47\xcd\x6f\x46\x9c\xfd\xdb\xef\x85\x09\x14\x69\xf3\xee\x3b\x48\x0f\x82\x82\xe6\xac\xc7\xf7\x98\x8f\xc9\x90\x9b\x2e\x8e\x20\x4b\xd4\xd6\xf8\xdb\xf0\xef\x2f\xe1\x4b\xf0\x7b\x6b\x24\x90\xb5\x2e\xc0\x93\x51\xe1\xc7\x6f\xe2\xad\xc0\x20\x19\x6d\x4e\x08\x6f\x24\x2c\x41\x82\x3e\x80\xba\x26\x9d\x4f\x78\xb8\x17\xe9\x12\x1e\xcb\xae\x3f\x3a\xdf\x4a\x14\xfb\x1c\x2e\x8a\x4a\x17\xce\x59\xe7\xc9\x37\x17\xce\xdf\x41\x98\x4b\x87\xbe\xb2\xc5\xf1\xc2\xfe\xe3\xf3\x87\x2b\xfb\xd3\x95\xfd\xb9\xb3\xbf\x94\x1c\xdd\x58\xb1\xae\x15\x39\x6e\xad\xc5\xe1\x39\x74\xed\x7e\xa7\x11\x0a\x32\x9b\x13\xfa\x8b\xb5\x08\x6e\x09\x8f\x75\x1d\x13\x42\xc8\x39\x6c\x55\x05\x48\xd7\x9e\xc1\x5f\x55\x60\x54\x5d\x5f\x64\x6a\x13\x74\xc3\xf6\xc2\xa0\x2c\x5e\x1d\x94\x97\x06\xe2\x59\x83\x3c\x2f\xa0\x6d\x97\x20\x6b\xd8\x3f\x3d\xe3\xc9\x37\x08\xf2\xbc\x90\x7f\xbd\x65\x0b\xe7\xfe\x9b\x2d\x63\x4a\x1f\x9a\x8b\xa8\x45\xf2\xee\xa6\x8e\xee\x46\xca\xca\xb2\x00\x83\x63\xea\x40\xa8\xd3\x68\x5b\x1a\xd9\x5c\x02\xa3\x31\xa9\xe2\x27\x65\xfd\xdd\xd5\xb6\xd9\xdd\xe8\xfe\xeb\xfe\x6a\xba\x1f\xd3\x87\xee\x71\x18\x05\x56\x14\x79\x70\x07\x70\x2b\xad\x60\x46\xd0\x95\xf0\x6d\xf0\x8a\x9d\x38\xce\xc8\x3d\x6b\x1f\x04\xe6\xda\x69\xf5\xac\x8f\xd4\xa0\x6a\x32\xfe\x31\x1e\x94\xd5\xda\xef\x51\xba\x78\xb7\x52\x08\xe7\xfd\xba\xce\x28\x1e\x3c\x83\xec\xfc\x98\xb1\xf6\x0f\xcd\x3f\x01\x00\x00\xff\xff\x60\xaf\x67\xb2\xe8\x08\x00\x00")

func webTemplateReportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webTemplateReportHtml,
		"web/template/report.html",
	)
}

func webTemplateReportHtml() (*asset, error) {
	bytes, err := webTemplateReportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/template/report.html", size: 2280, mode: os.FileMode(436), modTime: time.Unix(1594012999, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"web/compare-data.js":       webCompareDataJs,
	"web/compare-tables.js":     webCompareTablesJs,
	"web/compare.html":          webCompareHtml,
	"web/report-data.js":        webReportDataJs,
	"web/report-tables.js":      webReportTablesJs,
	"web/report.html":           webReportHtml,
	"web/template/compare.html": webTemplateCompareHtml,
	"web/template/report.html":  webTemplateReportHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"web": &bintree{nil, map[string]*bintree{
		"compare-data.js":   &bintree{webCompareDataJs, map[string]*bintree{}},
		"compare-tables.js": &bintree{webCompareTablesJs, map[string]*bintree{}},
		"compare.html":      &bintree{webCompareHtml, map[string]*bintree{}},
		"report-data.js":    &bintree{webReportDataJs, map[string]*bintree{}},
		"report-tables.js":  &bintree{webReportTablesJs, map[string]*bintree{}},
		"report.html":       &bintree{webReportHtml, map[string]*bintree{}},
		"template": &bintree{nil, map[string]*bintree{
			"compare.html": &bintree{webTemplateCompareHtml, map[string]*bintree{}},
			"report.html":  &bintree{webTemplateReportHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
