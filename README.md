## Noticador de mensajes - google meet 
Para este programa me apoye en utilidades ya programadas como ```github.com/gen2brain/beeep```. Pero tambien es posible hacer utilizando herramientas que provee cada sistema operativo, en linux por ejemplo: ````notify-send````

 
1. Run server on port 8000 (golang). 
```go 
go run main.go  
```

2. Prueba el endpoint data/ [POST]

```curl
curl -d '{"sender_name":"juan", "formatted_timestamp":"12:00", "messages":["Hello", "¿Puedes ves mis mensajes de google meet en otra ventana?"]}' http://localhost:8000/data
```

3. Script JS for google meet.

### Javascript que se ejecuta en google meet
Existen extensiones que permiten guardar el script en el navegador y ejecutarlo cada vez que ingrese a dicho sitio. Pero no lo recomiendo, dado que el proceso que suplanta es solamente copiar y pegar. 

<b>Nota:</b>
* El timer esta puesto a 5s, sientase libre de modificarlo.
* Goole meet puede cambiar el nombre etiquetas con el correr del tiempo, lo que hara que deba actualizarlo manualmente.

```js

// ejecución principal del programa.
var totalMensajesActuales = [];
var myInterval = setInterval(function(){
    let nuevosMensajesActuales = getMessages()

    diff = getDifference(totalMensajesActuales, nuevosMensajesActuales)
    if (diff.length > 0){
        sendToServer(diff)
        totalMensajesActuales = nuevosMensajesActuales 
    }


}, 5000)

// Obtiene los mensajes del panel de chat en google meet
function getMessages(){
    let nuevosMessages = []
    document.querySelectorAll('[data-sender-name]').forEach((element) => {
        let obj = {
            'sender_name':element.getAttribute('data-sender-name'),
            'formatted_timestamp': element.getAttribute('data-formatted-timestamp'),
            'messages': []
        }
        element.lastChild.querySelectorAll('[data-message-text]').forEach((msg) => {
            obj.messages.push(msg.getAttribute('data-message-text'))    
        })
        
        nuevosMessages.push(obj)
    })

    return nuevosMessages
}

// Obtiene los mensajes que no se han enviado aún
function getDifference(actuales, nuevos){
    var mensajesParaEnviar = [];
    // Corroboramos que hayan mensajes iguales entre arrays
    for(var i = 0; i < actuales.length; i++){
        if(actuales[i].messages.length === nuevos[i].messages.length){
            // console.log("mensajes con la misma cantidad ", actuales[i].messages.length, nuevos[i].messages.length);
        }else{
            // console.log("hay mensajes nuevos para enviar")
            // console.log("num: ", i, actuales[i].messages, nuevos[i].messages)
            // Si actual tiene 10 y nuevo tiene 12
            let nuevoObj = {
                'sender_name': actuales[i].sender_name,
                'formatted_timestamp': actuales[i].formatted_timestamp,
                'messages': []
            }

            for (var j = actuales[i].messages.length; j < nuevos[i].messages.length; j++){
                // console.log("Se agregará el mensaje: ", nuevos[i].messages[j])
                nuevoObj.messages.push(nuevos[i].messages[j]);
            }

            mensajesParaEnviar.push(nuevoObj);
        }
    }

    // console.log("Para enviar tenemos ", mensajesParaEnviar);
    // falta agregar los nuevos mensajes de usuarios
    if (actuales.length < nuevos.length) {
        for (var i = actuales.length; i < nuevos.length; i++){
            // console.log("Nuevos mensajes", nuevos[i])
            mensajesParaEnviar.push(nuevos[i])
        }
    }

    return mensajesParaEnviar
}


// Solo envia los mensajes que no se enviado previamente
function sendToServer(mensajesParaEnviar){
    mensajesParaEnviar.forEach((msg) => {
        fetch("http://localhost:8000/data", {
            headers: new Headers({ "content-type": "application/json" }),
            mode: 'no-cors',
            method: "POST",
            body: JSON.stringify(msg)   
        })
    })
}
```