package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prasanth-pn/go-grpc-auth-svc/pkg/models"
	"github.com/prasanth-pn/go-grpc-auth-svc/pkg/pb"
)

type Server struct {
	//H db.Handler
	//utils utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	fmt.Println(req.Email, req.Password, user)

	// if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
	// 	return &pb.RegisterResponse{
	// 		Status: http.StatusConflict,
	// 		Error:  "E-Mail already exists",
	// 	}, nil
	// }

	// user.Email = req.Email
	// user.Password = utils.HashPassword(req.Password)
	// s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

// func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
//     var user models.User
// fmt.Println(req.Email)
//     if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
//         return &pb.LoginResponse{
//             Status: http.StatusNotFound,
//             Error:  "User not found",
//         }, nil
//     }
// fmt.Println(req.Password,user.Password)
//     match := utils.CheckPasswordHash(req.Password, user.Password)

//     if !match {
//         return &pb.LoginResponse{
//             Status: http.StatusNotFound,
//             Error:  "User not found",
//         }, nil
//     }

//     token, _ := s.Jwt.GenerateToken(user)

//     return &pb.LoginResponse{
//         Status: http.StatusOK,
//         Token:  token,
//     }, nil
// }

// func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
//     claims, err := s.Jwt.ValidateToken(req.Token)

//     if err != nil {
//         return &pb.ValidateResponse{
//             Status: http.StatusBadRequest,
//             Error:  err.Error(),
//         }, nil
//     }

//     var user models.User

//     if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
//         return &pb.ValidateResponse{
//             Status: http.StatusNotFound,
//             Error:  "User not found",
//         }, nil
//     }

//     return &pb.ValidateResponse{
//         Status: http.StatusOK,
//         UserId: user.Id,
//     }, nil
// }
