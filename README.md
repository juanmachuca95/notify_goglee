## Read chat google meet

1. Script JS 
2. Run server on port 8000
3. Create endpoint data/ [POST]

https://github.com/martinlindhe/notify

Probando el correcto funcionamiento con curl:

```curl
curl -d '{"sender_name":"juan", "formatted_timestamp":"12:00", "messages":["hola yo tenia una consulta, no entiendo nada de lo se esta explicando en esta clase, podemos empezar de vuelta?", "quiero el reembolso"]}' http://localhost:8000/data
```

```js
// Script para obtener los mensajes del chat en google met
var totalMessagesEnviados = 0;
var url = 'http://localhost:8000/data';
var iterator = setInterval(function(){
    var actualTotalMessages = 0;

    // Todos los div con los mensajes por usuario
    dataSenders = document.querySelectorAll('[data-sender-name]');
    console.log("Hay ", dataSenders.length, " usuarios");
    for (var i=0; i<dataSenders.length; i++) {
        senderMessages = dataSenders[i].lastChild.querySelectorAll('[data-message-text]');
        actualTotalMessages+= senderMessages.length; // Total de mensajes individuales (son los que se envia en cada enter) 
    }

    console.log("Hay un total de " + actualTotalMessages +  " mensajes ")
    console.log("************** enviados ", totalMessagesEnviados," actuales", actualTotalMessages, "******************")
    if (totalMessagesEnviados < actualTotalMessages) {
        console.log("Hay mensajes "+ (actualTotalMessages - totalMessagesEnviados) +" para enviar")

        var msgArr = []; // array de mensajes
        for (var i=0; i<dataSenders.length; i++) {
            let messages = []
            let messagesSenders = dataSenders[i].lastChild.querySelectorAll('[data-message-text]');

            messagesSenders.forEach(function(ele) {messages.push(ele);})
            var obj = {
                'sender_name': dataSenders[i].getAttribute('data-sender-name'), // Nombre del usuario
                'formatted_timestamp': dataSenders[i].getAttribute('data-formatted-timestamp'), // Hora ej: 00:00
                'messages': messages, // Lista mensajes por usuario
            }
            console.log(obj)
            msgArr.push(obj)
        }


        let count_messages = 0;
        msgArr.forEach(function(obj){
            obj.forEach(function(msg){
                count_messages++;
                if(count_messages >= actualTotalMessages){
                    fetch(url, {
                        mode: 'no-cors',
                        headers: Headers({{ 'Content-Type', ' }})
                    })
                }
            })
        })


    }else{
        console.log("No hay mensajes para enviar")
    }
}, 2000)

```

Para envio de información 
```js
clearInterval(iterator);
```


### Enlaces útlies 

alinear ventanas https://parzibyte.me/blog/2021/08/17/alinear-ventana-tkinter-python/