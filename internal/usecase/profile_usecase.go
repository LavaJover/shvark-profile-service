package usecase

import (
	"context"
	"fmt"

	"github.com/LavaJover/shvark-profile-service/internal/domain"
)

type ProfileUsecase struct {
	repo domain.ProfileRepository
}

func NewProfileUsecase(repo domain.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{repo: repo}
}

func (u *ProfileUsecase) CreateProfile(ctx context.Context, username string) error {
	profile := &domain.Profile{
		Username: username,
	}

	fmt.Println(profile)

	return nil
}