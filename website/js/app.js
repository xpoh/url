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

    postData('http://localhost/butty/', { "url": document.getElementById('longLink').value, "PostCount":0, "GetCount":0 })
    // postData('https://url-butty.herokuapp.com/butty/', { "url": document.getElementById('longLink').value })
        .then((data) => {
            document.getElementById('longLink').value = data.url;
            document.getElementById('postCount').value = data.PostCount;
            document.getElementById('getCount').value = data.GetCount;

        });
}
const form = document.getElementById('urlSenderForm');
form.addEventListener('submit', logSubmit);