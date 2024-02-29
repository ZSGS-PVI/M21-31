package com.websocket;

import com.google.gson.Gson;
import com.google.gson.JsonObject;

import jakarta.websocket.OnClose;
import jakarta.websocket.OnError;
import jakarta.websocket.OnMessage;
import jakarta.websocket.OnOpen;
import jakarta.websocket.Session;
import jakarta.websocket.server.ServerEndpoint;

@ServerEndpoint("/serverws")
public class ServerWebSocket {

	@OnOpen
	public void onOpen(Session session) {
		System.out.println("write web socket con called!!");

	}

	@OnMessage
	public void onMessage(String message, Session session) throws Exception {
		addDataToQueue(message);
	}

	private void addDataToQueue(String message) {
		Gson gson = new Gson();

		JsonObject jsonObject = gson.fromJson(message, JsonObject.class);

		if (jsonObject.has("redis_log")) {
			System.out.println("Found redis_log key:");
			JsonDataQueue.redisDataQ.add(jsonObject.get("redis_log").toString());
		} 
		else if (jsonObject.has("njinix_logs")) {
			System.out.println("Found njinix_logs key:");
			JsonDataQueue.njinixDataQ.add(jsonObject.get("njinix_logs").toString());
		}
		else if (jsonObject.has("mysql_logs")) {
			System.out.println("Found mysql_logs key:");
			JsonDataQueue.mysqlDataQ.add(jsonObject.get("mysql_logs").toString());
		}
		else if (jsonObject.has("dns_logs")) {
			System.out.println("Found dns_logs key:");
			JsonDataQueue.dnsDataQ.add(jsonObject.get("dns_logs").toString());
		}
		else if (jsonObject.has("kvm_logs")) {
			System.out.println("Found kvm_logs key:");
			JsonDataQueue.kvmDataQ.add(jsonObject.get("kvm_logs").toString());
		}
		else {
			System.out.println("Expected logs key not found");
		}

	}

	@OnClose
	public void onClose(Session session) {
		// Handle closing of writer WebSocket connection
	}

	@OnError
	public void onError(Session session, Throwable throwable) {
		// Handle errors on writer WebSocket connection
	}
}
