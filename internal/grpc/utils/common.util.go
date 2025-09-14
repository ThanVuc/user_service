package utils

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/text/unicode/norm"
)

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ToUUID(id string) (pgtype.UUID, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}, nil
}

func RoundToTwoDecimal(val float64) float64 {
	return math.Round(val*100) / 100
}

func FromPgTypeTimeToUnix(t pgtype.Timestamp) *int64 {
	if !t.Valid {
		return nil
	}
	unixTime := t.Time.Unix()
	return &unixTime
}

func Difference[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(b))
	for _, item := range b {
		m[item] = struct{}{}
	}

	var diff []T
	for _, item := range a {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}

func ToBoolPointer(b bool) *bool {
	return &b
}

func ToStringPointer(s string) *string {
	return &s
}

func removeDiacritics(s string) string {
	t := norm.NFD.String(s)
	var result []rune
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		switch r {
		case 'đ':
			r = 'd'
		case 'Đ':
			r = 'D'
		}
		result = append(result, r)
	}
	return string(result)
}

func MakeSlug(fullname string, createdAt time.Time) string {
	s := strings.ToLower(fullname)
	s = removeDiacritics(s)
	s = strings.ReplaceAll(s, " ", "-")
	re := regexp.MustCompile(`[^a-z0-9-]+`)
	s = re.ReplaceAllString(s, "")
	s = strings.Trim(s, "-")

	return fmt.Sprintf("%s-%d", s, createdAt.UnixMilli())
}
