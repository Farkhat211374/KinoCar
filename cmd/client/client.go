package main

import (
	"context"
	"fmt"
	"github.com/farkhat/KinoCar/api/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	connection, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connection.Close() // Maybe this should be in a separate function and the error handled?

	//c1 := pb.NewCarServiceClient(connection)
	//c2 := pb.NewMovieServiceClient(connection)
	c3 := pb.NewFindcarServiceClient(connection)
	createFindcar(c3)

}

// Create Car
func createCar(c pb.CarServiceClient) {
	fmt.Println("Creating the blog")
	car := &pb.Car{
		Brand:        "Toyota",
		Model:        "Camry",
		Year:         2023,
		Body:         "Sedan",
		BrakeSystem:  "Disc brakes",
		Aspiration:   "Naturally aspirated",
		Horsepower:   350.0,
		Mpg:          32.0,
		Cylinders:    4,
		Acceleration: 6.1,
		Displacement: 2.5,
		Country:      "Japan",
	}
	createCarRes, err := c.CreateCar(context.Background(), &pb.CreateCarRequest{Car: car})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createCarRes)
	carID := createCarRes.GetCar().GetId()
	fmt.Println(carID)

}

// read Car
func readCar(c pb.CarServiceClient) {

	fmt.Println("Reading the blog")
	carID := "647c5ef8405f20f9e1f23404"

	_, err2 := c.ReadCar(context.Background(), &pb.ReadCarRequest{CarId: "647c5ef8405f20f9e1f23404"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readCarReq := &pb.ReadCarRequest{CarId: carID}
	readCarRes, readCarErr := c.ReadCar(context.Background(), readCarReq)
	if readCarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readCarErr)
	}

	fmt.Printf("Blog was read: %v \n", readCarRes)

}

// update Car
func updateCar(c pb.CarServiceClient) {

	carID := "647c38f0c9804f9b5b4cb076"
	newCar := &pb.Car{
		Id:           carID,
		Brand:        "Toyota",
		Model:        "Camry",
		Year:         1999,
		Body:         "Sedan",
		BrakeSystem:  "Disc brakes",
		Aspiration:   "Naturally aspirated",
		Horsepower:   350.0,
		Mpg:          32.0,
		Cylinders:    4,
		Acceleration: 6.1,
		Displacement: 2.5,
		Country:      "Japan",
	}
	updateRes, updateErr := c.UpdateCar(context.Background(), &pb.UpdateCarRequest{Car: newCar})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

}

// delete Car
func deleteCar(c pb.CarServiceClient) {
	carID := "647c38f0c9804f9b5b4cb076"
	deleteRes, deleteErr := c.DeleteCar(context.Background(), &pb.DeleteCarRequest{CarId: carID})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v \n", deleteRes)

}

// get list car
func listGetCar(c pb.CarServiceClient) {
	stream, err := c.ListCar(context.Background(), &pb.ListCarRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetCar())
	}
}

func createMovie(c pb.MovieServiceClient) {
	fmt.Println("Creating the blog")
	movie := &pb.Movie{
		Title:    "The Shawshank Redemption",
		Year:     1994,
		Director: "Frank Darabont",
		Genre:    []string{"Drama", "Crime"},
		Rating:   9.3,
		Duration: 142,
	}
	createMovieRes, err := c.CreateMovie(context.Background(), &pb.CreateMovieRequest{Movie: movie})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createMovieRes)
	carID := createMovieRes.GetMovie().GetId()
	fmt.Println(carID)

}

func readMovie(c pb.MovieServiceClient) {

	fmt.Println("Reading the blog")
	movieID := "647c5d3b917238ce35e490e8"

	_, err2 := c.ReadMovie(context.Background(), &pb.ReadMovieRequest{MovieId: "5bdc29e661b75adcac496cf4"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readMovieReq := &pb.ReadMovieRequest{MovieId: movieID}
	readMovieRes, readMovieErr := c.ReadMovie(context.Background(), readMovieReq)
	if readMovieErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readMovieErr)
	}

	fmt.Printf("Blog was read: %v \n", readMovieRes)

}

// update Movie
func updateMovie(c pb.MovieServiceClient) {

	movieID := "647c5d3b917238ce35e490e8"
	newMovie := &pb.Movie{
		Id:       movieID,
		Title:    "Koja resspect",
		Year:     2025,
		Director: "Frank Darasbonts",
		Genre:    []string{"Comedy"},
		Rating:   5.5,
		Duration: 102,
	}
	updateRes, updateErr := c.UpdateMovie(context.Background(), &pb.UpdateMovieRequest{Movie: newMovie})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

}

// delete Movie
func deleteMovie(c pb.MovieServiceClient) {
	movieID := "647c5d3b917238ce35e490e8"
	deleteRes, deleteErr := c.DeleteMovie(context.Background(), &pb.DeleteMovieRequest{MovieId: movieID})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v \n", deleteRes)

}

// get list Movie
func listGetMovie(c pb.MovieServiceClient) {
	stream, err := c.ListMovie(context.Background(), &pb.ListMovieRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetMovie())
	}
}

func createFindcar(c pb.FindcarServiceClient) {
	fmt.Println("Creating the blog")
	findcar := &pb.Findcar{
		Name:   "Toyota",
		Movies: []string{"film1", "film2"},
	}
	createCarRes, err := c.CreateFindcar(context.Background(), &pb.CreateFindcarRequest{Findcar: findcar})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createCarRes)
	carID := createCarRes.GetFindcar().GetName()
	fmt.Println(carID)

}
