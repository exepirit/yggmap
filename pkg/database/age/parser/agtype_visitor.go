// Code generated from ./Agtype.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Agtype

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AgtypeParser.
type AgtypeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AgtypeParser#agType.
	VisitAgType(ctx *AgTypeContext) interface{}

	// Visit a parse tree produced by AgtypeParser#agValue.
	VisitAgValue(ctx *AgValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#StringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#IntegerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#FloatValue.
	VisitFloatValue(ctx *FloatValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#TrueBoolean.
	VisitTrueBoolean(ctx *TrueBooleanContext) interface{}

	// Visit a parse tree produced by AgtypeParser#FalseBoolean.
	VisitFalseBoolean(ctx *FalseBooleanContext) interface{}

	// Visit a parse tree produced by AgtypeParser#NullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#ObjectValue.
	VisitObjectValue(ctx *ObjectValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#ArrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by AgtypeParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by AgtypeParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by AgtypeParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by AgtypeParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}
}
