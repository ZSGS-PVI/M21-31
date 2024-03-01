package com.websocket;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class JsonDataQueue {
	static BlockingQueue<String> redisDataQ = new LinkedBlockingQueue<>();
	static BlockingQueue<String> dnsDataQ = new LinkedBlockingQueue<>();
	static BlockingQueue<String> nginxDataQ = new LinkedBlockingQueue<>();
	static BlockingQueue<String> kvmDataQ = new LinkedBlockingQueue<>();
	static BlockingQueue<String> mysqlDataQ = new LinkedBlockingQueue<>();

}
