// Code generated from calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // calc
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by calcParser.
type calcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by calcParser#Op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by calcParser#Digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by calcParser#Paren.
	VisitParen(ctx *ParenContext) interface{}
}
