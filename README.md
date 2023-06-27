# Go Fiber Authentication API

Go Fiber Authentication API is a RESTful API built using the Fiber web framework.

## Installation

1. Clone the repository:
  ```shell
  git clone https://github.com/yusuftalhaklc/go-fiber-authentication.git
  ```

2. Navigate to the project directory:
  ```shell
  cd go-fiber-authentication
  ```
3. Install the dependencies:
  ```shell
  go mod tidy
  ```
4. Start the API server:
  ```shell
  go run main.go
  ```

**It is currently running on localhost port 8080.** [Postman Collection](https://red-shuttle-655108.postman.co/workspace/go-fiber-auth~1c48d0cc-5e90-4496-b2f0-c292446f90cf/collection/27159195-7a2c468b-a60e-4013-a4ee-0bd89310b1c7?action=share&creator=27159195)

## API Endpoints
### Signup

- **Endpoint:** `/api/signup`
- **Method:** `POST`
- **Description:** Signup.


### Login 

- **Endpoint:** `/api/user/login`
- **Method:** `POST`
- **Description:** Login.

### Logout 

- **Endpoint:** `/api/user/logout`
- **Method:** `GET`
- **Description:** Logout.

## Request Body and Response Examples

### Signup
- Request Body
```json
{
    "first_name":"Yusuf Talha",
    "last_name":"Kılıç",
    "password":"pass123",
    "email":"yusuftalhaklc@gmeil.com",
    "phone":"5555555555"
}
```
- Response
```json
{
    "data": {
        "ID": "649ae0cf306d78d5c350e496",
        "user_id": "649ae0cf306d78d5c350e496",
        "first_name": "Yusuf Talha",
        "last_name": "Kılıç",
        "password": "9b8769a4a742959a2d0298c36fb70623f2dfacda8436237df08d8dfd5b37374c",
        "email": "yusuftalhaklc@gmeil.com",
        "phone": "5555555555",
        "avatar": null,
        "token": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc4NzM0OTUsImlkIjoiNjQ5YWUwY2YzMDZkNzhkNWMzNTBlNDk2IiwibWFpbCI6Inl1c3VmdGFsaGFrbGNAZ21laWwuY29tIn0.zAMeCix0W2OX3bjl2owU9MTShdzRTbX19eDcdktCJpPaVCjTbdMFqgVC5qoNuoYkIkS5OXomGflCS19d4otmew",
        "created_at": "2023-06-27T16:14:55+03:00",
        "last_login_at": "0001-01-01T00:00:00Z",
        "logout_at": "0001-01-01T00:00:00Z",
        "deleted_at": "0001-01-01T00:00:00Z"
    },
    "message": "User has created",
    "status": "success"
}
```

### Login
- Request
```json
{
    "email":"yusuftalhaklc@gmeil.com",
    "password":"pass123"
}
```

- Response
```json
{
    "data": {
        "email": "yusuftalhaklc@gmeil.com",
        "token": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc4NzM1NTIsImlkIjoiMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwibWFpbCI6Inl1c3VmdGFsaGFrbGNAZ21laWwuY29tIn0.u60QhRQeiOvAEnVGtu0RPRbNGFPM41QBdPZbqqcntymbz096AwAO8jIEYiGQDZX-FPu2Kx9F2iamcsVbuI2Jww"
    },
    "message": "Successfully login",
    "status": "success"
}
```

### Logout
- Request <br>
Authorization = token
```http
GET /api/user/Logout
```

- Response
```json
{
    "message": "Successfully logout",
    "status": "success"
}
```
