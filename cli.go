package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	pb "github.com/muhammadhidayah/users-service/proto/users"
)

func main() {
	srv := micro.NewService(
		micro.Name("users.cli"),
		micro.Version("lates"),
	)

	srv.Init()

	client := pb.NewUsersService("users.service", microclient.DefaultClient)

	// name := "Muhammad Hidayah"
	// email := "muhammad30hidayahskom696@gmail.com"
	// personID := email
	// password := "12345678"
	// isActive := 1
	// id := "75199578-373b-4a07-a0e3-bb879181bfa0"

	// this is for create user
	// fmt.Println("===================== Create User ========================")
	// r, err := client.CreateUser(context.TODO(), &pb.User{
	// 	PersonName:     name,
	// 	PersonId:       personID,
	// 	PersonPassword: password,
	// 	IsActive:       int32(isActive),
	// })

	// if err != nil {
	// 	log.Printf("Could not create user: %v", err)
	// }

	// fmt.Printf("Created: %t", r.Created)
	// fmt.Println()

	// this is for update user
	// fmt.Println("===================== Update User ========================")
	// update, err := client.UpdateUser(context.TODO(), &pb.User{
	// 	PersonName:     "Muhammad Hidayah, S.KOM",
	// 	PersonPassword: "1234543321",
	// 	Id:             id,
	// })

	// if err != nil {
	// 	log.Printf("Could not update user: %v", err)
	// }

	// fmt.Printf("Updated: %t", update.Updated)
	// fmt.Println()

	// this is for delete user
	// fmt.Println("===================== Delete User ========================")
	// del, err := client.DeleteUser(context.TODO(), &pb.User{
	// 	Id: id,
	// })
	// if err != nil {
	// 	log.Printf("Could not delete user: %v", err)
	// }

	// fmt.Printf("Deleted: %t", del.Deleted)
	// fmt.Println()

	// this is for auth
	fmt.Println("===================== Auth User ========================")
	auth, err := client.GetUserByPersonIDAndPassword(context.TODO(), &pb.User{
		PersonId:       "muhammad30hidayah696@gmail.com",
		PersonPassword: "ssadsadsd23sdfdsf43scdfrfe",
	})
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}

	fmt.Println(auth.User)

}
