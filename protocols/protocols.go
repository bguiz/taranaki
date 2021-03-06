// Code generated by go-bindata.
// sources:
// .protocols/ApiKeys.proto
// DO NOT EDIT!

package protocols

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apikeysProto = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x41\x4b\xc3\x40\x10\x85\xef\xf3\x2b\x1e\x39\xe6\x52\xb4\xde\x42\x0e\x85\x4a\x11\x6f\xfe\x01\xd9\x26\x43\x1a\x9a\xee\xc4\x9d\x59\x30\x88\xff\x5d\xa2\xdd\x54\x21\x20\x6d\xe9\x69\x61\xf6\xbd\x37\x1f\x8f\xd1\xc1\x9b\x7b\x47\x89\xac\x0f\x62\xb2\xcc\x0a\xea\x5d\xb5\x77\x0d\x63\xd5\xb7\xcf\x3c\x68\x41\x24\xbd\xb5\xe2\xd1\xc8\x6b\xfa\x4b\xfa\x4a\xba\xac\x20\x3a\xb0\xea\xc9\x82\x0f\x02\xd4\x42\xeb\x1b\x04\xe9\x46\xf5\x5d\x71\x1a\x99\xec\xd9\xa3\xc4\xfd\x38\xdb\x8a\x74\x60\xef\xb6\x1d\xd7\x28\xb1\xfc\xa5\x8b\xca\xe1\x69\x8d\x12\x0f\x05\x7d\x12\x2d\xf2\x9c\x90\x63\x85\x17\x7e\x8b\xac\x06\xdb\x39\xc3\x20\x11\xba\x93\xd8\xd5\xa8\xa5\x8a\x07\xf6\x46\xc8\x17\x13\xd1\x86\x3d\x07\x67\x9c\x4c\xb3\x68\x7f\xd3\xb5\x17\xaf\x7c\x6e\xfc\xd1\x35\xe6\x1f\x5b\x70\x3f\xcf\xcc\x86\x33\xf8\x1f\xbf\xab\x99\xa1\x4f\x2d\x5e\x85\x9f\xd2\x27\xf8\x8b\x39\xd7\xad\xde\x12\x74\x8a\xbf\x9e\x74\xc3\x76\xb3\x63\xb0\xff\xef\xe0\x2b\x00\x00\xff\xff\x3e\x91\xe4\x32\x71\x03\x00\x00"

func apikeysProtoBytes() ([]byte, error) {
	return bindataRead(
		_apikeysProto,
		"ApiKeys.proto",
	)
}

func apikeysProto() (*asset, error) {
	bytes, err := apikeysProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ApiKeys.proto", size: 881, mode: os.FileMode(420), modTime: time.Unix(1516252639, 0)}
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
	"ApiKeys.proto": apikeysProto,
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
	"ApiKeys.proto": {apikeysProto, map[string]*bintree{}},
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
