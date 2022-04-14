package configs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	DB *mongo.Database
}

func InitDatabase(logger Logger, env Env) Database {

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(env.DBHost))
	if err != nil {
		logger.Zap.Info("host: ", env.DBHost)
		logger.Zap.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		logger.Zap.Info("host: ", env.DBHost)
		logger.Zap.Panic(err)
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		logger.Zap.Info("error ping mongo server: ", env.DBHost)
		logger.Zap.Panic(err)
	}

	logger.Zap.Info("✅ Database connection established")

	return Database{
		DB: mongoClient.Database(env.DBName),
	}
}
