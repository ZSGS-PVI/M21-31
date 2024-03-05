package com.websocket;

import jakarta.servlet.ServletException;
import jakarta.servlet.annotation.WebServlet;
import jakarta.servlet.http.HttpServlet;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.List;

import com.google.gson.Gson;

@WebServlet("/getlogs")
public class GetLogs extends HttpServlet {
	private static final long serialVersionUID = 1L;

	protected void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		String tableName = request.getParameter("table");

		List<String> logs = DBFunc.getLogs(tableName);

		Gson gson = new Gson();

		String jsonData = gson.toJson(logs);

		response.setContentType("application/json");

		PrintWriter out = response.getWriter();
		System.out.println(jsonData);
		out.println(jsonData);

	}

}
