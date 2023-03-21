package products

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ProductController struct {
	productModel *ProductModel
}

func NewProductController(pm *ProductModel) *ProductController {
	return &ProductController{pm}
}

func (pc *ProductController) HandleRequest() {
	
	var choice int
	
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Create Product")
		fmt.Println("2. Update Product")
		fmt.Println("3. Delete Product")
		fmt.Println("9. Exit")

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			pc.handleCreateProduct()
		case 2:
			fmt.Print("Exiting program...")
			return
		default:
			fmt.Print("Invalid choice")
			continue
		}
	}
}

func (pc *ProductController) handleCreateProduct() {
	var name string
	var price, stock, createdBy int

	fmt.Println("---------------------------")
	fmt.Println("Create Product")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter product name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty")
	}
	

	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	fmt.Print("Enter product stock: ")
	fmt.Scan(&stock)

	fmt.Print("Enter created by user ID: ")
	fmt.Scan(&createdBy)

	product, err := pc.productModel.CreateProduct(name, price, stock, createdBy)
	if err != nil {
		fmt.Println("Failed to create product:", err)
		return
	}

	fmt.Println("---------------------------")
	fmt.Printf("Product created successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d\nUpdated at\t: %s\n", product.ID, product.Name, product.Price, product.Stock, product.Created_by,product.Updated_at)
	fmt.Println("---------------------------")
}