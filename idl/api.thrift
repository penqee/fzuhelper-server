namespace go api
include "model.thrift"

//classroom
struct EmptyClassroomRequest {
    1: required string date
    2: required string campus
    3: required string startTime;//节数
    4: required string endTime;
}

struct EmptyClassroomResponse {
    1: required list<model.Classroom> classrooms
}

service ClassRoomService {
    EmptyClassroomResponse GetEmptyClassrooms(1: EmptyClassroomRequest request)(api.get="/api/v1/common/classroom/empty")
}

//user
struct GetLoginDataRequest {
    1: required string id
    2: required string password
}

struct GetLoginDataResponse {
    1: required string id
    2: required list<string> cookies
}

struct GetValidateCodeRequest{
    1:required string image,
}

struct GetValidateCodeResponse{
    1:model.BaseResp base,
    2:optional string code,
}

service UserService {
    GetLoginDataResponse GetLoginData(1: GetLoginDataRequest request)(api.get="/api/v1/jwch/user/login"),
    GetValidateCodeResponse GetValidateCode(1: GetValidateCodeRequest request)(api.post="/api/v1/jwch/user/validateCode"),

}
