package parser

import (
	"encoding/csv"
	"fmt"
	"io"
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
	r.FieldsPerRecord = -1
	_, err = r.Read()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("CSV empty or malformed")
		}
		return nil, err
	}

	// Second line
	_, err = r.Read()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("CSV empty after reading headers")
		}
		return nil, err
	}

	// Third line is the headers
	headers, err := r.Read()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("CSV empty afeter reading headers")
		}
		return nil, err
	}

	patterns := make(map[string][]*Pattern)
	lineNumber := 3 // Already read 2 lines before this, so start at 3

	for {
		lineNumber++

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Erro na leitura CSV na linha %d: %v\n", lineNumber, err)
			continue
		}

		if len(record) != len(headers) {
			fmt.Printf("Linha %d com número errado de campos. Esperado: %d, achado: %d\n", lineNumber, len(headers), len(record))
			fmt.Println("Conteúdo da linha:", record)
			continue // skip this record
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
	if res := matchUA(ua, p.patterns[prefix]); res != nil {
		p.cache.Add(ua, res)
		return res
	}

	for _, pats := range p.patterns {
		if res := matchUA(ua, pats); res != nil {
			p.cache.Add(ua, res)
			return res
		}
	}

	return nil
}
