package ig

type MediaType string

const (
    CAROUSEL    MediaType 		= "CAROUSEL"
    REELS       MediaType 	    = "REELS"
	STORIES     MediaType 	    = "STORIES"
)

// CheckMediaType checks if the given string matches any MediaType
func CheckMediaType(mType string) bool {
	switch MediaType(mType) {
	case CAROUSEL, REELS, STORIES:
		return true
	default:
		return false
	}
}
