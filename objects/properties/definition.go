package properties

import "errors"

type DefinitionType uint

var ErrDefinitionType = errors.New("the DefinitionType is neither tlp nor statement")

type TlpDefinition struct {
	Tlp string `json:"tlp,omitempty"`
}

type StatementDefinition struct {
	Statement string `json:"statement,omitempty"`
}

type DefinitionProperty struct {
	Definition interface{} `json:"definition"`
}

func (d *DefinitionProperty) SetDefinition(t, s string) (err error) {
	if t == "tlp" {
		d.Definition = TlpDefinition{Tlp: s}
		return nil
	}
	if t == "statement" {
		d.Definition = StatementDefinition{
			Statement: s,
		}
		return nil
	}
	return ErrDefinitionType
}
