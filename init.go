package main

import (
	"api-point-of-sales/config"
	"api-point-of-sales/handler/router"
	"api-point-of-sales/util"

	"github.com/muhammadluth/log"

	"api-point-of-sales/handler/middleware"

	authentication_mapper "api-point-of-sales/handler/app/authentication/mapper"
	authentication_repo "api-point-of-sales/handler/app/authentication/repository"
	authentication_usecase "api-point-of-sales/handler/app/authentication/usecase"

	user_management_mapper "api-point-of-sales/handler/app/user_management/mapper"
	user_management_repo "api-point-of-sales/handler/app/user_management/repository"
	user_management_usecase "api-point-of-sales/handler/app/user_management/usecase"
)

func RunningApplication() {
	properties := config.LoadConfig()

	log.SetupLogging(properties.LogPath)
	database := config.ConnectDatabase(properties.DBHost, properties.DBPort, properties.DBUser,
		properties.DBPassword, properties.DBName)
	expireAccessToken := config.ParseTimeDuration(properties.ExpireAccessToken)
	expireRefreshToken := config.ParseTimeDuration(properties.ExpireRefreshToken)
	privateKey, publicKey := config.LoadCredential(properties.PrivateKey, properties.PublicKey)

	// AUTHENTICATION
	iAuthenticationMapper := authentication_mapper.NewAuthenticationMapper()
	iAuthenticationRepo := authentication_repo.NewAuthenticationRepo(database)
	iAuthenticationReferenceRepo := authentication_repo.NewReferenceRepo(database)
	iAuthenticationCredentialUsecase := authentication_usecase.NewCredentialUsecase(privateKey,
		publicKey)
	iAuthenticationValidationUsecase := authentication_usecase.NewValidationUsecase(
		iAuthenticationReferenceRepo)
	iTokenUsecase := authentication_usecase.NewTokenUsecase(expireAccessToken, expireRefreshToken,
		privateKey, publicKey, iAuthenticationMapper)
	iRegisterUsecase := authentication_usecase.NewRegisterUsecase(iAuthenticationMapper,
		iAuthenticationRepo, iAuthenticationCredentialUsecase, iAuthenticationValidationUsecase)
	iLoginUsecase := authentication_usecase.NewLoginUsecase(iAuthenticationMapper,
		iAuthenticationRepo, iAuthenticationCredentialUsecase, iTokenUsecase)
	iForgetPasswordUsecase := authentication_usecase.NewForgetPasswordUsecase(iAuthenticationMapper,
		iAuthenticationRepo, iAuthenticationCredentialUsecase)

	// USER MANAGEMENT
	iUserManagementMapper := user_management_mapper.NewUserManagementMapper()
	iUserManagementReferenceMapper := user_management_mapper.NewReferenceMapper()
	iUserManagementRepo := user_management_repo.NewUserManagementRepo(database)
	iUserManagementReferenceRepo := user_management_repo.NewReferenceRepo(database)
	iUserManagementCredentialUsecase := user_management_usecase.NewCredentialUsecase(privateKey,
		publicKey)
	iUserManagementValidationUsecase := user_management_usecase.NewValidationUsecase(
		iUserManagementReferenceMapper, iUserManagementReferenceRepo)
	iRoleUsecase := user_management_usecase.NewRoleUsecase(iUserManagementMapper,
		iUserManagementRepo)
	iUserUsecase := user_management_usecase.NewUserUsecase(iUserManagementMapper,
		iUserManagementRepo, iUserManagementCredentialUsecase, iUserManagementValidationUsecase)

	// MIDDLEWARE
	iMiddleWare := middleware.NewMiddleware(iTokenUsecase)
	iSetupRouter := router.NewSetupRouter(properties, iMiddleWare, iRoleUsecase,
		iUserUsecase, iRegisterUsecase, iLoginUsecase, iForgetPasswordUsecase)

	util.RunningParallel(
		iSetupRouter.Router,
	)
}
