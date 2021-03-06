package handlers

import (
	"context"
	"ktrain/cmd/repository"
	"ktrain/pkg/config"
	"ktrain/pkg/storage"
	"ktrain/proto/pb"
	"log"
	"testing"

	"github.com/spf13/viper"
)

func bindConfig() {
	err := config.BindDefault("./config.yaml")
	if err != nil {
		log.Fatalf("Error when binding config, err: %v", err)
		return
	}
}
func initTest() *storage.MongoDBManager {
	bindConfig()
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("mongodb.timeout"))
	defer cancel()
	mongDB, err := storage.NewMongoDBManager(ctx)
	if err != nil {
		log.Fatalf("Error when connecting database, err: %v", err)
		return nil
	}
	return mongDB
}
func TestActivityLogHandler_CreateAction(t *testing.T) {
	mongDB := initTest()
	activityLogRepository := repository.NewActivityLogRepository(mongDB)
	h, err := NewActivityLogHandler(activityLogRepository)
	if err != nil {
		log.Fatalf("Error when creating new user handler, err: %v", err)
		return
	}
	_, err = h.CreateAction(context.Background(), &pb.CreateActionRequest{
		Id:  1,
		Log: "Create",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}
func TestActivityLogHandler_GetallLogAction(t *testing.T) {
	mongDB := initTest()
	activityLogRepository := repository.NewActivityLogRepository(mongDB)
	h, err := NewActivityLogHandler(activityLogRepository)
	if err != nil {
		log.Fatalf("Error when creating new user handler, err: %v", err)
		return
	}

	_, err = h.GetAllLogAction(context.Background(), &pb.GetLogActionRequest{
		Id: 2,
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}
