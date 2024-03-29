// Code generated from ./Agtype.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AgtypeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AgtypeLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agtypelexerLexerInit() {
	staticData := &AgtypeLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'true'", "'false'", "'null'", "'{'", "','", "'}'", "':'", "'['",
		"']'", "'::'", "'-'", "'Infinity'", "'NaN'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENT", "STRING",
		"INTEGER", "RegularFloat", "ExponentFloat", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "IDENT", "STRING", "ESC", "UNICODE",
		"HEX", "SAFECODEPOINT", "INTEGER", "INT", "RegularFloat", "ExponentFloat",
		"DECIMAL", "SCIENTIFIC", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 183, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 5, 13, 102, 8, 13, 10, 13, 12, 13, 105, 9, 13, 1, 14, 1,
		14, 1, 14, 5, 14, 110, 8, 14, 10, 14, 12, 14, 113, 9, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 3, 15, 120, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 3, 19, 133, 8, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 140, 8, 20, 10, 20, 12, 20, 143, 9,
		20, 3, 20, 145, 8, 20, 1, 21, 3, 21, 148, 8, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 3, 22, 154, 8, 22, 1, 22, 1, 22, 3, 22, 158, 8, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 4, 23, 164, 8, 23, 11, 23, 12, 23, 165, 1, 24, 1, 24, 3, 24,
		170, 8, 24, 1, 24, 4, 24, 173, 8, 24, 11, 24, 12, 24, 174, 1, 25, 4, 25,
		178, 8, 25, 11, 25, 12, 25, 179, 1, 25, 1, 25, 0, 0, 26, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 0, 33, 0, 35, 0, 37, 0, 39, 16, 41, 0, 43, 17,
		45, 18, 47, 0, 49, 0, 51, 19, 1, 0, 10, 3, 0, 65, 90, 95, 95, 97, 122,
		5, 0, 36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 8, 0, 34, 34, 47, 47, 92,
		92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70,
		97, 102, 3, 0, 0, 31, 34, 34, 92, 92, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0,
		69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 9, 10, 13, 13, 32, 32, 189,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53,
		1, 0, 0, 0, 3, 58, 1, 0, 0, 0, 5, 64, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9,
		71, 1, 0, 0, 0, 11, 73, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 77, 1, 0, 0,
		0, 17, 79, 1, 0, 0, 0, 19, 81, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 86,
		1, 0, 0, 0, 25, 95, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0, 29, 106, 1, 0, 0, 0,
		31, 116, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37, 129,
		1, 0, 0, 0, 39, 132, 1, 0, 0, 0, 41, 144, 1, 0, 0, 0, 43, 147, 1, 0, 0,
		0, 45, 153, 1, 0, 0, 0, 47, 161, 1, 0, 0, 0, 49, 167, 1, 0, 0, 0, 51, 177,
		1, 0, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 117,
		0, 0, 56, 57, 5, 101, 0, 0, 57, 2, 1, 0, 0, 0, 58, 59, 5, 102, 0, 0, 59,
		60, 5, 97, 0, 0, 60, 61, 5, 108, 0, 0, 61, 62, 5, 115, 0, 0, 62, 63, 5,
		101, 0, 0, 63, 4, 1, 0, 0, 0, 64, 65, 5, 110, 0, 0, 65, 66, 5, 117, 0,
		0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 108, 0, 0, 68, 6, 1, 0, 0, 0, 69, 70,
		5, 123, 0, 0, 70, 8, 1, 0, 0, 0, 71, 72, 5, 44, 0, 0, 72, 10, 1, 0, 0,
		0, 73, 74, 5, 125, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 5, 58, 0, 0, 76, 14,
		1, 0, 0, 0, 77, 78, 5, 91, 0, 0, 78, 16, 1, 0, 0, 0, 79, 80, 5, 93, 0,
		0, 80, 18, 1, 0, 0, 0, 81, 82, 5, 58, 0, 0, 82, 83, 5, 58, 0, 0, 83, 20,
		1, 0, 0, 0, 84, 85, 5, 45, 0, 0, 85, 22, 1, 0, 0, 0, 86, 87, 5, 73, 0,
		0, 87, 88, 5, 110, 0, 0, 88, 89, 5, 102, 0, 0, 89, 90, 5, 105, 0, 0, 90,
		91, 5, 110, 0, 0, 91, 92, 5, 105, 0, 0, 92, 93, 5, 116, 0, 0, 93, 94, 5,
		121, 0, 0, 94, 24, 1, 0, 0, 0, 95, 96, 5, 78, 0, 0, 96, 97, 5, 97, 0, 0,
		97, 98, 5, 78, 0, 0, 98, 26, 1, 0, 0, 0, 99, 103, 7, 0, 0, 0, 100, 102,
		7, 1, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0,
		0, 0, 103, 104, 1, 0, 0, 0, 104, 28, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0,
		106, 111, 5, 34, 0, 0, 107, 110, 3, 31, 15, 0, 108, 110, 3, 37, 18, 0,
		109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 114, 115, 5, 34, 0, 0, 115, 30, 1, 0, 0, 0, 116, 119, 5, 92,
		0, 0, 117, 120, 7, 2, 0, 0, 118, 120, 3, 33, 16, 0, 119, 117, 1, 0, 0,
		0, 119, 118, 1, 0, 0, 0, 120, 32, 1, 0, 0, 0, 121, 122, 5, 117, 0, 0, 122,
		123, 3, 35, 17, 0, 123, 124, 3, 35, 17, 0, 124, 125, 3, 35, 17, 0, 125,
		126, 3, 35, 17, 0, 126, 34, 1, 0, 0, 0, 127, 128, 7, 3, 0, 0, 128, 36,
		1, 0, 0, 0, 129, 130, 8, 4, 0, 0, 130, 38, 1, 0, 0, 0, 131, 133, 5, 45,
		0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 135, 3, 41, 20, 0, 135, 40, 1, 0, 0, 0, 136, 145, 5, 48, 0, 0, 137,
		141, 7, 5, 0, 0, 138, 140, 7, 6, 0, 0, 139, 138, 1, 0, 0, 0, 140, 143,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 145, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 144, 136, 1, 0, 0, 0, 144, 137, 1, 0, 0, 0,
		145, 42, 1, 0, 0, 0, 146, 148, 5, 45, 0, 0, 147, 146, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 3, 41, 20, 0, 150, 151,
		3, 47, 23, 0, 151, 44, 1, 0, 0, 0, 152, 154, 5, 45, 0, 0, 153, 152, 1,
		0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 3, 41, 20,
		0, 156, 158, 3, 47, 23, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 159, 1, 0, 0, 0, 159, 160, 3, 49, 24, 0, 160, 46, 1, 0, 0, 0, 161,
		163, 5, 46, 0, 0, 162, 164, 7, 6, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 48, 1, 0,
		0, 0, 167, 169, 7, 7, 0, 0, 168, 170, 7, 8, 0, 0, 169, 168, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 173, 7, 6, 0, 0, 172,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 50, 1, 0, 0, 0, 176, 178, 7, 9, 0, 0, 177, 176, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0,
		180, 181, 1, 0, 0, 0, 181, 182, 6, 25, 0, 0, 182, 52, 1, 0, 0, 0, 15, 0,
		103, 109, 111, 119, 132, 141, 144, 147, 153, 157, 165, 169, 174, 179, 1,
		6, 0, 0,
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

// AgtypeLexerInit initializes any static state used to implement AgtypeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAgtypeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgtypeLexerInit() {
	staticData := &AgtypeLexerLexerStaticData
	staticData.once.Do(agtypelexerLexerInit)
}

// NewAgtypeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAgtypeLexer(input antlr.CharStream) *AgtypeLexer {
	AgtypeLexerInit()
	l := new(AgtypeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AgtypeLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Agtype.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AgtypeLexer tokens.
const (
	AgtypeLexerT__0          = 1
	AgtypeLexerT__1          = 2
	AgtypeLexerT__2          = 3
	AgtypeLexerT__3          = 4
	AgtypeLexerT__4          = 5
	AgtypeLexerT__5          = 6
	AgtypeLexerT__6          = 7
	AgtypeLexerT__7          = 8
	AgtypeLexerT__8          = 9
	AgtypeLexerT__9          = 10
	AgtypeLexerT__10         = 11
	AgtypeLexerT__11         = 12
	AgtypeLexerT__12         = 13
	AgtypeLexerIDENT         = 14
	AgtypeLexerSTRING        = 15
	AgtypeLexerINTEGER       = 16
	AgtypeLexerRegularFloat  = 17
	AgtypeLexerExponentFloat = 18
	AgtypeLexerWS            = 19
)
