package parser

import (
	"path"
)

func (p *Pattern) Match(ua string) bool {
	match, err := path.Match(p.Pattern, ua)
	return err == nil && match
}
