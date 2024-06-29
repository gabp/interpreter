package repl

import (
	"fmt"
	"monkey/ast"

	"github.com/thediveo/go-asciitree"
)

type printableNode struct {
	Label    string          `asciitree:"label"`
	Children []printableNode `asciitree:"children"`
}

func PrintAST(node ast.Node) {

	p := printableNode{
		Label:    "AST",
		Children: toPrintableNodes(node.Children()),
	}

	fmt.Println(asciitree.RenderFancy(p))
}

func toPrintableNodes(nodes []ast.Node) []printableNode {
	n := make([]printableNode, len(nodes))
	for i, j := range nodes {
		n[i] = printableNode{
			Label:    j.TokenLiteral() + fmt.Sprintf(":  %T", j),
			Children: toPrintableNodes(j.Children()),
		}
	}

	return n
}
