package ctxvalues

import "context"

type structZeroValKey struct{}

func WithStructZeroVal(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, structZeroValKey{}, val)
}

func StructZeroVal(ctx context.Context) int {
	if val, ok := ctx.Value(structZeroValKey{}).(int); ok {
		return val
	}
	return 0
}

type structPointerType struct{}

var structPointerKey = &structPointerType{}

func WithStructPointer(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, structPointerKey, val)
}

func StructPointer(ctx context.Context) int {
	if val, ok := ctx.Value(structPointerKey).(int); ok {
		return val
	}
	return 0
}

type intType int

const intKey intType = iota

func WithInt(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, intKey, val)
}

func Int(ctx context.Context) int {
	if val, ok := ctx.Value(intKey).(int); ok {
		return val
	}
	return 0
}

type stringType string

const stringKey stringType = "mycoolstringkey"

func WithString(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, stringKey, val)
}

func String(ctx context.Context) int {
	if val, ok := ctx.Value(stringKey).(int); ok {
		return val
	}
	return 0
}
