package lily

import (
	"testing"
)
//
// Author João Nuno.
// 
// joaonrb@gmail.com
//

func TestGenerateBase64StringSize(t *testing.T) {
	for i, n := range []int{0, 1, 2, 5, 10, 25, 30, 50, 75, 100} {
		hex := GenerateBase64String(n)
		if len(hex) != n {
			t.Errorf("%d - Size of GenerateBase64String(%d) sould be %d, got %d instead.", i, n, n, len(hex))
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