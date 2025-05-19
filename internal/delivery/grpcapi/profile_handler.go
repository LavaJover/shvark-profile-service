package grpcapi

import (
	"context"

	"github.com/LavaJover/shvark-profile-service/internal/domain"
	profilepb "github.com/LavaJover/shvark-profile-service/proto/gen"
)

type ProfileHandler struct {
	profilepb.UnimplementedProfileServiceServer
	Uc domain.ProfileUsecase
}

func (h *ProfileHandler) CreateProfile(ctx context.Context, r *profilepb.CreateProfileRequest) (*profilepb.CreateProfileResponse, error) {
	profile := domain.Profile{
		UserID: r.UserId,
		AvatarURL: r.AvatarUrl,
		TgLink: r.AvatarUrl,
	}
	profileID, err := h.Uc.CreateProfile(&profile)
	if err != nil {
		return nil, err
	}

	return &profilepb.CreateProfileResponse{
		ProfileId: profileID,
	}, nil
}

func (h *ProfileHandler) UpdateProfile(ctx context.Context, r *profilepb.UpdateProfileRequest) (*profilepb.UpdateProfileResponse, error) {
	profileID := r.ProfileId
	profile := domain.Profile{
		ID: r.Profile.ProfileId,
		UserID: r.Profile.UserId,
		AvatarURL: r.Profile.AvatarUrl,
		TgLink: r.Profile.TgLink,
	}
	_, err := h.Uc.UpdateProfile(profileID, &profile, r.UpdateMask.Paths)
	if err != nil {
		return nil, err
	}

	return &profilepb.UpdateProfileResponse{}, nil
}

func (h *ProfileHandler) DeleteProfile(ctx context.Context, r *profilepb.DeleteProfileRequest) (*profilepb.DeleteProfileResponse, error) {
	profileID := r.ProfileId
	_, err := h.Uc.DeleteProfile(profileID)
	if err != nil {
		return nil, err
	}

	return &profilepb.DeleteProfileResponse{}, nil
}

func (h *ProfileHandler) GetProfileByID(ctx context.Context, r *profilepb.GetProfileByIDRequest) (*profilepb.GetProfileByIDResponse, error) {
	profileID := r.ProfileId
	respProfile, err := h.Uc.GetProfileByID(profileID)
	if err != nil {
		return nil, err
	}

	return &profilepb.GetProfileByIDResponse{
		ProfileId: respProfile.ID,
		UserId: respProfile.UserID,
		AvatarUrl: respProfile.AvatarURL,
		TgLink: respProfile.TgLink,
	}, nil
}

func (h *ProfileHandler) GetProfileByUserID(ctx context.Context, r *profilepb.GetProfileByUserIDRequest) (*profilepb.GetProfileByUserIDResponse, error) {
	userID := r.UserId
	respProfile, err := h.Uc.GetProfileByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &profilepb.GetProfileByUserIDResponse{
		ProfileId: respProfile.ID,
		UserId: respProfile.UserID,
		AvatarUrl: respProfile.AvatarURL,
		TgLink: respProfile.TgLink,
	}, nil
}