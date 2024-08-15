package main

import (
	"context"
	"github.com/devyk100/gengou-db/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	dsn := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	//fmt.Println(dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close(context.Background())

	queries := database.New(conn)

	//user, err := queries.InsertUser(ctx, database.InsertUserParams{
	//	Name:            "Yash",
	//	UserID:          "yashk095",
	//	EmailID:         "yash@dev.in",
	//	Phone:           "172873819247",
	//	PastExperiences: pgtype.Text{},
	//	UserType:        database.UserTypeLearner,
	//})
	//if err != nil {
	//	log.Fatal("failed to insert user", err.Error())
	//}
	//
	//fmt.Println(user.Name)
	err = queries.DeleteUser(ctx, "yashk095")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//Users, err := queries.GetUsers(ctx)
	//for _, user := range Users {
	//	fmt.Println(user.Name, user.UserID, user.EmailID)
	//}
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//User, err := queries.GetAUser(context.Background(), 1)
	//fmt.Println(User.UserID, err)
	//fmt.Println("Exiting")
	//var now time.Time
	//err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	//if err != nil {
	//	log.Fatal("failed to execute query", err)
	//}

	//fmt.Println(now)

}
