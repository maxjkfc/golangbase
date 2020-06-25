package main

import "testing"

func TestHello(t *testing.T) {
	str, err := Hello(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(str)
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Hello(true)
	}
	b.ReportAllocs()
}

