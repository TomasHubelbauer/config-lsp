// Code generated from Fstab.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type FstabLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FstabLexerLexerStaticData struct {
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

func fstablexerLexerInit() {
	staticData := &FstabLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "DIGITS", "WHITESPACE", "HASH", "STRING", "QUOTED_STRING", "ADFS",
		"AFFS", "BTRFS", "EXFAT",
	}
	staticData.RuleNames = []string{
		"DIGITS", "WHITESPACE", "HASH", "STRING", "QUOTED_STRING", "ADFS", "AFFS",
		"BTRFS", "EXFAT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 76, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 21,
		8, 0, 11, 0, 12, 0, 22, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 1, 2, 1,
		2, 1, 3, 4, 3, 33, 8, 3, 11, 3, 12, 3, 34, 1, 4, 1, 4, 3, 4, 39, 8, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 44, 8, 4, 10, 4, 12, 4, 47, 9, 4, 1, 4, 3, 4, 50,
		8, 4, 1, 4, 3, 4, 53, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 1, 0, 12, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 3, 0, 9, 9, 32, 32,
		35, 35, 2, 0, 65, 65, 97, 97, 2, 0, 68, 68, 100, 100, 2, 0, 70, 70, 102,
		102, 2, 0, 83, 83, 115, 115, 2, 0, 66, 66, 98, 98, 2, 0, 84, 84, 116, 116,
		2, 0, 82, 82, 114, 114, 2, 0, 69, 69, 101, 101, 2, 0, 88, 88, 120, 120,
		82, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0,
		0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0,
		0, 0, 0, 17, 1, 0, 0, 0, 1, 20, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 29, 1,
		0, 0, 0, 7, 32, 1, 0, 0, 0, 9, 36, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13,
		59, 1, 0, 0, 0, 15, 64, 1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 21, 7, 0, 0,
		0, 20, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23,
		1, 0, 0, 0, 23, 2, 1, 0, 0, 0, 24, 26, 7, 1, 0, 0, 25, 24, 1, 0, 0, 0,
		26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 4, 1, 0,
		0, 0, 29, 30, 5, 35, 0, 0, 30, 6, 1, 0, 0, 0, 31, 33, 8, 2, 0, 0, 32, 31,
		1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 8, 1, 0, 0, 0, 36, 38, 5, 34, 0, 0, 37, 39, 3, 3, 1, 0, 38, 37, 1,
		0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 45, 1, 0, 0, 0, 40, 41, 3, 7, 3, 0, 41,
		42, 3, 3, 1, 0, 42, 44, 1, 0, 0, 0, 43, 40, 1, 0, 0, 0, 44, 47, 1, 0, 0,
		0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 48, 50, 3, 7, 3, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0,
		50, 52, 1, 0, 0, 0, 51, 53, 5, 34, 0, 0, 52, 51, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 10, 1, 0, 0, 0, 54, 55, 7, 3, 0, 0, 55, 56, 7, 4, 0, 0, 56,
		57, 7, 5, 0, 0, 57, 58, 7, 6, 0, 0, 58, 12, 1, 0, 0, 0, 59, 60, 7, 3, 0,
		0, 60, 61, 7, 5, 0, 0, 61, 62, 7, 5, 0, 0, 62, 63, 7, 6, 0, 0, 63, 14,
		1, 0, 0, 0, 64, 65, 7, 7, 0, 0, 65, 66, 7, 8, 0, 0, 66, 67, 7, 9, 0, 0,
		67, 68, 7, 5, 0, 0, 68, 69, 7, 6, 0, 0, 69, 16, 1, 0, 0, 0, 70, 71, 7,
		10, 0, 0, 71, 72, 7, 11, 0, 0, 72, 73, 7, 5, 0, 0, 73, 74, 7, 3, 0, 0,
		74, 75, 7, 8, 0, 0, 75, 18, 1, 0, 0, 0, 8, 0, 22, 27, 34, 38, 45, 49, 52,
		0,
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

// FstabLexerInit initializes any static state used to implement FstabLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFstabLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FstabLexerInit() {
	staticData := &FstabLexerLexerStaticData
	staticData.once.Do(fstablexerLexerInit)
}

// NewFstabLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFstabLexer(input antlr.CharStream) *FstabLexer {
	FstabLexerInit()
	l := new(FstabLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FstabLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Fstab.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FstabLexer tokens.
const (
	FstabLexerDIGITS        = 1
	FstabLexerWHITESPACE    = 2
	FstabLexerHASH          = 3
	FstabLexerSTRING        = 4
	FstabLexerQUOTED_STRING = 5
	FstabLexerADFS          = 6
	FstabLexerAFFS          = 7
	FstabLexerBTRFS         = 8
	FstabLexerEXFAT         = 9
)
