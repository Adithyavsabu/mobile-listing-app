//script to perform edit functionality including fetching data from the backend and also uodating data in the backend 
function editProduct() {
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');   // Get the value of the 'id' parameter from the URL
        const editForm = document.getElementById('editForm');

        // Fetch product details using productId
        fetch(`http://localhost:8080/list/${productId}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(product => {
            
                document.getElementById('productname').value = product.name;
                document.getElementById('price').value = product.price;
                document.getElementById('specification').value = product.specification;

            
                editForm.addEventListener('submit', (e) => {
                    e.preventDefault();
 
                    const formData = new FormData(editForm);
                    const productName = formData.get('productname');
                    const price = formData.get('price');
                    const specification = formData.get('specification');
                    const image = document.getElementById('image').files[0];

                    // Create JSON object literal with form data
                    const productData = {
                        name: productName,
                        specification: specification,
                    };

                    const requestPayload = new FormData();
                    requestPayload.append("productData", JSON.stringify(productData));

                    if (image) {
                        requestPayload.append("image", image);
                    }
                    requestPayload.append("price", price);

                    fetch(`http://localhost:8080/update/${productId}`, {
                        method: "PUT",
                        body: requestPayload
                    })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Network response was not ok');
                        }
                        return response.json();
                    })
                    .then(data => {
                        console.log('Product updated successfully:', data);
                        window.location.href = "admin-home.html";
                    })
                    .catch(error => {
                        console.error('Error updating product:', error);
                    });
                });
            })
            .catch(error => {
                console.error('Error fetching product details:', error);
            });
}
editProduct();