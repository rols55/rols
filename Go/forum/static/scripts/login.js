
document.getElementById('login').addEventListener('click', function (event) {
  if (modal == null) return
  modal.classList.add('active')
  overlay.classList.add('active')
  var registerButton = document.getElementById('registerButton');
  var registerForm = document.getElementById('registerForm');

  registerButton.addEventListener('click', function () {
    if (registerForm.style.display === 'none') {
      registerForm.style.display = 'block';
      registerButton.style.display = 'none'; // Hide the initial register button
    } else {
      registerForm.style.display = 'none';
      registerButton.style.display = 'block'; // Show the initial register button
    }
  });

  loginForm.addEventListener('submit', function (event) {
    event.preventDefault();

    let username = document.querySelector('#loginForm input[type="text"]').value;
    let password = document.querySelector('#loginForm input[type="password"]').value;
    console.log(username, password)
    // Send the username and password to the server for login authentication
    var loginData = {
      username: username,
      password: password
    }
    // Send the login data to the server using Fetch
    fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(loginData)
    })
      .then(function (response) {
        // Handle the response from the server
        if (response.ok) {
          // Login successful
          console.log("Login successful")
          document.getElementById('logout').style.display = 'block'; // Show button 2
          document.getElementById('login').style.display = 'none'; // hide login
          // const usernameHeader = document.getElementById("usernameHeader");
          console.log(username);
          usernameHeader.textContent = username;
          document.getElementById('createPost').style.display = 'block';
          sortBySelect.style.display = "block";
          document.getElementById("myPosts").style.display = "block";
          modal.classList.remove('active')
          overlay.classList.remove('active')
          // Redirect the user to the desired page or perform other actions
          // After a successful login
        } else {
          // Login failed
          // Handle the error, display an error message, etc.
          errorContainer.textContent = 'Oops! Something went wrong. Please try again later.';
        }
      })
      .catch(function (error) {
        // Handle any error that occurred during the request
        console.error('Error:', error);
      });

    this.reset();
  });

  registerForm.addEventListener('submit', function (event) {
    event.preventDefault();

    let email = document.querySelector('#registerForm input[type="email"]').value;
    let username = document.querySelector('#registerForm input[type="text"]').value;
    let password = document.querySelector('#registerForm input[type="password"]').value;
    let repassword = document.getElementById('repassword').value;
    console.log(username, repassword)
    // Send the email, username, and password to the server for user registration
    // You can use AJAX, fetch, or other techniques to send the data to the server
    if (password !== repassword) {
      document.getElementById("repassword").setCustomValidity("Passwords do not match.");
      return
    } else {
      var registerData = {
        email: email,
        username: username,
        password: password
      }

      fetch('/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(registerData)
      })
        .then(function (response) {
          // Handle the response from the server
          if (response.ok) {
            // Registration successful
            console.log("Registration successful")
            // Redirect the user to the desired page or perform other actions
          } else {
            // Registration failed
            // Handle the error, display an error message, etc.
            alert("Username or email is not unique!");
          }
        })
        .catch(function (error) {
          // Handle any error that occurred during the request
          console.error('Error:', error);
        });
      this.reset();
    }
  })
})

document.getElementById('close-button').addEventListener('click', function (event) {
  if (modal == null) return
  modal.classList.remove('active')
  overlay.classList.remove('active')
})

document.getElementById('logout').addEventListener('click', function (event) {
  fetch('/logout', {
    method: 'POST',
    credentials: 'same-origin' // Send cookies along with the request
  })
    .then(response => {
      if (response.ok) {
        // Logout was successful, perform any necessary client-side cleanup
        // or redirect to the login page
        document.getElementById('logout').style.display = 'none'
        document.getElementById('login').style.display = 'block'; // Show button 1
        // const usernameHeader = document.getElementById("usernameHeader");
        usernameHeader.textContent = "";
        document.getElementById('createPost').style.display = 'none';
        sortBySelect.style.display = "none";
        document.getElementById("myPosts").style.display = "none";
      } else {
        // Logout failed, handle the error
        console.error('Logout failed:', response.statusText);
      }
    })
    .catch(error => {
      // An error occurred during the logout process
      console.error('Logout error:', error);
    });
})

