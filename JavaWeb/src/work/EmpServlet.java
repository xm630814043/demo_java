package work;

import java.io.IOException;
import java.util.List;
import javax.servlet.ServletException;
import javax.servlet.annotation.WebInitParam;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * Servlet implementation class EmpServlet
 *
/*
extends 是继承父类，只要那个类不是声明final或者定义为abstract就能继承,
JAVA中不支持多重继承，继承只能继承一个类，
但implements可以实现多个接口，用逗号分开就行了
 */
//@WebServlet(urlPatterns="/EmpServlet",initParams={@WebInitParam(name="charset",value="utf-8")})
//@WebServlet("/empServlet")
public class EmpServlet extends HttpServlet {
	private static final long serialVersionUID = 1L;
    /**
     * @see HttpServlet#HttpServlet()
     */
    public EmpServlet() {
        super();
        // TODO Auto-generated constructor stub
    }
	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		doPost(request, response);
	}
	/**
	 * @see HttpServlet#doPost(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		request.setCharacterEncoding("utf-8");
		EmpService empsevcice = new InEmpServiceImpl();
		/*
		 * request.getParameter,获取页面输入的内容
		 * */
        String status=request.getParameter("status");
        boolean flag=false;
        if("login".equals(status)){
            int empid=Integer.parseInt(request.getParameter("empid"));
            String pwd=request.getParameter("pwd");
            flag=empsevcice.login(empid, pwd);
            Emp emp=empsevcice.queryById(empid);
            if(flag){
                request.setAttribute("name", emp.getName());
                request.getRequestDispatcher("store.jsp").forward(request, response);
//				response.sendRedirect("index.jsp");
            }else{
                response.sendRedirect("index.jsp");
            }
        }
		else if("add".equals(status)){
			int empid=Integer.parseInt(request.getParameter("empid"));
			String name=request.getParameter("name");
			int sex=Integer.parseInt(request.getParameter("sex"));
			String pwd=request.getParameter("pwd");
			String tel=request.getParameter("tel");
			Emp emp = new Emp(empid, name, sex, pwd, tel);
			flag=empsevcice.addemp(emp);
			if(flag) {
				response.sendRedirect("empServlet?status=queryAll");
			}
			else{
				response.sendRedirect("emp/lista.jsp");
			}
		}
		else if("queryAll".equals(status)){
			List<Emp> list=empsevcice.queryall();
			request.setAttribute("list",list);
			request.getRequestDispatcher("emp/lista.jsp").forward(request, response);
		}
		else if("queryById".equals(status)){
			int empid=Integer.parseInt(request.getParameter("empid"));
			Emp emp = new Emp();
			emp=empsevcice.queryById(empid);
			request.setAttribute("lista",emp);
			request.getRequestDispatcher("emp/update.jsp").forward(request, response);
		}
		else if("update".equals(status)){
			int empid=Integer.parseInt(request.getParameter("empid"));
			String name=request.getParameter("name");
			int sex=Integer.parseInt(request.getParameter("sex"));
			String pwd=request.getParameter("pwd");
			String tel=request.getParameter("tel");
			Emp emp = new Emp(empid, name, sex, pwd, tel);
			flag=empsevcice.updateemp(emp);
			if(flag) {
				response.sendRedirect("empServlet?status=queryAll");
			}
			else{
				response.sendRedirect("emp/lista.jsp");
			}
		}
		else if("delete".equals(status)){
			int empid=Integer.parseInt(request.getParameter("empid"));
			flag=empsevcice.deleteemp(empid);
			if(flag) {
				response.sendRedirect("empServlet?status=queryAll");
			}
			else{
				response.sendRedirect("emp/lista.jsp");
			}
		}
	}
}
