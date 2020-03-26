package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rockspore/monkey-interpreter/evaluator"
	"github.com/rockspore/monkey-interpreter/lexer"
	"github.com/rockspore/monkey-interpreter/parser"
)

// Start : entry point of the REPL
func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
