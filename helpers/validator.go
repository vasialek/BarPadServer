package helpers

// IsValidUniqueID returns true for valid table IDs
func IsValidUniqueID(uid string) bool {
	if len(uid) == 32 {
		return true
	}
	return false
}
