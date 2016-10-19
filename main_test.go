package main

import (
	"testing"
)

func BenchmarkMut(b *testing.B){
	mut()
	for i:=0; i < b.N; i++ { // .N - iterations
		mut()
	}
}
