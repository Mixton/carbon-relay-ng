// Code generated by go-bindata. DO NOT EDIT.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html

package web

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataAdminhttpassetsAppcss = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x88\xce\x4b\xd7\x85" +
		"\xb3\x53\x12\x4b\x12\x75\x91\x05\x2a\x90\x79\x7a\x30\xb6\x8e\x82\x1e\x42\x02\x22\x9e\x91\x99\x92\xaa\x50\xcd\xa5" +
		"\xa0\xa0\xa0\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62" +
		"\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func bindataAdminhttpassetsAppcssBytes() ([]byte, error) {
	return bindataRead(
		_bindataAdminhttpassetsAppcss,
		"admin_http_assets/app.css",
	)
}

func bindataAdminhttpassetsAppcss() (*asset, error) {
	bytes, err := bindataAdminhttpassetsAppcssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "admin_http_assets/app.css",
		size:        122,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1498236484, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataAdminhttpassetsAppjs = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4d\x8f\xdb\x36\x13\xbe\xfb\x57\x0c\x84\x05\x2c\x21\x8a\x9c\xf7" +
		"\x3d\x14\xa8\x16\x46\xea\xa6\x0d\x9a\xc3\x02\xc5\x36\x6d\x0f\x5b\x07\xe0\x4a\x63\x2d\x61\x9a\x14\x48\xca\x6b\xc3" +
		"\xd1\x7f\x2f\x48\x7d\x51\xb4\xb3\xd9\x78\xdd\x56\x17\x5b\xe4\x7c\x3c\xf3\x70\x86\x1c\x6a\x4b\x24\x90\xb2\x84\x39" +
		"\x70\x7c\x04\xc2\x8b\x8a\x11\x99\x6c\x44\x5e\x31\x0c\x83\x8c\xc8\x7b\xc1\x5f\x4b\x64\x64\xff\x9a\x17\x41\x0c\x77" +
		"\x01\x2f\x6e\x51\x89\x4a\x66\x18\xc4\x10\x54\x34\xb9\x17\x42\x2b\x2d\x49\x19\x2c\xa3\xeb\xc9\x84\x94\x65\x92\x09" +
		"\xae\xa5\x60\x0c\x65\x18\xdc\x10\xca\xdf\x69\x66\x75\xaf\x54\x26\x4a\xab\x77\x25\x1d\x23\x57\x1b\x91\x13\x23\xb1" +
		"\xaa\x78\xa6\xa9\xe0\x61\x23\x18\x43\x2f\x16\x43\x23\x14\x1d\x26\x00\xcd\x6c\x42\x18\x4a\xad\x60\x0e\x77\xcb\xeb" +
		"\x09\x80\x09\xe5\x9d\xe0\x2b\x5a\xc0\x7c\x50\x0c\x83\x59\x66\x07\x67\x41\xd4\x49\x7d\x24\xf7\x0c\x3d\x21\x6d\xc6" +
		"\x1c\x99\x5b\x7c\x94\x54\xa3\xf4\xc4\x64\x3b\xac\x66\x29\xe5\x39\xee\x06\x85\x1f\x19\xc9\xd6\x8c\x2a\xed\x69\xdc" +
		"\x77\xe3\xc7\x2a\x8b\xa2\x90\x58\x10\x2d\x7c\x2f\xa4\x9f\x38\x56\xba\x15\x95\xf6\xc1\x4b\x33\xa6\x66\xe9\x1a\xf7" +
		"\x41\x0c\x87\x35\xee\x53\x98\xfe\xb0\xc6\xfd\xb4\x8e\xe1\x50\xf7\xba\x3f\xa1\xd2\x94\x13\x43\xf1\x97\x2d\xcc\xf2" +
		"\x41\xca\x75\x3f\x19\x98\xdf\x12\x46\xf3\x45\x9e\x4b\x54\x86\xff\xd9\xa7\xbb\x4f\xe9\xf2\xd5\x5f\xe9\xdd\x9b\xd7" +
		"\xdf\x2f\x5f\x85\xa9\x7d\x8d\xde\x5e\xcd\xae\x3d\x1d\x0b\xfe\xe3\xbe\x44\xab\x15\x2a\xe4\x79\xb8\x60\xec\xf3\x7b" +
		"\x2a\x95\x8e\x6e\x88\xce\x1e\xa2\xcf\x61\x26\xb8\xa2\x4a\x23\xd7\xbf\x10\xf5\x40\x79\x11\xcd\x7c\x3b\x58\xe0\x0e" +
		"\xe6\x10\xf6\x19\x13\x81\xc9\x0c\xf3\x48\xd4\x95\xe4\xfd\xab\x79\x34\x2a\x9d\x0e\xd9\xb5\x25\xac\xc2\x68\x24\x61" +
		"\x1e\xc3\x10\x55\x7f\x18\xfb\x30\x07\x2d\x2b\xbc\xf6\x24\xb4\xdc\x1f\x69\x81\xad\x9c\x5b\x2c\x7e\xde\x95\xad\x65" +
		"\x5f\xad\x86\xcc\x44\x16\x9e\xf0\x09\x8e\xc7\x15\x61\xea\xc8\x65\xed\xbd\xb7\xd1\xb5\x5a\xae\x74\x27\x59\x9b\xc1" +
		"\x3a\x0a\xcd\x92\x79\x2b\x56\x14\xef\x2b\x9e\x5d\x9a\x37\xba\x82\x66\x0a\xe6\x73\x08\xc8\xb6\x08\x4e\xc5\xd9\x3b" +
		"\x38\x45\xac\x1f\xe5\xd8\x64\x8e\x4c\x93\x4b\x1b\xcd\x44\xc5\xf5\xe5\x91\x4a\xba\xc5\x4b\x5b\x65\x44\x5d\x1c\xe9" +
		"\x86\xec\x2e\x6e\x92\xf2\x4b\x9b\x54\x3a\xc7\xed\xc5\x8d\x56\x9b\x73\x4c\x9e\x2e\xc4\xa3\xa2\x7d\xb2\x0c\xdb\xf3" +
		"\xa1\xaf\x28\x9a\xef\xa2\x06\x88\x3d\x93\x92\x02\xf5\x50\x9a\x39\xd1\x24\xea\x60\xb6\x06\x74\x7b\x74\x99\xb9\xc6" +
		"\x6b\xb3\xbb\xd7\xbe\x97\xd0\x0e\x37\xe7\xe1\xd8\x6c\xb6\x2a\xba\xe8\x5b\xf9\xac\x3b\x35\xb3\x55\xd1\xa0\x76\xad" +
		"\x71\x7c\x74\x0e\xc3\x66\xb3\x6b\x5e\x1b\x1f\xdd\x69\x9c\xe7\x8e\xd8\xd1\xfe\x72\xfa\xcc\x3e\xe5\x24\xb9\x52\x64" +
		"\x8b\x61\x94\xe8\x07\xe4\x03\x6c\x89\xaa\x1c\x56\xed\x99\xd8\x1c\xd1\x81\x14\x80\x3a\xee\xcd\xa2\x94\x47\x56\x07" +
		"\x90\x87\x8d\x2a\x52\x40\x29\x13\x43\x78\x82\x52\x0a\x59\x2f\x3d\xe2\x7b\x3d\x89\x1b\xb1\xc5\x53\x2c\x0c\xeb\x6c" +
		"\x52\xd1\xf2\x2d\x37\xe1\x74\x21\x11\xf6\xa2\x02\x55\xb5\x7f\x1e\x09\xd7\xa0\x05\xe4\xc8\x50\x23\x74\xdd\x06\x20" +
		"\x37\x47\x0f\x17\x09\x4c\xe1\x15\x18\x6b\x03\xe8\x9e\xb7\x46\x29\x3c\x4c\xed\x91\x3d\x4d\x69\xbe\xab\x9f\x64\xc1" +
		"\xcf\x1b\x8e\x8f\x8b\xa2\x68\x89\x1c\x9a\x93\xf0\x60\x4e\xec\x14\x02\x52\x98\x16\xf0\x03\xd7\x28\xb7\x84\xa5\xdf" +
		"\xbd\x89\xe1\x4f\x42\x75\x0a\xff\xfb\xff\x9b\xda\xcb\x85\x51\x6f\x73\x46\x36\x2c\x8a\xe2\x1b\x13\xe1\x85\xd0\xff" +
		"\x8b\x54\x39\x4d\xd2\x99\xc9\x32\x34\x8d\x4f\xa6\xcb\xe0\xf3\x22\x09\xd3\xb5\xa3\xb6\xf6\xcc\xff\xf0\xf0\x2b\xcd" +
		"\xd6\x0c\xd3\x66\x6b\x8c\xe1\xb7\x52\x08\x96\xda\x4d\x35\x06\x6f\xae\x5d\x1d\xd3\x07\x2e\x18\xb3\xfd\x5f\xe0\x67" +
		"\x52\xe7\xe1\x9c\x2d\xc5\xa8\x7e\xeb\x7e\xf2\x0f\x05\xf4\xaf\xe4\xd7\x40\x9c\x89\xd9\x25\xcd\x36\xf8\x5f\x63\xae" +
		"\xe1\xcb\xd2\x75\xa8\x63\xb0\x3a\xf1\x88\xf8\x71\x04\x50\xc7\x0d\x5c\x2f\x86\x67\xa3\x3f\x85\xbc\x29\x0d\xf7\x3a" +
		"\xf5\xd2\xca\xe8\xaf\x60\x4f\x16\x46\xef\xf1\xe5\x75\xd1\x1e\x04\x7e\xe2\xae\x71\xff\x35\xf8\x6f\xa7\xee\xce\x6e" +
		"\x97\xa3\x47\x63\xee\x73\xe6\x82\x76\x0e\x96\xf1\xcd\xcf\x45\x14\xc3\x33\x58\x1d\xc1\x72\xaf\x91\xc7\xe8\x62\x38" +
		"\x9b\x36\x51\x22\xf7\x49\x03\xbb\xe8\xad\x6f\x73\x43\xb3\x9f\x01\x3e\x70\xa5\x09\xcf\xec\x3d\xd8\x0e\x58\xdd\xb0" +
		"\x43\xa8\x71\x53\x32\xa2\xf1\x77\xc9\x52\x98\x56\x65\x4e\x74\xb3\x1a\x37\x56\xf6\x41\x6f\xd8\xb4\xcd\x5c\x58\xe3" +
		"\xfe\x5e\x10\x99\x77\x25\xdc\x0e\x0f\x9f\x30\x52\x07\x4b\xff\x5d\x62\x84\xa2\xad\x14\xb7\xa7\xec\xc8\x6f\xa3\xb1" +
		"\xbf\xd7\xfe\xac\x58\x8f\x02\x1d\xf7\xa4\x63\x17\x49\xc6\x84\xc2\xd0\x35\xeb\x5c\x33\xeb\x23\xd3\x99\x51\x62\xcf" +
		"\x37\x9f\x53\xb5\xa1\x4a\x85\xd3\x46\x71\x7a\xca\x78\x57\xec\x60\x5d\xa4\xad\xab\xb8\xbf\x3e\x2a\xc1\xb6\x98\x3a" +
		"\x5e\x2c\xcc\xf4\x8b\x10\xda\x0e\xba\xfb\xe8\x94\x89\x72\x3f\x0a\x50\xdd\xd1\x7c\xb7\x74\x91\x4c\xdc\xdf\xb6\x4b" +
		"\x85\x71\x46\x24\x12\x55\xc5\xf4\x78\xb3\x87\xd0\x5b\x20\x67\x93\x0c\x5d\x32\x87\xed\xa8\x36\x8e\xff\x0e\x00\x00" +
		"\xff\xff\xe0\x8e\x73\x4e\x1e\x13\x00\x00")

func bindataAdminhttpassetsAppjsBytes() ([]byte, error) {
	return bindataRead(
		_bindataAdminhttpassetsAppjs,
		"admin_http_assets/app.js",
	)
}

func bindataAdminhttpassetsAppjs() (*asset, error) {
	bytes, err := bindataAdminhttpassetsAppjsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "admin_http_assets/app.js",
		size:        4894,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1556619149, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataAdminhttpassetsIndexhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x1b\x6b\x6f\xdb\xb6\xf6\x7b\x7f\x05\x23\x0c\x48\x8b\x1b\x49\x89\xbb" +
		"\x61\x5b\x60\x1b\xc8\xb2\xf6\xae\x77\xeb\x5a\xa4\xdb\x7d\x60\xd8\x1d\x68\xf1\x58\x62\x42\x91\x2a\x49\x39\xf1\x0d" +
		"\xf6\xdf\x2f\x48\x89\x96\x64\xeb\xe5\xe6\xd5\xf5\x43\x23\x4a\x87\x87\xe7\xfd\x20\xcd\xe9\xc1\xf7\xef\xce\x7f\xf9" +
		"\xcf\xfb\x57\x28\xd1\x29\x9b\x3f\x9b\x9a\x3f\x88\x61\x1e\xcf\x3c\xe0\x1e\xe2\xb1\x8f\xb3\x6c\xe6\x45\x58\x2e\x04" +
		"\xf7\x25\x30\xbc\xf6\x79\xec\x19\x48\xc0\x64\xfe\x0c\xa1\x69\x0a\x1a\xa3\x28\xc1\x52\x81\x9e\x79\xb9\x5e\xfa\xdf" +
		"\x78\xd5\x87\x44\xeb\xcc\x87\x8f\x39\x5d\xcd\xbc\x7f\xfb\xbf\x9e\xf9\xe7\x22\xcd\xb0\xa6\x0b\x06\x1e\x8a\x04\xd7" +
		"\xc0\xf5\xcc\x7b\xf3\x6a\x06\x24\x86\xda\x3c\x8e\x53\x98\x79\x2b\x0a\xd7\x99\x90\xba\x06\x7a\x4d\x89\x4e\x66\x04" +
		"\x56\x34\x02\xdf\x0e\x8e\x10\xe5\x54\x53\xcc\x7c\x15\x61\x06\xb3\x93\x1d\x34\x04\x54\x24\x69\xa6\xa9\xe0\x35\x4c" +
		"\x3b\x60\x38\xd7\x89\x90\x3b\x10\x8c\xf2\x2b\x24\x81\xcd\x3c\x1a\x19\x04\x89\x84\xe5\xcc\x0b\x82\x30\x08\xc2\x25" +
		"\x5e\x99\x97\x01\x8d\x84\x37\x7f\x66\xa0\x35\xd5\x0c\xe6\xe7\x85\xc0\x2e\xac\xc0\x7e\xfe\x3b\xc2\x24\xa5\x7c\x1a" +
		"\x16\x1f\x2d\xdc\x81\xef\xa3\x9f\xb0\x06\xa5\x51\x24\xd2\x8c\x32\x20\x08\x73\x82\x52\xca\xe9\x92\x02\x41\xe7\x1f" +
		"\x3e\x20\xdf\xdf\xa2\x40\xe9\x35\x03\x95\x00\x68\x47\x47\x18\xa6\xf8\x26\x22\x3c\x58\x08\xa1\x95\x96\x38\x33\x83" +
		"\x48\xa4\xe1\xe6\x45\xf8\x32\x98\x04\xc7\x61\xa4\x54\xf5\x2e\x48\x29\x0f\x22\xa5\xbc\x8a\x9a\x77\x56\x40\x98\x21" +
		"\x9d\x40\x0a\x0f\xb8\xb6\x6f\x17\xd8\xa2\xa0\x77\x1d\x9c\x65\x5b\xc4\xfe\xf0\xcb\xdb\x9f\xbe\x42\x2a\xa1\xa9\x95" +
		"\xda\x05\xa8\x4c\x70\x12\x5c\x2a\xf4\xe6\xd5\x37\x48\xe5\x99\x31\x1b\x24\x96\x25\x20\x30\x48\x81\x6b\x55\x88\x18" +
		"\x08\xc5\xe8\x63\x0e\x92\x82\x72\x7c\x1e\xf8\xfe\x6f\x74\x89\x98\x46\x6f\x5e\xa1\x6f\x7f\xb7\xef\x0a\xab\x41\x4a" +
		"\x46\x33\xcf\x18\xb2\x3a\x0d\x43\xa1\x54\x50\x72\x6d\x18\x35\x0e\xf3\x95\x4a\xe8\x2a\x7c\x19\x7c\x1d\x4c\xaa\xb1" +
		"\x65\xef\x52\x79\xf3\x69\x58\xa0\x19\x8b\x51\x16\xac\x84\x27\xc1\x97\xc1\xc4\x8d\xda\xb1\x1d\xfc\x06\x9c\xd0\xe5" +
		"\xef\x86\x85\x69\x58\x78\xe4\xb3\xe9\x42\x90\x35\x92\x82\x19\xc3\x17\x51\x6e\xf8\xf6\x50\x25\xb9\xd7\xf4\x06\x08" +
		"\xe2\x78\xb5\xc0\xd2\x31\x4f\xe8\x0a\x45\x0c\x2b\x35\xf3\xca\x0f\xc5\x1f\x9f\xf2\x15\x48\x05\x5e\x89\x8f\xe3\x15" +
		"\x8d\xb1\xf5\x23\x33\xaf\x39\xd3\xb8\x0d\xa6\x1c\x64\xf9\xad\x0d\xaf\x6f\x88\xac\x41\x20\x34\xc5\x5b\x10\x0b\x89" +
		"\x39\xd9\x68\xde\xdb\x76\xa5\x69\x88\x37\xe8\x43\x42\x57\x3d\x6b\x45\x82\x31\x9c\x29\x40\xee\xa1\xbe\x6c\xce\x6a" +
		"\xd0\x8e\x5d\x8e\x57\x35\x18\x6b\x95\x0e\x0a\x47\x9a\xae\xa0\xf1\xd5\x12\xbf\xa1\xf3\x07\x91\x42\x8d\xb8\x82\x40" +
		"\x46\xb7\xd0\x75\xcd\x5f\x60\xf2\x16\xb4\xa4\x91\x0a\x27\x5f\x26\xc1\xa5\x32\x22\xfe\x0e\x1b\x63\xb5\x6f\x7b\x31" +
		"\x4f\xc3\x9c\xb5\x0b\xe5\xc0\xf7\xc3\x80\xe3\x55\x25\x0b\xdf\x9f\x57\x30\xe5\xc3\xb3\x2e\x45\x96\x6a\x4f\x31\x2d" +
		"\x92\x81\xf9\x22\x05\x63\x20\x67\xde\x5b\x4c\xf9\xb9\x66\x35\x43\x68\x53\x85\x14\xd7\xc5\x4c\x26\xf0\x55\x5d\xeb" +
		"\x0c\xa4\x36\x1f\x24\x64\x80\xf5\xcc\x2b\x5e\x50\x8e\xec\x83\xf2\xe6\xb7\xb7\xf6\x29\x48\x55\xfc\xe7\x9f\xd3\xd0" +
		"\x0e\x6a\x08\x1a\xf4\x32\x3f\x25\xfe\xc9\xa4\xa9\xbb\x64\x32\xff\x27\x66\x94\x58\x7b\x9d\x86\xc9\x64\x4b\xf6\x1a" +
		"\x2f\x18\x38\x1c\xc5\xc0\xfe\x6f\xb8\x24\xc0\x15\x90\x2d\x6d\x9b\x39\x2e\xed\x6d\xbf\x97\xbb\x2f\x2d\x78\x69\xbe" +
		"\xd3\x50\x27\x5d\x10\xa5\xde\xd1\x24\x38\xee\x03\x2b\x79\x01\x24\x24\x01\xd9\x0e\x39\x0d\x77\x09\x31\x90\x2d\x44" +
		"\x4f\xb5\x89\x13\x7b\xb0\x42\xe6\xb7\xb7\x91\xe0\x4b\x1a\x07\x95\x58\xff\x60\xb0\x02\xf6\x07\x83\x18\x47\x6b\xa3" +
		"\x26\xdd\x22\x9d\xa1\xd9\xe9\xe4\x78\xbf\xa9\xf0\x87\x95\x41\xd7\xa4\x0e\x29\x6c\xf3\x3b\x0d\xad\xba\xeb\x6e\x54" +
		"\x33\xe2\x31\x26\x56\x18\xd9\x77\x0c\x47\x57\x8c\x2a\xfd\x64\x36\xf6\x16\xeb\x28\x41\x99\x84\x25\xbd\xe9\xb5\x34" +
		"\x0b\xa7\xf2\x85\xd2\x92\xf2\x78\x18\x54\x42\x0c\xbd\x18\xcf\x22\xa3\x47\x75\x4f\xd6\x58\x8f\x07\x0b\x13\x0b\xac" +
		"\xac\x82\x85\x13\xf0\x8e\xac\x3a\xa5\xe2\xec\x66\x11\x14\x52\xe9\xb6\xaf\x0a\x72\x23\x97\x31\xc0\x56\x32\x03\x80" +
		"\x1b\xbd\xc3\x8d\xf6\x23\xe0\xda\x24\xbe\x29\x2e\xc2\x21\x8d\xae\x66\x9e\x84\x54\xac\x60\x63\x40\xcf\xbf\xa0\x9c" +
		"\xc0\xcd\x0b\x6f\x3e\xdd\xe4\x9d\x98\xad\xb3\xc4\xd4\x98\x68\xf3\xe4\x17\xd3\xfc\x88\xca\x88\x81\x17\xce\x4d\x72" +
		"\xb8\x93\x2b\xb4\x38\x03\x2a\xad\xfb\x02\xae\x25\xd5\x20\x55\x8b\x75\x2f\x85\x4c\xcb\xda\x59\x96\x60\xaf\x85\x4c" +
		"\x3d\x47\xbb\xf9\xee\x63\x42\x90\x7d\x88\xa5\xc8\x33\x94\x60\xe5\x2f\x01\xc8\x02\x47\x57\x2e\xbf\x2c\xed\x24\x1e" +
		"\xfb\x2a\x5f\xa4\xd4\x24\x03\x42\xdc\xba\xcf\x5f\x78\x88\x8b\x55\xe9\xf8\x4f\xe2\x5e\xef\x18\xe9\xf3\x81\x9f\xe1" +
		"\xba\xdf\x93\x9e\xc8\x83\x64\xe5\x41\x4e\x3b\x6a\x7f\x0f\x92\x81\x60\x64\xd8\x23\x64\xc0\xe1\x7a\x0c\x58\x8a\xef" +
		"\xcb\x6d\x36\x16\xf2\x74\x5e\x83\xf6\x4f\xa4\x0d\x16\xbb\xfc\xa2\x7d\xa2\x9d\x4c\x79\x96\xdb\xc2\x29\x15\xc4\x74" +
		"\x4b\x1c\xae\x9d\x20\x82\x77\x8c\x78\xa5\x3f\x0a\xf3\x58\x5f\xa5\x2c\xdd\x3c\x94\x31\x1c\x41\x22\x18\x31\x45\x9c" +
		"\x05\x93\xa6\x3f\x97\xd0\xa1\x11\xe4\x52\xa1\xf1\xcf\x44\x5c\x37\x7d\xdd\x58\x47\xf0\x05\xe5\xd6\x41\x7b\xe8\xb6" +
		"\x58\x54\x86\x79\x1f\x1a\x90\x52\xc8\x20\xc3\x5a\x83\xe4\xde\xfc\xd5\x4d\x06\x91\x36\xcd\x8a\xe0\x3e\xa4\x99\x5e" +
		"\x23\x97\xb8\x0c\xa6\x1e\x72\x9b\x89\xbc\xf9\x69\x84\xe5\xdd\xb3\x5a\x7e\x86\x6b\xa7\x16\x6e\x1e\x87\xd5\x62\xc0" +
		"\x9e\x82\x7e\xbd\xce\x0c\x91\x79\xba\x30\xc5\x7f\x3b\x37\x6f\xf1\x8d\xe3\x26\x35\x8f\xad\xdc\xdc\xc9\xa4\x52\x7c" +
		"\x73\x1f\x26\x65\xd1\x74\x99\x14\x46\x05\x97\x48\x0b\x44\x80\x0b\x0d\x28\xc5\x37\x48\x82\xd5\x43\xb1\x57\xf0\x1c" +
		"\xa3\x4c\x28\x6a\xfa\xbe\x12\xfa\x08\x09\x89\xfc\x13\x93\xcf\x10\x17\x88\xd1\x94\xea\x17\x0f\x65\x8d\x1d\x1f\x16" +
		"\xb9\xd6\x82\x3b\xb1\x2f\x34\x47\x0b\xcd\x7d\x95\xda\x3f\x99\xa4\x29\x96\x6b\xfb\xbc\x60\xc2\xa4\xd8\x42\xa7\x45" +
		"\x66\xb5\x3a\x25\x54\x99\xa4\x40\xb6\xc4\x55\x49\xfc\x8c\x90\x69\x58\x2c\xb3\x0f\xd9\x6d\x51\x73\xaf\x6a\x23\x34" +
		"\x36\xb4\x5b\x81\x9c\xc5\xb1\x84\x18\x6b\x31\x54\x83\xe0\x38\xbe\xb7\xf2\xa3\x5a\xf4\x33\x28\x40\x7e\x84\x75\x5f" +
		"\x09\xf1\x3a\xe7\x51\xd1\xe3\x76\xc3\x5c\x0c\x55\xf2\xef\x72\x6d\x22\x80\x91\x05\xd6\x7d\x80\xe7\x38\x4a\xa0\x0f" +
		"\xe0\x8d\xc9\xd8\x2b\xcc\xd0\x73\x05\xd1\x8b\x3e\xc8\x7f\x61\xaa\x87\xa1\xbe\x97\x22\xbb\xc0\xbd\x35\xd6\xc3\x15" +
		"\x51\xb8\x2a\xa2\x70\x65\x87\xfb\x97\x51\x38\xf8\x11\x7a\x1a\xe4\x0a\x6c\x99\xf3\x31\x60\x63\xfa\x0f\x0b\xf8\x2e" +
		"\xd7\xaf\x53\x3d\x06\xd2\x2a\x76\x0c\xa0\x53\xf0\x18\x58\xa3\xe2\x31\x70\xa5\x92\xef\xa9\x34\xac\x79\xef\xa3\x14" +
		"\x87\x9f\x75\x79\x78\x16\xc7\xc1\xeb\x9c\xbb\xa4\xbd\x34\x8f\xc3\x25\x88\x0b\x2a\x55\x2e\x37\x58\xcb\x44\x3a\xf3" +
		"\x6c\x3c\x3c\x8b\x63\x03\xd6\x47\x4f\x23\xcb\x97\x01\xda\xd8\xf8\x27\x26\xf8\x06\x86\xce\xdc\x5e\x6a\x9f\x0a\x8e" +
		"\x96\x25\x1b\xa7\x08\xaf\xe2\x23\x14\x89\x9c\xeb\x23\x44\x80\x69\x6c\xfe\x48\xba\x82\x23\xc4\xb0\xd2\x47\x26\xff" +
		"\x1f\xa1\x94\xf2\x23\xa4\x34\x81\x95\x49\xf5\x2a\x4f\xff\x1a\xb5\xa6\xd1\xb1\x0d\xf1\xde\xa6\x1f\xb7\x83\x61\x3d" +
		"\x97\x80\x3b\xba\x2d\xb0\xed\xab\x59\x8b\xed\x8e\xba\x2d\x71\x74\x69\xd7\x62\x46\x12\xe2\x9c\x61\x89\xe0\x26\x93" +
		"\xa0\x94\xcd\x7e\x7f\x15\x45\x15\x21\x79\xd3\xa9\xe5\x7a\x69\x46\x23\x9a\xb5\x5c\x0f\x55\xd6\x0f\xc2\x49\xa3\x39" +
		"\x88\x12\x88\xae\x16\xe2\xc6\xdb\xe5\xcb\x26\x10\xc7\x56\x39\x68\xe3\xea\x33\xeb\x6c\x0c\xe9\x2e\xa5\x39\xea\xe9" +
		"\x66\xdc\xaa\x96\xba\xb7\x84\xff\xfd\xed\xd8\xff\xf6\xf7\xbf\x7d\x11\xee\xed\x2d\x6e\x95\x3b\x3a\x4c\x85\x66\xb0" +
		"\xdb\x79\x4e\x39\x52\x60\x8a\x53\xf5\xa2\xd6\xfa\x7c\xcc\x31\xd7\xf4\x7f\x94\xc7\xc8\x21\xfb\x6c\xdd\x69\x40\x93" +
		"\xa6\xe0\x70\x5a\xbc\xb6\xcf\x0f\xa9\x41\xb3\xc2\x1d\xb5\x57\xa0\xf8\x54\xcd\x99\xa6\xd5\x60\xf8\xdc\xd4\x35\x36" +
		"\x66\x94\x75\x9f\xd3\x98\x19\x22\x3b\xbe\xa7\xc0\xf1\xf0\x4d\xb4\xd3\xe3\x67\xdb\x3f\x5f\x88\x5c\xdb\x4d\xb3\xde" +
		"\xfd\x7b\x91\x6b\xb8\xbf\xcd\x7b\x83\xed\x89\x1b\xe7\xa2\x33\xb4\x2d\x61\xf1\xc3\x87\x6b\x38\x64\x0c\x19\xab\x34" +
		"\x15\xbf\x42\x09\x48\xf7\x2b\x97\xb6\x99\x96\x07\x74\xd5\xdf\x7a\x17\x40\xc6\x3a\x86\x0f\xd3\x9e\xe4\x7c\x8e\x10" +
		"\x53\x1e\xf5\x81\x7c\xc8\x84\x60\x7d\x00\xef\x69\x74\xc5\xfa\xf8\x43\x91\x60\x26\xfa\xcc\x26\xe8\x91\x8e\x33\x8c" +
		"\xd4\xf7\x68\xc2\x4d\x1c\x28\xa2\x6f\x4f\x07\x98\x31\xbc\xf6\x10\x96\x14\xfb\x09\x25\x04\xf8\xcc\xd3\x32\x07\xfb" +
		"\x93\x1b\x13\x5a\x7b\x0e\xa8\x1d\x5a\xca\x97\xc2\xb3\xa7\x1c\x57\x7d\xcd\x7e\xeb\x0c\x63\x43\x7b\x4e\x49\x8d\x0d" +
		"\x80\x1c\x3c\xe2\xec\x9d\x3c\xe2\xd4\xb3\x77\xfe\xc0\x46\xc4\xf6\xdc\x8d\xad\x78\x2f\xbd\xd1\x22\xad\x26\x4d\xda" +
		"\x4f\x82\x6c\xb8\xb1\x52\x7f\xf8\x46\x5f\xcb\xba\x55\x12\x63\x95\x32\x20\xa0\x34\xe5\xf5\x5f\x46\xb5\xf0\x34\x37" +
		"\x61\x68\xd0\x0e\xa3\x04\x56\xd2\x10\x4a\xe3\x44\xf7\x19\xa4\xef\x77\x0a\xb0\x5c\xee\xd3\xbf\x16\x22\xb6\x34\xde" +
		"\x1e\x12\xcc\x63\x90\x87\xa7\xe8\x80\x04\x82\x33\xca\xe1\x08\x1d\x1a\xc5\x1c\x9e\x22\xf7\xe6\x4f\x63\x15\x64\xb4" +
		"\x49\xde\xcb\x22\x63\x0f\xec\xef\xba\xce\xa8\xb3\xfe\x4f\x5c\x03\x17\x01\xfa\xfe\xb1\xb7\x6e\x95\xb5\xe3\xb7\x06" +
		"\xb8\xa9\x4b\x49\xa0\x4c\x46\xf0\xfa\x6c\x34\x21\xc4\x0b\x5b\x3d\xb7\xc7\xa3\x6b\x3c\x20\xc7\x04\x7a\x38\x2e\x32" +
		"\x9b\xb6\x7a\xd9\x10\x57\xbe\xa2\x31\x7f\x2c\x56\xfa\x83\xd8\xf7\x55\x08\x29\x42\xd9\xd1\x53\x9e\x6b\x7f\xc2\xcf" +
		"\xc3\x1e\x23\xc3\xee\xd7\x86\xb4\x9e\xca\x9a\x5c\x11\xfc\x08\x6b\xd7\x72\xfc\xb2\xce\x3a\xf6\x29\xaa\x9d\xcf\xe6" +
		"\x8e\x99\x2b\x0c\x5b\x57\x7d\x44\x26\x0a\xca\x07\xb9\x68\x10\xaf\x80\x93\x33\xc6\x6c\xf9\xd8\xb7\xb5\x6b\x57\xb0" +
		"\x48\x3b\xa8\x6a\x1e\xdf\xba\xee\xc1\xd2\x34\xdc\x10\x6f\x1f\xdb\x6e\x4d\xef\x6a\x86\xeb\xc4\x1f\xd9\xd1\x6b\x2a" +
		"\x95\x2e\xc7\x42\x22\x53\xd7\x53\xa5\x81\xeb\x1f\xb0\x4a\x06\x7e\x29\xd0\xd9\x12\x3f\xa2\x0a\xdf\xdb\x3c\xe9\x94" +
		"\xe8\x46\xc3\x6a\x2c\x21\x9f\x98\xfa\x0f\x2e\x01\x3b\x06\x6a\x2f\x86\x79\xa8\x80\x9f\x98\x8d\xc6\xc6\xf9\xc5\xd8" +
		"\x8d\xf3\x8b\xbd\x37\xce\x1f\x91\xa5\xb2\xef\x73\x4c\x6d\x86\x7b\x84\xb9\x0a\xc5\xce\xb1\x4f\xf9\xe5\x31\x79\xec" +
		"\xdd\x48\x2a\xad\xb1\x28\x5b\x4a\x4b\x6c\xd4\x30\x83\xbb\x48\x4f\x48\xf6\xfb\xb2\x4e\x29\x43\x40\xb3\x6a\x19\x49" +
		"\x78\x47\x46\x7e\xf0\xbd\xaf\x2a\x70\x8f\xde\xfd\xba\xff\xdf\xa9\x6e\xef\x7b\x35\x6f\x40\x54\x83\xda\x65\x87\xcd" +
		"\xfd\x87\x03\xdf\x47\xe1\xe6\xbe\x83\xdd\x0d\x72\xaf\xbf\x73\xd7\x95\x50\x24\x24\xa0\x7f\xe0\x15\xfe\x60\xaf\xde" +
		"\x58\x64\xb3\xbd\xff\xd5\x2e\x1a\xa1\xf7\xc6\xd1\x08\xc2\x1a\xe9\x04\x10\x70\x82\xc4\xd2\x3e\xba\x0b\x3b\x48\x09" +
		"\x3b\xce\x70\x0c\x0a\x31\x81\x09\x5a\x62\xa5\x61\x73\x63\xa7\xed\x22\x11\xbe\xc4\x37\x41\x2c\x44\xcc\x00\x67\x54" +
		"\xd9\xdb\x44\xe6\x5d\xc8\xe8\x42\x85\x97\x1f\x73\x90\xeb\xf0\x24\x38\x39\x09\x4e\xca\xd1\xf0\x25\xa5\xf1\x57\xbc" +
		"\x2e\xb7\x6f\x97\x35\xf1\x76\x11\x1d\x09\x02\x01\xe6\xf6\x34\xef\x52\x05\x42\xc6\xe1\x49\x30\x09\x4e\x8e\xc3\xf2" +
		"\xe5\xf8\x8b\x54\x83\xa8\x7c\x09\x4a\xe4\x32\x82\x31\x7c\x47\x84\x5f\xaa\x20\x62\x22\x27\x4b\x86\x25\x6c\x89\xd3" +
		"\xa1\xcc\xa9\x5f\x49\xe2\xd8\x08\xf7\x38\xac\xbf\xf3\x75\xc6\xd4\x80\x2c\xec\xfd\xb6\x2e\x72\x0a\xf7\x33\x0d\x50" +
		"\xc8\x63\x5f\x43\x9a\x31\xac\xc1\x43\x94\xcc\xbc\x3c\x23\x58\x17\xdb\x1f\x6f\x05\xc1\x2c\x48\x74\xca\x5a\xae\x66" +
		"\xa5\xe6\xe3\xf6\xdd\xab\x69\xf2\xb2\xf9\xdd\xde\x4e\xf4\xe6\xbf\x5a\xa4\xc8\xfa\xf6\x29\xba\xbd\xb5\x0f\x6e\x47" +
		"\x2b\x79\xd9\x70\x26\xb4\xb5\x8d\xbc\xbc\xc3\x0e\xf2\xce\x56\xf1\x2e\x07\x26\x1e\xd4\x2f\x71\xd5\x00\xaa\x35\xb6" +
		"\x6e\x70\xe1\x05\x30\x43\xc1\xcc\x33\x59\xcc\x9b\xbf\x2f\x72\xd9\x34\xb4\x5f\x1a\xb0\xdb\xe9\xb4\xe0\xdc\x4c\x70" +
		"\xe1\xd9\xa2\x18\x53\x9f\x95\xb5\xeb\x98\xf2\xa0\x59\x4a\x1b\xa4\xb6\xf4\xed\xaa\xa2\xb7\x6a\xe7\x1a\x7c\xff\x89" +
		"\x79\xdf\x49\xf9\xf6\xcd\x97\xee\x8b\x30\x23\xc4\x8c\x09\x91\x5e\xb5\xf3\x3c\x56\xcc\x66\x82\x13\xb3\x45\x31\x42" +
		"\xcc\x27\x93\xaf\x83\xe3\xe0\x38\x38\x39\x9d\x1c\x1f\x7f\xd9\xfb\x53\x95\x96\x9a\xa5\x45\xf0\x66\xe1\x7d\x04\x5f" +
		"\xc0\x77\x08\xfe\x14\x1d\x26\x42\xe9\xd3\x4c\x48\x7d\xb8\x9f\xd0\x3b\xaf\x36\x16\x7e\xb0\x14\xa2\xb9\x13\xd2\x96" +
		"\xd4\x4d\x0a\x27\xb0\xc4\x39\xd3\x5e\x6d\xb7\x21\xc2\x3c\x02\xf6\xfc\x85\x37\x3f\x67\x42\xc1\x6e\xae\xee\x28\x10" +
		"\xca\xca\x60\xab\x02\x58\x36\x92\x7f\x6d\x19\x71\x65\x96\x28\x62\xc9\xf6\x1a\x8d\xa4\xec\x92\x77\x15\xfb\xa6\x61" +
		"\x91\xf7\xa7\x61\x71\x2b\xfd\xff\x01\x00\x00\xff\xff\xc6\x45\xd2\x30\xa6\x3e\x00\x00")

func bindataAdminhttpassetsIndexhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAdminhttpassetsIndexhtml,
		"admin_http_assets/index.html",
	)
}

func bindataAdminhttpassetsIndexhtml() (*asset, error) {
	bytes, err := bindataAdminhttpassetsIndexhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "admin_http_assets/index.html",
		size:        16038,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1557215618, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"admin_http_assets/app.css":    bindataAdminhttpassetsAppcss,
	"admin_http_assets/app.js":     bindataAdminhttpassetsAppjs,
	"admin_http_assets/index.html": bindataAdminhttpassetsIndexhtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"admin_http_assets": {Func: nil, Children: map[string]*bintree{
		"app.css":    {Func: bindataAdminhttpassetsAppcss, Children: map[string]*bintree{}},
		"app.js":     {Func: bindataAdminhttpassetsAppjs, Children: map[string]*bintree{}},
		"index.html": {Func: bindataAdminhttpassetsIndexhtml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
