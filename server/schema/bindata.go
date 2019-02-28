// Code generated by go-bindata.
// sources:
// schema.graphql
// types/chat.graphql
// types/circle.graphql
// types/message.graphql
// types/user.graphql
// types/user_settings.graphql
// DO NOT EDIT!

package schema

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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x31\x8f\x83\x30\x0c\x85\x77\xff\x0a\xdf\xc6\xfd\x85\xac\xb0\x30\x30\xdc\xc0\x74\xea\x60\x81\xd5\x20\x15\x4a\x13\x67\x40\x15\xff\xbd\x22\x4e\x18\x2a\x01\x13\xcf\xbc\x2f\x2f\x2f\xf6\x9d\xe5\x91\xf0\x0d\x88\xaf\xc0\x6e\x31\xf8\xb7\x7d\x00\x71\x0c\x42\x32\x3c\x27\x83\x4d\x52\xb0\x02\xc8\x32\xb3\x22\xf1\x4c\xf0\xec\x8a\x61\x32\xd8\x7a\x76\xf5\x34\x07\xf9\x55\x0d\x88\xdd\xe0\xba\x07\x47\xb7\x8c\x32\xfb\x3a\x6d\x84\x25\x51\xdf\x92\xec\xae\x25\xd9\xae\x67\xef\xe9\xce\x3e\xfa\x4d\x1a\x32\xf3\x9f\x7e\xdc\xf6\x4a\xb9\x63\x6c\x45\x7d\xdf\x9e\x14\x0b\x73\x4f\xc2\xc7\x84\x26\x94\x97\xf5\x35\xe7\x8a\x4b\x69\x27\x4f\x4d\x39\x87\x84\x26\xe4\x1d\x14\xba\xd7\xba\x32\x58\x57\x3f\xdf\x29\xe7\xd4\x0a\x9f\x00\x00\x00\xff\xff\xdd\xe8\x14\x6a\xf0\x01\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 496, mode: os.FileMode(420), modTime: time.Unix(1550867808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesChatGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\xce\x48\x2c\x51\xa8\xe6\x52\x50\xc8\x4c\xb1\x52\xf0\x74\x51\xe4\x52\x50\xc8\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x07\xf1\x93\x8b\x52\x13\x4b\x52\x53\x1c\x4b\xac\x14\x42\x32\x73\x53\xc1\x42\x99\x45\xc9\x39\xa9\xc5\x56\x0a\xd1\xce\x60\x56\x2c\x57\x2d\x17\x17\xdc\x3c\xcf\xbc\x82\x52\x64\x43\xd1\xcc\xe4\xaa\xe5\x02\x04\x00\x00\xff\xff\xa0\x44\xd9\xe4\x7b\x00\x00\x00")

func typesChatGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesChatGraphql,
		"types/chat.graphql",
	)
}

func typesChatGraphql() (*asset, error) {
	bytes, err := typesChatGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/chat.graphql", size: 123, mode: os.FileMode(420), modTime: time.Unix(1550867857, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesCircleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\xce\x2c\x4a\xce\x49\x55\xa8\xe6\x52\x50\xc8\x4c\xb1\x52\xf0\x74\x51\xe4\x52\x50\x48\xce\x48\x2c\xf1\x74\x81\xf1\xf2\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xc1\xb2\x45\xa9\x89\x25\xa9\x29\x8e\x25\x56\x0a\x21\x99\xb9\xa9\x20\xa1\xdc\xd4\xe2\xe2\xc4\xf4\xd4\x62\x2b\x85\x68\x5f\x08\x33\x96\xab\x96\x8b\x2b\x33\xaf\xa0\xb4\x04\x6a\x85\x27\x98\x8d\xb0\x07\xc5\x1a\x34\x5b\xb8\x6a\xb9\x00\x01\x00\x00\xff\xff\x55\xef\xa9\xc0\x9d\x00\x00\x00")

func typesCircleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesCircleGraphql,
		"types/circle.graphql",
	)
}

func typesCircleGraphql() (*asset, error) {
	bytes, err := typesCircleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/circle.graphql", size: 157, mode: os.FileMode(420), modTime: time.Unix(1550867849, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesMessageGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x3d\x0a\xc2\x40\x14\x04\xe0\xfe\x9d\x62\x72\x8d\xed\x84\x34\x5b\x58\xe9\x05\x96\xec\x10\x1f\xe8\x33\xec\x8e\x85\x84\xdc\x5d\x08\xf8\x83\xe5\x37\x03\x33\x7a\x2e\xc4\x91\xbd\x97\x99\x58\x0d\xf0\x9a\x90\xc7\xc1\x80\xe9\x52\x94\xc7\x8f\xbc\x4d\x57\x7e\xdd\x19\x95\xed\xa7\xbf\x87\x18\x4a\x38\xa9\x79\xcc\x7b\xd4\x58\xc4\x7a\x50\xc2\xd9\x6f\x1c\x6c\x33\xf3\x58\x1e\x7a\x1f\xf6\xbc\x6b\xfd\x9b\xb7\xcd\x5e\x01\x00\x00\xff\xff\x33\xfd\x3d\xae\x97\x00\x00\x00")

func typesMessageGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesMessageGraphql,
		"types/message.graphql",
	)
}

func typesMessageGraphql() (*asset, error) {
	bytes, err := typesMessageGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/message.graphql", size: 151, mode: os.FileMode(420), modTime: time.Unix(1550867889, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\xb1\x0a\xc2\x30\x14\x85\xe1\xfd\x3e\xc5\xf1\x35\xb2\x89\x2e\x5d\xba\xa8\x93\x38\x5c\xcd\xc1\x04\x4c\x0c\xc9\xed\x50\xa4\xef\x2e\x5d\x4a\x29\x1d\xcf\xcf\x81\xcf\xc6\x42\xdc\x1a\x2b\x7e\x02\x44\xef\xd0\x9d\x0f\x02\x64\x4d\x74\xb8\x58\x8d\xf9\x3d\xef\x12\xbe\x99\xfd\x90\x9e\xac\xeb\xec\x63\x2b\x1f\x1d\xfb\xcd\xfb\x55\xa9\x46\x7f\x34\x87\x6b\x4c\x9c\x4b\x50\x6b\x0e\xf7\x53\x50\x7b\xc8\x24\xb2\xc0\x5d\x2e\x83\xad\xf4\x0d\xbe\x6f\xef\xd3\x32\xc9\x3f\x00\x00\xff\xff\x6f\x3f\xea\x14\xd0\x00\x00\x00")

func typesUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesUserGraphql,
		"types/user.graphql",
	)
}

func typesUserGraphql() (*asset, error) {
	bytes, err := typesUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/user.graphql", size: 208, mode: os.FileMode(420), modTime: time.Unix(1550867875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesUser_settingsGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2d\x4e\x2d\x0a\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\xa8\xe6\x52\x50\xc8\x4c\xb1\x52\xf0\x74\x51\xe4\x52\x50\x28\x2d\x4e\x2d\xf2\x74\x81\xf1\x92\x8b\x52\x13\x4b\x52\x53\x1c\x4b\xac\x14\x42\x32\x73\x53\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\x34\x7b\x92\x8b\x40\x00\x00\x00")

func typesUser_settingsGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesUser_settingsGraphql,
		"types/user_settings.graphql",
	)
}

func typesUser_settingsGraphql() (*asset, error) {
	bytes, err := typesUser_settingsGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/user_settings.graphql", size: 64, mode: os.FileMode(420), modTime: time.Unix(1550867867, 0)}
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
	"schema.graphql": schemaGraphql,
	"types/chat.graphql": typesChatGraphql,
	"types/circle.graphql": typesCircleGraphql,
	"types/message.graphql": typesMessageGraphql,
	"types/user.graphql": typesUserGraphql,
	"types/user_settings.graphql": typesUser_settingsGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"types": &bintree{nil, map[string]*bintree{
		"chat.graphql": &bintree{typesChatGraphql, map[string]*bintree{}},
		"circle.graphql": &bintree{typesCircleGraphql, map[string]*bintree{}},
		"message.graphql": &bintree{typesMessageGraphql, map[string]*bintree{}},
		"user.graphql": &bintree{typesUserGraphql, map[string]*bintree{}},
		"user_settings.graphql": &bintree{typesUser_settingsGraphql, map[string]*bintree{}},
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
