package service

import (
	"context"
	"errors"
	"log"
	"math/rand"

	pb "github.com/carousell/chope-assignment/proto"
	"github.com/carousell/chope-assignment/store"
)

type Svc struct {
	Storage store.StorageService
	pb.UnimplementedUserLoginServer
}
type UserLoginServer interface {
	pb.UserLoginServer
}

//Login - Handles the following resp
//Validates the Input,Handles the ratelimit
//Decrypts the DB password value and checks if password match
//if yes , creates token . Else returns increments counter
func (s *Svc) Login(ctx context.Context, req *pb.LogInRequest) (resp *pb.LogInResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return &pb.LogInResponse{Token: ""}, errors.New("Error : Please fill in username and password")
	}
	count, err := s.Storage.GetLoginAttempts(ctx, req.GetUsername())
	if err != nil {
		// handle redis error by defaulting count to 0
		//	return &pb.LogInResponse{Token: ""}, errors.New("Error : Error Fetching in User rate limit info from Redis")
		count = 0
	}
	if count > 5 {
		return &pb.LogInResponse{Token: ""}, errors.New("Error : Rate limitted please try again in some time")
	}

	//jhandle no user returned
	users, err := s.Storage.GetAccountsUser(ctx, req.GetUsername())
	if err != nil {
		// DB error
		return &pb.LogInResponse{Token: ""}, errors.New("Error : Error Fetching in User from DB")
	}

	if len(users) == 0 {
		return &pb.LogInResponse{Token: ""}, errors.New("User Not Found : User not found , Please Sign up")
	}

	user := users[0]

	if IsUserActive(user, req.GetPassword()) {
		// Generate Tokens for Inside and 3rd Party
		token, err := s.Storage.StoreInHouseToken(ctx, randSeq(rand.Intn(100)), user.ID.String, "8h")
		if err != nil {
			return &pb.LogInResponse{Token: ""}, errors.New("Error : Cannot generate Inhouse token")
		}
		//store login activity
		err = s.Storage.StoreLoginActivity(ctx, user.ID.String, "InHouse", true)
		if err != nil {
			log.Println("error whilr storing log activity")
		}
		return &pb.LogInResponse{Token: token}, errors.New("Error : Rate limitted please try again in some time")
	} else {
		s.Storage.IncrementRedisRetryCounter(ctx, user.Username.String)
		err = s.Storage.StoreLoginActivity(ctx, user.ID.String, "InHouse", false)
		if err != nil {
			log.Println("error while storing log activity")
		}

	}

	//	err := s.Storage.CreateUser(ctx, &user)
	//	if err != nil {
	//	return &pb.LogInResponse{Token: ""}, errors.New("Error : Failed to save user to the database")
	//}
	return &pb.LogInResponse{Token: "NAE"}, nil
}

func (s *Svc) mustEmbedUnimplementedUserLoginServer() {
}
