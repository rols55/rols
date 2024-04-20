export default class Socket {
    constructor(name, url) {
        this.socket = new WebSocket(url ? url : `ws://localhost:8080/ws?name=${name}`);
        this.socket.onopen = (e) => this.handleOnOpen(e);
        this.socket.onclose = (e) => this.handleOnClose(e);
        this.socket.onerror = (e) => this.handleOnError(e);
        this.socket.onmessage = (e) => {
            var data = JSON.parse(e.data)
            console.log("WebSocket received:",data);
            switch (data.type) {
                case "message":
                    this.handleOnMessage(data)
                    break;
                case "connected":
                    this.handleOnConnected(data);
                    break;
                case "disconnected":
                    this.handleOnDisconnected(data);
                    break;
                case "move":
                    this.handleOnMove(data);
                    break;
                case "BOMB":
                    this.handleOnBomb(data);
                    break;
                case "update":
                    this.handleOnUpdate(data);
                case "timer":
                    this.handleOnTimer(data);
                    break;
                case "pupdate":
                    this.handleOnPlayerUpdate(data);
                    break;
                case "start": 
                    this.handleOnStart(data);
                    break;
                default:
                    this.handleOnDefault(data);
            }
        }
        this.handleOnStart = (data) => {
            console.log("handleOnStart recived data:", data);
        }
        this.handleOnPlayerUpdate = (data) => {
            console.log("handleOnPlayerUpdate recived data:", data);
        }

        this.handleOnUpdate = (data) => {
            console.log("handleOnUpdate recived data:", data);
        }

        this.handleOnTimer = (data) => {
            console.log("handleOnUpdate recived data:", data);
        }
        
        this.handleOnBomb = (data) => {
            console.log("handleOnBomb recived data:", data);
        }

        this.handleOnMove = (data) => {
            console.log("handleOnMove recived data:", data);
        }

        this.handleOnOpen = () => {
            console.log("Connected to WebSocket server");
        }

        this.handleOnClose = () => {
            console.log("Disconnected from WebSocket server");
        }

        this.handleOnMessage = (data) => {
            console.log("handleOnMessage received data:", data);
        }

        this.handleOnConnected = (data) => {
            console.log("handleOnConnected received data:", data);
        }

        this.handleOnDisconnected = (data) => {
            console.log("handleOnDisconnected received data:", data);
        }

        this.handleOnGetOnline = (data) => {
            console.log("handleOnGetOnline received data:", data);
        }

        this.handleOnDefault = (data) => {
            console.log("handleOnDefault received data:", data);
        }

        this.handleOnError = (error) => {
            console.log("WebSocket Error: ", error);
        };
    }

    send(data) {
        if (this.socket.readyState === WebSocket.OPEN) {
            console.log("Sending data: " + JSON.stringify(data));
            const sendData = JSON.stringify(data);
            this.socket.send(sendData);
        } else {
            console.log(`Unable to send data: ${data}, WebSocket ready state: ${this.socket.readyState}`);
        }
    }

    close() {
        this.socket.close();
    }
}