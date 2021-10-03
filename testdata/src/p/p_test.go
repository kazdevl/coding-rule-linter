package p_test

import (
	"fmt"
	"testing"
	"time"
)

func TestMissingUnderScore(t *testing.T) {} // want "test func name should follow the coding rules"

func Test_Success(t *testing.T) {}

func NotTarget_HasResults() string {
	return ""
}

func NotTarget_HasTwoArgument(a int, b string) {
	fmt.Printf("%d: %s", a, b)
}

func NotTarget_HasTwoArgument_InOneGroup(a, b string) {
	fmt.Println(a + b)
}

func NotTarget_IsnotPointerArgument(a string) {
	fmt.Println(a)
}

func NotTarget_IsnotSelectorExpr(a *string) {
	fmt.Println(*a)
}

func NotTarget_IsnotTargetArgument(a *time.Time) {
	fmt.Println(*a)
}
