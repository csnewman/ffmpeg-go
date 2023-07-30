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
		"", "'...'", "'='", "'*'", "'-'", "','", "", "':'", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "'const'", "'unsigned'", "'void'", "'enum'", "'struct'",
		"'static'", "'inline'", "'typedef'", "'attribute_deprecated'",
	}
	staticData.SymbolicNames = []string{
		"", "", "Assign", "Star", "Minus", "Comma", "Semi", "Colon", "LParen",
		"RParen", "LBracket", "RBracket", "LBrace", "RBrace", "Const", "Unsigned",
		"Void", "Enum", "Struct", "Static", "Inline", "Typedef", "AttributeDeprecated",
		"Ifndef", "Ifdir", "Endif", "ElseDir", "Define", "Undef", "Pragma",
		"Include", "Identifier", "Whitespace", "Newline", "Constant", "DigitSequence",
		"StringLiteral", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"unit", "topLevel", "elseCond", "enumBody", "enumItem", "structItem",
		"constant", "params", "param", "arraySize", "arraySizeDim", "type",
		"typeInner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 349, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9, 1, 1,
		1, 3, 1, 43, 8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 48, 8, 1, 10, 1, 12, 1, 51,
		9, 1, 1, 1, 3, 1, 54, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 63, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 68, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 85, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 92, 8, 1, 10, 1, 12, 1,
		95, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 101, 8, 1, 1, 1, 1, 1, 3, 1, 105,
		8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 110, 8, 1, 10, 1, 12, 1, 113, 9, 1, 1, 1,
		3, 1, 116, 8, 1, 1, 1, 1, 1, 3, 1, 120, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 128, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 142, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 153, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 161, 8, 1, 1, 2, 1, 2, 5, 2, 165, 8, 2, 10, 2, 12, 2, 168,
		9, 2, 1, 3, 1, 3, 3, 3, 172, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 178, 8,
		3, 11, 3, 12, 3, 179, 1, 3, 1, 3, 3, 3, 184, 8, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 189, 8, 3, 3, 3, 191, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 200, 8, 4, 1, 5, 3, 5, 203, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 208, 8,
		5, 10, 5, 12, 5, 211, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 217, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 225, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 234, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 246, 8, 5, 10, 5, 12, 5, 249, 9, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 254, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 266, 8, 5, 1, 6, 3, 6, 269, 8, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 277, 8, 7, 10, 7, 12, 7, 280, 9, 7, 1, 7,
		1, 7, 3, 7, 284, 8, 7, 3, 7, 286, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 305, 8, 8, 1, 9, 4, 9, 308, 8, 9, 11, 9, 12, 9, 309, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 320, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 330, 8, 11, 10, 11, 12,
		11, 333, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12,
		342, 8, 12, 1, 12, 1, 12, 1, 12, 3, 12, 347, 8, 12, 1, 12, 0, 1, 22, 13,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 0, 406, 0, 29, 1, 0,
		0, 0, 2, 160, 1, 0, 0, 0, 4, 162, 1, 0, 0, 0, 6, 190, 1, 0, 0, 0, 8, 199,
		1, 0, 0, 0, 10, 265, 1, 0, 0, 0, 12, 268, 1, 0, 0, 0, 14, 285, 1, 0, 0,
		0, 16, 304, 1, 0, 0, 0, 18, 307, 1, 0, 0, 0, 20, 319, 1, 0, 0, 0, 22, 321,
		1, 0, 0, 0, 24, 346, 1, 0, 0, 0, 26, 28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0,
		28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1,
		0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34,
		38, 5, 23, 0, 0, 35, 37, 3, 2, 1, 0, 36, 35, 1, 0, 0, 0, 37, 40, 1, 0,
		0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 41, 43, 3, 4, 2, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 161, 5, 25, 0, 0, 45, 49, 5, 24, 0, 0, 46, 48,
		3, 2, 1, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 54, 3,
		4, 2, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55,
		161, 5, 25, 0, 0, 56, 161, 5, 27, 0, 0, 57, 161, 5, 28, 0, 0, 58, 161,
		5, 30, 0, 0, 59, 161, 5, 29, 0, 0, 60, 62, 5, 17, 0, 0, 61, 63, 5, 31,
		0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65,
		5, 12, 0, 0, 65, 67, 3, 6, 3, 0, 66, 68, 5, 5, 0, 0, 67, 66, 1, 0, 0, 0,
		67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5, 13, 0, 0, 70, 71, 5,
		6, 0, 0, 71, 161, 1, 0, 0, 0, 72, 73, 5, 21, 0, 0, 73, 74, 5, 17, 0, 0,
		74, 75, 5, 12, 0, 0, 75, 77, 3, 6, 3, 0, 76, 78, 5, 5, 0, 0, 77, 76, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 13, 0, 0, 80,
		81, 5, 31, 0, 0, 81, 82, 5, 6, 0, 0, 82, 161, 1, 0, 0, 0, 83, 85, 5, 22,
		0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87,
		5, 21, 0, 0, 87, 88, 5, 18, 0, 0, 88, 89, 5, 31, 0, 0, 89, 93, 5, 12, 0,
		0, 90, 92, 3, 10, 5, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0,
		96, 97, 5, 13, 0, 0, 97, 98, 5, 31, 0, 0, 98, 161, 5, 6, 0, 0, 99, 101,
		5, 22, 0, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0,
		0, 0, 102, 104, 5, 18, 0, 0, 103, 105, 5, 22, 0, 0, 104, 103, 1, 0, 0,
		0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 115, 5, 31, 0, 0, 107,
		111, 5, 12, 0, 0, 108, 110, 3, 10, 5, 0, 109, 108, 1, 0, 0, 0, 110, 113,
		1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0,
		0, 0, 113, 111, 1, 0, 0, 0, 114, 116, 5, 13, 0, 0, 115, 107, 1, 0, 0, 0,
		115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 161, 5, 6, 0, 0, 118,
		120, 5, 22, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 122, 5, 21, 0, 0, 122, 123, 3, 22, 11, 0, 123, 124, 5,
		31, 0, 0, 124, 125, 5, 6, 0, 0, 125, 161, 1, 0, 0, 0, 126, 128, 5, 22,
		0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 130, 5, 21, 0, 0, 130, 131, 3, 22, 11, 0, 131, 132, 5, 8, 0, 0, 132,
		133, 5, 3, 0, 0, 133, 134, 5, 31, 0, 0, 134, 135, 5, 9, 0, 0, 135, 136,
		5, 8, 0, 0, 136, 137, 3, 14, 7, 0, 137, 138, 5, 9, 0, 0, 138, 139, 5, 6,
		0, 0, 139, 161, 1, 0, 0, 0, 140, 142, 5, 22, 0, 0, 141, 140, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 3, 22, 11, 0, 144,
		145, 5, 31, 0, 0, 145, 146, 5, 8, 0, 0, 146, 147, 3, 14, 7, 0, 147, 148,
		5, 9, 0, 0, 148, 149, 5, 6, 0, 0, 149, 161, 1, 0, 0, 0, 150, 152, 5, 19,
		0, 0, 151, 153, 5, 20, 0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 154, 155, 3, 22, 11, 0, 155, 156, 5, 31, 0, 0, 156,
		157, 5, 8, 0, 0, 157, 158, 3, 14, 7, 0, 158, 159, 5, 9, 0, 0, 159, 161,
		1, 0, 0, 0, 160, 34, 1, 0, 0, 0, 160, 45, 1, 0, 0, 0, 160, 56, 1, 0, 0,
		0, 160, 57, 1, 0, 0, 0, 160, 58, 1, 0, 0, 0, 160, 59, 1, 0, 0, 0, 160,
		60, 1, 0, 0, 0, 160, 72, 1, 0, 0, 0, 160, 84, 1, 0, 0, 0, 160, 100, 1,
		0, 0, 0, 160, 119, 1, 0, 0, 0, 160, 127, 1, 0, 0, 0, 160, 141, 1, 0, 0,
		0, 160, 150, 1, 0, 0, 0, 161, 3, 1, 0, 0, 0, 162, 166, 5, 26, 0, 0, 163,
		165, 3, 2, 1, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 5, 1, 0, 0, 0, 168, 166, 1, 0, 0,
		0, 169, 171, 5, 27, 0, 0, 170, 172, 3, 6, 3, 0, 171, 170, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 191, 1, 0, 0, 0, 173, 177, 5, 24, 0, 0, 174, 175,
		3, 8, 4, 0, 175, 176, 5, 5, 0, 0, 176, 178, 1, 0, 0, 0, 177, 174, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0,
		180, 181, 1, 0, 0, 0, 181, 183, 5, 25, 0, 0, 182, 184, 3, 6, 3, 0, 183,
		182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 191, 1, 0, 0, 0, 185, 188,
		3, 8, 4, 0, 186, 187, 5, 5, 0, 0, 187, 189, 3, 6, 3, 0, 188, 186, 1, 0,
		0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190, 169, 1, 0, 0, 0,
		190, 173, 1, 0, 0, 0, 190, 185, 1, 0, 0, 0, 191, 7, 1, 0, 0, 0, 192, 200,
		5, 31, 0, 0, 193, 194, 5, 31, 0, 0, 194, 195, 5, 2, 0, 0, 195, 200, 3,
		12, 6, 0, 196, 197, 5, 31, 0, 0, 197, 198, 5, 2, 0, 0, 198, 200, 5, 31,
		0, 0, 199, 192, 1, 0, 0, 0, 199, 193, 1, 0, 0, 0, 199, 196, 1, 0, 0, 0,
		200, 9, 1, 0, 0, 0, 201, 203, 5, 22, 0, 0, 202, 201, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 209, 3, 22, 11, 0, 205, 206, 5,
		31, 0, 0, 206, 208, 5, 5, 0, 0, 207, 205, 1, 0, 0, 0, 208, 211, 1, 0, 0,
		0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211,
		209, 1, 0, 0, 0, 212, 213, 5, 31, 0, 0, 213, 214, 5, 6, 0, 0, 214, 266,
		1, 0, 0, 0, 215, 217, 5, 22, 0, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1, 0,
		0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 3, 22, 11, 0, 219, 220, 5, 31, 0,
		0, 220, 221, 3, 18, 9, 0, 221, 222, 5, 6, 0, 0, 222, 266, 1, 0, 0, 0, 223,
		225, 5, 22, 0, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226,
		1, 0, 0, 0, 226, 227, 3, 22, 11, 0, 227, 228, 5, 31, 0, 0, 228, 229, 5,
		7, 0, 0, 229, 230, 5, 34, 0, 0, 230, 231, 5, 6, 0, 0, 231, 266, 1, 0, 0,
		0, 232, 234, 5, 22, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 236, 3, 22, 11, 0, 236, 237, 5, 8, 0, 0, 237, 238,
		5, 3, 0, 0, 238, 239, 5, 31, 0, 0, 239, 240, 5, 9, 0, 0, 240, 241, 3, 18,
		9, 0, 241, 242, 5, 6, 0, 0, 242, 266, 1, 0, 0, 0, 243, 247, 5, 24, 0, 0,
		244, 246, 3, 10, 5, 0, 245, 244, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 250, 266, 5, 25, 0, 0, 251, 266, 5, 27, 0, 0, 252, 254, 5,
		22, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0,
		0, 255, 256, 3, 22, 11, 0, 256, 257, 5, 8, 0, 0, 257, 258, 5, 3, 0, 0,
		258, 259, 5, 31, 0, 0, 259, 260, 5, 9, 0, 0, 260, 261, 5, 8, 0, 0, 261,
		262, 3, 14, 7, 0, 262, 263, 5, 9, 0, 0, 263, 264, 5, 6, 0, 0, 264, 266,
		1, 0, 0, 0, 265, 202, 1, 0, 0, 0, 265, 216, 1, 0, 0, 0, 265, 224, 1, 0,
		0, 0, 265, 233, 1, 0, 0, 0, 265, 243, 1, 0, 0, 0, 265, 251, 1, 0, 0, 0,
		265, 253, 1, 0, 0, 0, 266, 11, 1, 0, 0, 0, 267, 269, 5, 4, 0, 0, 268, 267,
		1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 5, 34,
		0, 0, 271, 13, 1, 0, 0, 0, 272, 286, 5, 16, 0, 0, 273, 274, 3, 16, 8, 0,
		274, 275, 5, 5, 0, 0, 275, 277, 1, 0, 0, 0, 276, 273, 1, 0, 0, 0, 277,
		280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 283,
		1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 284, 3, 16, 8, 0, 282, 284, 5, 1,
		0, 0, 283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0, 284, 286, 1, 0, 0, 0,
		285, 272, 1, 0, 0, 0, 285, 278, 1, 0, 0, 0, 286, 15, 1, 0, 0, 0, 287, 288,
		3, 22, 11, 0, 288, 289, 5, 31, 0, 0, 289, 305, 1, 0, 0, 0, 290, 291, 3,
		22, 11, 0, 291, 292, 5, 31, 0, 0, 292, 293, 3, 18, 9, 0, 293, 305, 1, 0,
		0, 0, 294, 305, 3, 22, 11, 0, 295, 296, 3, 22, 11, 0, 296, 297, 5, 8, 0,
		0, 297, 298, 5, 3, 0, 0, 298, 299, 5, 31, 0, 0, 299, 300, 5, 9, 0, 0, 300,
		301, 5, 8, 0, 0, 301, 302, 3, 14, 7, 0, 302, 303, 5, 9, 0, 0, 303, 305,
		1, 0, 0, 0, 304, 287, 1, 0, 0, 0, 304, 290, 1, 0, 0, 0, 304, 294, 1, 0,
		0, 0, 304, 295, 1, 0, 0, 0, 305, 17, 1, 0, 0, 0, 306, 308, 3, 20, 10, 0,
		307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309,
		310, 1, 0, 0, 0, 310, 19, 1, 0, 0, 0, 311, 312, 5, 10, 0, 0, 312, 320,
		5, 11, 0, 0, 313, 314, 5, 10, 0, 0, 314, 315, 5, 31, 0, 0, 315, 320, 5,
		11, 0, 0, 316, 317, 5, 10, 0, 0, 317, 318, 5, 34, 0, 0, 318, 320, 5, 11,
		0, 0, 319, 311, 1, 0, 0, 0, 319, 313, 1, 0, 0, 0, 319, 316, 1, 0, 0, 0,
		320, 21, 1, 0, 0, 0, 321, 322, 6, 11, -1, 0, 322, 323, 3, 24, 12, 0, 323,
		331, 1, 0, 0, 0, 324, 325, 10, 2, 0, 0, 325, 330, 5, 3, 0, 0, 326, 327,
		10, 1, 0, 0, 327, 328, 5, 3, 0, 0, 328, 330, 5, 14, 0, 0, 329, 324, 1,
		0, 0, 0, 329, 326, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0,
		0, 331, 332, 1, 0, 0, 0, 332, 23, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334,
		347, 5, 31, 0, 0, 335, 336, 5, 17, 0, 0, 336, 347, 5, 31, 0, 0, 337, 338,
		5, 18, 0, 0, 338, 347, 5, 31, 0, 0, 339, 341, 5, 15, 0, 0, 340, 342, 5,
		31, 0, 0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 347, 1, 0, 0,
		0, 343, 347, 5, 16, 0, 0, 344, 345, 5, 14, 0, 0, 345, 347, 3, 24, 12, 0,
		346, 334, 1, 0, 0, 0, 346, 335, 1, 0, 0, 0, 346, 337, 1, 0, 0, 0, 346,
		339, 1, 0, 0, 0, 346, 343, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 347, 25, 1,
		0, 0, 0, 45, 29, 38, 42, 49, 53, 62, 67, 77, 84, 93, 100, 104, 111, 115,
		119, 127, 141, 152, 160, 166, 171, 179, 183, 188, 190, 199, 202, 209, 216,
		224, 233, 247, 253, 265, 268, 278, 283, 285, 304, 309, 319, 329, 331, 341,
		346,
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
	CParserT__0                = 1
	CParserAssign              = 2
	CParserStar                = 3
	CParserMinus               = 4
	CParserComma               = 5
	CParserSemi                = 6
	CParserColon               = 7
	CParserLParen              = 8
	CParserRParen              = 9
	CParserLBracket            = 10
	CParserRBracket            = 11
	CParserLBrace              = 12
	CParserRBrace              = 13
	CParserConst               = 14
	CParserUnsigned            = 15
	CParserVoid                = 16
	CParserEnum                = 17
	CParserStruct              = 18
	CParserStatic              = 19
	CParserInline              = 20
	CParserTypedef             = 21
	CParserAttributeDeprecated = 22
	CParserIfndef              = 23
	CParserIfdir               = 24
	CParserEndif               = 25
	CParserElseDir             = 26
	CParserDefine              = 27
	CParserUndef               = 28
	CParserPragma              = 29
	CParserInclude             = 30
	CParserIdentifier          = 31
	CParserWhitespace          = 32
	CParserNewline             = 33
	CParserConstant            = 34
	CParserDigitSequence       = 35
	CParserStringLiteral       = 36
	CParserBlockComment        = 37
	CParserLineComment         = 38
)

// CParser rules.
const (
	CParserRULE_unit         = 0
	CParserRULE_topLevel     = 1
	CParserRULE_elseCond     = 2
	CParserRULE_enumBody     = 3
	CParserRULE_enumItem     = 4
	CParserRULE_structItem   = 5
	CParserRULE_constant     = 6
	CParserRULE_params       = 7
	CParserRULE_param        = 8
	CParserRULE_arraySize    = 9
	CParserRULE_arraySizeDim = 10
	CParserRULE_type         = 11
	CParserRULE_typeInner    = 12
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4193239040) != 0 {
		{
			p.SetState(26)
			p.TopLevel()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
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

func (s *TopLevelIfdirContext) ElseCond() IElseCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseCondContext)
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

type TopLevelEnumTypedefContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelEnumTypedefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelEnumTypedefContext {
	var p = new(TopLevelEnumTypedefContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelEnumTypedefContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelEnumTypedefContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelEnumTypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelEnumTypedefContext) Typedef() antlr.TerminalNode {
	return s.GetToken(CParserTypedef, 0)
}

func (s *TopLevelEnumTypedefContext) Enum() antlr.TerminalNode {
	return s.GetToken(CParserEnum, 0)
}

func (s *TopLevelEnumTypedefContext) LBrace() antlr.TerminalNode {
	return s.GetToken(CParserLBrace, 0)
}

func (s *TopLevelEnumTypedefContext) EnumBody() IEnumBodyContext {
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

func (s *TopLevelEnumTypedefContext) RBrace() antlr.TerminalNode {
	return s.GetToken(CParserRBrace, 0)
}

func (s *TopLevelEnumTypedefContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *TopLevelEnumTypedefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelEnumTypedefContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
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

func (s *TopLevelStructContext) LBrace() antlr.TerminalNode {
	return s.GetToken(CParserLBrace, 0)
}

func (s *TopLevelStructContext) RBrace() antlr.TerminalNode {
	return s.GetToken(CParserRBrace, 0)
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

func (s *TopLevelIfndefContext) ElseCond() IElseCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseCondContext)
}

type TopLevelStaticFunctionContext struct {
	TopLevelContext
	itemname antlr.Token
}

func NewTopLevelStaticFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopLevelStaticFunctionContext {
	var p = new(TopLevelStaticFunctionContext)

	InitEmptyTopLevelContext(&p.TopLevelContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelContext))

	return p
}

func (s *TopLevelStaticFunctionContext) GetItemname() antlr.Token { return s.itemname }

func (s *TopLevelStaticFunctionContext) SetItemname(v antlr.Token) { s.itemname = v }

func (s *TopLevelStaticFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelStaticFunctionContext) Static() antlr.TerminalNode {
	return s.GetToken(CParserStatic, 0)
}

func (s *TopLevelStaticFunctionContext) Type_() ITypeContext {
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

func (s *TopLevelStaticFunctionContext) LParen() antlr.TerminalNode {
	return s.GetToken(CParserLParen, 0)
}

func (s *TopLevelStaticFunctionContext) Params() IParamsContext {
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

func (s *TopLevelStaticFunctionContext) RParen() antlr.TerminalNode {
	return s.GetToken(CParserRParen, 0)
}

func (s *TopLevelStaticFunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TopLevelStaticFunctionContext) Inline() antlr.TerminalNode {
	return s.GetToken(CParserInline, 0)
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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTopLevelIfndefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(CParserIfndef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4193239040) != 0 {
			{
				p.SetState(35)
				p.TopLevel()
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserElseDir {
			{
				p.SetState(41)
				p.ElseCond()
			}

		}
		{
			p.SetState(44)
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
			p.SetState(45)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4193239040) != 0 {
			{
				p.SetState(46)
				p.TopLevel()
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserElseDir {
			{
				p.SetState(52)
				p.ElseCond()
			}

		}
		{
			p.SetState(55)
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
			p.SetState(56)
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
			p.SetState(57)
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
			p.SetState(58)
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
			p.SetState(59)
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
			p.SetState(60)
			p.Match(CParserEnum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserIdentifier {
			{
				p.SetState(61)

				var _m = p.Match(CParserIdentifier)

				localctx.(*TopLevelEnumContext).itemname = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(64)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.EnumBody()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(66)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(69)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewTopLevelEnumTypedefContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(CParserEnum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.EnumBody()
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(76)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(79)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelEnumTypedefContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewTopLevelStructTypeDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(83)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(86)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructTypeDefContext).structname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Match(CParserLBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2303180800) != 0 {
			{
				p.SetState(90)
				p.StructItem()
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(96)
			p.Match(CParserRBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructTypeDefContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewTopLevelStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(99)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(102)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(103)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(106)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStructContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserLBrace {
			{
				p.SetState(107)
				p.Match(CParserLBrace)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2303180800) != 0 {
				{
					p.SetState(108)
					p.StructItem()
				}

				p.SetState(113)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(114)
				p.Match(CParserRBrace)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

	case 11:
		localctx = NewTopLevelTypeDefSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(118)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(121)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.type_(0)
		}
		{
			p.SetState(123)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelTypeDefSimpleContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewTopLevelTypeDefFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(126)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(129)
			p.Match(CParserTypedef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.type_(0)
		}
		{
			p.SetState(131)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelTypeDefFuncContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Params()
		}
		{
			p.SetState(137)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewTopLevelFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(140)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(143)
			p.type_(0)
		}
		{
			p.SetState(144)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelFunctionContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Params()
		}
		{
			p.SetState(147)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewTopLevelStaticFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(150)
			p.Match(CParserStatic)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserInline {
			{
				p.SetState(151)
				p.Match(CParserInline)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(154)
			p.type_(0)
		}
		{
			p.SetState(155)

			var _m = p.Match(CParserIdentifier)

			localctx.(*TopLevelStaticFunctionContext).itemname = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Params()
		}
		{
			p.SetState(158)
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

// IElseCondContext is an interface to support dynamic dispatch.
type IElseCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElseDir() antlr.TerminalNode
	AllTopLevel() []ITopLevelContext
	TopLevel(i int) ITopLevelContext

	// IsElseCondContext differentiates from other interfaces.
	IsElseCondContext()
}

type ElseCondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseCondContext() *ElseCondContext {
	var p = new(ElseCondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_elseCond
	return p
}

func InitEmptyElseCondContext(p *ElseCondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CParserRULE_elseCond
}

func (*ElseCondContext) IsElseCondContext() {}

func NewElseCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseCondContext {
	var p = new(ElseCondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_elseCond

	return p
}

func (s *ElseCondContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseCondContext) ElseDir() antlr.TerminalNode {
	return s.GetToken(CParserElseDir, 0)
}

func (s *ElseCondContext) AllTopLevel() []ITopLevelContext {
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

func (s *ElseCondContext) TopLevel(i int) ITopLevelContext {
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

func (s *ElseCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *CParser) ElseCond() (localctx IElseCondContext) {
	localctx = NewElseCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_elseCond)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(CParserElseDir)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4193239040) != 0 {
		{
			p.SetState(163)
			p.TopLevel()
		}

		p.SetState(168)
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

func (p *CParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_enumBody)
	var _la int

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CParserDefine:
		localctx = NewEnumBodyDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(CParserDefine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2298478592) != 0 {
			{
				p.SetState(170)
				p.EnumBody()
			}

		}

	case CParserIfdir:
		localctx = NewEnumBodyIfdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == CParserIdentifier {
			{
				p.SetState(174)
				p.EnumItem()
			}
			{
				p.SetState(175)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(181)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2298478592) != 0 {
			{
				p.SetState(182)
				p.EnumBody()
			}

		}

	case CParserIdentifier:
		localctx = NewEnumBodyItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.EnumItem()
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(186)
				p.Match(CParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(187)
				p.EnumBody()
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

type EnumItemMappedContext struct {
	EnumItemContext
}

func NewEnumItemMappedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumItemMappedContext {
	var p = new(EnumItemMappedContext)

	InitEmptyEnumItemContext(&p.EnumItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*EnumItemContext))

	return p
}

func (s *EnumItemMappedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemMappedContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CParserIdentifier)
}

func (s *EnumItemMappedContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, i)
}

func (s *EnumItemMappedContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
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
	p.EnterRule(localctx, 8, CParserRULE_enumItem)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEnumItemNoValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
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
			p.SetState(193)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(CParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Constant()
		}

	case 3:
		localctx = NewEnumItemMappedContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(CParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(CParserIdentifier)
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

type StructItemBitfieldContext struct {
	StructItemContext
}

func NewStructItemBitfieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructItemBitfieldContext {
	var p = new(StructItemBitfieldContext)

	InitEmptyStructItemContext(&p.StructItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructItemContext))

	return p
}

func (s *StructItemBitfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructItemBitfieldContext) Type_() ITypeContext {
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

func (s *StructItemBitfieldContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructItemBitfieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(CParserColon, 0)
}

func (s *StructItemBitfieldContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *StructItemBitfieldContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StructItemBitfieldContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
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

func (s *StructItemFunctionContext) AttributeDeprecated() antlr.TerminalNode {
	return s.GetToken(CParserAttributeDeprecated, 0)
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
	p.EnterRule(localctx, 10, CParserRULE_structItem)
	var _la int

	var _alt int

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructItemMemberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(201)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(204)
			p.type_(0)
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(205)
					p.Match(CParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(206)
					p.Match(CParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStructItemArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(215)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(218)
			p.type_(0)
		}
		{
			p.SetState(219)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.ArraySize()
		}
		{
			p.SetState(221)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStructItemBitfieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(223)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(226)
			p.type_(0)
		}
		{
			p.SetState(227)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(CParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(CParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStructItemArrayPtrContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(232)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(235)
			p.type_(0)
		}
		{
			p.SetState(236)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.ArraySize()
		}
		{
			p.SetState(241)
			p.Match(CParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStructItemIfdirContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(243)
			p.Match(CParserIfdir)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2303180800) != 0 {
			{
				p.SetState(244)
				p.StructItem()
			}

			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(250)
			p.Match(CParserEndif)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStructItemDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(251)
			p.Match(CParserDefine)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewStructItemFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CParserAttributeDeprecated {
			{
				p.SetState(252)
				p.Match(CParserAttributeDeprecated)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(255)
			p.type_(0)
		}
		{
			p.SetState(256)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Params()
		}
		{
			p.SetState(262)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
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
	p.EnterRule(localctx, 12, CParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CParserMinus {
		{
			p.SetState(267)

			var _m = p.Match(CParserMinus)

			localctx.(*ConstantContext).min = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(270)
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
	vararg antlr.Token
}

func NewParamsValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamsValuesContext {
	var p = new(ParamsValuesContext)

	InitEmptyParamsContext(&p.ParamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamsContext))

	return p
}

func (s *ParamsValuesContext) GetVararg() antlr.Token { return s.vararg }

func (s *ParamsValuesContext) SetVararg(v antlr.Token) { s.vararg = v }

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
	p.EnterRule(localctx, 14, CParserRULE_params)
	var _alt int

	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamsVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(CParserVoid)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewParamsValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(273)
					p.Param()
				}
				{
					p.SetState(274)
					p.Match(CParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CParserConst, CParserUnsigned, CParserVoid, CParserEnum, CParserStruct, CParserIdentifier:
			{
				p.SetState(281)
				p.Param()
			}

		case CParserT__0:
			{
				p.SetState(282)

				var _m = p.Match(CParserT__0)

				localctx.(*ParamsValuesContext).vararg = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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
	p.EnterRule(localctx, 16, CParserRULE_param)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.type_(0)
		}
		{
			p.SetState(288)
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
			p.SetState(290)
			p.type_(0)
		}
		{
			p.SetState(291)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.ArraySize()
		}

	case 3:
		localctx = NewParamAnonContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(294)
			p.type_(0)
		}

	case 4:
		localctx = NewParamFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(295)
			p.type_(0)
		}
		{
			p.SetState(296)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(CParserStar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(CParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(CParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Params()
		}
		{
			p.SetState(302)
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
	p.EnterRule(localctx, 18, CParserRULE_arraySize)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CParserLBracket {
		{
			p.SetState(306)
			p.ArraySizeDim()
		}

		p.SetState(309)
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

type ArraySizeDimUnsizedContext struct {
	ArraySizeDimContext
}

func NewArraySizeDimUnsizedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArraySizeDimUnsizedContext {
	var p = new(ArraySizeDimUnsizedContext)

	InitEmptyArraySizeDimContext(&p.ArraySizeDimContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArraySizeDimContext))

	return p
}

func (s *ArraySizeDimUnsizedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeDimUnsizedContext) LBracket() antlr.TerminalNode {
	return s.GetToken(CParserLBracket, 0)
}

func (s *ArraySizeDimUnsizedContext) RBracket() antlr.TerminalNode {
	return s.GetToken(CParserRBracket, 0)
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
	p.EnterRule(localctx, 20, CParserRULE_arraySizeDim)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArraySizeDimUnsizedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.Match(CParserLBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(CParserRBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewArraySizeDimIdenContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(CParserLBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(CParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(CParserRBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewArraySizeDimConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.Match(CParserLBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.Match(CParserConstant)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, CParserRULE_type, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewTypePassContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(322)
		p.TypeInner()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewTypePtrContext(p, NewTypeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_type)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(325)
					p.Match(CParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewTypePtrConstContext(p, NewTypeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_type)
				p.SetState(326)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(327)
					p.Match(CParserStar)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(328)
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
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, CParserRULE_typeInner)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		localctx = NewTypeInnerIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
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
			p.SetState(335)
			p.Match(CParserEnum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
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
			p.SetState(337)
			p.Match(CParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
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
			p.SetState(339)
			p.Match(CParserUnsigned)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(340)
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
			p.SetState(343)
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
			p.SetState(344)
			p.Match(CParserConst)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
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
	case 11:
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
