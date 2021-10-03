package p

var _ int = 1
var sample int = 2
var sample1 int = 3 // want "var identifier should follow the coding rules"
var Sample int = 4  // want "var identifier should follow the coding rules"

const SAMPLE int = 1
const SAMPLE_SAMPLE int = 2
const _SAMPLE int = 3
const SamPle int = 4  // want "const identifier should follow the coding rules"
const SAMPLE1 int = 5 // want "const identifier should follow the coding rules"
