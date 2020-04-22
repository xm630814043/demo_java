<%--
  Created by IntelliJ IDEA.
  User: zpt
  Date: 2019/12/9
  Time: 16:48
  To change this template use File | Settings | File Templates.
--%>
<%@ page language="java" contentType="text/html; charset=utf-8"
         pageEncoding="utf-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
  <title>Insert title here</title>
</head>
<body>
欢迎<%=request.getAttribute("name") %>登录！
<table border="0" width="50%">
  <tr>
    <td><a href="emp/add.jsp">添加</a></td>
  </tr>
  <tr>
    <td><a href="empServlet?status=queryAll" target="main">查找</a></td>
  </tr>
</table>
</body>
</html>