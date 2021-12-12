package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	id := primitive.NewObjectID()

	fmt.Println(id.Hex())
	fmt.Println(id.String())
	fmt.Println(id.Timestamp())
}
