package products

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
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
		now := time.Now()
		fmt.Println(" Product Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
		pc.handleListProduct() // menggunakan lastest update_at
		fmt.Println("1. Create Product")
		fmt.Println("2. Update Product Name")
		fmt.Println("3. Update Product Price")
		fmt.Println("4. Update Stock")
		fmt.Println("5. Delete Product") // -- soft delete product --
		fmt.Println("9. Exit")           // -- read product by stock availability and delete at is not null untuk keperluan transaksi--
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice) // -- read product by id -- sku untuk dapat akses nama etc and delete at is null --

		switch choice {
		case 1:
			pc.handleCreateProduct()
		case 2:

		case 3:

		case 4:

		case 5:

		case 9:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func (pc *ProductController) handleListProduct() {
	products, err := pc.productModel.ListProduct("", 0, 0, "", 0)
	if err != nil {
		fmt.Println("Failed to retrieve product list:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Stock", "Username", "Latest Update"})

	for _, p := range products {
		table.Append([]string{fmt.Sprint(p.ID), p.Name, fmt.Sprintf("%d", p.Price), fmt.Sprintf("%d", p.Stock), p.Username, p.Updated_at})
	}

	table.Render()
}

func (pc *ProductController) handleCreateProduct() {
	var name string
	var price, stock, createdBy int

	fmt.Println("---------------------------")
	fmt.Println("Create Product")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter product name: ")
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
	fmt.Printf("Product created successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d\nUpdated at\t: %s\n", product.ID, product.Name, product.Price, product.Stock, product.Created_by, product.Updated_at)
	fmt.Println("---------------------------")
}
