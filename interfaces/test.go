syntax = "proto3";

package test;

option go_package = "github.com/interfaces";

import "google/api/annotations.proto";

message OperationRequest {
  string greeting = 1;
}

message OperationResponse {
 string greeting = 2;
}

service test {
rpc Greeting(OperationRequest) returns (OperationResponse) {
   option (google.api.http) = {
       get: "/api/v1/getGreeting"
    };
}
