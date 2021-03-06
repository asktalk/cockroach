package securitytest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _test_certs_readme_md = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\x41\x4e\xc3\x40\x0c\x45\xf7\x73\x8a\xbf\x1e\x11\xc4\x86\x05\xd9\xa1\x4a\x9c\x80\x03\x64\x3a\x71\x88\xc5\x60\x57\xb6\x5b\xda\xdb\x13\x12\x84\xa2\x2e\xed\xf7\xdf\xd7\xcf\x19\xef\xe4\x81\x4a\x16\x3c\x71\x2d\x41\x8e\x91\x8d\x6a\xa8\xdd\x90\x73\x4a\x07\x95\x28\x2c\x8e\x98\x09\x93\xb6\xa6\xdf\x2c\x1f\x98\xb8\x91\xf7\x29\x65\xd4\xf2\x58\x2d\x7a\x1c\x5e\xf7\x35\x1b\xf8\xa4\xdb\x0a\x4e\xc6\x97\xe5\x89\xe5\x5e\x80\xe8\x48\x9b\xe3\x64\x17\xb2\x3b\x6f\xc5\xab\xf9\x87\xf7\x76\x7a\x53\x43\xc1\x7c\xfe\x2a\xd2\x19\x95\xb1\x1c\x1b\x61\x49\x39\xab\x40\xa7\x75\xe6\xae\xef\x01\x76\x96\x3e\x0d\xc3\x70\x2c\x3e\x27\x3d\x91\xb8\x37\x5c\x9f\x9f\x5e\xd0\xb1\xfc\x6f\x41\x17\x74\x8d\xdf\x5c\xfa\x09\x00\x00\xff\xff\xe0\x4c\xe7\x9d\x15\x01\x00\x00")

func test_certs_readme_md_bytes() ([]byte, error) {
	return bindata_read(
		_test_certs_readme_md,
		"test_certs/README.md",
	)
}

func test_certs_readme_md() (*asset, error) {
	bytes, err := test_certs_readme_md_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_certs/README.md", size: 277, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certs_ca_crt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\x4f\x77\x9a\x40\x14\xc5\xf7\x7c\x8a\xee\x73\x7a\x40\xa5\x4a\x17\x59\xbc\x37\x33\x0c\x63\x32\x34\xc3\x1f\x23\xec\x4c\x05\x61\x50\x51\x4b\x19\xf0\xd3\xb7\xda\xd3\x2c\xda\xbc\xe5\x6f\x71\xdf\xef\xdc\xfb\xf9\x76\xc8\xb8\x08\x3f\x11\x16\x25\xc2\x17\x04\x12\x76\xa7\x96\x14\x02\x63\x4d\xd0\x3b\x03\x85\x1d\x23\x50\x01\x72\x59\xed\xdc\xe6\x44\xf8\xdc\x53\x9a\x2d\xa6\xaa\xf1\x8d\x84\x96\x13\x72\xe6\xb1\x74\xbf\x22\x48\x22\x51\x0d\x4c\x83\xb2\x70\x17\xae\x10\xda\x84\xa4\x61\x9f\x4d\xbb\xf1\x6d\xea\xeb\x0d\x85\xc2\x37\xce\x10\x26\xe0\x48\xcd\x86\xf0\x2a\x26\x92\x56\x9b\x3b\xd3\x37\x06\xef\xcc\xfa\x9b\xf4\x71\x10\xe6\x12\x25\xc7\xf1\xcf\xe7\xdf\x86\xef\x16\x60\x58\x00\x8e\x00\x0b\xb9\x7a\xa9\xb6\xb0\x7b\xed\xf2\x72\x11\xa5\xe2\xc1\xc3\x1f\x01\x39\x16\xeb\x4b\x55\x07\xe9\x91\xcf\xea\xd0\x73\xfd\xe0\xea\xd3\xf6\x90\x2c\x7e\x8e\x2f\xd9\x35\x20\x87\x72\xa4\xe9\xda\xe3\xb2\xb0\xad\xba\x5f\xda\xac\xd3\x63\xb6\xef\x5e\x35\xa7\x45\x3a\x6f\xf2\x62\xbe\xd1\x4b\x0d\x8d\x04\x97\xc3\x24\xdd\x52\xc3\xd0\x36\x8a\x81\x11\xf0\x85\x42\x7c\xb3\x0d\x22\x89\x50\x7a\xcc\x22\x14\x38\xa8\xc0\xbe\x15\xf8\x6f\x51\xe0\x7c\x07\xc9\x14\x11\x64\xc9\xbb\xa6\xd8\xcb\xa7\xe7\xd2\x3e\x43\x3b\xd0\x88\x45\x01\x3b\x8f\xf1\xfa\xc1\xba\x98\xf0\x64\x4f\x4f\x99\x39\x16\xfd\x9b\xca\x5d\xa8\xb1\x2b\x5b\x99\xf5\xc7\x67\x9c\xf5\x97\xcd\xc4\x9f\x1c\xca\x3c\x80\x55\x5c\xd4\x87\x61\xb6\x67\xe3\xb7\xd5\x7e\xfb\x74\x0d\xd1\x91\x83\x79\x7c\xb4\xee\x63\xb2\x90\xfe\x3f\xf0\xaf\x00\x00\x00\xff\xff\x62\xa7\x34\xb7\xfd\x01\x00\x00")

func test_certs_ca_crt_bytes() ([]byte, error) {
	return bindata_read(
		_test_certs_ca_crt,
		"test_certs/ca.crt",
	)
}

func test_certs_ca_crt() (*asset, error) {
	bytes, err := test_certs_ca_crt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_certs/ca.crt", size: 509, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certs_ca_key = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x31\x53\x83\x30\x18\x40\x77\x7e\x05\x7b\xcf\x6b\xab\xd6\x90\xa1\x43\x48\x3f\x42\x4c\x03\x04\x8a\x1e\x6c\x2d\x10\x8e\x80\x4d\x39\xe5\xe2\xf9\xeb\x3d\x3b\xfb\xd6\xb7\xbc\xf7\xf0\x47\x08\x8c\x27\x3e\x50\x3f\xcb\xf9\x1b\x39\x81\x2f\xa0\xba\x0b\x4f\xc6\x0d\x25\x0a\x80\xa7\xf5\x4d\x47\xe6\x82\x3f\x51\x70\x9d\x3b\x81\xc7\x58\xa5\x2f\x78\x99\xc6\x60\xf8\x18\x8f\x2d\xd5\x9a\xd8\xc8\x1c\x37\xad\x0b\x2d\xb1\x8c\xd2\x99\x15\xf2\x19\x7b\xc4\x41\x6c\x4b\x75\x50\x3d\x81\x9a\xac\xa2\x0d\x0d\xcf\xdb\x69\xfd\x15\x29\x83\x1c\x73\x8d\x68\x77\xdd\xc2\x78\x5b\x34\x97\x8e\x3f\xfd\xf4\xa5\x96\x65\x3a\xd4\xd9\x82\x38\x36\xb2\x11\x5e\xbd\xe6\x49\xae\x5d\xf5\x8d\x56\x22\xb8\x06\xc5\x63\xf6\x3a\xbc\x6f\xcf\xb2\x4a\x76\xa7\x39\x9f\xd0\xad\xdf\xef\xbd\x7b\x2c\x24\x87\x7f\x1f\x7e\x03\x00\x00\xff\xff\xce\x4c\x11\x5d\xe3\x00\x00\x00")

func test_certs_ca_key_bytes() ([]byte, error) {
	return bindata_read(
		_test_certs_ca_key,
		"test_certs/ca.key",
	)
}

func test_certs_ca_key() (*asset, error) {
	bytes, err := test_certs_ca_key_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_certs/ca.key", size: 227, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certs_node_crt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\xcd\x72\xa2\x40\x14\x85\xf7\x3c\xc5\xbc\xc0\x14\xa8\x24\xe2\x22\x8b\x7b\xfb\x07\x1b\x6d\x48\x23\x88\x64\x87\x18\x1b\x10\x7f\x48\xc6\x69\xf4\xe9\x47\xad\x4c\xcd\x62\x2a\x77\xf9\x2d\xbe\x73\xea\xd6\xf9\x79\x3f\x64\xbe\x08\x7f\x10\x16\x27\x82\x0b\x02\x09\x7b\x50\x4b\x0a\x81\xfa\x4a\x08\x2c\x3a\x0d\x46\x20\x68\xa1\xd2\x40\xe4\xbf\xa5\xec\xa3\x90\x76\xa3\xf9\xc0\x44\x82\x73\x9b\xc2\x0c\xb5\xee\xaa\x5d\x13\xbd\x2a\x45\xa1\x81\x54\xc6\xc2\x58\x0c\x72\xba\x54\x6a\xc6\x4c\x4b\xd7\xc3\xf0\xa3\xdc\x4f\xaa\x7c\xa8\xcd\xb4\x2a\x43\x99\xa4\x26\xa4\xa2\x97\x0d\x18\x49\xe1\x29\x7b\xb0\xfc\xce\xcc\x5f\x66\x65\x5f\xa6\xef\x44\x59\x02\x09\xea\xf2\x2b\x59\xe0\xbf\x16\x12\xd1\x84\xc4\x02\x48\x4e\x83\xcb\xb6\x6c\x52\x3e\x91\x05\xd2\xeb\x7e\x5e\xe6\x98\xce\x2e\xd7\x43\x4b\xdb\x55\xc1\x73\xbb\xb0\xcb\xa8\xad\xc7\xd1\xf8\x1c\x10\x67\xe0\x79\xc5\x75\x97\x39\xdd\xbb\xe3\x38\xae\x08\x2c\x7b\xeb\x69\x77\xc8\x90\x06\xd5\x78\x2d\xab\x7a\x45\xcf\x4f\x75\x9f\xc1\x71\xe0\x9a\x15\x85\x08\x75\xb8\x9c\x2a\x0f\x61\xeb\x31\x04\x49\x20\x02\x33\x55\xb7\xb6\xb1\xd3\x22\xe6\xc6\xe2\x90\x8b\x99\xc9\x11\x55\x3a\x05\xc3\x7c\x42\x3e\x7d\x50\x29\x47\x23\x89\x04\xcf\x87\x41\xba\x61\x86\xa1\x6d\x14\x97\x77\x81\x12\xc6\xd7\x0f\x41\x8c\x28\x6f\x4f\xcc\x45\xb0\xf6\x27\x4d\x9e\xf5\xc7\xf5\x28\x74\x2a\x13\xdb\x00\x80\x12\x8e\x37\x59\xe7\x2f\xa4\x3b\x79\x24\x3b\x25\x48\xa6\x88\x20\x67\xfe\x3c\xd2\xd9\xeb\xaa\x10\x9f\x1f\xbf\xa8\x15\x19\x6e\x8a\xa0\xf6\x76\xe5\xa2\x3e\xf5\xe5\x69\x26\xdf\x7c\x1a\xbc\xd5\x17\xe2\x08\xa8\x01\x3b\x7a\x18\x0e\xcb\xe7\x25\x87\xf7\xe5\xee\x14\xc4\xbd\xbb\x27\x87\x76\x95\xd8\x61\x73\xd9\x1c\x27\xa9\xb5\xec\xdc\xdd\x9c\x6c\xf8\x42\xbf\xbc\x58\x8f\x69\xb0\x90\xfe\x3f\x97\x3f\x01\x00\x00\xff\xff\xb9\x94\x37\x92\x4b\x02\x00\x00")

func test_certs_node_crt_bytes() ([]byte, error) {
	return bindata_read(
		_test_certs_node_crt,
		"test_certs/node.crt",
	)
}

func test_certs_node_crt() (*asset, error) {
	bytes, err := test_certs_node_crt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_certs/node.crt", size: 587, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certs_node_key = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\xc1\x6e\x82\x40\x10\x40\xef\x7c\x05\x3f\xd0\xb4\x65\xdd\xaa\x07\x0f\xc3\x32\x1d\x27\x84\x95\x69\x29\x66\x8f\xd5\x96\x95\xd8\xb0\x1a\x40\x52\xbf\xbe\xa9\x67\xdf\xf5\x5d\xde\x7b\xf8\x27\x45\x62\x1b\xa3\x89\xcb\x37\xae\xa1\xc2\x38\x47\x77\x13\x51\xb1\xde\x1b\x10\x44\xce\xb9\x3c\xab\xcc\x6b\x29\x82\x7f\xea\x02\xed\xf0\xe8\x4e\x10\x76\xd3\x71\xeb\x67\xcb\x8a\xf3\xb6\x5a\xa4\x94\x5e\x3e\x61\x08\x10\xc8\x98\x33\xbd\x17\xb3\x65\x04\x13\xae\xc3\x87\x64\xe2\x01\x5f\xbe\xf6\x9d\xe2\xe7\xb4\xa9\xc8\xcb\x42\xb7\x8a\xa0\x36\x7d\xaf\xb5\xe8\x3a\x39\x6c\xcb\xe4\xa2\xb2\x93\x1b\xaf\xe3\xbc\x95\xc1\x36\x25\xf5\xfa\x35\x1a\xf2\x6e\xb0\x76\x63\x4c\xa3\x1e\x79\x63\x0f\x20\xbf\xee\x3b\xb9\xb2\xfb\x99\xe6\xa3\xeb\x6b\x0f\xab\x55\x74\x8b\x45\x9b\xdd\x7d\xf8\x0b\x00\x00\xff\xff\x53\x94\x2a\x24\xe3\x00\x00\x00")

func test_certs_node_key_bytes() ([]byte, error) {
	return bindata_read(
		_test_certs_node_key,
		"test_certs/node.key",
	)
}

func test_certs_node_key() (*asset, error) {
	bytes, err := test_certs_node_key_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_certs/node.key", size: 227, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
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
	"test_certs/README.md": test_certs_readme_md,
	"test_certs/ca.crt":    test_certs_ca_crt,
	"test_certs/ca.key":    test_certs_ca_key,
	"test_certs/node.crt":  test_certs_node_crt,
	"test_certs/node.key":  test_certs_node_key,
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

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test_certs": {nil, map[string]*_bintree_t{
		"README.md": {test_certs_readme_md, map[string]*_bintree_t{}},
		"ca.crt":    {test_certs_ca_crt, map[string]*_bintree_t{}},
		"ca.key":    {test_certs_ca_key, map[string]*_bintree_t{}},
		"node.crt":  {test_certs_node_crt, map[string]*_bintree_t{}},
		"node.key":  {test_certs_node_key, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
