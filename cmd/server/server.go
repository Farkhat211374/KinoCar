package main

import (
	"context"
	"fmt"
	"github.com/farkhat/KinoCar/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net"
	"os"
	"os/signal"
)

var CollectionCar *mongo.Collection
var CollectionMovie *mongo.Collection
var CollectionFindcar *mongo.Collection

type Server struct {
	pb.CarServiceServer
	pb.MovieServiceServer
	pb.FindcarServiceServer
}

type Findcar struct {
	Name   string   `bson:"carName,omitempty"`
	Movies []string `bson:"movies,omitempty"`
}

type Movie struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Rating   float32            `bson:"rating,omitempty"`
	Year     int32              `bson:"year,omitempty"`
	Director string             `bson:"director,omitempty"`
	Genre    []string           `bson:"genre,omitempty"`
	Duration int32              `bson:"duration,omitempty"`
}

type Car struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Brand        string             `bson:"brand,omitempty"`
	Model        string             `bson:"model,omitempty"`
	Year         int32              `bson:"year,omitempty"`
	Body         string             `json:"body,omitempty"`
	BrakeSystem  string             `json:"brake_system,omitempty"`
	Aspiration   string             `json:"aspiration,omitempty"`
	Horsepower   float32            `json:"horsepower,omitempty"`
	Mpg          float32            `json:"mpg,omitempty"`
	Cylinders    int32              `json:"cylinders,omitempty"`
	Acceleration float32            `json:"acceleration,omitempty"`
	Displacement float32            `json:"displacement,omitempty"`
	Country      string             `json:"country,omitempty"`
}

func (*Server) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarResponse, error) {
	fmt.Println("Create car request")
	car := req.GetCar()

	data := Car{
		Brand:        car.GetBrand(),
		Model:        car.GetModel(),
		Year:         car.GetYear(),
		Body:         car.GetBody(),
		BrakeSystem:  car.GetBrakeSystem(),
		Aspiration:   car.GetAspiration(),
		Horsepower:   car.GetHorsepower(),
		Mpg:          car.GetMpg(),
		Cylinders:    car.GetCylinders(),
		Acceleration: car.GetAcceleration(),
		Displacement: car.GetDisplacement(),
		Country:      car.GetCountry(),
	}

	res, err := CollectionCar.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &pb.CreateCarResponse{
		Car: &pb.Car{
			Id:           oid.Hex(),
			Brand:        car.GetBrand(),
			Model:        car.GetModel(),
			Year:         car.GetYear(),
			Body:         car.GetBody(),
			BrakeSystem:  car.GetBrakeSystem(),
			Aspiration:   car.GetAspiration(),
			Horsepower:   car.GetHorsepower(),
			Mpg:          car.GetMpg(),
			Cylinders:    car.GetCylinders(),
			Acceleration: car.GetAcceleration(),
			Displacement: car.GetDisplacement(),
			Country:      car.GetCountry(),
		},
	}, nil

}

func (*Server) ReadCar(ctx context.Context, req *pb.ReadCarRequest) (*pb.ReadCarResponse, error) {
	fmt.Println("Read blog request")

	blogID := req.GetCarId()
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &Car{}
	filter := bson.M{"_id": oid}

	res := CollectionCar.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &pb.ReadCarResponse{
		Car: dataToCarPb(data),
	}, nil
}

func dataToCarPb(data *Car) *pb.Car {
	return &pb.Car{
		Id:           data.ID.Hex(),
		Brand:        data.Brand,
		Model:        data.Model,
		Year:         data.Year,
		Body:         data.Body,
		BrakeSystem:  data.BrakeSystem,
		Aspiration:   data.Aspiration,
		Horsepower:   data.Horsepower,
		Mpg:          data.Mpg,
		Cylinders:    data.Cylinders,
		Acceleration: data.Acceleration,
		Displacement: data.Displacement,
		Country:      data.Country,
	}
}

func (*Server) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.UpdateCarResponse, error) {
	fmt.Println("Update blog request")
	car := req.GetCar()
	oid, err := primitive.ObjectIDFromHex(car.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &Car{}
	filter := bson.M{"_id": oid}

	res := CollectionCar.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	// we update our internal struct
	data.Brand = car.GetBrand()
	data.Model = car.GetModel()
	data.Year = car.GetYear()
	data.Displacement = car.GetDisplacement()
	data.Body = car.GetBody()
	data.Mpg = car.GetMpg()
	data.Cylinders = car.GetCylinders()
	data.Horsepower = car.GetHorsepower()
	data.BrakeSystem = car.GetBrakeSystem()
	data.Aspiration = car.GetAspiration()
	data.Acceleration = car.GetAcceleration()
	data.Country = car.GetCountry()

	_, updateErr := CollectionCar.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &pb.UpdateCarResponse{
		Car: dataToCarPb(data),
	}, nil

}

func (*Server) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteCarResponse, error) {
	fmt.Println("Delete blog request")
	oid, err := primitive.ObjectIDFromHex(req.GetCarId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := CollectionCar.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog in MongoDB: %v", err),
		)
	}

	return &pb.DeleteCarResponse{CarId: req.GetCarId()}, nil
}

func (*Server) ListCar(_ *pb.ListCarRequest, stream pb.CarService_ListCarServer) error {
	fmt.Println("List blog request")

	cur, err := CollectionCar.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background()) // Should handle err
	for cur.Next(context.Background()) {
		data := &Car{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}
		stream.Send(&pb.ListCarResponse{Car: dataToCarPb(data)}) // Should handle err
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")) //
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Blog Service Started")

	CollectionCar = client.Database("mydb").Collection("car")
	CollectionMovie = client.Database("mydb").Collection("movie")
	CollectionFindcar = client.Database("mydb").Collection("findcar")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCarServiceServer(grpcServer, &Server{})
	pb.RegisterMovieServiceServer(grpcServer, &Server{})
	pb.RegisterFindcarServiceServer(grpcServer, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	go func() {
		fmt.Println("Starting Server...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	grpcServer.Stop()
	fmt.Println("End of Program")
}

func (*Server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Println("Create movie request")
	car := req.GetMovie()

	data := Movie{
		Title:    car.GetTitle(),
		Rating:   car.GetRating(),
		Year:     car.GetYear(),
		Director: car.GetDirector(),
		Genre:    car.GetGenre(),
		Duration: car.GetDuration(),
	}

	res, err := CollectionMovie.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:       oid.Hex(),
			Title:    car.GetTitle(),
			Rating:   car.GetRating(),
			Year:     car.GetYear(),
			Director: car.GetDirector(),
			Genre:    car.GetGenre(),
			Duration: car.GetDuration(),
		},
	}, nil

}

func (*Server) ReadMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read blog request")

	blogID := req.GetMovieId()
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &Movie{}
	filter := bson.M{"_id": oid}

	res := CollectionMovie.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &pb.ReadMovieResponse{
		Movie: dataToMoviePb(data),
	}, nil
}

func dataToMoviePb(data *Movie) *pb.Movie {
	return &pb.Movie{
		Id:       data.ID.Hex(),
		Title:    data.Title,
		Rating:   data.Rating,
		Year:     data.Year,
		Director: data.Director,
		Genre:    data.Genre,
		Duration: data.Duration,
	}
}

func (*Server) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Println("Update blog request")
	movie := req.GetMovie()
	oid, err := primitive.ObjectIDFromHex(movie.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &Movie{}
	filter := bson.M{"_id": oid}

	res := CollectionMovie.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	// we update our internal struct
	data.Title = movie.GetTitle()
	data.Rating = movie.GetRating()
	data.Year = movie.GetYear()
	data.Director = movie.GetDirector()
	data.Genre = movie.GetGenre()
	data.Duration = movie.GetDuration()

	_, updateErr := CollectionCar.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &pb.UpdateMovieResponse{
		Movie: dataToMoviePb(data),
	}, nil

}

func (*Server) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Println("Delete blog request")
	oid, err := primitive.ObjectIDFromHex(req.GetMovieId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := CollectionMovie.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog in MongoDB: %v", err),
		)
	}

	return &pb.DeleteMovieResponse{MovieId: req.GetMovieId()}, nil
}

func (*Server) ListMovie(_ *pb.ListMovieRequest, stream pb.MovieService_ListMovieServer) error {
	fmt.Println("List blog request")

	cur, err := CollectionMovie.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background()) // Should handle err
	for cur.Next(context.Background()) {
		data := &Movie{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}
		stream.Send(&pb.ListMovieResponse{Movie: dataToMoviePb(data)}) // Should handle err
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}

func (*Server) CreateFindcar(ctx context.Context, req *pb.CreateFindcarRequest) (*pb.CreateFindcarResponse, error) {
	fmt.Println("Create findcar request")
	findcar := req.GetFindcar()

	data := Findcar{
		Name:   findcar.GetName(),
		Movies: findcar.GetMovies(),
	}

	res, err := CollectionFindcar.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &pb.CreateFindcarResponse{
		Findcar: &pb.Findcar{
			Name:   findcar.GetName(),
			Movies: findcar.GetMovies(),
		},
	}, nil

}
