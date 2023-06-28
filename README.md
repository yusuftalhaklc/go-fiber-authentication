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
- **Method:** `POST`
- **Description:** Logout.

### Get User

- **Endpoint:** `/api/user/`
- **Method:** `GET`
- **Description:** Get user details.

## Request Body and Response Examples

### Signup
- Request Body
```json
{
    "first_name":"User First Name",
    "last_name":"User Last Name",
    "password":"password",
    "email":"username@example.com",
    "phone":"5555555555"
}
```
- Response
```json
{
    "data": {
        "ID": "649c8813a9299b295118e33c",
        "user_id": "649c8813a9299b295118e33c",
        "first_name": "User First Name",
        "last_name": "User Last Name",
        "password": "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
        "email": "username@example.com",
        "phone": "5555555555",
        "avatar": null,
        "token": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc5ODE4NTEsImlkIjoiNjQ5Yzg4MTNhOTI5OWIyOTUxMThlMzNjIiwibWFpbCI6InVzZXJuYW1lQGV4YW1wbGUuY29tIn0._WnQuFPkJBr3s49EpU8BJo6ndVizOmz1OM5hp9uzh2Jz-NhOvcEByZK9EYJKoNRNBoJN2uKtPn1MyIdB-Nus-w",
        "created_at": "2023-06-28T22:20:51+03:00",
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
    "email": "username@example.com",
    "password": "password"
}
```

- Response
```json
{
    "data": {
        "email": "username@example.com",
        "token": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc5ODIwOTEsImlkIjoiMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwibWFpbCI6InVzZXJuYW1lQGV4YW1wbGUuY29tIn0.QY9WFwJdTi4tod8S8bnh3gRGt6SzwVsf3RXOzRwQlHhPsfkOv9KiK4l3BX9FpBu_kM1aSWzkEO7Mx5Y_vxEH3A"
    },
    "message": "Successfully login",
    "status": "success"
}
```

### Logout
- Request <br>
Authorization = token
```http
POST /api/user/Logout
```

- Response
```json
{
    "message": "Successfully logout",
    "status": "success"
}
```

### Get user
- Request <br>
Authorization = token
```http
GET /api/user/
```

- Response
```json
{
    "data": {
        "first_name": "User First Name",
        "last_name": "User Last Name",
        "password": "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
        "email": "username@example.com",
        "phone": "5555555555",
        "avatar": null,
        "created_at": "2023-06-28T19:20:51Z",
        "last_login_at": "2023-06-28T19:38:27Z",
        "logout_at": "2023-06-28T19:38:19Z"
    },
    "message": "Successfully found",
    "status": "success"
}
```
