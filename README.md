## Read chat google meet

1. Script JS 
2. Run server on port 8000
3. Create endpoint data/ [POST]

```js
// Script para obtener los mensajes del chat en google met
dataSenders = document.querySelectorAll('[data-sender-name]');
for (var i=0; i<dataSenders.length; i++) {
    var msgArr = [];
    dataSenders[i].lastChild.querySelectorAll('[data-message-text]').forEach(function(ele) {
        msgArr.push(ele.innerHTML);
    });

    var obj = {
        'sender_name': dataSenders[i].getAttribute('data-sender-name'),
        'formatted_timestamp': dataSenders[i].getAttribute('data-formatted-timestamp'),
        'messages': msgArr 
    }

    console.log(JSON.stringify(obj))
    var url = 'http://localhost:8000/data';
    fetch(url, {
        headers: new Headers({ "content-type": "application/json" }),
        mode: 'no-cors',
        method: "POST",
        body: JSON.stringify(obj)
    })
}

```