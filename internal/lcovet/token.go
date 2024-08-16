package lcovet

type Token int

const (
	ILLEGAL Token = iota

	LITERAL

	// Delimiter
	COLLON
	COMMA

	// Identifiers
	TN   // test name
	SF   // source file path
	FN   // line number,function name
	FNDA // function data
	FNF  //  number functions found
	FNH  // number hit
	BRDA // branch data: line, block, (expressions,count)+
	BRF  // branches found
	BRH  // branches hit
	DA   // line number, hit count
	LF   // lines found
	LH   //  lines hit.
	EOF  // end_of_file
)
