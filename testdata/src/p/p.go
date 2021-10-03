package p

import "fmt"

var _ int = 1
var sample int = 2
var Sample int = 3
var sample1 int = 4 // want "var identifier should follow the coding rules"

const SAMPLE int = 1
const SAMPLE_SAMPLE int = 2
const _SAMPLE int = 3
const SamPle int = 4  // want "const identifier should follow the coding rules"
const SAMPLE1 int = 5 // want "const identifier should follow the coding rules"

var ( // want "var identifier should follow the coding rules"
	sampleInVarBlock  int = 1
	sample1InVarBlock int = 1
)

const ( // want "const identifier should follow the coding rules"
	SAMPLE_IN_CONST_BLOCK  int = 1
	SAMPLE1_IN_CONST_BLOCK int = 1
)

func SampleFunc() {
	var _ int = 1
	var sampleInFunc int = 2
	var SampleInFunc int = 3
	var sample1InFunc int = 4 // want "var identifier should follow the coding rules"

	const SAMPLEINFUNC int = 1
	const SAMPLE_SAMPLE_INFUNC int = 2
	const _SAMPLE_INFUNC int = 3
	const SamPle_INFUNC int = 4  // want "const identifier should follow the coding rules"
	const SAMPLE1_INFUNC int = 5 // want "const identifier should follow the coding rules"

	fmt.Println(sampleInFunc)
	fmt.Println(sample1InFunc)
	fmt.Println(SampleInFunc)
}
