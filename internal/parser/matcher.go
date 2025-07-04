package parser

func getPrefix(ua string) string {
	if len(ua) >= 6 {
		return ua[:6]
	}
	return ua
}

func matchUA(ua string, patterns []*Pattern) *Result {
	for _, pat := range patterns {
		if pat.Match(ua) {
			return &Result{
				Browser:    pat.Browser,
				Version:    pat.Version,
				Platform:   pat.Platform,
				DeviceType: pat.DeviceType,
				Matched:    pat.Pattern,
			}
		}
	}
	return nil
}
