

function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "52.90.55.175:8001/users/login/admin" + username);
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  
  xhttp.send(JSON.stringify({
    "id": "-1",
    "username": username,
    "password": password
  }));
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4) {
      console.log('Status:', this.status);
      console.log('Response:', this.responseText);
  
      if (this.status == 200 && this.responseText) {
        const objects = JSON.parse(this.responseText);
        console.log(objects);
  
        if (objects['err']) {
          Swal.fire({
            text: objects["err"],
            icon: 'error',
            confirmButtonText: 'OK'
          });
        } else {
          localStorage.setItem("jwt", "true")
          Swal.fire({
            text: objects['message'],
            icon: 'success',
            confirmButtonText: 'OK'
          }).then((result) => {
            if (result.isConfirmed) {
              window.location.href = './index.html';
            }
          });
        }
      } else {
        console.error('The request did not return a successful response.');
      }
    }
  };
  return false;
}