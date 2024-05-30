package Auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// SignedDetails
type SignedDetails struct {
	Username string //`json:"name"`
	jwt.StandardClaims
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("webexec").Collection("user")

	return collection
}
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client = DBinstance()

var userCollection *mongo.Collection = OpenCollection(Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both the detailed token and refresh token
func GenerateAllTokens(Username string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Username: Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		} else {

			temp := strings.Split(authHeader, "Bearer ")
			fmt.Println(temp)
			if len(temp) < 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
				return
			}

			tokenString := temp[1]
			token, err := ValidateToken(tokenString)
			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
		}
		c.Next()
	}
}
