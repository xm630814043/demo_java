<%--
  Created by IntelliJ IDEA.
  User: zpt
  Date: 2019/12/10
  Time: 16:03
  To change this template use File | Settings | File Templates.
--%>
<%@page import="work.Emp"%>
<%@page import="work.EmpDao"%>
<%@page import="java.util.ArrayList"%>
<%@page import="java.util.*"%>
<%@ page import="work.IEmpDaoImpl" %>
<%@ page contentType="text/html;charset=UTF-8" language="java" pageEncoding="utf-8"%>
html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
    <title>Insert title here</title>
</head>
<body>
<script type="text/javascript">
    function check(empid){
        var flag=confirm("您确定要删除吗?");
        if(flag){
            location.href="../empServlet?status=delete&empid="+empid;
        }
    }

</script>
<table border="0" width="50%">
    <tr>
        <th>编号</th>
        <th>姓名</th>
        <th>性别</th>
        <th>密码</th>
        <th>电话</th>
        <th>操作</th>
    </tr>
    <%
        EmpDao empdao=new IEmpDaoImpl();
        List<Emp> empslista = empdao.queryall();
        for( Emp emp : empslista ){
    %>
        <tr>
            <td><%= emp.getEmpid() %></td>
            <td><%= emp.getName() %></td>
            <td><%= emp.getSex() == 0 ? "男":"女"%></td>
            <td><%= emp.getPwd() %></td>
            <td><%= emp.getTel() %></td>
            <td>
                <a href="../empServlet?status=queryById&empid=<%=emp.getEmpid()%>">编辑</a>
                <a href="javascript:check(<%=emp.getEmpid()%>)">删除</a>
            </td>
        </tr>
    <%
        }
    %>
</table>
</body>
</html>
