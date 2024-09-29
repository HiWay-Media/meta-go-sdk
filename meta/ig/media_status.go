package ig

type MediaStatus string

const (
    EXPIRED     MediaStatus 		= "EXPIRED"
    ERROR       MediaStatus 	    = "ERROR"
	FINISHED    MediaStatus 	    = "FINISHED"
    IN_PROGRESS MediaStatus 	    = "IN_PROGRESS"
	PUBLISHED   MediaStatus 	    = "PUBLISHED"
)

// CheckMediaStatus checks if the given string matches any MediaStatus
func CheckMediaStatus(status string) bool {
	switch MediaStatus(status) {
	case EXPIRED, ERROR, FINISHED, IN_PROGRESS, PUBLISHED:
		return true
	default:
		return false
	}
}
