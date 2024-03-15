package cors

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// Config struct to hold the allowed origins
type Config struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
}

func DefinesAllowOrigins(webEngine *gin.Engine) {
	// create cors config
	config := cors.DefaultConfig()

	// Read allowed origins from a YAML file
	allowedOrigins := readAllowedOriginsFromFile("cors.yml")

	// config cross origins
	config.AllowOrigins = allowedOrigins

	// add cors rules
	webEngine.Use(cors.New(config))
}

// Function to read allowed origins from a YAML file
func readAllowedOriginsFromFile(filename string) []string {
	var config Config

	// Read YAML file
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read YAML file: %v", err)
		return []string{}
	}

	// Unmarshal YAML data into Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Failed to unmarshal YAML: %v", err)
		return []string{}
	}

	return config.AllowedOrigins
}
