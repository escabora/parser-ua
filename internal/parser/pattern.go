package parser

import (
	"regexp"
	"strings"
)

func (p *Pattern) Match(ua string) bool {

	// Regex do pattern
	escaped := regexp.QuoteMeta(p.Pattern)
	regexPattern := strings.ReplaceAll(escaped, `\*`, ".*")
	regexPattern = "^" + regexPattern + "$"

	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return false
	}

	if !re.MatchString(ua) {
		return false
	}

	if p.Browser != "" && !strings.Contains(ua, p.Browser) {
		return false
	}

	if p.Version != "" && !strings.Contains(ua, p.Version) {
		return false
	}

	if p.Platform != "" && !strings.Contains(ua, p.Platform) {
		return false
	}

	return true
}

func EnrichResultFromUA(res *Result, ua string) {
	// Browser + Version
	if res.Browser == "" {
		switch {
		case strings.Contains(ua, "Chrome/"):
			res.Browser = "Chrome"
			re := regexp.MustCompile(`Chrome/([\d\.]+)`)
			m := re.FindStringSubmatch(ua)
			if len(m) >= 2 {
				res.Version = m[1]
			}
		case strings.Contains(ua, "Firefox/"):
			res.Browser = "Firefox"
			re := regexp.MustCompile(`Firefox/([\d\.]+)`)
			m := re.FindStringSubmatch(ua)
			if len(m) >= 2 {
				res.Version = m[1]
			}
		case strings.Contains(ua, "Safari/") && strings.Contains(ua, "Version/"):
			res.Browser = "Safari"
			re := regexp.MustCompile(`Version/([\d\.]+)`)
			m := re.FindStringSubmatch(ua)
			if len(m) >= 2 {
				res.Version = m[1]
			}
		case strings.Contains(ua, "Edg/"):
			res.Browser = "Edge"
			re := regexp.MustCompile(`Edg/([\d\.]+)`)
			m := re.FindStringSubmatch(ua)
			if len(m) >= 2 {
				res.Version = m[1]
			}
		}
	}

	// Platform
	if res.Platform == "" {
		switch {
		case strings.Contains(ua, "Mac OS X"):
			res.Platform = "MacOS"
		case strings.Contains(ua, "Windows NT"):
			res.Platform = "Windows"
		case strings.Contains(ua, "Linux"):
			res.Platform = "Linux"
		case strings.Contains(ua, "Android"):
			res.Platform = "Android"
		case strings.Contains(ua, "iPhone") || strings.Contains(ua, "iPad"):
			res.Platform = "iOS"
		}
	}

	// DeviceType
	if res.DeviceType == "" {
		switch {
		case strings.Contains(ua, "Mobile") || strings.Contains(ua, "Android") || strings.Contains(ua, "iPhone"):
			res.DeviceType = "Mobile"
		case strings.Contains(ua, "iPad"):
			res.DeviceType = "Tablet"
		default:
			res.DeviceType = "Desktop"
		}
	}
}
