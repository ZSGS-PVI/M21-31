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
			String data = jsonObject.get("redis_log").toString();
			JsonDataQueue.redisDataQ.add(data);
			DBFunc.insertLogToDb(data, "redis_logs");
		} 
		else if (jsonObject.has("nginx_logs")) {
			System.out.println("Found nginx_logs key:");
			String data = jsonObject.get("nginx_logs").toString();
			JsonDataQueue.nginxDataQ.add(data);
			DBFunc.insertLogToDb(data, "nginx_logs");
		}
		else if (jsonObject.has("mysql_logs")) {
			System.out.println("Found mysql_logs key:");
			String data = jsonObject.get("mysql_logs").toString();
			JsonDataQueue.mysqlDataQ.add(data);
			DBFunc.insertLogToDb(data, "mysql_logs");
		}
		else if (jsonObject.has("dns_logs")) {
			System.out.println("Found dns_logs key:");
			String data = jsonObject.get("dns_logs").toString();
			JsonDataQueue.dnsDataQ.add(data);
			DBFunc.insertLogToDb(data, "dns_logs");
		}
		else if (jsonObject.has("kvm_logs")) {
			System.out.println("Found kvm_logs key:");
			String data = jsonObject.get("kvm_logs").toString();
			JsonDataQueue.kvmDataQ.add(data);
			DBFunc.insertLogToDb(data, "kvm_logs");
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
