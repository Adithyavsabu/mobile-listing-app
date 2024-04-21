package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mobile_listing_app/data"
	"mobile_listing_app/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

//ceate new product

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}
	var product data.Product
	// Extract the JSON data
	err = json.Unmarshal([]byte(r.FormValue("productData")), &product)
	if err != nil {
		fmt.Println("unable to parse product data", err)
		return
	}

	// Extract the image file

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to retrieve image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//Extract price from requestpayload
	var strprice = r.FormValue("price")
	var price float64
	price, err = strconv.ParseFloat(strprice, 64)
	if err != nil {
		fmt.Println("unable to parse price", err)
		return
	}
	product.Price = price

	// Save the image file to the server
	fileExt := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExt
	fmt.Println("new filename", fileName)
	uploadDir := "../frontend/images/"
	imagePath := filepath.Join(uploadDir, fileName)
	f, err := os.Create(imagePath)
	if err != nil {
		fmt.Println("image error :Unable to save image file", err)
		return
	}
	product.ImagePath = "../images/" + fileName //displaying in template using ../
	product.Image = fileName
	defer f.Close()
	io.Copy(f, file)

	// Database operation to insert product details
	_, err = db.Exec("INSERT INTO product(name, price, specification, image) VALUES($1, $2, $3, $4)", product.ProductName, product.Price, product.Specification, product.Image)
	if err != nil {
		fmt.Println("Error adding details to the db:", err) // Print error message

		return
	}

	// Return the created product as JSON response
	json.NewEncoder(w).Encode(product)
}

// List all products
func GetAllProduct(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		fmt.Println("Error processing the query", err)

		return
	}
	defer rows.Close()

	var products []data.Product
	for rows.Next() {
		var product data.Product
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Specification, &product.Image)

		if err != nil {
			fmt.Println("Error scanning the query", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(product.Image)
		product.ImagePath = "../images/" + product.Image
		products = append(products, product)

	}
	json.NewEncoder(w).Encode(products)
}

// Get one product based on id
func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	productID := utils.FetchId(w, r)

	row := db.QueryRow("SELECT productid, name, price, specification, image FROM product WHERE productid = $1", productID)
	var product data.Product
	err := row.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Specification, &product.Image)
	if err != nil {
		fmt.Println("Error processing the query:Product not found", err)

		return
	}
	product.ImagePath = "../images/" + product.Image

	// Return the product details
	json.NewEncoder(w).Encode(product)
}

// Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	// Extract product ID from the request URL
	productID := utils.FetchId(w, r)

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	//var updatedData data.UpdatedProduct

	var product data.Product
	// Extract the JSON data
	err = json.Unmarshal([]byte(r.FormValue("productData")), &product)
	if err != nil {
		fmt.Println("Error parsing data: ", err)
		http.Error(w, "Unable to parse updated product data", http.StatusBadRequest)
		return
	}

	//Extract price from requestpayload
	var strprice = r.FormValue("price")
	var price float64
	price, err = strconv.ParseFloat(strprice, 64)
	if err != nil {
		fmt.Println("unable to parse price", err)
		return
	}
	product.Price = price

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to retrieve image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExt := filepath.Ext(handler.Filename)
	fileName := uuid.New().String() + fileExt
	uploadDir := "../frontend/images/"
	imagePath := filepath.Join(uploadDir, fileName)
	f, err := os.Create(imagePath)
	if err != nil {
		fmt.Println(" error creating image path", err)
		http.Error(w, "Unable to save image file", http.StatusInternalServerError)
		return
	}
	product.Image = fileName
	product.ImagePath = "../images/" + fileName //displaying in template using ../
	defer f.Close()
	io.Copy(f, file)

	_, err = db.Exec("UPDATE product SET name=$1, price=$2, specification=$3, image=$4 WHERE productid=$5",
		product.ProductName, product.Price, product.Specification, product.Image, productID)
	if err != nil {
		fmt.Println("Error  updating data: ", err)
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	fmt.Println("product update data:", product)

	// Return success message as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})

}

// Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := utils.FetchId(w, r)
	var image string
	err := db.QueryRow("SELECT image FROM product WHERE productid=$1", productID).Scan(&image)
	if err != nil {
		fmt.Println("Error querying", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var imagePath = "../frontend/images/" + image

	// Delete the image file from the server
	err = os.Remove(imagePath)
	if err != nil {
		fmt.Println("Error removing the image ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the product from the database
	_, err = db.Exec("DELETE FROM product WHERE productid=$1", productID)
	if err != nil {
		fmt.Println("Error deleting the product", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	// Return success message as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})

}

// Search for products based on a search query
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the search query from the request parameters
	params := mux.Vars(r)
	query := params["query"]

	// Execute the database query to search for products
	rows, err := db.Query("SELECT * FROM product WHERE name ILIKE '%' || $1 || '%'", query)
	if err != nil {
		fmt.Println("Error searching products:", err)
		http.Error(w, "Error searching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Parse the query results into Product structs
	var products []data.Product
	for rows.Next() {
		var product data.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Specification, &product.Image); err != nil {
			fmt.Println("Error scanning row:", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		product.ImagePath = "../images/" + product.Image
		products = append(products, product)
	}

	// Return the search results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
