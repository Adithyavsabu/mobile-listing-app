document.addEventListener("DOMContentLoaded", function() {
    if (window.location.pathname.includes("admin-home.html")) {
        fetchProducts();
        document.querySelector(".add-product-btn").addEventListener("click", redirectToAddProduct);
        document.querySelector(".search-box form").addEventListener("submit", handleSearch);
       
    }

});
function fetchProducts(){
    fetch("http://localhost:8080/list")
    .then(response=>{return response.json()})
    .then(products => displayProducts(products))

    .catch((err)=>console.log("Could not fetch details of products",err));
}

function displayProducts(productList) {
    const productContainer = document.querySelector(".product");
    if (!productContainer) {
        console.error("Product container not found in the DOM");
        return;
    }
    let output = "";
    for (let i = 0; i < productList.length; i++) {
        let product = productList[i];
        output += `
        <div class="product-card">
        <div class="product-img"><img src="${product.image_path}" ></div>
        <div class="product-info">
            <h3>${product.name}</h3>
            <p>Rs ${product.price}</p>
            <div class="btn-group">
            <a href="#" onclick="window.location.href = 'product-view.html?id=${product.product_id}';"> 
            <button class="btn btn-view">View</button>
                </a>
                <a href="#" onclick="window.location.href = 'edit-product.html?id=${product.product_id}';"> 
                <button class="btn btn-edit">Edit</button>
            </a>   
            <a href="#" onclick="deleteProduct('${product.product_id}')"> 
            <button class="btn btn-delete">Delete</button>
            </a> 
            </div>
        </div>
    </div>     `;


    }
    productContainer.innerHTML = output;
}

function redirectToAddProduct() {
    window.location.href = "add-product.html";
    addProduct();

}
 
function addProduct(){
    const createForm = document.getElementById('createForm');
    createForm.addEventListener('submit', (e) => {
        e.preventDefault()
        const formData = new FormData(createForm);
        const productname = formData.get('productname');
        const price = formData.get('price');
        const specification = formData.get('specification');
        const image = document.getElementById('image').files[0];
    
    
         // Create JSON object with form data except image 
         const productData = {
            name: productname,
            price: price,
            specification: specification,
        };
    
        const requestPayLoad = new FormData();
        requestPayLoad.append("productData",JSON.stringify(productData));
        requestPayLoad.append( "image", image);
    
    
        fetch("http://localhost:8080/create",{
         method:"POST",
         body:requestPayLoad
         
         })
         .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Product created successfully:', data);
            window.location.href="admin-home.html";
           
        })
        .catch(error => {
            console.error('Error adding product:', error);
           
        });
    
        })
    } 

   //to delete product from db  
    
function deleteProduct(productId) { 
    if (confirm("Are you sure you want to delete this product?")) {
        fetch(`http://localhost:8080/delete/${productId}`, {
            
            method: "DELETE",
            headers: { 'Content-Type': 'application/json' },
            
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            alert(data.message);
           location.reload() ;
        })
        .catch(error => {
            console.error('Error deleting product:', error);
        });
    }
}

//to perform search functionality

function handleSearch(event){
            event.preventDefault();
            const searchInput = document.querySelector('input[name="search"]');
            const searchQuery = searchInput.value.trim();
            if (searchQuery !== '') {
                performSearch(searchQuery);
            }
        }



function performSearch(searchQuery) {
    fetch(`http://localhost:8080/search/${searchQuery}`)
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(products => {
        displayProducts(products);})
        .catch(error => {
            console.error('Error searching:', error);
        });
}



