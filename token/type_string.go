// Code generated by "stringer -type Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LeftParen-0]
	_ = x[RightParen-1]
	_ = x[LeftBrace-2]
	_ = x[RightBrace-3]
	_ = x[Comma-4]
	_ = x[Dot-5]
	_ = x[Minus-6]
	_ = x[Plus-7]
	_ = x[Semicolon-8]
	_ = x[Slash-9]
	_ = x[Star-10]
	_ = x[Bang-11]
	_ = x[BangEqual-12]
	_ = x[Equal-13]
	_ = x[EqualEqual-14]
	_ = x[Greater-15]
	_ = x[GreaterEqual-16]
	_ = x[Less-17]
	_ = x[LessEqual-18]
	_ = x[Identifier-19]
	_ = x[String-20]
	_ = x[Number-21]
	_ = x[And-22]
	_ = x[Class-23]
	_ = x[Else-24]
	_ = x[False-25]
	_ = x[Fun-26]
	_ = x[For-27]
	_ = x[If-28]
	_ = x[Nil-29]
	_ = x[Or-30]
	_ = x[Print-31]
	_ = x[Return-32]
	_ = x[Super-33]
	_ = x[This-34]
	_ = x[True-35]
	_ = x[Var-36]
	_ = x[While-37]
	_ = x[EOF-38]
}

const _Type_name = "LeftParenRightParenLeftBraceRightBraceCommaDotMinusPlusSemicolonSlashStarBangBangEqualEqualEqualEqualGreaterGreaterEqualLessLessEqualIdentifierStringNumberAndClassElseFalseFunForIfNilOrPrintReturnSuperThisTrueVarWhileEOF"

var _Type_index = [...]uint8{0, 9, 19, 28, 38, 43, 46, 51, 55, 64, 69, 73, 77, 86, 91, 101, 108, 120, 124, 133, 143, 149, 155, 158, 163, 167, 172, 175, 178, 180, 183, 185, 190, 196, 201, 205, 209, 212, 217, 220}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
