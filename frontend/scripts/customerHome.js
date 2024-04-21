//includes scripts that are related to the components in the customer home page 
//fetch product list from the db
document.addEventListener("DOMContentLoaded", function() {
    if (window.location.pathname.includes("home.html")) {
        fetchProducts();
        document.querySelector(".search-box form").addEventListener("submit", handleSearch);  
    }

});
function fetchProducts(){
    fetch("http://localhost:8080/list")
    .then(response=>{return response.json()})
    .then(data => {displayProducts(data);console.log(data)})
    .catch((err)=>console.log("Could not fetch details of products",err));
}

// to display thr fetched products
function displayProducts(productList) {
    const productContainer = document.querySelector(".product");
    let output = "";
    for (let i = 0; i < productList.length; i++) {
        let product = productList[i];
        console.log(product.image)
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
            </div>
        </div>
    </div>   `;
    }
    productContainer.innerHTML = output;
}

function handleSearch(event){

    event.preventDefault();
    const searchInput = document.querySelector('input[name="search"]');
    const searchQuery = searchInput.value.trim();
    if (searchQuery !== '') {
        // Call a function to handle the search action
        performSearch(searchQuery);
    }
}



function performSearch(searchQuery) {
// You can make an AJAX request to your Go backend here
// Example: Fetch data from the backend using the searchQuery
fetch(`http://localhost:8080/search/${searchQuery}`)
.then(response => {
if (!response.ok) {
    throw new Error('Network response was not ok');
}
return response.json();
})
.then(products => {
// Handle the search results
displayProducts(products);})
.catch(error => {
    console.error('Error searching:', error);
});
}





