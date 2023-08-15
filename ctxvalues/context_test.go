package ctxvalues

import (
	"context"
	"math"
	"runtime"
	"testing"
)

func BenchmarkWithStructZeroVal(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx = WithStructZeroVal(ctx, math.MaxInt)
	}
	runtime.KeepAlive(ctx)
}

func BenchmarkStructZeroVal(b *testing.B) {
	ctx := WithStructZeroVal(context.Background(), math.MaxInt)
	var val int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val = StructZeroVal(ctx)
	}
	runtime.KeepAlive(val)
}

func BenchmarkWithStructPointer(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx = WithStructPointer(ctx, math.MaxInt)
	}
	runtime.KeepAlive(ctx)
}

func BenchmarkStructPointer(b *testing.B) {
	ctx := WithStructZeroVal(context.Background(), math.MaxInt)
	var val int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val = StructPointer(ctx)
	}
	runtime.KeepAlive(val)
}

func BenchmarkWithInt(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx = WithInt(ctx, math.MaxInt)
	}
	runtime.KeepAlive(ctx)
}

func BenchmarkInt(b *testing.B) {
	ctx := WithStructZeroVal(context.Background(), math.MaxInt)
	var val int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val = Int(ctx)
	}
	runtime.KeepAlive(val)
}

func BenchmarkWithString(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx = WithString(ctx, math.MaxInt)
	}
	runtime.KeepAlive(ctx)
}

func BenchmarkString(b *testing.B) {
	ctx := WithStructZeroVal(context.Background(), math.MaxInt)
	var val int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val = String(ctx)
	}
	runtime.KeepAlive(val)
}
