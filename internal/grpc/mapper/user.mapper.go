package mapper

import (
	"user_service/internal/grpc/database"
	"user_service/proto/user"
)

type userMapper struct{}

func (u *userMapper) ConvertDbUserPrifileRowToGrpcUser(userPr *[]database.GetUserProfileRow) *user.UserProfileItem {
	resp := &user.UserProfileItem{}
	if userPr == nil || len(*userPr) == 0 {
		return nil
	}

	userProfileData := (*userPr)[0]
	resp.Id = userProfileData.UserID.String()
	resp.Fullname = userProfileData.Fullname.String
	resp.Email = userProfileData.Email
	resp.AvatarUrl = userProfileData.AvatarUrl.String
	resp.Bio = userProfileData.Bio.String
	resp.Slug = userProfileData.Slug.String
	resp.Gender = userProfileData.Gender.Bool
	resp.Sentence = userProfileData.Sentence.String
	resp.Author = userProfileData.Author.String

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