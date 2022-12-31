package usecase

import (
	"context"
	"logger/internal/domain"
	"logger/internal/repository"
)

// UserUseCase represent the user's usecase contract
type LogUseCase interface {
	InsertLog(ctx context.Context, login *domain.Log) error
}

type LogUseCaseImpl struct {
	LogRepo repository.LogRepository
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewLogUseCase(LogRepo repository.LogRepository) LogUseCase {
	return &LogUseCaseImpl{
		LogRepo: LogRepo,
	}
}

func (uc *LogUseCaseImpl) InsertLog(ctx context.Context, login *domain.Log) error {
	err := uc.LogRepo.Insert(ctx, login)
	return err
}
