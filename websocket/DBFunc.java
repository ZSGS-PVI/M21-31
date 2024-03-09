package com.websocket;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;

public class DBFunc {

	public static void insertLogToDb(String data, String tableName) {
		Connection con = DBCon.getInstance().getConnection();

		String sqlQuery = "INSERT INTO " + tableName + "(log_details) VALUES(?);";

		try (PreparedStatement ps = con.prepareStatement(sqlQuery)) {

			ps.setString(1, data);

			int row = ps.executeUpdate();
			// System.out.println(row);

		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

	}

	public static List<String> getLogs(String tableName, int offset) {
		Connection con = DBCon.getInstance().getConnection();

		String sqlQ = "SELECT * FROM " + tableName +" order by id desc limit 10 offset ?;";

		List<String> dataList = new ArrayList<>();
		try (PreparedStatement pst = con.prepareStatement(sqlQ)) {

			
			pst.setInt(1, offset);

			ResultSet rs = pst.executeQuery();

			while (rs.next()) {
				dataList.add(rs.getString(2));
			}

		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

		return dataList;
	}

	
}
