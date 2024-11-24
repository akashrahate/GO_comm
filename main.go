package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Login struct {
    Name     string `json:"name"`
    Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
    var user Login
    // Bind JSON input to the `user` struct
    if err := c.BindJSON(&user); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    // Check if the user exists or perform authentication
    if user.Name == "" || user.Password == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Name and password are required"})
        return
    }
    // Dummy validation (replace with actual logic)
    if user.Name == "admin" && user.Password == "pass" {
        c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
    } else {
        c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
    }
}

func main() {
    router := gin.Default()
    router.POST("/login", loginHandler)
    // Bind to all interfaces (0.0.0.0) to expose the API outside the container
    router.Run(":88")
}
