package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	// Import the generated Person struct
	pb "github.com/jaygaha/go-beginner/cmd/22_microservices/22_8_protocol_buffers/proto"
)

func main() {
	person := &pb.Person{
		Name: "Jay",
		Age:  30,
		SocialFollowers: &pb.SocialFollowers{
			Youtube: 999,
			Twitter: 9999,
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)
	newPerson := &pb.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(newPerson.GetName())
	fmt.Println(newPerson.GetAge())
	fmt.Println(newPerson.GetSocialFollowers().GetYoutube())
	fmt.Println(newPerson.GetSocialFollowers().GetTwitter())
}
