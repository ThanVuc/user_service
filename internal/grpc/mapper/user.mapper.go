package mapper

import (
	"user_service/internal/grpc/database"
	"user_service/proto/user"
)

type userMapper struct{}

func (u *userMapper) ConvertDbUserPrifileRowToGrpcUser(userPr *database.GetUserProfileRow) *user.UserProfileItem {
	resp := &user.UserProfileItem{}
	if userPr == nil {
		return nil
	}

	userProfileData := userPr

	percentage := 0

	if userProfileData.Fullname.Valid && userProfileData.Fullname.String != "" {
		percentage += 25
	}
	if userProfileData.DateOfBirth.Valid {
		percentage += 15
	}
	if userProfileData.Gender.Valid {
		percentage += 10
	}
	if userProfileData.Bio.Valid && userProfileData.Bio.String != "" {
		percentage += 20
	}
	if userProfileData.Sentence.Valid && userProfileData.Sentence.String != "" {
		percentage += 15
	}
	if userProfileData.Author.Valid && userProfileData.Author.String != "" {
		percentage += 15
	}

	resp.Id = userProfileData.UserID.String()
	resp.Fullname = userProfileData.Fullname.String
	resp.Email = userProfileData.Email
	resp.AvatarUrl = userProfileData.AvatarUrl.String
	resp.Bio = userProfileData.Bio.String
	resp.Slug = userProfileData.Slug.String
	resp.Gender = userProfileData.Gender.Bool
	resp.Sentence = userProfileData.Sentence.String
	resp.Author = userProfileData.Author.String
	resp.ProfileCompletedPercentage = int32(percentage)
	if userProfileData.CreatedAt.Valid {
		timestamp := userProfileData.CreatedAt.Time.Unix()
		resp.CreatedAt = timestamp
	}

	if userProfileData.UpdatedAt.Valid {
		timestamp := userProfileData.UpdatedAt.Time.Unix()
		resp.UpdatedAt = timestamp
	}

	if userProfileData.DateOfBirth.Valid {
		timestamp := userProfileData.DateOfBirth.Time.Unix()
		resp.DateOfBirth = timestamp
	}

	return resp
}
