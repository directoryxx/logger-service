package repository

import (
	"context"
	"logger/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository represent the user's repository contract
type LogRepository interface {
	Insert(ctx context.Context, log *domain.Log) error
}

type LogRepositoryImpl struct {
	DB *mongo.Client
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewLogRepository(db *mongo.Client) LogRepository {
	return &LogRepositoryImpl{
		DB: db,
	}
}

func (logRepo *LogRepositoryImpl) Insert(ctx context.Context, log *domain.Log) (err error) {
	coll := logRepo.DB.Database("logs").Collection("access_log")
	_, err = coll.InsertOne(ctx, log)
	if err != nil {
		return err
	}

	return nil
}
