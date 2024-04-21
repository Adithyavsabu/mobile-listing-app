//onclick submit button of login
const form = document.getElementById('loginForm');
const errorDiv = document.getElementById("error");
form.addEventListener("submit", function(e){
    e.preventDefault(); 

    
const formData = new FormData(form);
const username = formData.get('username');
const password = formData.get('password');

   // Perform client-side validation (e.g., check if fields are not empty)
 if (!username || !password) {
    errorDiv.textContent = "Please enter both username and password.";
    return;
}
console.log(`Username: ${username}, Password: ${password}`);

    fetch("http://localhost:8080/UserValidation",{
method: "POST",
body: JSON.stringify({"username" : username, "password": password})

    })
    .then(function(response){
    if(response.ok) { 
        console.log("Succesful login");
        return response.json()}

    else {
        throw new Error('Server Response not OK')}

})
    .then(function(data){console.log(data);
        if(data=="admin"){
            window.location.href="admin-home.html";
        }
        else if(data=="customer"){
            window.location.href="home.html";
        }
    })
    .catch(function (error) {
        console.log('Could not post data ', error);
    });
})







