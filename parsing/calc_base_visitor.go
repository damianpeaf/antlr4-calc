// Code generated from calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // calc
import "github.com/antlr4-go/antlr/v4"

type BasecalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecalcVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalcVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalcVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}
