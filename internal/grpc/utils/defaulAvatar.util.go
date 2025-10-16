package utils


const (
	DefaultMaleAvatar = "https://pub-49aeefce5382467589ef388921404348.r2.dev/default/ea8a5596-8088-403e-9050-e30826a5a1db.webp"
	DefaultFemaleAvatar = "https://pub-49aeefce5382467589ef388921404348.r2.dev/default/0cdaa125-f05c-4098-bb4d-d9ce183a6fbf.webp"
)

func GetDefaultAvatarUrl(isMale bool) (string, error) {
	if isMale {
		return DefaultMaleAvatar, nil
	}
	return DefaultFemaleAvatar, nil
}