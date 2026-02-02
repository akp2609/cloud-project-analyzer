package analysis

func SeverityFromDeviation(pct float64) string {
	switch {
	case pct >= 100:
		return "HIGH"
	case pct >= 50:
		return "MEDIUM"
	default:
		return "LOW"
	}
}
