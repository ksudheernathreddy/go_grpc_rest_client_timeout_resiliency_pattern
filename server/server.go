package main

import(
  "net"
  "net/http"
  "context"
  "fmt"
  pb "github.com/interfaces"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
)

const (
  grpcPort    = ":50051"
  restPort    = ":50052"
)


func startGRPCServer(address string) error {
  lis, err := net.Listen("tcp", address)
  if err != nil {
   return err
  }
  s := grpc.NewServer()
  pb.RegisterTestServer(s, &server{})
  fmt.Printf("Starting HTTP/2 GRPC server on %s", address)
  if err := s.Serve(lis); err != nil {
   return err
  }
  return nil
}

func startRESTServer(restAddress, grpcAddress string) error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := pb.RegisterTestHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
  if err != nil {
    return err
  }
  fmt.Printf("Starting HTTP/1.1 REST server on %s", restAddress)
  http.ListenAndServe(restAddress, mux)
  return nil
}

func main(){
  err := startGRPCServer(grpcPort)
  if err != nil {
    fmt.Printf("Failed to create GRPC server: %v", err)
  }
  err := startRESTServer(restPort)
  if err != nil {
     fmt.Printf("Failed to create REST server: %v", err)
  }
  select{}
}
