// Code generated from Match.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type MatchLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MatchLexerLexerStaticData struct {
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

func matchlexerLexerInit() {
	staticData := &MatchLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "','",
	}
	staticData.SymbolicNames = []string{
		"", "USER", "GROUP", "HOST", "LOCALADDRESS", "LOCALPORT", "RDOMAIN",
		"ADDRESS", "COMMA", "STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"USER", "GROUP", "HOST", "LOCALADDRESS", "LOCALPORT", "RDOMAIN", "ADDRESS",
		"COMMA", "STRING", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 88, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 80, 8, 8,
		11, 8, 12, 8, 81, 1, 9, 4, 9, 85, 8, 9, 11, 9, 12, 9, 86, 0, 0, 10, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 1, 0, 18,
		2, 0, 85, 85, 117, 117, 2, 0, 83, 83, 115, 115, 2, 0, 69, 69, 101, 101,
		2, 0, 82, 82, 114, 114, 2, 0, 71, 71, 103, 103, 2, 0, 79, 79, 111, 111,
		2, 0, 80, 80, 112, 112, 2, 0, 72, 72, 104, 104, 2, 0, 84, 84, 116, 116,
		2, 0, 76, 76, 108, 108, 2, 0, 67, 67, 99, 99, 2, 0, 65, 65, 97, 97, 2,
		0, 68, 68, 100, 100, 2, 0, 77, 77, 109, 109, 2, 0, 73, 73, 105, 105, 2,
		0, 78, 78, 110, 110, 5, 0, 9, 10, 13, 13, 32, 32, 35, 35, 44, 44, 2, 0,
		9, 9, 32, 32, 89, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0,
		0, 0, 3, 26, 1, 0, 0, 0, 5, 32, 1, 0, 0, 0, 7, 37, 1, 0, 0, 0, 9, 50, 1,
		0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15, 76, 1, 0, 0, 0, 17,
		79, 1, 0, 0, 0, 19, 84, 1, 0, 0, 0, 21, 22, 7, 0, 0, 0, 22, 23, 7, 1, 0,
		0, 23, 24, 7, 2, 0, 0, 24, 25, 7, 3, 0, 0, 25, 2, 1, 0, 0, 0, 26, 27, 7,
		4, 0, 0, 27, 28, 7, 3, 0, 0, 28, 29, 7, 5, 0, 0, 29, 30, 7, 0, 0, 0, 30,
		31, 7, 6, 0, 0, 31, 4, 1, 0, 0, 0, 32, 33, 7, 7, 0, 0, 33, 34, 7, 5, 0,
		0, 34, 35, 7, 1, 0, 0, 35, 36, 7, 8, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 7,
		9, 0, 0, 38, 39, 7, 5, 0, 0, 39, 40, 7, 10, 0, 0, 40, 41, 7, 11, 0, 0,
		41, 42, 7, 9, 0, 0, 42, 43, 7, 11, 0, 0, 43, 44, 7, 12, 0, 0, 44, 45, 7,
		12, 0, 0, 45, 46, 7, 3, 0, 0, 46, 47, 7, 2, 0, 0, 47, 48, 7, 1, 0, 0, 48,
		49, 7, 1, 0, 0, 49, 8, 1, 0, 0, 0, 50, 51, 7, 9, 0, 0, 51, 52, 7, 5, 0,
		0, 52, 53, 7, 10, 0, 0, 53, 54, 7, 11, 0, 0, 54, 55, 7, 9, 0, 0, 55, 56,
		7, 6, 0, 0, 56, 57, 7, 5, 0, 0, 57, 58, 7, 3, 0, 0, 58, 59, 7, 8, 0, 0,
		59, 10, 1, 0, 0, 0, 60, 61, 7, 3, 0, 0, 61, 62, 7, 12, 0, 0, 62, 63, 7,
		5, 0, 0, 63, 64, 7, 13, 0, 0, 64, 65, 7, 11, 0, 0, 65, 66, 7, 14, 0, 0,
		66, 67, 7, 15, 0, 0, 67, 12, 1, 0, 0, 0, 68, 69, 7, 11, 0, 0, 69, 70, 7,
		12, 0, 0, 70, 71, 7, 12, 0, 0, 71, 72, 7, 3, 0, 0, 72, 73, 7, 2, 0, 0,
		73, 74, 7, 1, 0, 0, 74, 75, 7, 1, 0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 5,
		44, 0, 0, 77, 16, 1, 0, 0, 0, 78, 80, 8, 16, 0, 0, 79, 78, 1, 0, 0, 0,
		80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 18, 1,
		0, 0, 0, 83, 85, 7, 17, 0, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 20, 1, 0, 0, 0, 3, 0, 81, 86, 0,
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

// MatchLexerInit initializes any static state used to implement MatchLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMatchLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchLexerInit() {
	staticData := &MatchLexerLexerStaticData
	staticData.once.Do(matchlexerLexerInit)
}

// NewMatchLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMatchLexer(input antlr.CharStream) *MatchLexer {
	MatchLexerInit()
	l := new(MatchLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MatchLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Match.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MatchLexer tokens.
const (
	MatchLexerUSER         = 1
	MatchLexerGROUP        = 2
	MatchLexerHOST         = 3
	MatchLexerLOCALADDRESS = 4
	MatchLexerLOCALPORT    = 5
	MatchLexerRDOMAIN      = 6
	MatchLexerADDRESS      = 7
	MatchLexerCOMMA        = 8
	MatchLexerSTRING       = 9
	MatchLexerWHITESPACE   = 10
)
