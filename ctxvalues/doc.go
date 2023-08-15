// Package ctxvalues make some attempts at demystifying the comment on
// https://pkg.go.dev/context#WithValue about avoiding allocations:
// To avoid allocating when assigning to an interface{}, context keys often
// have concrete type struct{}. Alternatively, exported context key variables'
// static type should be a pointer or interface.
package ctxvalues
