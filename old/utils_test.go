package old

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLineIterator(t *testing.T) {
	tmp := "/tmp/line_iterator.txt"
	content := "line1\n1\n"
	defer os.Remove(tmp)
	err := ioutil.WriteFile(tmp, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Tmp file couldn't be written becauser error %s", err.Error())
		os.Exit(1)
	}
	iter, err := NewLineIterator(tmp)
	defer iter.Close()
	switch {
	case err != nil:
		t.Error(err.Error())
	case iter.HasNext() && "line1" != iter.Next():
		t.Error("Line 1 is not expected(line1).")
	case iter.HasNext() && iter.Next() != "1":
		t.Error("Line 2 is not expected(1).")
	case iter.HasNext() && iter.Next() != "":
		t.Error("Line 3 is not expected().")
	case iter.HasNext() && iter.Next() != "" && !iter.HasNext():
		t.Error("Line 4 is not expected at all.")
	}
}

func TestLineIteratorFileDontExist(t *testing.T) {
	tmp := "/tmp/line_iterator_not.txt"
	iter, err := NewLineIterator(tmp)
	switch {
	case err == nil:
		t.Error("Error do not happened")
		iter.Close()
	}
}

func TestGenerateBase64StringSize(t *testing.T) {
	for i, n := range []int{0, 1, 2, 5, 10, 25, 30, 50, 75, 100} {
		hex := GenerateBase64String(n)
		if len(hex) != n {
			t.Errorf("%d - Size of GenerateBase64String(%d) should be %d, got %d instead.", i, n, n, len(hex))
			return
		}
	}
}

func BGenerateBase64StringN(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateBase64String(n)
	}
}

func BenchmarkGenerateBase64String1(b *testing.B) {
	BGenerateBase64StringN(1, b)
}

func BenchmarkGenerateBase64String5(b *testing.B) {
	BGenerateBase64StringN(5, b)
}

func BenchmarkGenerateBase64String10(b *testing.B) {
	BGenerateBase64StringN(10, b)
}

func BenchmarkGenerateBase64String25(b *testing.B) {
	BGenerateBase64StringN(25, b)
}

func BenchmarkGenerateBase64String50(b *testing.B) {
	BGenerateBase64StringN(50, b)
}

func BenchmarkGenerateBase64String75(b *testing.B) {
	BGenerateBase64StringN(75, b)
}

func BenchmarkGenerateBase64String100(b *testing.B) {
	BGenerateBase64StringN(100, b)
}
