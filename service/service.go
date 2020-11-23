package service

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/carousell/gcp-self-study/model"
	pb "github.com/carousell/gcp-self-study/proto"
	"github.com/carousell/gcp-self-study/store"
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
		count = 0
	}
	if count > 5 {
		return &pb.LogInResponse{Token: ""}, errors.New("Error : Rate limitted please try again in some time")
	}

	//handle no user returned
	user, err := s.CheckifExistingUser(ctx, req.GetUsername(), "")
	if err != nil {
		return &pb.LogInResponse{Token: ""}, err
	}
	if user == nil {
		return &pb.LogInResponse{Token: ""}, errors.New("User does not exist")
	}
	token := ""
	// We can have more business logic on to check validity of user login here
	isActive, err := IsUserActive(user, req.GetPassword())
	if isActive {
		// Generate Tokens for Inside and 3rd Party
		token, err = s.Storage.StoreInHouseToken(ctx, randSeq(rand.Intn(100)), user.ID.String, "8h")
		if err != nil {
			return &pb.LogInResponse{Token: ""}, errors.New("Error : Cannot generate Inhouse token")
		}
		//store login activity
		err = s.Storage.StoreActivity(ctx, user.ID.String, "InHouse", true, "Login")
		if err != nil {
			log.Println("error whilr storing log activity")
		}
		return &pb.LogInResponse{Token: token}, nil
	} else {
		s.Storage.IncrementRedisRetryCounter(ctx, user.Username.String)
		s.Storage.StoreActivity(ctx, user.ID.String, "InHouse", false, "Login")

		return &pb.LogInResponse{Token: token}, err

	}
	return &pb.LogInResponse{Token: token}, nil
}

// Register function is used to create the user record , and generate a token .
func (s *Svc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp := &pb.RegisterResponse{}
	if req.GetApiKey() == "" || req.GetSecret() == "" {
		return &pb.RegisterResponse{Token: ""}, errors.New("Bad Request : Please provide Api/Secret")
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		return &pb.RegisterResponse{Token: ""}, errors.New("Error : Please fill in username and password")
	}

	existingUser, err := s.CheckifExistingUser(ctx, req.GetUsername(), req.GetEmail())

	if err != nil {
		return &pb.RegisterResponse{Token: ""}, errors.New("Existing User: Error Fetching existing User")
	}
	if existingUser != nil {
		return &pb.RegisterResponse{Token: ""}, errors.New("Existing User: The user is existing in the user")
	}

	modelUser := model.AccountsUser{}
	modelUser.Email.Scan(req.GetEmail())
	modelUser.Username.Scan(req.GetUsername())
	modelUser.Passowrd.Scan(RotEn(req.GetPassword()))
	modelUser.IsActive.Scan(true)
	user, err := s.Storage.CreateUser(ctx, &modelUser)
	if err != nil {
		return &pb.RegisterResponse{Token: ""}, errors.New("DB Error: Error Saving DB to the User")
	}
	token := ""
	token, err = s.Storage.StoreInHouseToken(ctx, randSeq(rand.Intn(100)), user.ID.String, "8h")
	if err != nil {
		return &pb.RegisterResponse{Token: ""}, errors.New("Token Error: Error generating token")
	}
	resp.Token = token
	resp.User = &pb.User{}
	resp.User.Username = modelUser.Username.String
	resp.User.Email = modelUser.Email.String
	err = s.Storage.StoreActivity(ctx, modelUser.ID.String, "InHouse", true, "Register")
	if err != nil {
		log.Println("error whilr storing log activity")
	}
	return resp, nil
}

//Logout removes the token and inserts activity
func (s *Svc) Logout(ctx context.Context, req *pb.LogoutRequest) (resp *pb.LogoutResponse, err error) {
	userID, err := s.Storage.GetInHouseToken(ctx, req.GetToken())
	if userID != req.GetToken() {
		return &pb.LogoutResponse{IsLoggedOut: false}, nil
	}
	err = s.Storage.RemoveInHouseToken(ctx, req.GetToken())
	if err != nil {
		return &pb.LogoutResponse{IsLoggedOut: false}, errors.New("Logout User : Failed to logout user ,deleting token failed")
	}
	err = s.Storage.StoreActivity(ctx, req.GetUserId(), "InHouse", true, "Logout")
	if err != nil {
		log.Println("error whilr storing log activity")
	}
	return &pb.LogoutResponse{IsLoggedOut: true}, nil
}

// CheckifExistingUser implements logic to check before registering
func (s *Svc) CheckifExistingUser(ctx context.Context, username string, email string) (user *model.AccountsUser, err error) {
	users, err := s.Storage.GetUsersByUsernameOrEmail(ctx, username, email)
	if err != nil {
		// DB error
		return nil, errors.New("Error : Error Fetching in User from DB")
	}
	if len(users) == 0 {
		return nil, nil
	}
	user = users[0]
	return user, nil
}

func (s *Svc) mustEmbedUnimplementedUserLoginServer() {
}
