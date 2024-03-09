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




export const getLogs = async (tableName, offset) => {

    try {
        const response = await axios.get(`http://localhost:8087/RedisWebSoc/getlogs?table=${tableName}&offset=${offset}`);
        const data = await response.data;
        //console.log(data)
        return data;
    } catch (error) {
        console.error(error);
        throw error; // re-throw the error to propagate it to the caller
    }

}

