package main

import (
	"context"
	"fmt"
	"log"

	"fibo/api/cli"
	"fibo/api/http"
	authImpl "fibo/internal/auth/impl"
	cryptoImpl "fibo/internal/base/crypto/impl"
	databaseImpl "fibo/internal/base/database/impl"
	postImpl "fibo/internal/post/impl"
	userImpl "fibo/internal/user/impl"
)

func main() {
	ctx := context.Background()
	parser := cli.NewParser()
	//
	conf, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := databaseImpl.NewClient(ctx, conf.Database())

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)

	defer dbClient.Close()

	crypto := cryptoImpl.NewCrypto()
	dbService := databaseImpl.NewService(dbClient)

	userRepositoryOpts := userImpl.UserRepositoryOpts{
		ConnManager: dbService,
	}
	userRepository := userImpl.NewUserRepository(userRepositoryOpts)

	authServiceOpts := authImpl.AuthServiceOpts{
		Crypto:         crypto,
		Config:         conf.Auth(),
		UserRepository: userRepository,
	}
	authService := authImpl.NewAuthService(authServiceOpts)

	userUsecasesOpts := userImpl.UserUsecasesOpts{
		TxManager:      dbService,
		UserRepository: userRepository,
		Crypto:         crypto,
	}
	userUsecases := userImpl.NewUserUsecases(userUsecasesOpts)

	postRepositoryOpts := postImpl.PostRepositoryOpts{
		ConnManager: dbService,
	}

	postRepository := postImpl.NewPostRepository(postRepositoryOpts)

	postUsecasesOpts := postImpl.PostUsecaseOpts{
		PostRepository: postRepository,
		TxManager:      dbService,
	}

	postUsecases := postImpl.NewPostUsecase(postUsecasesOpts)

	serverOpts := http.ServerOpts{
		UserUsecases: userUsecases,
		AuthService:  authService,
		Crypto:       crypto,
		Config:       conf.HTTP(),
		Post:         postUsecases,
	}
	server := http.NewServer(serverOpts)

	log.Fatal(server.Listen())
}
