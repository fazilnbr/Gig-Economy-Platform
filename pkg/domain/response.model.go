package domain

import "github.com/golang-jwt/jwt/v4"

type AdminResponse struct {
	ID           int    `json:"id_login"`
	Username     string `json:"email"`
	Password     string `json:"password,omitempty"`
	Role         int    `json:"role"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	UserName     string `json:"first_name"`
	Password     string `json:"password"`
	Verification bool   `json:"verification"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type WorkerResponse struct {
	ID           int    `json:"id"`
	UserName     string `json:"first_name"`
	Password     string `json:"password"`
	Verification bool   `json:"verification"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type SignedDetails struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	Source   string `json:"source"`
	jwt.StandardClaims
}

type ChangePassword struct {
	Email       string `json:"email" binding:"required"`
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}

type WorkerJob struct {
	Id          int
	JobTitile   string
	Wage        int64
	Description string
}

type ListJobsWithWorker struct {
	IdJob        int    `gorm:"primaryKey;autoIncrement:true;unique"`
	CategoryName string `json:"categoryname"`
	WorkerName   string `json:"workername"`
}

type ListFavorite struct {
	FavoriteId  int
	Name        string
	Photo       string
	JobCategory string
	Wage        int
	Description string
}
