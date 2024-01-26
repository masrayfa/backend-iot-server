package helper

func convertBoolToInt(b bool) int64 {
	if b {
		return 1
	}

	return 0
}