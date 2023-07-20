// Code generated from calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // calc
import "github.com/antlr4-go/antlr/v4"

// BasecalcListener is a complete listener for a parse tree produced by calcParser.
type BasecalcListener struct{}

var _ calcListener = &BasecalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOp is called when production Op is entered.
func (s *BasecalcListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production Op is exited.
func (s *BasecalcListener) ExitOp(ctx *OpContext) {}

// EnterDigit is called when production Digit is entered.
func (s *BasecalcListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production Digit is exited.
func (s *BasecalcListener) ExitDigit(ctx *DigitContext) {}

// EnterParen is called when production Paren is entered.
func (s *BasecalcListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BasecalcListener) ExitParen(ctx *ParenContext) {}
