// Code generated from calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // calc
import "github.com/antlr4-go/antlr/v4"

// calcListener is a complete listener for a parse tree produced by calcParser.
type calcListener interface {
	antlr.ParseTreeListener

	// EnterOp is called when entering the Op production.
	EnterOp(c *OpContext)

	// EnterDigit is called when entering the Digit production.
	EnterDigit(c *DigitContext)

	// EnterParen is called when entering the Paren production.
	EnterParen(c *ParenContext)

	// ExitOp is called when exiting the Op production.
	ExitOp(c *OpContext)

	// ExitDigit is called when exiting the Digit production.
	ExitDigit(c *DigitContext)

	// ExitParen is called when exiting the Paren production.
	ExitParen(c *ParenContext)
}
