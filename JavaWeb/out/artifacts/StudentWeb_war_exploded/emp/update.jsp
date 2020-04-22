<%@ page language="java" contentType="text/html; charset=utf-8"
         pageEncoding="utf-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Insert title here</title>
</head>
<body>
<form action="../empServlet?status=update" method="post">
    <table border="0" width="50%">
        <tr>
            <td>编号</td>
            <td><input type="text" name="empid" value="${emp.empid}" ></td>
        </tr>
        <tr>
            <td>姓名</td>
            <td><input type="text" name="name" value="${emp.name}" ></td>
        </tr>
        <tr>
            <td>性别</td>
            <td><input type="radio" name="sex" value="0" ${emp.sex=='0'?"checked=\"checked\"":""} >男
                <input type="radio" name="sex" value="1" ${emp.sex=='1'?"checked=\"checked\"":""} >女</td>
        </tr>
        <tr>
            <td>密码</td>
            <td><input type="text" name="pwd" value="${emp.pwd}" ></td>
        </tr>
        <tr>
            <td>电话</td>
            <td><input type="text" name="tel" value="${emp.tel}" ></td>
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