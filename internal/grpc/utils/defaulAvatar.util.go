package utils

const (
	DefaultMaleAvatar   = "https://pub-49aeefce5382467589ef388921404348.r2.dev/default/ea8a5596-8088-403e-9050-e30826a5a1db.webp"
	DefaultFemaleAvatar = "https://pub-49aeefce5382467589ef388921404348.r2.dev/default/79c38322-cf71-4f60-bbd5-ca6861bffcb7.webp"
)

func GetDefaultAvatarUrl(isMale bool) (string, error) {
	if isMale {
		return DefaultMaleAvatar, nil
	}
	return DefaultFemaleAvatar, nil
}
