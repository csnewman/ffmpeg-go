// Code generated from C.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // C
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CParser struct {
	*antlr.BaseParser
}

var CParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cParserInit() {
	staticData := &CParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'*'", "'-'", "','", "';'", "'('", "')'", "'['", "']'", "'{'",
		"'}'", "'const'", "'unsigned'", "'void'", "'enum'", "'struct'", "'typedef'",
		"'attribute_deprecated'",
	}
	staticData.SymbolicNames = []string{
		"", "Assign", "Star", "Minus", "Comma", "Semi", "LParen", "RParen",
		"LBracket", "RBracket", "LBrace", "RBrace", "Const", "Unsigned", "Void",
		"Enum", "Struct", "Typedef", "AttributeDeprecated", "Ifndef", "Ifdir",
		"Endif", "Define", "Undef", "Pragma", "Include", "Identifier", "Whitespace",
		"Newline", "Constant", "DigitSequence", "StringLiteral", "BlockComment",
		"LineComment",
	}
	staticData.RuleNames = []string{
		"unit", "topLevel", "enumBody", "enumItem", "structItem", "constant",
		"params", "param", "arraySize", "arraySizeDim", "type", "typeInner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 288, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 5, 1, 35, 8, 1, 10, 1, 12, 1, 38, 9, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 43, 8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 55, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 73, 8,
		1, 10, 1, 12, 1, 76, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 1,
		1, 1, 3, 1, 86, 8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94,
		9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 99, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 107, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 121, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 130, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 138,
		8, 2, 11, 2, 12, 2, 139, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 150, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 156, 8, 3, 1, 4, 3, 4, 159,
		8, 4, 1, 4, 1, 4, 1, 4, 5, 4, 164, 8, 4, 10, 4, 12, 4, 167, 9, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 173, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 181, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 5, 4, 193, 8, 4, 10, 4, 12, 4, 196, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 210, 8, 4, 1, 5,
		3, 5, 213, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 221, 8, 6, 10,
		6, 12, 6, 224, 9, 6, 1, 6, 3, 6, 227, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 246, 8, 7, 1, 8, 4, 8, 249, 8, 8, 11, 8, 12, 8, 250, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 259, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 269, 8, 10, 10, 10, 12, 10, 272, 9,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 281, 8, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 286, 8, 11, 1, 11, 0, 1, 20, 12, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 330, 0, 27, 1, 0, 0, 0, 2, 129, 1,
		0, 0, 0, 4, 149, 1, 0, 0, 0, 6, 155, 1, 0, 0, 0, 8, 209, 1, 0, 0, 0, 10,
		212, 1, 0, 0, 0, 12, 226, 1, 0, 0, 0, 14, 245, 1, 0, 0, 0, 16, 248, 1,
		0, 0, 0, 18, 258, 1, 0, 0, 0, 20, 260, 1, 0, 0, 0, 22, 285, 1, 0, 0, 0,
		24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1,
		0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30,
		31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0, 32, 36, 5, 19, 0, 0, 33, 35, 3, 2, 1,
		0, 34, 33, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 39, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 130, 5, 21, 0,
		0, 40, 44, 5, 20, 0, 0, 41, 43, 3, 2, 1, 0, 42, 41, 1, 0, 0, 0, 43, 46,
		1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0,
		46, 44, 1, 0, 0, 0, 47, 130, 5, 21, 0, 0, 48, 130, 5, 22, 0, 0, 49, 130,
		5, 23, 0, 0, 50, 130, 5, 25, 0, 0, 51, 130, 5, 24, 0, 0, 52, 54, 5, 15,
		0, 0, 53, 55, 5, 26, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55,
		56, 1, 0, 0, 0, 56, 57, 5, 10, 0, 0, 57, 59, 3, 4, 2, 0, 58, 60, 5, 4,
		0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62,
		5, 11, 0, 0, 62, 63, 5, 5, 0, 0, 63, 130, 1, 0, 0, 0, 64, 66, 5, 18, 0,
		0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68,
		5, 17, 0, 0, 68, 69, 5, 16, 0, 0, 69, 70, 5, 26, 0, 0, 70, 74, 5, 10, 0,
		0, 71, 73, 3, 8, 4, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		77, 78, 5, 11, 0, 0, 78, 79, 5, 26, 0, 0, 79, 130, 5, 5, 0, 0, 80, 82,
		5, 18, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 85, 5, 16, 0, 0, 84, 86, 5, 18, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 5, 26, 0, 0, 88, 92, 5, 10, 0, 0,
		89, 91, 3, 8, 4, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95,
		96, 5, 11, 0, 0, 96, 130, 5, 5, 0, 0, 97, 99, 5, 18, 0, 0, 98, 97, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 17, 0, 0, 101,
		102, 3, 20, 10, 0, 102, 103, 5, 26, 0, 0, 103, 104, 5, 5, 0, 0, 104, 130,
		1, 0, 0, 0, 105, 107, 5, 18, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 5, 17, 0, 0, 109, 110, 3, 20, 10,
		0, 110, 111, 5, 6, 0, 0, 111, 112, 5, 2, 0, 0, 112, 113, 5, 26, 0, 0, 113,
		114, 5, 7, 0, 0, 114, 115, 5, 6, 0, 0, 115, 116, 3, 12, 6, 0, 116, 117,
		5, 7, 0, 0, 117, 118, 5, 5, 0, 0, 118, 130, 1, 0, 0, 0, 119, 121, 5, 18,
		0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0,
		122, 123, 3, 20, 10, 0, 123, 124, 5, 26, 0, 0, 124, 125, 5, 6, 0, 0, 125,
		126, 3, 12, 6, 0, 126, 127, 5, 7, 0, 0, 127, 128, 5, 5, 0, 0, 128, 130,
		1, 0, 0, 0, 129, 32, 1, 0, 0, 0, 129, 40, 1, 0, 0, 0, 129, 48, 1, 0, 0,
		0, 129, 49, 1, 0, 0, 0, 129, 50, 1, 0, 0, 0, 129, 51, 1, 0, 0, 0, 129,
		52, 1, 0, 0, 0, 129, 65, 1, 0, 0, 0, 129, 81, 1, 0, 0, 0, 129, 98, 1, 0,
		0, 0, 129, 106, 1, 0, 0, 0, 129, 120, 1, 0, 0, 0, 130, 3, 1, 0, 0, 0, 131,
		132, 5, 22, 0, 0, 132, 150, 3, 4, 2, 0, 133, 137, 5, 20, 0, 0, 134, 135,
		3, 6, 3, 0, 135, 136, 5, 4, 0, 0, 136, 138, 1, 0, 0, 0, 137, 134, 1, 0,
		0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 142, 5, 21, 0, 0, 142, 143, 3, 4, 2, 0, 143,
		150, 1, 0, 0, 0, 144, 145, 3, 6, 3, 0, 145, 146, 5, 4, 0, 0, 146, 147,
		3, 4, 2, 0, 147, 150, 1, 0, 0, 0, 148, 150, 3, 6, 3, 0, 149, 131, 1, 0,
		0, 0, 149, 133, 1, 0, 0, 0, 149, 144, 1, 0, 0, 0, 149, 148, 1, 0, 0, 0,
		150, 5, 1, 0, 0, 0, 151, 156, 5, 26, 0, 0, 152, 153, 5, 26, 0, 0, 153,
		154, 5, 1, 0, 0, 154, 156, 3, 10, 5, 0, 155, 151, 1, 0, 0, 0, 155, 152,
		1, 0, 0, 0, 156, 7, 1, 0, 0, 0, 157, 159, 5, 18, 0, 0, 158, 157, 1, 0,
		0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 165, 3, 20, 10,
		0, 161, 162, 5, 26, 0, 0, 162, 164, 5, 4, 0, 0, 163, 161, 1, 0, 0, 0, 164,
		167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168,
		1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 5, 26, 0, 0, 169, 170, 5, 5,
		0, 0, 170, 210, 1, 0, 0, 0, 171, 173, 5, 18, 0, 0, 172, 171, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 3, 20, 10, 0, 175,
		176, 5, 26, 0, 0, 176, 177, 3, 16, 8, 0, 177, 178, 5, 5, 0, 0, 178, 210,
		1, 0, 0, 0, 179, 181, 5, 18, 0, 0, 180, 179, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 3, 20, 10, 0, 183, 184, 5, 6, 0,
		0, 184, 185, 5, 2, 0, 0, 185, 186, 5, 26, 0, 0, 186, 187, 5, 7, 0, 0, 187,
		188, 3, 16, 8, 0, 188, 189, 5, 5, 0, 0, 189, 210, 1, 0, 0, 0, 190, 194,
		5, 20, 0, 0, 191, 193, 3, 8, 4, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0,
		0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0,
		196, 194, 1, 0, 0, 0, 197, 210, 5, 21, 0, 0, 198, 210, 5, 22, 0, 0, 199,
		200, 3, 20, 10, 0, 200, 201, 5, 6, 0, 0, 201, 202, 5, 2, 0, 0, 202, 203,
		5, 26, 0, 0, 203, 204, 5, 7, 0, 0, 204, 205, 5, 6, 0, 0, 205, 206, 3, 12,
		6, 0, 206, 207, 5, 7, 0, 0, 207, 208, 5, 5, 0, 0, 208, 210, 1, 0, 0, 0,
		209, 158, 1, 0, 0, 0, 209, 172, 1, 0, 0, 0, 209, 180, 1, 0, 0, 0, 209,
		190, 1, 0, 0, 0, 209, 198, 1, 0, 0, 0, 209, 199, 1, 0, 0, 0, 210, 9, 1,
		0, 0, 0, 211, 213, 5, 3, 0, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0,
		0, 213, 214, 1, 0, 0, 0, 214, 215, 5, 29, 0, 0, 215, 11, 1, 0, 0, 0, 216,
		227, 5, 14, 0, 0, 217, 218, 3, 14, 7, 0, 218, 219, 5, 4, 0, 0, 219, 221,
		1, 0, 0, 0, 220, 217, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0,
		0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0,
		225, 227, 3, 14, 7, 0, 226, 216, 1, 0, 0, 0, 226, 222, 1, 0, 0, 0, 227,
		13, 1, 0, 0, 0, 228, 229, 3, 20, 10, 0, 229, 230, 5, 26, 0, 0, 230, 246,
		1, 0, 0, 0, 231, 232, 3, 20, 10, 0, 232, 233, 5, 26, 0, 0, 233, 234, 3,
		16, 8, 0, 234, 246, 1, 0, 0, 0, 235, 246, 3, 20, 10, 0, 236, 237, 3, 20,
		10, 0, 237, 238, 5, 6, 0, 0, 238, 239, 5, 2, 0, 0, 239, 240, 5, 26, 0,
		0, 240, 241, 5, 7, 0, 0, 241, 242, 5, 6, 0, 0, 242, 243, 3, 12, 6, 0, 243,
		244, 5, 7, 0, 0, 244, 246, 1, 0, 0, 0, 245, 228, 1, 0, 0, 0, 245, 231,
		1, 0, 0, 0, 245, 235, 1, 0, 0, 0, 245, 236, 1, 0, 0, 0, 246, 15, 1, 0,
		0, 0, 247, 249, 3, 18, 9, 0, 248, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0,
		250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 17, 1, 0, 0, 0, 252, 253,
		5, 8, 0, 0, 253, 254, 5, 26, 0, 0, 254, 259, 5, 9, 0, 0, 255, 256, 5, 8,
		0, 0, 256, 257, 5, 29, 0, 0, 257, 259, 5, 9, 0, 0, 258, 252, 1, 0, 0, 0,
		258, 255, 1, 0, 0, 0, 259, 19, 1, 0, 0, 0, 260, 261, 6, 10, -1, 0, 261,
		262, 3, 22, 11, 0, 262, 270, 1, 0, 0, 0, 263, 264, 10, 2, 0, 0, 264, 269,
		5, 2, 0, 0, 265, 266, 10, 1, 0, 0, 266, 267, 5, 2, 0, 0, 267, 269, 5, 12,
		0, 0, 268, 263, 1, 0, 0, 0, 268, 265, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0,
		270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 21, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 273, 286, 5, 26, 0, 0, 274, 275, 5, 15, 0, 0, 275, 286, 5,
		26, 0, 0, 276, 277, 5, 16, 0, 0, 277, 286, 5, 26, 0, 0, 278, 280, 5, 13,
		0, 0, 279, 281, 5, 26, 0, 0, 280, 279, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0,
		281, 286, 1, 0, 0, 0, 282, 286, 5, 14, 0, 0, 283, 284, 5, 12, 0, 0, 284,
		286, 3, 22, 11, 0, 285, 273, 1, 0, 0, 0, 285, 274, 1, 0, 0, 0, 285, 276,
		1, 0, 0, 0, 285, 278, 1, 0, 0, 0, 285, 282, 1, 0, 0, 0, 285, 283, 1, 0,
		0, 0, 286, 23, 1, 0, 0, 0, 33, 27, 36, 44, 54, 59, 65, 74, 81, 85, 92,
		98, 106, 120, 129, 139, 149, 155, 158, 165, 172, 180, 194, 209, 212, 222,
		226, 245, 250, 258, 268, 270, 280, 285,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CParserInit initializes any static state used to implement CParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CParserInit() {
	staticData := &CParserStaticData
	staticData.once.Do(cParserInit)
}

// NewCParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCParser(input antlr.TokenStream) *CParser {
	CParserInit()
	this := new(CParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "C.g4"

	return this
}

// CParser tokens.
const (
	CParserEOF                 = antlr.TokenEOF
	CParserAssign              = 1
	CParserStar                = 2
	CParserMinus               = 3
	CParserComma               = 4
	CParserSemi                = 5
	CParserLParen              = 6
	CParserRParen              = 7
	CParserLBracket            = 8
	CParserRBracket            = 9
	CParserLBrace              = 10
	CParserRBrace              = 11
	CParserConst               = 12
	CParserUnsigned            = 13
	CParserVoid                = 14
	CParserEnum                = 15
	CParserStruct              = 16
	CParserTypedef             = 17
	CParserAttributeDeprecated = 18
	CParserIfndef              = 19
	CParserIfdir               = 20
	CParserEndif               = 21
	CParserDefine              = 22
	CParserUndef               = 23
	CParserPragma              = 24
	CParserInclude             = 25
	CParserIdentifier          = 26
	CParserWhitespace          = 27
	CParserNewline             = 28
	CParserConstant            = 29
	CParserDigitSequence       = 30
	CParserStringLiteral       = 31
	CParserBlockComment        = 32
	CParserLineComment         = 33
)

// CParser rules.
const (
	CParserRULE_unit         = 0
	CParserRULE_topLevel     = 1
	CParserRULE_enumBody     = 2
	CParserRULE_enumItem     = 3
	CParserRULE_structItem   = 4
	CParserRULE_constant     = 5
	CParserRULE_params       = 6
	CParserRULE_param        = 7
	CParserRULE_arraySize    = 8
	CParserRULE_arraySizeDim = 9
	CParserRULE_type         = 10
	CParserRULE_typeInner    = 11
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTopLevel() []ITopLevelContext
	TopLevel(i int) ITopLevelContext

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(CParserEOF, 0)
}

func (s *UnitContext) AllTopLevel() []ITopLevelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelContext); ok {
			tst[i] = t.(ITopLevelContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) TopLevel(i int) ITopLevelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *CParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CParserRULE_unit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132116480) != 0 {
		{
			p.SetState(24)
			p.TopLevel()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(CParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopLevelContext is an interface to support dynamic dispatch.
type ITopLevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopLevelContext differentiates from other interfaces.
	IsTopLevelContext()
}

type TopLevelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelContext() *TopLevelContext {
	var p = new(TopLevelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_topLevel
	return p
}

func InitEmptyTopLevelContext(p *TopLevelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_topLevel
}

func (*TopLevelContext) IsTopLevelContext() {}

func NewTopLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelContext {
	var p = new(TopLevelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_topLevel

	return p
}

func (s *TopLevelContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelContext) CopyAll(ctx *TopLevelContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TopLevelIfdirContext struct {
	TopLevelContext
}

func NewTopLevelIfdirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelIfdirContext {
	var p = new(TopLevelIfdirContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelIfdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelIfdirContext) Ifdir() antlr.TerminalNode {
	return s.GetToken(CParserIfdir, 0)
}

func (s *TopLevelIfdirContext) Endif() antlr.TerminalNode {
	return s.GetToken(CParserEndif, 0)
}

func (s *TopLevelIfdirContext) AllTopLevel() []ITopLevelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelContext); ok {
			tst[i] = t.(ITopLevelContext)
			i++
		}
	}

	return tst
}

func (s *TopLevelIfdirContext) TopLevel(i int) ITopLevelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelContext)
}

type TopLevelIncludeContext struct {
	TopLevelContext
}

func NewTopLevelIncludeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelIncludeContext {
	var p = new(TopLevelIncludeContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelIncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelIncludeContext) Include() antlr.TerminalNode {
	return s.GetToken(CParserInclude, 0)
}

type TopLevelUndefContext struct {
	TopLevelContext
}

func NewTopLevelUndefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelUndefContext {
	var p = new(TopLevelUndefContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelUndefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelUndefContext) Undef() antlr.TerminalNode {
	return s.GetToken(CParserUndef, 0)
}

type TopLevelStructContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelStructContext {
	var p = new(TopLevelStructContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelStructContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelStructContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelStructContext) Struct() antlr.TerminalNode {
	return s.GetToken(CParserStruct, 0)
}

func (s *TopLevelStructContext) LBrace() antlr.TerminalNode {
	return s.GetToken(CParserLBrace, 0)
}

func (s *TopLevelStructContext) RBrace() antlr.TerminalNode {
	return s.GetToken(CParserRBrace, 0)
}

func (s *TopLevelStructContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelStructContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelStructContext) AllAttributeDeprecated() []antlr.TerminalNode {
	return s.GetTokens(CParserAttributeDeprecated)
}

func (s *TopLevelStructContext) AttributeDeprecated(i int) antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, i)
}

func (s *TopLevelStructContext) AllStructItem() []IStructItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructItemContext); ok {
			len++
		}
	}

	tst := make([]IStructItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructItemContext); ok {
			tst[i] = t.(IStructItemContext)
			i++
		}
	}

	return tst
}

func (s *TopLevelStructContext) StructItem(i int) IStructItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructItemContext)
}

type TopLevelIfndefContext struct {
	TopLevelContext
}

func NewTopLevelIfndefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelIfndefContext {
	var p = new(TopLevelIfndefContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelIfndefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelIfndefContext) Ifndef() antlr.TerminalNode {
	return s.GetToken(CParserIfndef, 0)
}

func (s *TopLevelIfndefContext) Endif() antlr.TerminalNode {
	return s.GetToken(CParserEndif, 0)
}

func (s *TopLevelIfndefContext) AllTopLevel() []ITopLevelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelContext); ok {
			tst[i] = t.(ITopLevelContext)
			i++
		}
	}

	return tst
}

func (s *TopLevelIfndefContext) TopLevel(i int) ITopLevelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelContext)
}

type TopLevelDefineContext struct {
	TopLevelContext
}

func NewTopLevelDefineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelDefineContext {
	var p = new(TopLevelDefineContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefineContext) Define() antlr.TerminalNode {
	return s.GetToken(CParserDefine, 0)
}

type TopLevelPragmaContext struct {
	TopLevelContext
}

func NewTopLevelPragmaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelPragmaContext {
	var p = new(TopLevelPragmaContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelPragmaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelPragmaContext) Pragma() antlr.TerminalNode {
	return s.GetToken(CParserPragma, 0)
}

type TopLevelFunctionContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelFunctionContext {
	var p = new(TopLevelFunctionContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelFunctionContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelFunctionContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelFunctionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TopLevelFunctionContext) LParen() antlr.TerminalNode {
	return s.GetToken(CParserLParen, 0)
}

func (s *TopLevelFunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *TopLevelFunctionContext) RParen() antlr.TerminalNode {
	return s.GetToken(CParserRParen, 0)
}

func (s *TopLevelFunctionContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelFunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelFunctionContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

type TopLevelTypeDefFuncContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelTypeDefFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelTypeDefFuncContext {
	var p = new(TopLevelTypeDefFuncContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelTypeDefFuncContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelTypeDefFuncContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelTypeDefFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelTypeDefFuncContext) Typedef() antlr.TerminalNode {
	return s.GetToken(CParserTypedef, 0)
}

func (s *TopLevelTypeDefFuncContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TopLevelTypeDefFuncContext) AllLParen() []antlr.TerminalNode {
	return s.GetTokens(CParserLParen)
}

func (s *TopLevelTypeDefFuncContext) LParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserLParen, i)
}

func (s *TopLevelTypeDefFuncContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *TopLevelTypeDefFuncContext) AllRParen() []antlr.TerminalNode {
	return s.GetTokens(CParserRParen)
}

func (s *TopLevelTypeDefFuncContext) RParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserRParen, i)
}

func (s *TopLevelTypeDefFuncContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *TopLevelTypeDefFuncContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelTypeDefFuncContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelTypeDefFuncContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

type TopLevelEnumContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelEnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelEnumContext {
	var p = new(TopLevelEnumContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelEnumContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelEnumContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelEnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelEnumContext) Enum() antlr.TerminalNode {
	return s.GetToken(CParserEnum, 0)
}

func (s *TopLevelEnumContext) LBrace() antlr.TerminalNode {
	return s.GetToken(CParserLBrace, 0)
}

func (s *TopLevelEnumContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TopLevelEnumContext) RBrace() antlr.TerminalNode {
	return s.GetToken(CParserRBrace, 0)
}

func (s *TopLevelEnumContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelEnumContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *TopLevelEnumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type TopLevelStructTypeDefContext struct {
	TopLevelContext
	structname antlr.Token
	itemname   antlr.Token
}

func NewTopLevelStructTypeDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelStructTypeDefContext {
	var p = new(TopLevelStructTypeDefContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelStructTypeDefContext) GetStructname() antlr.Token { return s.structname }

func (s *TopLevelStructTypeDefContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelStructTypeDefContext) SetStructname(v antlr.Token) { s.structname = v }

func (s *TopLevelStructTypeDefContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelStructTypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelStructTypeDefContext) Typedef() antlr.TerminalNode {
	return s.GetToken(CParserTypedef, 0)
}

func (s *TopLevelStructTypeDefContext) Struct() antlr.TerminalNode {
	return s.GetToken(CParserStruct, 0)
}

func (s *TopLevelStructTypeDefContext) LBrace() antlr.TerminalNode {
	return s.GetToken(CParserLBrace, 0)
}

func (s *TopLevelStructTypeDefContext) RBrace() antlr.TerminalNode {
	return s.GetToken(CParserRBrace, 0)
}

func (s *TopLevelStructTypeDefContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelStructTypeDefContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CParserIdentifier)
}

func (s *TopLevelStructTypeDefContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, i)
}

func (s *TopLevelStructTypeDefContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

func (s *TopLevelStructTypeDefContext) AllStructItem() []IStructItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructItemContext); ok {
			len++
		}
	}

	tst := make([]IStructItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructItemContext); ok {
			tst[i] = t.(IStructItemContext)
			i++
		}
	}

	return tst
}

func (s *TopLevelStructTypeDefContext) StructItem(i int) IStructItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructItemContext)
}

type TopLevelTypeDefSimpleContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelTypeDefSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelTypeDefSimpleContext {
	var p = new(TopLevelTypeDefSimpleContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelTypeDefSimpleContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelTypeDefSimpleContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelTypeDefSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelTypeDefSimpleContext) Typedef() antlr.TerminalNode {
	return s.GetToken(CParserTypedef, 0)
}

func (s *TopLevelTypeDefSimpleContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TopLevelTypeDefSimpleContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelTypeDefSimpleContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelTypeDefSimpleContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

func (p *CParser) TopLevel() (localctx ITopLevelContext) {
	localctx = NewTopLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CParserRULE_topLevel)
	var _la int

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTopLevelIfndefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Match(CParserIfndef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132116480) != 0 {
			{
				p.SetState(33)
				p.TopLevel()
			}

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(39)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTopLevelIfdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132116480) != 0 {
			{
				p.SetState(41)
				p.TopLevel()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(47)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTopLevelDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(CParserDefine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTopLevelUndefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(CParserUndef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewTopLevelIncludeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(CParserInclude)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewTopLevelPragmaContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.Match(CParserPragma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewTopLevelEnumContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(52)
			p.Match(CParserEnum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserIdentifier {
			{
				p.SetState(53)

				var _m = p.Match(CParserIdentifier)

				localctx.(*TopLevelEnumContext).itemname = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(56)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.EnumBody()
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(58)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(61)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewTopLevelStructTypeDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(64)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(67)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructTypeDefContext).structname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72740864) != 0 {
			{
				p.SetState(71)
				p.StructItem()
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructTypeDefContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewTopLevelStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(80)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(83)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(84)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(87)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72740864) != 0 {
			{
				p.SetState(89)
				p.StructItem()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewTopLevelTypeDefSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(97)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(100)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.type_(0)
		}
		{
			p.SetState(102)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelTypeDefSimpleContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewTopLevelTypeDefFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(105)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(108)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.type_(0)
		}
		{
			p.SetState(110)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelTypeDefFuncContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Params()
		}
		{
			p.SetState(116)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewTopLevelFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(119)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(122)
			p.type_(0)
		}
		{
			p.SetState(123)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelFunctionContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Params()
		}
		{
			p.SetState(126)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) CopyAll(ctx *EnumBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumBodyItemContext struct {
	EnumBodyContext
}

func NewEnumBodyItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumBodyItemContext {
	var p = new(EnumBodyItemContext)

	InitEmptyEnumBodyContext(&p.EnumBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumBodyItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyItemContext) EnumItem() IEnumItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumItemContext)
}

func (s *EnumBodyItemContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *EnumBodyItemContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

type EnumBodyEndContext struct {
	EnumBodyContext
}

func NewEnumBodyEndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumBodyEndContext {
	var p = new(EnumBodyEndContext)

	InitEmptyEnumBodyContext(&p.EnumBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumBodyEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyEndContext) EnumItem() IEnumItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumItemContext)
}

type EnumBodyDefineContext struct {
	EnumBodyContext
}

func NewEnumBodyDefineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumBodyDefineContext {
	var p = new(EnumBodyDefineContext)

	InitEmptyEnumBodyContext(&p.EnumBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumBodyDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyDefineContext) Define() antlr.TerminalNode {
	return s.GetToken(CParserDefine, 0)
}

func (s *EnumBodyDefineContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

type EnumBodyIfdirContext struct {
	EnumBodyContext
}

func NewEnumBodyIfdirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumBodyIfdirContext {
	var p = new(EnumBodyIfdirContext)

	InitEmptyEnumBodyContext(&p.EnumBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumBodyContext))

	return p
}

func (s *EnumBodyIfdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyIfdirContext) Ifdir() antlr.TerminalNode {
	return s.GetToken(CParserIfdir, 0)
}

func (s *EnumBodyIfdirContext) Endif() antlr.TerminalNode {
	return s.GetToken(CParserEndif, 0)
}

func (s *EnumBodyIfdirContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumBodyIfdirContext) AllEnumItem() []IEnumItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumItemContext); ok {
			len++
		}
	}

	tst := make([]IEnumItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumItemContext); ok {
			tst[i] = t.(IEnumItemContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyIfdirContext) EnumItem(i int) IEnumItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumItemContext)
}

func (s *EnumBodyIfdirContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(CParserComma)
}

func (s *EnumBodyIfdirContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(CParserComma, i)
}

func (p *CParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_enumBody)
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEnumBodyDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(CParserDefine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.EnumBody()
		}

	case 2:
		localctx = NewEnumBodyIfdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == CParserIdentifier {
			{
				p.SetState(134)
				p.EnumItem()
			}
			{
				p.SetState(135)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(141)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.EnumBody()
		}

	case 3:
		localctx = NewEnumBodyItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.EnumItem()
		}
		{
			p.SetState(145)
			p.Match(CParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.EnumBody()
		}

	case 4:
		localctx = NewEnumBodyEndContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.EnumItem()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumItemContext is an interface to support dynamic dispatch.
type IEnumItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEnumItemContext differentiates from other interfaces.
	IsEnumItemContext()
}

type EnumItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumItemContext() *EnumItemContext {
	var p = new(EnumItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_enumItem
	return p
}

func InitEmptyEnumItemContext(p *EnumItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_enumItem
}

func (*EnumItemContext) IsEnumItemContext() {}

func NewEnumItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumItemContext {
	var p = new(EnumItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_enumItem

	return p
}

func (s *EnumItemContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumItemContext) CopyAll(ctx *EnumItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EnumItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumItemNoValueContext struct {
	EnumItemContext
}

func NewEnumItemNoValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumItemNoValueContext {
	var p = new(EnumItemNoValueContext)

	InitEmptyEnumItemContext(&p.EnumItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumItemContext))

	return p
}

func (s *EnumItemNoValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemNoValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type EnumItemConstantContext struct {
	EnumItemContext
}

func NewEnumItemConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumItemConstantContext {
	var p = new(EnumItemConstantContext)

	InitEmptyEnumItemContext(&p.EnumItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumItemContext))

	return p
}

func (s *EnumItemConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemConstantContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *EnumItemConstantContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *EnumItemConstantContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (p *CParser) EnumItem() (localctx IEnumItemContext) {
	localctx = NewEnumItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_enumItem)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEnumItemNoValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewEnumItemConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(CParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Constant()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructItemContext is an interface to support dynamic dispatch.
type IStructItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructItemContext differentiates from other interfaces.
	IsStructItemContext()
}

type StructItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructItemContext() *StructItemContext {
	var p = new(StructItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_structItem
	return p
}

func InitEmptyStructItemContext(p *StructItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_structItem
}

func (*StructItemContext) IsStructItemContext() {}

func NewStructItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructItemContext {
	var p = new(StructItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_structItem

	return p
}

func (s *StructItemContext) GetParser() antlr.Parser { return s.parser }

func (s *StructItemContext) CopyAll(ctx *StructItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructItemIfdirContext struct {
	StructItemContext
}

func NewStructItemIfdirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemIfdirContext {
	var p = new(StructItemIfdirContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemIfdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemIfdirContext) Ifdir() antlr.TerminalNode {
	return s.GetToken(CParserIfdir, 0)
}

func (s *StructItemIfdirContext) Endif() antlr.TerminalNode {
	return s.GetToken(CParserEndif, 0)
}

func (s *StructItemIfdirContext) AllStructItem() []IStructItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructItemContext); ok {
			len++
		}
	}

	tst := make([]IStructItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructItemContext); ok {
			tst[i] = t.(IStructItemContext)
			i++
		}
	}

	return tst
}

func (s *StructItemIfdirContext) StructItem(i int) IStructItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructItemContext)
}

type StructItemFunctionContext struct {
	StructItemContext
}

func NewStructItemFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemFunctionContext {
	var p = new(StructItemFunctionContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemFunctionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructItemFunctionContext) AllLParen() []antlr.TerminalNode {
	return s.GetTokens(CParserLParen)
}

func (s *StructItemFunctionContext) LParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserLParen, i)
}

func (s *StructItemFunctionContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *StructItemFunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructItemFunctionContext) AllRParen() []antlr.TerminalNode {
	return s.GetTokens(CParserRParen)
}

func (s *StructItemFunctionContext) RParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserRParen, i)
}

func (s *StructItemFunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *StructItemFunctionContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

type StructItemDefineContext struct {
	StructItemContext
}

func NewStructItemDefineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemDefineContext {
	var p = new(StructItemDefineContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemDefineContext) Define() antlr.TerminalNode {
	return s.GetToken(CParserDefine, 0)
}

type StructItemMemberContext struct {
	StructItemContext
}

func NewStructItemMemberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemMemberContext {
	var p = new(StructItemMemberContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemMemberContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructItemMemberContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CParserIdentifier)
}

func (s *StructItemMemberContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, i)
}

func (s *StructItemMemberContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StructItemMemberContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

func (s *StructItemMemberContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(CParserComma)
}

func (s *StructItemMemberContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(CParserComma, i)
}

type StructItemArrayContext struct {
	StructItemContext
}

func NewStructItemArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemArrayContext {
	var p = new(StructItemArrayContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemArrayContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructItemArrayContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructItemArrayContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *StructItemArrayContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StructItemArrayContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

type StructItemArrayPtrContext struct {
	StructItemContext
}

func NewStructItemArrayPtrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemArrayPtrContext {
	var p = new(StructItemArrayPtrContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemArrayPtrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemArrayPtrContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructItemArrayPtrContext) LParen() antlr.TerminalNode {
	return s.GetToken(CParserLParen, 0)
}

func (s *StructItemArrayPtrContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *StructItemArrayPtrContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructItemArrayPtrContext) RParen() antlr.TerminalNode {
	return s.GetToken(CParserRParen, 0)
}

func (s *StructItemArrayPtrContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *StructItemArrayPtrContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StructItemArrayPtrContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
}

func (p *CParser) StructItem() (localctx IStructItemContext) {
	localctx = NewStructItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_structItem)
	var _la int

	var _alt int

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructItemMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(157)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(160)
			p.type_(0)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(161)
					p.Match(CParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(162)
					p.Match(CParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStructItemArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(171)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(174)
			p.type_(0)
		}
		{
			p.SetState(175)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.ArraySize()
		}
		{
			p.SetState(177)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStructItemArrayPtrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(179)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(182)
			p.type_(0)
		}
		{
			p.SetState(183)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.ArraySize()
		}
		{
			p.SetState(188)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStructItemIfdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(190)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72740864) != 0 {
			{
				p.SetState(191)
				p.StructItem()
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStructItemDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(198)
			p.Match(CParserDefine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStructItemFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(199)
			p.type_(0)
		}
		{
			p.SetState(200)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Params()
		}
		{
			p.SetState(206)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMin returns the min token.
	GetMin() antlr.Token

	// SetMin sets the min token.
	SetMin(antlr.Token)

	// Getter signatures
	Constant() antlr.TerminalNode
	Minus() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	min    antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetMin() antlr.Token { return s.min }

func (s *ConstantContext) SetMin(v antlr.Token) { s.min = v }

func (s *ConstantContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *ConstantContext) Minus() antlr.TerminalNode {
	return s.GetToken(CParserMinus, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *CParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CParserMinus {
		{
			p.SetState(211)

			var _m = p.Match(CParserMinus)

			localctx.(*ConstantContext).min = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(214)
		p.Match(CParserConstant)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) CopyAll(ctx *ParamsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamsValuesContext struct {
	ParamsContext
}

func NewParamsValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamsValuesContext {
	var p = new(ParamsValuesContext)

	InitEmptyParamsContext(&p.ParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamsContext))

	return p
}

func (s *ParamsValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsValuesContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsValuesContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsValuesContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(CParserComma)
}

func (s *ParamsValuesContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(CParserComma, i)
}

type ParamsVoidContext struct {
	ParamsContext
}

func NewParamsVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamsVoidContext {
	var p = new(ParamsVoidContext)

	InitEmptyParamsContext(&p.ParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamsContext))

	return p
}

func (s *ParamsVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsVoidContext) Void() antlr.TerminalNode {
	return s.GetToken(CParserVoid, 0)
}

func (p *CParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_params)
	var _alt int

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamsVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Match(CParserVoid)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewParamsValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(217)
					p.Param()
				}
				{
					p.SetState(218)
					p.Match(CParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Param()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) CopyAll(ctx *ParamContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamArrayContext struct {
	ParamContext
}

func NewParamArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamArrayContext {
	var p = new(ParamArrayContext)

	InitEmptyParamContext(&p.ParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamContext))

	return p
}

func (s *ParamArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamArrayContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamArrayContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *ParamArrayContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

type ParamAnonContext struct {
	ParamContext
}

func NewParamAnonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamAnonContext {
	var p = new(ParamAnonContext)

	InitEmptyParamContext(&p.ParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamContext))

	return p
}

func (s *ParamAnonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamAnonContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

type ParamSimpleContext struct {
	ParamContext
}

func NewParamSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamSimpleContext {
	var p = new(ParamSimpleContext)

	InitEmptyParamContext(&p.ParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamContext))

	return p
}

func (s *ParamSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamSimpleContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamSimpleContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type ParamFunctionContext struct {
	ParamContext
}

func NewParamFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamFunctionContext {
	var p = new(ParamFunctionContext)

	InitEmptyParamContext(&p.ParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamContext))

	return p
}

func (s *ParamFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamFunctionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamFunctionContext) AllLParen() []antlr.TerminalNode {
	return s.GetTokens(CParserLParen)
}

func (s *ParamFunctionContext) LParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserLParen, i)
}

func (s *ParamFunctionContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *ParamFunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *ParamFunctionContext) AllRParen() []antlr.TerminalNode {
	return s.GetTokens(CParserRParen)
}

func (s *ParamFunctionContext) RParen(i int) antlr.TerminalNode {
	return s.GetToken(CParserRParen, i)
}

func (s *ParamFunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (p *CParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_param)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.type_(0)
		}
		{
			p.SetState(229)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewParamArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.type_(0)
		}
		{
			p.SetState(232)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.ArraySize()
		}

	case 3:
		localctx = NewParamAnonContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.type_(0)
		}

	case 4:
		localctx = NewParamFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(236)
			p.type_(0)
		}
		{
			p.SetState(237)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Params()
		}
		{
			p.SetState(243)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArraySizeContext is an interface to support dynamic dispatch.
type IArraySizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArraySizeDim() []IArraySizeDimContext
	ArraySizeDim(i int) IArraySizeDimContext

	// IsArraySizeContext differentiates from other interfaces.
	IsArraySizeContext()
}

type ArraySizeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySizeContext() *ArraySizeContext {
	var p = new(ArraySizeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_arraySize
	return p
}

func InitEmptyArraySizeContext(p *ArraySizeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_arraySize
}

func (*ArraySizeContext) IsArraySizeContext() {}

func NewArraySizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySizeContext {
	var p = new(ArraySizeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_arraySize

	return p
}

func (s *ArraySizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySizeContext) AllArraySizeDim() []IArraySizeDimContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArraySizeDimContext); ok {
			len++
		}
	}

	tst := make([]IArraySizeDimContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArraySizeDimContext); ok {
			tst[i] = t.(IArraySizeDimContext)
			i++
		}
	}

	return tst
}

func (s *ArraySizeContext) ArraySizeDim(i int) IArraySizeDimContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeDimContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeDimContext)
}

func (s *ArraySizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *CParser) ArraySize() (localctx IArraySizeContext) {
	localctx = NewArraySizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CParserRULE_arraySize)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CParserLBracket {
		{
			p.SetState(247)
			p.ArraySizeDim()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArraySizeDimContext is an interface to support dynamic dispatch.
type IArraySizeDimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArraySizeDimContext differentiates from other interfaces.
	IsArraySizeDimContext()
}

type ArraySizeDimContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySizeDimContext() *ArraySizeDimContext {
	var p = new(ArraySizeDimContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_arraySizeDim
	return p
}

func InitEmptyArraySizeDimContext(p *ArraySizeDimContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_arraySizeDim
}

func (*ArraySizeDimContext) IsArraySizeDimContext() {}

func NewArraySizeDimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySizeDimContext {
	var p = new(ArraySizeDimContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_arraySizeDim

	return p
}

func (s *ArraySizeDimContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySizeDimContext) CopyAll(ctx *ArraySizeDimContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArraySizeDimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeDimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArraySizeDimConstContext struct {
	ArraySizeDimContext
}

func NewArraySizeDimConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArraySizeDimConstContext {
	var p = new(ArraySizeDimConstContext)

	InitEmptyArraySizeDimContext(&p.ArraySizeDimContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArraySizeDimContext))

	return p
}

func (s *ArraySizeDimConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeDimConstContext) LBracket() antlr.TerminalNode {
	return s.GetToken(CParserLBracket, 0)
}

func (s *ArraySizeDimConstContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *ArraySizeDimConstContext) RBracket() antlr.TerminalNode {
	return s.GetToken(CParserRBracket, 0)
}

type ArraySizeDimIdenContext struct {
	ArraySizeDimContext
}

func NewArraySizeDimIdenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArraySizeDimIdenContext {
	var p = new(ArraySizeDimIdenContext)

	InitEmptyArraySizeDimContext(&p.ArraySizeDimContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArraySizeDimContext))

	return p
}

func (s *ArraySizeDimIdenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeDimIdenContext) LBracket() antlr.TerminalNode {
	return s.GetToken(CParserLBracket, 0)
}

func (s *ArraySizeDimIdenContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *ArraySizeDimIdenContext) RBracket() antlr.TerminalNode {
	return s.GetToken(CParserRBracket, 0)
}

func (p *CParser) ArraySizeDim() (localctx IArraySizeDimContext) {
	localctx = NewArraySizeDimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CParserRULE_arraySizeDim)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArraySizeDimIdenContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(CParserLBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(CParserRBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewArraySizeDimConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Match(CParserLBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(CParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Match(CParserRBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypePtrContext struct {
	TypeContext
}

func NewTypePtrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypePtrContext {
	var p = new(TypePtrContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *TypePtrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypePtrContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypePtrContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

type TypePassContext struct {
	TypeContext
}

func NewTypePassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypePassContext {
	var p = new(TypePassContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *TypePassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypePassContext) TypeInner() ITypeInnerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeInnerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeInnerContext)
}

type TypePtrConstContext struct {
	TypeContext
}

func NewTypePtrConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypePtrConstContext {
	var p = new(TypePtrConstContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *TypePtrConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypePtrConstContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypePtrConstContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *TypePtrConstContext) Const() antlr.TerminalNode {
	return s.GetToken(CParserConst, 0)
}

func (p *CParser) Type_() (localctx ITypeContext) {
	return p.type_(0)
}

func (p *CParser) type_(_p int) (localctx ITypeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, CParserRULE_type, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewTypePassContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(261)
		p.TypeInner()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewTypePtrContext(p, NewTypeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_type)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(264)
					p.Match(CParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewTypePtrConstContext(p, NewTypeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_type)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(266)
					p.Match(CParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(267)
					p.Match(CParserConst)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeInnerContext is an interface to support dynamic dispatch.
type ITypeInnerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeInnerContext differentiates from other interfaces.
	IsTypeInnerContext()
}

type TypeInnerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeInnerContext() *TypeInnerContext {
	var p = new(TypeInnerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_typeInner
	return p
}

func InitEmptyTypeInnerContext(p *TypeInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_typeInner
}

func (*TypeInnerContext) IsTypeInnerContext() {}

func NewTypeInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeInnerContext {
	var p = new(TypeInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typeInner

	return p
}

func (s *TypeInnerContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeInnerContext) CopyAll(ctx *TypeInnerContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeInnerIdentContext struct {
	TypeInnerContext
}

func NewTypeInnerIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerIdentContext {
	var p = new(TypeInnerIdentContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerIdentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type TypeInnerVoidContext struct {
	TypeInnerContext
}

func NewTypeInnerVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerVoidContext {
	var p = new(TypeInnerVoidContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerVoidContext) Void() antlr.TerminalNode {
	return s.GetToken(CParserVoid, 0)
}

type TypeInnerUnsignedContext struct {
	TypeInnerContext
}

func NewTypeInnerUnsignedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerUnsignedContext {
	var p = new(TypeInnerUnsignedContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerUnsignedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerUnsignedContext) Unsigned() antlr.TerminalNode {
	return s.GetToken(CParserUnsigned, 0)
}

func (s *TypeInnerUnsignedContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type TypeInnerConstContext struct {
	TypeInnerContext
}

func NewTypeInnerConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerConstContext {
	var p = new(TypeInnerConstContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerConstContext) Const() antlr.TerminalNode {
	return s.GetToken(CParserConst, 0)
}

func (s *TypeInnerConstContext) TypeInner() ITypeInnerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeInnerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeInnerContext)
}

type TypeInnerEnumContext struct {
	TypeInnerContext
}

func NewTypeInnerEnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerEnumContext {
	var p = new(TypeInnerEnumContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerEnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerEnumContext) Enum() antlr.TerminalNode {
	return s.GetToken(CParserEnum, 0)
}

func (s *TypeInnerEnumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

type TypeInnerStructContext struct {
	TypeInnerContext
}

func NewTypeInnerStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeInnerStructContext {
	var p = new(TypeInnerStructContext)

	InitEmptyTypeInnerContext(&p.TypeInnerContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeInnerContext))

	return p
}

func (s *TypeInnerStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInnerStructContext) Struct() antlr.TerminalNode {
	return s.GetToken(CParserStruct, 0)
}

func (s *TypeInnerStructContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (p *CParser) TypeInner() (localctx ITypeInnerContext) {
	localctx = NewTypeInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_typeInner)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		localctx = NewTypeInnerIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CParserEnum:
		localctx = NewTypeInnerEnumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(CParserEnum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CParserStruct:
		localctx = NewTypeInnerStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CParserUnsigned:
		localctx = NewTypeInnerUnsignedContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(278)
			p.Match(CParserUnsigned)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(279)
				p.Match(CParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case CParserVoid:
		localctx = NewTypeInnerVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.Match(CParserVoid)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CParserConst:
		localctx = NewTypeInnerConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(283)
			p.Match(CParserConst)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.TypeInner()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *TypeContext = nil
		if localctx != nil {
			t = localctx.(*TypeContext)
		}
		return p.Type__Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CParser) Type__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
