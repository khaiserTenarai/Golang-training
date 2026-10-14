fetch("/api/hello")
  .then(function (response) { return response.json(); })
  .then(function (data) {
    document.getElementById("message").textContent = data.message;
  });
