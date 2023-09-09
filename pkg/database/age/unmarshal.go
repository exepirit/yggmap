package age

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/exepirit/yggmap/pkg/database/age/parser"
	"strconv"
	"strings"
)

func newUnmarshaller() *unmarshaller {
	return &unmarshaller{}
}

type unmarshaller struct{}

func (u *unmarshaller) unmarshal(text string) (Agtype, error) {
	ageParser := parser.NewAgtypeParser(nil)
	input := antlr.NewInputStream(text)
	lexer := parser.NewAgtypeLexer(input)
	ageParser.SetInputStream(
		antlr.NewCommonTokenStream(lexer, 0),
	)

	var visitor unmarshallerVisitor
	tree := ageParser.AgType()
	result := tree.Accept(&visitor)

	return result.(Agtype), nil
}

type unmarshallerVisitor struct{}

func (v unmarshallerVisitor) VisitAgType(ctx *parser.AgTypeContext) interface{} {
	agtypeValue := ctx.AgValue()
	if agtypeValue != nil {
		return agtypeValue.Accept(v)
	}
	return nil
}

func (v unmarshallerVisitor) VisitAgValue(ctx *parser.AgValueContext) interface{} {
	typeAnnotationCtx := ctx.TypeAnnotation()
	valueCtx := ctx.Value()

	if typeAnnotationCtx != nil {
		typeAnnotation := typeAnnotationCtx.(*parser.TypeAnnotationContext).IDENT().GetText()

		switch typeAnnotation {
		case "vertex":
			dictionary := valueCtx.Accept(v).(map[string]any)
			return newVertex(dictionary)
		default:
			return nil
		}
	}

	return valueCtx.Accept(v)
}

func (v unmarshallerVisitor) VisitStringValue(ctx *parser.StringValueContext) interface{} {
	return ctx.GetText()
}

func (v unmarshallerVisitor) VisitIntegerValue(ctx *parser.IntegerValueContext) interface{} {
	value, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return value
}

func (v unmarshallerVisitor) VisitFloatValue(ctx *parser.FloatValueContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v unmarshallerVisitor) VisitTrueBoolean(ctx *parser.TrueBooleanContext) interface{} {
	return true
}

func (v unmarshallerVisitor) VisitFalseBoolean(ctx *parser.FalseBooleanContext) interface{} {
	return false
}

func (v unmarshallerVisitor) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return nil
}

func (v unmarshallerVisitor) VisitObjectValue(ctx *parser.ObjectValueContext) interface{} {
	return ctx.GetChild(0).(*parser.ObjContext).Accept(v)
}

func (v unmarshallerVisitor) VisitArrayValue(ctx *parser.ArrayValueContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v unmarshallerVisitor) VisitObj(ctx *parser.ObjContext) interface{} {
	props := make(map[string]interface{})
	for _, pairCtx := range ctx.AllPair() {
		pairCtx.Accept(v)
		pair := pairCtx.(*parser.PairContext)
		key := strings.Trim(pair.STRING().GetText(), "\"")
		value := pair.AgValue().Accept(v)
		props[key] = value
	}

	return props
}

func (v unmarshallerVisitor) VisitPair(ctx *parser.PairContext) interface{} {
	return nil
}

func (v unmarshallerVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v unmarshallerVisitor) VisitTypeAnnotation(ctx *parser.TypeAnnotationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v unmarshallerVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v unmarshallerVisitor) Visit(tree antlr.ParseTree) interface{} {
	return nil
}

func (v unmarshallerVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for _, c := range node.GetChildren() {
		pt := c.(antlr.ParseTree)
		result = pt.Accept(v)
	}
	return result
}

func (v unmarshallerVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v unmarshallerVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}
