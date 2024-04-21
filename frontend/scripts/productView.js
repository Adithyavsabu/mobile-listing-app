function viewProduct(){
    const urlParams = new URLSearchParams(window.location.search);
     const productId = urlParams.get('id');   // Get the value of the 'id' parameter from the URL
    fetch(`http://localhost:8080/list/${productId}`)
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data=>displaySingleProduct(data))
    .catch(error => {
        console.error('Error fetching product details:', error)});

}

function displaySingleProduct(product){
    const container= document.querySelector('.product-details');
    
    let output=``;
    output+=`   <div class="product-img"><img src="${product.image_path}"></div>
    <div class="product-info">
        <h1>${product.name}</h1>
        <p>Rs ${product.price}</p>
        <div class="product-specs">
            <div class="spec-item">
            ${product.specification}
            </div>
        </div>        </div>`;
       
    container.innerHTML=output;
}
viewProduct();
