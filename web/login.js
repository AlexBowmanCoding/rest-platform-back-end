

function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://127.0.0.1:8001/users/login/" + username);
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  
  xhttp.send(JSON.stringify({
    "id": "-1",
    "username": username,
    "password": password
  }));
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4) {
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
    }
  };
  return false;
}