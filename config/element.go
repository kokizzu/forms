package config

type Element struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Label      string                 `json:"label"`
	LabelCols  int                    `json:"labelCols,omitempty"`
	FieldCols  int                    `json:"fieldCols,omitempty"`
	Value      string                 `json:"value"`
	HelpText   string                 `json:"helpText"`
	Template   string                 `json:"template"`
	Valid      string                 `json:"valid"`
	Attributes [][]string             `json:"attributes"`
	Choices    []*Choice              `json:"choices"`
	Elements   []*Element             `json:"elements"`
	Format     string                 `json:"format"`
	Languages  []*Language            `json:"languages,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
}

func (e *Element) Clone() *Element {
	r := *e
	return &r
}
