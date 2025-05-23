document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("emailForm").addEventListener('submit', function (event) {
        event.preventDefault();
        let url = "https://uo9lbqpwoi.execute-api.eu-central-1.amazonaws.com/newsletter-api/";
        const formData = new FormData(this);
        let jsonData = {};

        formData.forEach((value, key) => {
            jsonData[key] = value;
        });

        let action = event.submitter.value;
        let method = action === "create" ? "POST" : "DELETE";

        fetch(url + action, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData)
        })
            .then(response => {
                if (response.ok) {
                    return response.text();
                } else {
                    throw new Error('Network response was not ok');
                }
            })
            .then(data => alert("Your request was successfully send"))
            .catch(error => {
                console.error('Fetch error:', error);
            });
    });
});
