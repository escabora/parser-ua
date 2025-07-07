package parser

func getPrefix(ua string) string {
	for i, c := range ua {
		if c == ' ' {
			return ua[:i]
		}
	}
	return ua
}

func matchUA(ua string, patterns []*Pattern) *Result {
	for _, pat := range patterns {
		if pat.Match(ua) {
			r := &Result{
				Browser:    pat.Browser,
				Version:    pat.Version,
				Platform:   pat.Platform,
				DeviceType: pat.DeviceType,
				Matched:    pat.Pattern,
			}
			EnrichResultFromUA(r, ua)
			return r

		}
	}

	return nil
}
