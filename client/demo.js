/* eslint-env browser */

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

let pc = new RTCPeerConnection({
    iceServers: [
        {
            urls: 'stun:84.201.128.121:3478'
        },
        {
            urls: 'stun:stun.l.google.com:19302'
        }
    ]
})
let log = msg => {
    document.getElementById('logs').innerHTML += msg + '<br>'
}

console.log("here")

let sendChannel = pc.createDataChannel('foo')
sendChannel.onclose = () => console.log('sendChannel has closed')
sendChannel.onopen = () => console.log('sendChannel has opened')
sendChannel.onmessage = e => log(`Message from DataChannel '${sendChannel.label}' payload '${e.data}'`)

pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
pc.onicecandidate = async event => {
    // for (; ;) {
    if (event.candidate === null) {

        let socket = new WebSocket("ws://localhost:8100");
        // let socket = new WebSocket("ws://84.201.181.0:8100");

        socket.onopen = function (e) {
            console.log('Connected to server!')
            socket.send(JSON.stringify(pc.localDescription))
        };
        //for (; ;) {
        console.log("Here and Now!")
        socket.onmessage = function (event) {
            console.log('data from server: ', event.data)
        };


        socket.onerror = function (error) {
            console.log('Error is: ', error.message)
        };

        // document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
        // await sleep(1000);
        //}

    }
    // await sleep(3000);
    // }

}

pc.onnegotiationneeded = e =>
    pc.createOffer().then(d => pc.setLocalDescription(d)).catch(log)

window.sendMessage = () => {
    let message = document.getElementById('message').value
    if (message === '') {
        return alert('Message must not be empty')
    }

    sendChannel.send(message)
}

window.startSession = () => {
    let sd = document.getElementById('remoteSessionDescription').value
    if (sd === '') {
        return alert('Session Description must not be empty')
    }

    try {
        pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
    } catch (e) {
        alert(e)
    }
}
