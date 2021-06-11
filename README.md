# email-grpc
A C/S architecture by grpc to support sending email.
# Quickly Start
## Clone this Repository
```shell script
git clone https://github.com/Pivot-Studio/email-grpc.git
```
## Start Server
- cd server path
```shell script
#here you are in blabla/email-grpc
cd server
```
- Add a `Config.json` to store email Settings  
Create a `Config.json` file at `email-grpc/server` like this:
```json
{
  "EmailSenderSettings":{
    "email":"youremail@your-email.com",
    "password":"your email token,replace it",
    "servername":"your email host"
  }
}
```
> Port 465 is default port
- run server
```shell script
go run main.go
```

## Start Client
- cd client  
Here I assume you're in path `email-grpc`
```shell script
cd client
```
- run client
```shell script
go run main.go -addr localhost:xxxx -email yourtestemail
```
 Done all of them,your server will be called to send a email to you from you client by gRPC.
