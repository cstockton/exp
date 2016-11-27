package pathy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

var breaker = false
var sep = string(os.PathSeparator)
var lsep = string(os.PathListSeparator)
var dblSep = sep + sep
var P = func(p string) Path {
	return Path(p)
}

// @TODO Create one or two main structs with Paths that describe all of the
// components, i.e. name, stem, path, isAbs, etc and drive tests from that.

func ToSlash(p Path) Path {
	return Path(filepath.ToSlash(string(p)))
}

func ToColon(p PathList) PathList {
	if ':' == os.PathListSeparator {
		return p
	}
	return PathList(strings.Replace(string(p), lsep, ":", -1))
}

func FromSlash(p Path) Path {
	return Path(filepath.FromSlash(string(p)))
}

func FromColon(p PathList) PathList {
	if ':' == os.PathListSeparator {
		return p
	}
	return PathList(strings.Replace(string(p), ":", lsep, -1))
}

func StartPython() bool {
	return false
}

func CheckAgainstPathLib() bool {

	cmd := exec.Command("python", "-c", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	return false
}

const tmpDirName = "tmp_testing_"

var dirOnce sync.Once
var fileOnce sync.Once
var sharedTmpDir string
var sharedTmpFile *os.File
var tmpDirs []string
var tmpFiles []*os.File

func setup() {

}

func init() {

}

func cfatal(t *testing.T, err error) {
	// fmt.Println(err)
	if err != nil {
		fatal(t, err)
	}
}

func cfatalp(t *testing.T, err error, p Path) {
	if err != nil {
		fatalp(t, err, p)
	}
}

func fatal(t *testing.T, err error) {
	t.Fatalf("Fatal Error: err(%s)", err)
}

func fatalp(t *testing.T, err error, p Path) {
	t.Fatalf("Fatal Error: err(%s) pathy(%v)", err, p)
}

func shutdown() {
	for i := range tmpFiles {
		tmpFiles[i].Close()
	}
	for i := range tmpDirs {
		err := os.RemoveAll(tmpDirs[i])

		if err != nil {
			fmt.Println("Unable to remove path({0}) reason({1})", tmpDirs[i], err)
		}
	}
}

func createTempFile(t *testing.T) *os.File {
	tmpDir := getTempDir(t)
	tmpFile, err := ioutil.TempFile(tmpDir, "pathy_test_")
	if err != nil {
		t.Fatal(err)
	}

	tmpFiles = append(tmpFiles, tmpFile)

	return tmpFile
}

func createTempDir(t *testing.T) string {
	cwdDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tmpDir, err := ioutil.TempDir(cwdDir, tmpDirName)
	if err != nil {
		t.Fatal(err)
	}

	tmpDirs = append(tmpDirs, tmpDir)

	return tmpDir
}

func getTempDir(t *testing.T) string {
	onceBody := func() {
		sharedTmpDir = createTempDir(t)
	}
	dirOnce.Do(onceBody)
	return sharedTmpDir
}

func getTempFile(t *testing.T) *os.File {
	onceBody := func() {
		sharedTmpFile = createTempFile(t)
	}
	fileOnce.Do(onceBody)
	return sharedTmpFile
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}
