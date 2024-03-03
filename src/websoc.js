import axios from 'axios';


let socket = null;
let onDataReceivedCallback = null;

export const connectWebSocket = (callback, params) => {
    onDataReceivedCallback = callback;

    if (!socket) {
        socket = new WebSocket(`ws://localhost:8087/RedisWebSoc/clientws?type=${params}`);

        socket.onopen = () => {
            console.log("Connected to WebSocket");
        };

        socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            onDataReceivedCallback(data);
        };

        socket.onclose = () => {
            console.log("WebSocket connection closed");
            socket = null;
        };

        socket.onerror = (error) => {
            console.error("WebSocket error:", error);
        };
    }
};

export const sendWebSocketMessage = (message) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(message));
    } else {
        console.warn("WebSocket connection is not open.");
    }
};

export const closeWebSocket = () => {
    if (socket) {
        socket.close();
    }
};




export const getLogs = (tableName) => {

    axios.get(`http://localhost:8087/RedisWebSoc/getlogs?table=${tableName}`)
        .then(response => {
            console.log(response.data)
            const dataArray = [];
            dataArray.push(JSON.parse(response.data));
            console.log(dataArray);
            return dataArray;
        })
        .catch(e => {
            console.log(e);
        });
}
