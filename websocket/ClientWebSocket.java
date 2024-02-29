package com.websocket;

import java.io.IOException;

import jakarta.websocket.OnClose;
import jakarta.websocket.OnError;
import jakarta.websocket.OnMessage;
import jakarta.websocket.OnOpen;
import jakarta.websocket.Session;

import jakarta.websocket.server.ServerEndpoint;

@ServerEndpoint("/clientws")
public class ClientWebSocket {

	private Session session;

	@OnOpen
	public void onOpen(Session session) {

		this.session = session;

		System.out.println("WebSocket connection opened: " + session.getId());

		Thread senderThread = new Thread(() -> sendDataToClient());
		senderThread.start();

	}

	@OnClose
	public void onClose(Session session) {
		System.out.println("WebSocket connection closed: " + session.getId());
	}

	@OnError
	public void onError(Session session, Throwable throwable) {
		System.out.println("WebSocket error: " + session.getId());
		throwable.printStackTrace();
	}

	@OnMessage
	public void onMessage(String message, Session session) {

		// You can process the received JSON message here

		// Example: Convert JSON message back to object
		/*
		 * ObjectMapper mapper = new ObjectMapper(); try { LogEntry logEntry =
		 * mapper.readValue(message, LogEntry.class); // Process the log entry object as
		 * needed } catch (IOException e) { e.printStackTrace(); }
		 */
	}

	private void sendDataToClient() {
		String queryString = session.getQueryString();
		System.out.println(queryString);
		String paramV = queryString.split("=")[1];

		while (true) {
			try {
				String data = getData(paramV);
				if (data != null) {
					session.getBasicRemote().sendText(data);
				} else {
					// Handle the case where data is null
					System.out.println("No data available for " + paramV);
				}
			} catch (IOException | InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

	private String getData(String paramV) throws InterruptedException, IOException {
		String data = null;

		if (paramV.equals("redis-log")) {
			data = JsonDataQueue.redisDataQ.poll();
		} else if (paramV.equals("njinix_logs")) {
			data = JsonDataQueue.njinixDataQ.poll();
		} else if (paramV.equals("mysql_logs")) {
			data = JsonDataQueue.mysqlDataQ.poll();
		} else if (paramV.equals("dns_logs")) {
			data = JsonDataQueue.dnsDataQ.poll();
		} else if (paramV.equals("kvm_logs")) {
			data = JsonDataQueue.kvmDataQ.poll();
		}

		return data;
	}

}
