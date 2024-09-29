package adapter

// PrintBanner is struct
type PrintBanner struct {
	*banner
}

// NewPrintBanner func for initializing PrintBanner
func NewPrintBanner(str string) *PrintBanner {
	return &PrintBanner{
		banner: &banner{str: str},
	}
}

// PrintWeak func for formatting with paren
func (p *PrintBanner) PrintWeak() {
	p.showWithParen()
}

// PrintString func for formatting with aster
func (p *PrintBanner) PrintString() {
	p.showWithAster()
}
