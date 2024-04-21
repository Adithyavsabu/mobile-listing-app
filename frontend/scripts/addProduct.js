
 
function addProduct(){
const createForm = document.getElementById('createForm');
createForm.addEventListener('submit', (e) => {
    e.preventDefault()
    const formData = new FormData(createForm);
    const productname = formData.get('productname');
    const price = formData.get('price');
    const specification = formData.get('specification');
    const image = document.getElementById('image').files[0];


     // Create JSON object with form data
     const productData = {
        name: productname,
        specification: specification,
    };

    const requestPayLoad = new FormData();
    requestPayLoad.append("productData",JSON.stringify(productData));
    requestPayLoad.append( "image", image);
    requestPayLoad.append( "price", price);
    



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
addProduct();







