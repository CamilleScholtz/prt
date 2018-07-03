package ports

import "testing"

func BenchmarkLocate(b *testing.B) {
	PrtDir = "/usr/srt/prt"
	Order = []string{"opt", "6c37-git", "contrib", "6c37"}
	a, _ := All()

	for i := 0; i < b.N; i++ {
		Locate(a, "firefox")
	}
}

func BenchmarkAll(b *testing.B) {
	PrtDir = "/usr/srt/prt"

	for i := 0; i < b.N; i++ {
		All()
	}
}
