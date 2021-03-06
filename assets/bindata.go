package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

func templatesIndexHTMLBytes() ([]byte, error) {
	html, err := ioutil.ReadFile("./templates/index.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return html, err

	//return _templatesIndexHTML, nil
}

func templatesIndexHTML() (*asset, error) {
	bytes, err := templatesIndexHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func templatesNavigationBarHTMLBytes() ([]byte, error) {
	html, err := ioutil.ReadFile("./templates/navigation_bar.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return html, err

	//return _templatesNavigationBarHTML, nil
}

func templatesNavigationBarHTML() (*asset, error) {
	bytes, err := templatesNavigationBarHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/navigation_bar.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func pageFormSigninHTML() (*asset, error) {
	bytes, err := templatesNavigationBarHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/form_signin.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func templatesSecondViewHTMLBytes() ([]byte, error) {
	html, err := ioutil.ReadFile("./templates/second_view.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return html, err

	//return _templatesSecondViewHTML, nil
}

func templatesSecondViewHTML() (*asset, error) {
	bytes, err := templatesSecondViewHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/second_view.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func templatesThirdViewHTMLBytes() ([]byte, error) {
	html, err := ioutil.ReadFile("./templates/third_view.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return html, err
	//return _templatesThirdViewHTML, nil
}

func templatesThirdViewHTML() (*asset, error) {
	bytes, err := templatesThirdViewHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/third_view.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticNavigationBarCSS = []byte(``)

func staticNavigationBarCSSBytes() ([]byte, error) {
	return _staticNavigationBarCSS, nil
}

func staticNavigationBarCSS() (*asset, error) {
	bytes, err := staticNavigationBarCSSBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/navigation_bar.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func staticStyleCSSBytes() ([]byte, error) {
	css, err := ioutil.ReadFile("/static/style.css") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return css, err
	//return _staticStyleCSS, nil
}

func staticStyleCSS() (*asset, error) {
	bytes, err := staticStyleCSSBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "/static/style.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

func staticThirdViewCSSBytes() ([]byte, error) {
	css, err := ioutil.ReadFile("/static/third_view.css") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return css, err
	//return _staticThirdViewCSS, nil
}

func staticThirdViewCSS() (*asset, error) {
	bytes, err := staticThirdViewCSSBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/third_view.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/index.html":          templatesIndexHTML,
	"templates/navigation_bar.html": templatesNavigationBarHTML,
	"templates/second_view.html":    templatesSecondViewHTML,
	"templates/third_view.html":     templatesThirdViewHTML,
	"static/navigation_bar.css":     staticNavigationBarCSS,
	"static/style.css":              staticStyleCSS,
	"static/third_view.css":         staticThirdViewCSS,
	"pages/form_signin.html":        pageFormSigninHTML,
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
	"static": &bintree{nil, map[string]*bintree{
		"navigation_bar.css": &bintree{staticNavigationBarCSS, map[string]*bintree{}},
		"style.css":          &bintree{staticStyleCSS, map[string]*bintree{}},
		"third_view.css":     &bintree{staticThirdViewCSS, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"index.html":          &bintree{templatesIndexHTML, map[string]*bintree{}},
		"navigation_bar.html": &bintree{templatesNavigationBarHTML, map[string]*bintree{}},
		"second_view.html":    &bintree{templatesSecondViewHTML, map[string]*bintree{}},
		"third_view.html":     &bintree{templatesThirdViewHTML, map[string]*bintree{}},
	}},
	"pages": &bintree{nil, map[string]*bintree{
		"form_signin.html": &bintree{pageFormSigninHTML, map[string]*bintree{}},
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
