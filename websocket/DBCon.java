package com.websocket;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class DBCon {
    private static DBCon instance;
    private Connection connection;

    private String url = "jdbc:mysql://localhost:3306/m21_31";
	private String username = "root";
	private String password = "Password@123";
    
    private DBCon() {
        try {
            connection = DriverManager.getConnection(url, username, password);
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public static synchronized DBCon getInstance() {
        if (instance == null) {
            instance = new DBCon();
        }
        return instance;
    }

    public Connection getConnection() {
        return connection;
    }
}