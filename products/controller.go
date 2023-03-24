package products

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (pc *ProductController) HandleRequest(id int) error {

	var choice int

	for choice != 9 {
		now := time.Now()
		fmt.Println(" Product Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
		pc.HandleListProduct() // menggunakan lastest update_at
		fmt.Println("1. Create Product")
		fmt.Println("2. Update Product Name")
		fmt.Println("3. Update Product Price")
		fmt.Println("4. Update Stock")
		fmt.Println("5. Remove Product") // -- soft delete product --
		fmt.Println("9. Exit")           // -- read product by stock availability and delete at is not null untuk keperluan transaksi--
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice) // -- read product by id -- sku untuk dapat akses nama etc and delete at is null --

		switch choice {
		case 1:
			err := pc.handleCreateProduct(id)
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			err := pc.handleUpdateProductNameByID()
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			err := pc.handleUpdateProductPriceByID()
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			err := pc.handleUpdateProductStockByID()
			if err != nil {
				log.Fatal(err)
			}
		case 5:
			err := pc.handleRemoveProductByID()
			if err != nil {
				log.Fatal(err)
			}
		case 9:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
	return nil
}

func (pc *ProductController) HandleListProduct() error {
	products, err := pc.productModel.ListProduct("", 0, 0, 0)
	if err != nil {
		fmt.Println("Failed to retrieve product list:", err)
		return err
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Price", "Stock", "Username"})

	for _, p := range products {
		table.Append([]string{fmt.Sprint(p.ID), p.Name, fmt.Sprintf("%d", p.Price), fmt.Sprintf("%d", p.Stock), p.Username})
	}

	table.Render()
	return nil
}

func (pc *ProductController) handleCreateProduct(id int) error {
	var name string
	var price, stock int

	fmt.Println("---------------------------")
	fmt.Println("Create Product")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter product name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(strings.ToLower(name))

		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty")
	}

	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	fmt.Print("Enter product stock: ")
	fmt.Scan(&stock)

	product, err := pc.productModel.CreateProduct(name, price, stock, id)
	if err != nil {
		return err
	}

	fmt.Println("---------------------------")
	fmt.Printf("Product created successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d", product.ID, product.Name, product.Price, product.Stock, id)
	fmt.Println("---------------------------")
	return nil
}

func (pc *ProductController) handleUpdateProductNameByID() error {
	var name string
	var createdBy int

	fmt.Println("---------------------------")
	fmt.Println("Update Product Name")
	fmt.Println("---------------------------")

	fmt.Print("Enter updated by user ID: ")
	fmt.Scan(&createdBy)

	fmt.Print("Enter product ID: ")
	var ID int
	fmt.Scan(&ID)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter product name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(strings.ToLower(name))

		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty")
	}

	product, err := pc.productModel.UpdateProductNameByID(ID, createdBy, name)
	if err != nil {
		fmt.Println("Failed to update product name:", err)
		return err
	}

	fmt.Println("---------------------------")
	fmt.Printf("Product updated successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d\n", product.ID, product.Name, product.Price, product.Stock, product.Created_by)
	fmt.Println("---------------------------")
	return nil
}

func (pc *ProductController) handleUpdateProductPriceByID() error {
	var price int
	var createdBy int

	fmt.Println("---------------------------")
	fmt.Println("Update Product Price")
	fmt.Println("---------------------------")

	fmt.Print("Enter updated by user ID: ")
	fmt.Scan(&createdBy)

	fmt.Print("Enter product ID: ")
	var ID int
	fmt.Scan(&ID)

	for {
		fmt.Print("Enter product price: ")
		var input string
		fmt.Scan(&input)
		parsedPrice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please input a valid number")
			continue
		}
		if parsedPrice < 0 {
			fmt.Println("There is no minus price")
			continue
		}
		price = parsedPrice
		break
	}

	product, err := pc.productModel.UpdateProductPriceByID(ID, createdBy, price)
	if err != nil {
		fmt.Println("Failed to update product price:", err)
		return err
	}

	fmt.Println("---------------------------")
	fmt.Printf("Product updated successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d\n", product.ID, product.Name, product.Price, product.Stock, product.Created_by)
	fmt.Println("---------------------------")
	return nil
}

func (pc *ProductController) handleUpdateProductStockByID() error {
	var stock int
	var createdBy int

	fmt.Println("---------------------------")
	fmt.Println("Update Product Price")
	fmt.Println("---------------------------")

	fmt.Print("Enter updated by user ID: ")
	fmt.Scan(&createdBy)

	fmt.Print("Enter product ID: ")
	var ID int
	fmt.Scan(&ID)

	for {
		fmt.Print("Enter product price: ")
		var input string
		fmt.Scan(&input)
		parsedStock, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please input a valid number")
			continue
		}
		if parsedStock < 0 {
			fmt.Println("There is no minus price")
			continue
		}
		stock = parsedStock
		break
	}

	product, err := pc.productModel.UpdateProductPriceByID(ID, createdBy, stock)
	if err != nil {
		fmt.Println("Failed to update product price:", err)
		return err
	}

	fmt.Println("---------------------------")
	fmt.Printf("Product updated successfully with details:\nID\t: %d\nName\t: %s\nPrice\t: %d\nStock\t: %d\nCreated by\t: %d\n", product.ID, product.Name, product.Price, product.Stock, product.Created_by)
	fmt.Println("---------------------------")
	return nil
}

func (pc *ProductController) handleRemoveProductByID() error {
	fmt.Println("---------------------------")
	fmt.Println("Remove Product ID")
	fmt.Println("---------------------------")

	// make map untuk nampung IDs yang sudah ada.
	productIDs := make(map[int]bool)

	// memproses map untuk nampung IDs yang sudah ada, jika ada maka true dan sebalikanya.
	products, err := pc.productModel.ListProduct("", 0, 0, 0)
	if err != nil {
		fmt.Println("failed to retrieve products:", err)
		return err
	}
	for _, product := range products {
		productIDs[product.ID] = true
	}

	for {
		var idProduct int
		fmt.Print("Enter id product: ")
		fmt.Scan(&idProduct)

		if !productIDs[idProduct] {
			fmt.Println("id product", idProduct, "does not exist, please check again.")
			continue
		}

		err := pc.productModel.RemoveProductByID(idProduct)
		if err != nil {
			fmt.Println("failed to remove id product", idProduct, "error:", err)
		} else {
			fmt.Println("id product", idProduct, "has been removed.")
		}
		break
	}
	return nil
}
