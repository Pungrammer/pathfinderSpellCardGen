package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func filter(filterExp string, spells []*Spell) ([]*Spell, error) {
	comp, err := expr.Compile(filterExp, expr.Env(&Spell{}))
	if err != nil {
		return nil, fmt.Errorf("compile filter: %s", err)
	}

	filtered := make([]*Spell, 0)
	for _, spell := range spells {
		run, err := expr.Run(comp, spell)
		if err != nil {
			return nil, fmt.Errorf("run filter: %s", err)
		}
		fits, ok := run.(bool)
		if !ok {
			return nil, fmt.Errorf("filter must be a boolean expression")
		}

		if fits {
			filtered = append(filtered, spell)
		}
	}

	return filtered, nil
}
