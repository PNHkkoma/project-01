package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect() *mongo.Collection {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	collection := client.Database("sessionData").Collection("sessionData")
	return collection
}
func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")) // tạo đường dẫn tới
	if err != nil {
		log.Println("không tạo được đường dẫn")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //tạo kênh context với 10s
	defer cancel()
	err = client.Connect(ctx) //kết nối client với ctx đó
	if err != nil {
		log.Println("không tạo được kết nối")
	}

	err = client.Ping(context.TODO(), nil) //gửi 1 phương thức (1 yêu cầu ping) tới db để kiểm tra kết nối
	//context.todo trả về một ngữ cảnh rỗng, tức là ko chứa bất kỳ gt nào, thường dùng khi ko cần thiết lập bất kỳ gt ngữ cảnh, nil là cho biết ko có tỳ chọn nào đc cung cấp cho hoạt động ping
	if err != nil {
		log.Println("ping ko trả về mẹ gi")
		return nil
	}
	//này cần xóa này
	collection := client.Database("sessionData").Collection("sessionData")
	log.Println(collection)
	fmt.Println("Successfully Connected to the mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func SessionData(client *mongo.Client, CollectionName string) *mongo.Collection { //tạo 2 biến data tồn tại xuyên suốt hệ thống ...
	var collection *mongo.Collection = client.Database("sessionData").Collection("sessionData")
	return collection

}
