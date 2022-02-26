function logSubmit(event) {
    event.preventDefault();

    async function postData(url = '', data = {}) {
        const response = await fetch(url, {
            method: 'POST',
            cache: 'no-cache',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        });
        return await response.json();
    }

    postData('https://url-butty.herokuapp.com/butty/', { "url": document.getElementById('longLink').value })
        .then((data) => {
            document.getElementById('longLink').value = data.url;
        });
}
const form = document.getElementById('urlSenderForm');
form.addEventListener('submit', logSubmit);