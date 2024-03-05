package com.websocket;

import java.io.IOException;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.annotation.WebFilter;
import jakarta.servlet.http.HttpFilter;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;



@WebFilter(urlPatterns = "/*")
public class CorsFilter extends HttpFilter {
 
    private static final long serialVersionUID = 1085085035971195525L;

	@Override
    public void doFilter(HttpServletRequest request, HttpServletResponse response, FilterChain chain) throws IOException, ServletException {
        // Allow requests from all origins
        
    	System.out.println("filter called1!");
    	response.setHeader("Access-Control-Allow-Origin", "*");
        
        // Allow specific HTTP methods
        response.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS");
        
        // Allow specific headers
        response.setHeader("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
        
        // Allow credentials
        response.setHeader("Access-Control-Allow-Credentials", "true");

        // Continue the filter chain
        chain.doFilter(request, response);
    }
}
