package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/adoringonion/golang_monkey/evaluator"
	"github.com/adoringonion/golang_monkey/lexer"
	"github.com/adoringonion/golang_monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
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
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const yaruo = `
            ＿＿＿_
     	／⌒       ⌒ ＼
　　　／  （●） 　（●）  ＼
　 ／  ::⌒ (__人__) ⌒::     ＼ 
　 |　　　　　|r┬-|　　　　　|
　 ＼ 　　 　  ー' 　 　   ／
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, yaruo)
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
