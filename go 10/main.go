package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")
	f1(ctx)

	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go getProductDetailsFromExternalService(10)

	select {
	case productDetail := <-productChannel:
		fmt.Println("Ürün bilgileri başarı ile getirildi!", productDetail)
	case <-ctx.Done():
		fmt.Println("Timeout occurred, context cancelled!")
	}
	fmt.Println("End of the main.")*/

}

func f1(a context.Context) {
	fmt.Println("F1", a.Value("correlation-id"))
	f2(a)
}

func f2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	f3(ctx)
}

func f3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}

func getProductDetailsFromExternalService(productid int64) {
	time.Sleep(2 * time.Second)
	productChannel <- Product{
		Id:   10,
		Name: "Çamaşır Makinesi",
	}
}
