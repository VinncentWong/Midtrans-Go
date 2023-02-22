# Midtrans-Go
This repository demonstrate how to use Midtrans with Go

# How to install and run the project
1. Make sure you have do registration on [Midtrans](https://midtrans.com)
2. Clone this repository
3. Type `cd Midtrans-Go`
4. Make a .env file and copy all properties from .env.example file to your .env file.
5. Make sure too that you read the notes that provided on .env.example file
6. Then copy this command on your terminal `go build`
7. Finally, you can run your application with `go run main.go`
8. Happy coding :dizzy:

# Endpoint
### This repository only has one endpoint that is 
`1. /midtrans?paymentType=?`
**Notes: paymentType query only accept two values which are gopay/shopeepay**

# Limitation
### This application only demonstrate how to use midtrans with **gopay/shopeepay** with dummy data that already provided. However, you should have understand how it works if you read the `InitDummyData` method carefully, and can change it to make data dynamically based on User input.

# Contact
Feel free to contact me from my [Instagram Account](https://instagram.com/centwong_) or [Linkedin](https://www.linkedin.com/in/vinncent-alexander-wong-493759213/)
