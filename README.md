### Publisher and Subscriber in GoLang with RabbitMQ and MySQL

**How it works?**

* Run RabbitMQ in local with default username and password
  
        `username: guest` `password: guest` `port: 5672`
    * We can change this value in `queue.go`

    
* Run MySQL in local with default username and password
  
        `username: root` `password: root` `port:3306`
    * We can change this value in `database.go`
    * Create _DB / Schema_ with a name **Booking**
  
        `create schema booking;`
* Then run the `main.go` using go command `go run main.go`


**Note:**   
* This project support publisher and consumer both but for best interest Publisher portion is commented.
* Use this linked project to publish message [Publisher Project - GitHub](https://github.com/sathishkumar-manogaran/publisher-in-golang).

