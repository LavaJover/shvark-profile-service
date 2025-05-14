package main

import (
	"github.com/LavaJover/shvark-profile-service/internal/infrastructure/kafka"
	"github.com/LavaJover/shvark-profile-service/internal/usecase"
)

func main(){
	broker := "localhost:9092"
	kafka.StartConsumer(broker, "user.created", &usecase.ProfileUsecase{})
}