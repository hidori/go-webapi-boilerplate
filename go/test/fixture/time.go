package fixture

import "time"

// Now は、現在の日時のフィクスチャ（2000-01-01T00:00:00+00:00）を返します。
func Now() time.Time {
	return time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
}

// Yesterday は、昨日の日時のフィクスチャ（1999-12-31T00:00:00+00:00）を返します。
func Yesterday() time.Time {
	return Now().AddDate(0, 0, -1)
}

// Tomorrow は、明日の日時のフィクスチャ（2000-01-02T00:00:00+00:00）を返します。
func Tomorrow() time.Time {
	return Now().AddDate(0, 0, 1)
}
