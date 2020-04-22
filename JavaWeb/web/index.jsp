<%--
  Created by IntelliJ IDEA.
  User: zpt
  Date: 2019/12/10
  Time: 16:10
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
<form action="empServlet?status=login" method="post">
    <h1>欢迎登录</h1>
    <table border="0" width="30%">
        <tr>
            <td>账号</td>
            <td>
                <input type="text" name="empid" >
            </td>
        </tr>

        <tr>
            <td>密码</td>
            <td>
                <input type="password" name="pwd" >
            </td>
        </tr>

        <tr>
            <td colspan="2">
                <input type="submit" value="提交">
                <input type="reset" value="重置">
            </td>
        </tr>

    </table>
</form>
</body>
</html>