// Code generated from ./Agtype.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Agtype

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

type AgtypeParser struct {
	*antlr.BaseParser
}

var AgtypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agtypeParserInit() {
	staticData := &AgtypeParserStaticData
	staticData.LiteralNames = []string{
		"", "'true'", "'false'", "'null'", "'{'", "','", "'}'", "':'", "'['",
		"']'", "'::'", "'-'", "'Infinity'", "'NaN'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENT", "STRING",
		"INTEGER", "RegularFloat", "ExponentFloat", "WS",
	}
	staticData.RuleNames = []string{
		"agType", "agValue", "value", "obj", "pair", "array", "typeAnnotation",
		"floatLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 80, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3,
		1, 22, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 32,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 38, 8, 3, 10, 3, 12, 3, 41, 9, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 66, 8, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 74, 8, 7,
		1, 7, 1, 7, 3, 7, 78, 8, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0,
		0, 87, 0, 16, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 46,
		1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 65, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0,
		14, 77, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0,
		0, 0, 19, 21, 3, 4, 2, 0, 20, 22, 3, 12, 6, 0, 21, 20, 1, 0, 0, 0, 21,
		22, 1, 0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 32, 5, 15, 0, 0, 24, 32, 5, 16,
		0, 0, 25, 32, 3, 14, 7, 0, 26, 32, 5, 1, 0, 0, 27, 32, 5, 2, 0, 0, 28,
		32, 5, 3, 0, 0, 29, 32, 3, 6, 3, 0, 30, 32, 3, 10, 5, 0, 31, 23, 1, 0,
		0, 0, 31, 24, 1, 0, 0, 0, 31, 25, 1, 0, 0, 0, 31, 26, 1, 0, 0, 0, 31, 27,
		1, 0, 0, 0, 31, 28, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0,
		32, 5, 1, 0, 0, 0, 33, 34, 5, 4, 0, 0, 34, 39, 3, 8, 4, 0, 35, 36, 5, 5,
		0, 0, 36, 38, 3, 8, 4, 0, 37, 35, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37,
		1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		42, 43, 5, 6, 0, 0, 43, 47, 1, 0, 0, 0, 44, 45, 5, 4, 0, 0, 45, 47, 5,
		6, 0, 0, 46, 33, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48,
		49, 5, 15, 0, 0, 49, 50, 5, 7, 0, 0, 50, 51, 3, 2, 1, 0, 51, 9, 1, 0, 0,
		0, 52, 53, 5, 8, 0, 0, 53, 58, 3, 2, 1, 0, 54, 55, 5, 5, 0, 0, 55, 57,
		3, 2, 1, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0,
		58, 59, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 5,
		9, 0, 0, 62, 66, 1, 0, 0, 0, 63, 64, 5, 8, 0, 0, 64, 66, 5, 9, 0, 0, 65,
		52, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 11, 1, 0, 0, 0, 67, 68, 5, 10,
		0, 0, 68, 69, 5, 14, 0, 0, 69, 13, 1, 0, 0, 0, 70, 78, 5, 17, 0, 0, 71,
		78, 5, 18, 0, 0, 72, 74, 5, 11, 0, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0,
		0, 0, 74, 75, 1, 0, 0, 0, 75, 78, 5, 12, 0, 0, 76, 78, 5, 13, 0, 0, 77,
		70, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 77, 76, 1, 0, 0,
		0, 78, 15, 1, 0, 0, 0, 8, 21, 31, 39, 46, 58, 65, 73, 77,
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

// AgtypeParserInit initializes any static state used to implement AgtypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAgtypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgtypeParserInit() {
	staticData := &AgtypeParserStaticData
	staticData.once.Do(agtypeParserInit)
}

// NewAgtypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAgtypeParser(input antlr.TokenStream) *AgtypeParser {
	AgtypeParserInit()
	this := new(AgtypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AgtypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Agtype.g4"

	return this
}

// AgtypeParser tokens.
const (
	AgtypeParserEOF           = antlr.TokenEOF
	AgtypeParserT__0          = 1
	AgtypeParserT__1          = 2
	AgtypeParserT__2          = 3
	AgtypeParserT__3          = 4
	AgtypeParserT__4          = 5
	AgtypeParserT__5          = 6
	AgtypeParserT__6          = 7
	AgtypeParserT__7          = 8
	AgtypeParserT__8          = 9
	AgtypeParserT__9          = 10
	AgtypeParserT__10         = 11
	AgtypeParserT__11         = 12
	AgtypeParserT__12         = 13
	AgtypeParserIDENT         = 14
	AgtypeParserSTRING        = 15
	AgtypeParserINTEGER       = 16
	AgtypeParserRegularFloat  = 17
	AgtypeParserExponentFloat = 18
	AgtypeParserWS            = 19
)

// AgtypeParser rules.
const (
	AgtypeParserRULE_agType         = 0
	AgtypeParserRULE_agValue        = 1
	AgtypeParserRULE_value          = 2
	AgtypeParserRULE_obj            = 3
	AgtypeParserRULE_pair           = 4
	AgtypeParserRULE_array          = 5
	AgtypeParserRULE_typeAnnotation = 6
	AgtypeParserRULE_floatLiteral   = 7
)

// IAgTypeContext is an interface to support dynamic dispatch.
type IAgTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AgValue() IAgValueContext
	EOF() antlr.TerminalNode

	// IsAgTypeContext differentiates from other interfaces.
	IsAgTypeContext()
}

type AgTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgTypeContext() *AgTypeContext {
	var p = new(AgTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_agType
	return p
}

func InitEmptyAgTypeContext(p *AgTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_agType
}

func (*AgTypeContext) IsAgTypeContext() {}

func NewAgTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgTypeContext {
	var p = new(AgTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_agType

	return p
}

func (s *AgTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AgTypeContext) AgValue() IAgValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAgValueContext)
}

func (s *AgTypeContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgtypeParserEOF, 0)
}

func (s *AgTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterAgType(s)
	}
}

func (s *AgTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitAgType(s)
	}
}

func (s *AgTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitAgType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) AgType() (localctx IAgTypeContext) {
	localctx = NewAgTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgtypeParserRULE_agType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.AgValue()
	}
	{
		p.SetState(17)
		p.Match(AgtypeParserEOF)
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

// IAgValueContext is an interface to support dynamic dispatch.
type IAgValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	TypeAnnotation() ITypeAnnotationContext

	// IsAgValueContext differentiates from other interfaces.
	IsAgValueContext()
}

type AgValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgValueContext() *AgValueContext {
	var p = new(AgValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_agValue
	return p
}

func InitEmptyAgValueContext(p *AgValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_agValue
}

func (*AgValueContext) IsAgValueContext() {}

func NewAgValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgValueContext {
	var p = new(AgValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_agValue

	return p
}

func (s *AgValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AgValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AgValueContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *AgValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterAgValue(s)
	}
}

func (s *AgValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitAgValue(s)
	}
}

func (s *AgValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitAgValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) AgValue() (localctx IAgValueContext) {
	localctx = NewAgValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgtypeParserRULE_agValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Value()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgtypeParserT__9 {
		{
			p.SetState(20)
			p.TypeAnnotation()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullValueContext struct {
	ValueContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ObjectValueContext struct {
	ValueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (s *ObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerValueContext struct {
	ValueContext
}

func NewIntegerValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerValueContext {
	var p = new(IntegerValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AgtypeParserINTEGER, 0)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (s *IntegerValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitIntegerValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueBooleanContext struct {
	ValueContext
}

func NewTrueBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueBooleanContext {
	var p = new(TrueBooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *TrueBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterTrueBoolean(s)
	}
}

func (s *TrueBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitTrueBoolean(s)
	}
}

func (s *TrueBooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitTrueBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseBooleanContext struct {
	ValueContext
}

func NewFalseBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseBooleanContext {
	var p = new(FalseBooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *FalseBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFalseBoolean(s)
	}
}

func (s *FalseBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFalseBoolean(s)
	}
}

func (s *FalseBooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFalseBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatValueContext struct {
	ValueContext
}

func NewFloatValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatValueContext {
	var p = new(FloatValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (s *FloatValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFloatValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringValueContext struct {
	ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgtypeParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayValueContext struct {
	ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AgtypeParserRULE_value)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AgtypeParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(AgtypeParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserINTEGER:
		localctx = NewIntegerValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(AgtypeParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__10, AgtypeParserT__11, AgtypeParserT__12, AgtypeParserRegularFloat, AgtypeParserExponentFloat:
		localctx = NewFloatValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.FloatLiteral()
		}

	case AgtypeParserT__0:
		localctx = NewTrueBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
			p.Match(AgtypeParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__1:
		localctx = NewFalseBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(27)
			p.Match(AgtypeParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__2:
		localctx = NewNullValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(28)
			p.Match(AgtypeParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__3:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(29)
			p.Obj()
		}

	case AgtypeParserT__7:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(30)
			p.Array()
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

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPair() []IPairContext
	Pair(i int) IPairContext

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_obj
	return p
}

func InitEmptyObjContext(p *ObjContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_obj
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitObj(s)
	}
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AgtypeParserRULE_obj)
	var _la int

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(AgtypeParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Pair()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == AgtypeParserT__4 {
			{
				p.SetState(35)
				p.Match(AgtypeParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(36)
				p.Pair()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(AgtypeParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(AgtypeParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(AgtypeParserT__5)
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

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	AgValue() IAgValueContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgtypeParserSTRING, 0)
}

func (s *PairContext) AgValue() IAgValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAgValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AgtypeParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(AgtypeParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(AgtypeParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.AgValue()
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

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAgValue() []IAgValueContext
	AgValue(i int) IAgValueContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllAgValue() []IAgValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAgValueContext); ok {
			len++
		}
	}

	tst := make([]IAgValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAgValueContext); ok {
			tst[i] = t.(IAgValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) AgValue(i int) IAgValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgValueContext); ok {
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

	return t.(IAgValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AgtypeParserRULE_array)
	var _la int

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(AgtypeParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.AgValue()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == AgtypeParserT__4 {
			{
				p.SetState(54)
				p.Match(AgtypeParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(55)
				p.AgValue()
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(61)
			p.Match(AgtypeParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(AgtypeParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(AgtypeParserT__8)
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

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_typeAnnotation
	return p
}

func InitEmptyTypeAnnotationContext(p *TypeAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_typeAnnotation
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(AgtypeParserIDENT, 0)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AgtypeParserRULE_typeAnnotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(AgtypeParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(AgtypeParserIDENT)
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

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RegularFloat() antlr.TerminalNode
	ExponentFloat() antlr.TerminalNode

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_floatLiteral
	return p
}

func InitEmptyFloatLiteralContext(p *FloatLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgtypeParserRULE_floatLiteral
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) RegularFloat() antlr.TerminalNode {
	return s.GetToken(AgtypeParserRegularFloat, 0)
}

func (s *FloatLiteralContext) ExponentFloat() antlr.TerminalNode {
	return s.GetToken(AgtypeParserExponentFloat, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AgtypeParserRULE_floatLiteral)
	var _la int

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AgtypeParserRegularFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(AgtypeParserRegularFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserExponentFloat:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(AgtypeParserExponentFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__10, AgtypeParserT__11:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AgtypeParserT__10 {
			{
				p.SetState(72)
				p.Match(AgtypeParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(75)
			p.Match(AgtypeParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgtypeParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Match(AgtypeParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
