document.getElementById("emailForm").addEventListener('submit', function (event) {
    event.preventDefault();
    let url = "https://5rrn370t09.execute-api.eu-central-1.amazonaws.com/default/test-cors";
    const formData = new FormData(this)
    let action = document.activeElement.value
    let method = action === "create" ? "POST" : "DELETE"
    fetch(url + action, {
        method: method,
        body: formData
    })
        .then(r => r.text())
        .then(data => console.log(data))
        .catch(error => console.log(error));
})