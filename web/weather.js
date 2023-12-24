var jwt = localStorage.getItem("jwt");
if (jwt != null) {
  window.location.href = './index.html'
}

function getWeather() {
  const zipCode = document.getElementById("zipcode").value;
  const country = document.getElementById("countries").value;
  const tempUnit = document.getElementById("tempUnit").value;


  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://127.0.0.1:8001/weather");
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  
  xhttp.send(JSON.stringify({
    "zipCode": zipCode,
    "country": country,
    "tempType": tempUnit
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
        localStorage.setItem("temp", objects["main"]["temp"] )
        localStorage.setItem("minTemp", objects["main"]["temp_min"] )
        localStorage.setItem("maxTemp", objects["main"]["temp_max"] )
        localStorage.setItem("name", objects["name"] )
        Swal.fire({
            text: objects['message'],
            icon: 'success',
            confirmButtonText: 'OK'
          }).then((result) => {
            if (result.isConfirmed) {
              window.location.href = './weather_result.html';
            }
          });
      }
    }
  };
  return false;
}