package parser

import (
	"github.com/yassinebenaid/ryuko/ast"
	"github.com/yassinebenaid/ryuko/token"
)

func (p *parser) parseAssignement() ast.ParameterAssignement {
	var assigns ast.ParameterAssignement

	for {
		if !(p.curr.Type == token.WORD && p.next.Type == token.ASSIGN) {
			break
		}
		assignment := ast.Assignement{Name: p.curr.Literal}
		p.proceed()
		p.proceed()
		assignment.Value = p.parseExpression()
		assigns = append(assigns, assignment)

		if p.curr.Type == token.BLANK {
			p.proceed()
		}
	}

	return assigns
}
