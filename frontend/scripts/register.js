function registerUser() {
    const registerForm = document.getElementById('registerform');
registerForm.addEventListener('submit', (e) => {
    e.preventDefault()
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
 
    const user = {
        username: username,
        password: password
    };
    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(user)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to register user');
        }
        return response.json();
    })
    .then(data => {
        console.log('User registered successfully:', data);
        window.location.href = "login.html";
    })
    .catch(error => {
        console.error('Error registering user:', error);
   
    });
    });
}

registerUser();