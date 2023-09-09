// Code generated from ./Agtype.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Agtype

import "github.com/antlr4-go/antlr/v4"

type BaseAgtypeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAgtypeVisitor) VisitAgType(ctx *AgTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitAgValue(ctx *AgValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFloatValue(ctx *FloatValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitTrueBoolean(ctx *TrueBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFalseBoolean(ctx *FalseBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitObjectValue(ctx *ObjectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
