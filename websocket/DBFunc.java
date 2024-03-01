package com.websocket;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.sql.Statement;

public class DBFunc {

	public static void insertLogToDb(String data, String tableName) {
		Connection con = DBCon.getInstance().getConnection();

		String sqlQuery = "INSERT INTO " + tableName +"(log_details) VALUES(?);";

		try (PreparedStatement ps = con.prepareStatement(sqlQuery)) {
			
			ps.setString(1, data);
			
			int row = ps.executeUpdate();
			System.out.println(row);

		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

	}
	
	


}
