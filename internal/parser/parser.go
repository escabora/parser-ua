package parser

import (
	"encoding/csv"
	"os"
)

type Parser struct {
	patterns map[string][]*Pattern
	cache    *UACache
}

func NewParser(csvPath string) (*Parser, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	headers, err := r.Read()
	if err != nil {
		return nil, err
	}

	patterns := make(map[string][]*Pattern)
	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		fieldMap := map[string]string{}
		for i, h := range headers {
			fieldMap[h] = record[i]
		}

		patternStr := record[0]
		prefix := getPrefix(patternStr)

		pat := &Pattern{
			Pattern:    patternStr,
			Browser:    fieldMap["Browser"],
			Version:    fieldMap["Version"],
			Platform:   fieldMap["Platform"],
			DeviceType: fieldMap["Device_Type"],
			RawFields:  fieldMap,
		}

		patterns[prefix] = append(patterns[prefix], pat)
	}

	return &Parser{
		patterns: patterns,
		cache:    NewUACache(10000),
	}, nil
}

func (p *Parser) Parse(ua string) *Result {
	if res, ok := p.cache.Get(ua); ok {
		return res
	}

	prefix := getPrefix(ua)
	res := matchUA(ua, p.patterns[prefix])
	if res != nil {
		p.cache.Add(ua, res)
	}
	return res
}
