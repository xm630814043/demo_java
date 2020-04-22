<%--
  Created by IntelliJ IDEA.
  User: zpt
  Date: 2019/12/10
  Time: 15:59
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" pageEncoding="utf-8"%>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
    <title>Insert title here</title>
</head>
<body>
<form action="../empServlet?status=add" method="post">
    <table border="0" width="50%">
        <tr>
            <td>编号</td>
            <td><input type="text" name="empid" ></td>
        </tr>
        <tr>
            <td>姓名</td>
            <td><input type="text" name="name" ></td>
        </tr>
        <tr>
            <td>性别</td>
            <td><input type="radio" name="sex" value="0">男
                <input type="radio" name="sex" value="1">女</td>
        </tr>
        <tr>
            <td>密码</td>
            <td><input type="text" name="pwd" ></td>
        </tr>
        <tr>
            <td>电话</td>
            <td><input type="text" name="tel"></td>
        </tr>
        <tr>
            <td>
                <input type="submit" value="提交">
                <input type="reset" value="重置">
            </td>
        </tr>
    </table>
</form>
</body>
</html>