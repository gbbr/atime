package atime

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"
)

type testcase struct {
	setup []string
	touch []string
	want  []string
}

func TestTime(t *testing.T) {
	tmp := os.TempDir()
	dir := filepath.Join(tmp, "atime-test")
	wipe := func() error { return os.RemoveAll(dir) }
	err := os.Mkdir(dir, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer wipe()
	for i, test := range [...]testcase{
		0: {
			setup: []string{"1", "2", "3", "4", "5"},
			want:  []string{"1", "2", "3", "4", "5"},
		},
		1: {
			setup: []string{"1", "2", "3", "4", "5"},
			touch: []string{"2", "3"},
			want:  []string{"1", "4", "5", "2", "3"},
		},
		2: {
			setup: []string{"1", "2", "3", "4", "5"},
			touch: []string{"5", "4", "3", "2", "1"},
			want:  []string{"5", "4", "3", "2", "1"},
		},
	} {
		// create files
		for _, fname := range test.setup {
			absPath := filepath.Join(dir, fname)
			f, err := os.Create(absPath)
			if err != nil {
				t.Fatal(err)
			}
			f.Close()
		}
		// touch files
		for j, fname := range test.touch {
			absPath := filepath.Join(dir, fname)
			newTime := time.Now().Add(time.Duration((j+1)*10) * time.Minute)
			if err := os.Chtimes(absPath, newTime, time.Now()); err != nil {
				t.Fatal(err)
			}
		}
		// check order
		f, err := os.Open(dir)
		if err != nil {
			t.Fatal(err)
		}
		fis, err := f.Readdir(0)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := len(fis), len(test.want); got != want {
			t.Fatalf("#%d: expected %d files, got %d", i, want, got)
		}
		sort.Sort(Ascending(fis))
		got := make([]string, len(fis))
		for j, fi := range fis {
			got[j] = fi.Name()
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("#%d: wanted order %#v, got %#v", i, test.want, got)
		}
	}
}
