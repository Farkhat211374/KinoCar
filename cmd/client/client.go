package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/farkhat/KinoCar/api/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strings"
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
	//c3 := pb.NewFindcarServiceClient(connection)
	//c4 := pb.NewUserServiceClient(connection)
	//c5 := pb.NewPublishServiceClient(connection)
	c6 := pb.NewReceiveServiceClient(connection)
	//createFindcar(c3)
	receive(c6)

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

func readFindcar(c pb.FindcarServiceClient) {

	fmt.Println("Reading the blog")
	carName := "Toyota"

	_, err2 := c.ReadFindcar(context.Background(), &pb.ReadFindcarRequest{CarName: "bmv"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readFindcarReq := &pb.ReadFindcarRequest{CarName: carName}
	readFindcarRes, readFindcarErr := c.ReadFindcar(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}

	fmt.Printf("Blog was read: %v \n", readFindcarRes)

}

// update Movie
func updateFindcar(c pb.FindcarServiceClient) {
	carName := "Toyota"
	newFindcar := &pb.Findcar{
		Name:   carName,
		Movies: []string{"Pump Fiction", "Fast and Furious"},
	}
	updateRes, updateErr := c.UpdateFindcar(context.Background(), &pb.UpdateFindcarRequest{Findcar: newFindcar})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

}

// delete Movie
func deleteFindcar(c pb.FindcarServiceClient) {
	carName := "Toyota"
	deleteRes, deleteErr := c.DeleteFindcar(context.Background(), &pb.DeleteFindcarRequest{CarName: carName})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v \n", deleteRes)

}

func getMoviesByCar(c pb.UserServiceClient) {

	fmt.Println("Loading the GetMoviesByCarRequest")
	carName := "Toyota"

	readFindcarReq := &pb.GetMoviesByCarRequest{CarName: carName}
	readFindcarRes, readFindcarErr := c.GetMoviesByCar(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("If you want to get specil info enter 'yes', else any other words: ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	if strings.TrimRight(answer, "\r\n") == "yes" {
		movies := readFindcarRes.Movies
		i := 1
		for _, movie := range movies {
			request := &pb.GetMovieInfoRequest{MovieName: movie}
			answer, err := c.GetMovieInfo(context.Background(), request)
			if err != nil {
				fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
				fmt.Printf("movie-%v : Movie Not Found \n", i)
			} else {
				fmt.Printf("movie-%v : %s \n", i, answer)
			}
			i++
		}

	}

}

func getCarsByMovie(c pb.UserServiceClient) {

	fmt.Println("Loading ...")
	movieName := "film1"

	readFindcarReq := &pb.GetCarsByMovieRequest{MovieName: movieName}
	readFindcarRes, readFindcarErr := c.GetCarsByMovie(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("If you want to get specil info enter 'yes', else any other words: ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}

	if strings.TrimRight(answer, "\r\n") == "yes" {
		cars := readFindcarRes.Cars
		i := 1
		for _, car := range cars {
			//fmt.Println(car)
			request := &pb.GetCarInfoRequest{CarName: car}
			answer, err := c.GetCarInfo(context.Background(), request)
			if err != nil {
				fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
				fmt.Printf("car-%v : Car Not Found \n", i)
			} else {
				fmt.Printf("car-%v : %s \n", i, answer)
			}
			i++
		}

	}

}

func getMovieInfo(c pb.UserServiceClient) {

	fmt.Println("Loading ...")
	movieName := "The Shawshank Redemption"

	readFindcarReq := &pb.GetMovieInfoRequest{MovieName: movieName}
	readFindcarRes, readFindcarErr := c.GetMovieInfo(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

}

func getCarInfo(c pb.UserServiceClient) {

	fmt.Println("Loading ...")
	carName := "Toyota"

	readFindcarReq := &pb.GetCarInfoRequest{CarName: carName}
	readFindcarRes, readFindcarErr := c.GetCarInfo(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

}

func publish(c pb.PublishServiceClient) {

	fmt.Println("Loading ...")
	carName := "Toyota"
	movies := []string{"toy story", "star wars"}

	readFindcarReq := &pb.PublishRequest{CarName: carName, Movies: movies}
	readFindcarRes, readFindcarErr := c.Publish(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

}

func receive(c pb.ReceiveServiceClient) {

	fmt.Println("Loading ...")
	carName := "Toyota"

	readFindcarReq := &pb.ReceiveRequest{CarName: carName}
	readFindcarRes, readFindcarErr := c.Receive(context.Background(), readFindcarReq)
	if readFindcarErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readFindcarErr)
	}
	fmt.Println("==================================")
	fmt.Printf("%v \n", readFindcarRes)
	fmt.Println("==================================")

}
