package work;

import java.sql.ResultSet;
import java.util.ArrayList;
import java.util.List;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;

public class IEmpDaoImpl implements EmpDao {
	Connection con=null;
	PreparedStatement ps=null;
	ResultSet rs=null;		//用于保存查询所得的结果集
	String sql="";			//输入sql的值
	boolean flag=false;

	@Override
	public boolean addemp(Emp emp) {
		MyConn myconn = new MyConn();
	    sql="insert into emp values(?,?,?,?,?)";
		try {
				con=myconn.getCon();		//用于完成对特定数据库的连接；
				ps=con.prepareStatement(sql);	//用于执行预编译的SQL语句；
				ps.setInt(1,emp.getEmpid());		//调取emp表中的声明属性
				ps.setString(2,emp.getName());
				ps.setInt(3,emp.getSex());
				ps.setString(4, emp.getPwd());
				ps.setString(5,emp.getTel());
				int result=ps.executeUpdate();
				if(result>0){
					flag=true;
				}
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return flag;
	}

	@Override
	public boolean updateemp(Emp emp) {
		MyConn myconn = new MyConn();
	    sql="update emp set empid=?,name=?,sex=?,pwd=?,tel=? where empid=?";
		try {
				con=myconn.getCon();
				ps=con.prepareStatement(sql);
				ps.setInt(1,emp.getEmpid());
				ps.setString(2,emp.getName());
				ps.setInt(3,emp.getSex());
				ps.setString(4, emp.getPwd());
				ps.setString(5,emp.getTel());
				ps.setInt(6,emp.getEmpid());
				int result=ps.executeUpdate();
				if(result>0){
					flag=true;
			}
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return flag;
	}

	@Override
	public boolean deleteemp(int empid) {
		MyConn myconn = new MyConn();
	    sql="delete from emp where empid=?";
		try {
			con=myconn.getCon();
			ps=con.prepareStatement(sql);
			ps.setInt(1,empid);
			int result=ps.executeUpdate();
			if(result>0){
				flag=true;
			}
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return flag;
	}

	@Override
	public Emp queryById(int empid) {
		MyConn myconn = new MyConn();
		Emp emp=new Emp();
	    sql="select * from emp where empid=?";
		try {
			con=myconn.getCon();
			ps=con.prepareStatement(sql);
			ps.setInt(1,empid);
			rs=ps.executeQuery();
			if(rs.next()){
				emp.setEmpid(rs.getInt(1));
				emp.setName(rs.getString(2));
				emp.setSex(rs.getInt(3));
				emp.setPwd(rs.getString(4));
				emp.setTel(rs.getString(5));
			}	
				
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return emp;
	}

	@Override
	public List<Emp> queryall() {
		List<Emp> list=new ArrayList<Emp>();
		MyConn myconn = new MyConn();
	    sql="select * from emp ";
		try {
			con=myconn.getCon();
			ps=con.prepareStatement(sql);
			rs=ps.executeQuery();
			while(rs.next()){
				Emp emp=new Emp();
				emp.setEmpid(rs.getInt(1));
				emp.setName(rs.getString(2));
				emp.setSex(rs.getInt(3));
				emp.setPwd(rs.getString(4));
				emp.setTel(rs.getString(5));
				list.add(emp);
			}	
				
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return list;
	}

	@Override
	public boolean login(int empid, String pwd) {
		MyConn myconn = new MyConn();
		sql="select * from emp where empid=? and pwd=?";
//	    sql="delete from Emp where empid=? and pwd=?";
		try {
			con=myconn.getCon();
			ps=con.prepareStatement(sql);
			ps.setInt(1,empid);
			ps.setString(2,pwd);
			rs=ps.executeQuery();
//			int result=ps.executeUpdate();
			if(rs.next()){
				flag=true;
			}
		} catch (SQLException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}finally{
			myconn.closeALL(con, ps, rs);
		}
		return flag;
	}
	@Override
	public List<Emp> queryByLike(String eName) {
		List<Emp> list = new ArrayList<Emp>();
		MyConn mycon = new MyConn();
		sql="select * from emp where  name  like ?";
		try {
			con = mycon.getCon();//链接数据路
			ps = con.prepareStatement(sql);//预编译SQL语句
			ps.setString(1, "%"+eName+"%");//定义条件
			rs = ps.executeQuery();//执行SQL，并返回结果值
			while(rs.next()){
				Emp emp = new Emp();
				emp.setEmpid(rs.getInt(1));
				emp.setName(rs.getString(2));
				emp.setSex(rs.getInt(3));
				emp.setPwd(rs.getString(4));
				emp.setTel(rs.getString(5));
				list.add(emp);
			}
		} catch (SQLException e) {
			// TODO 自动生成的 catch 块
			e.printStackTrace();
		}finally{
			mycon.closeALL(con, ps, rs);
		}
		// TODO 自动生成的方法存根
		return list;
	}
	/**
	 * ResultSet 对象具有指向其当前数据行的光标。
	 * 最初，光标被置于第一行之前。next 方法将光标移动到下一行；
	 * 因为该方法在 ResultSet 对象没有下一行时返回 false，
	 * 所以可以在 while 循环中使用它来迭代结果集。
	 */
	@Override
	public int getCount() {
		int count=0;
			MyConn myco = new MyConn();
		sql="select count(*) from emp";
		try {
			con = myco.getCon();//链接数据库
			ps =con.prepareStatement(sql);//用于执行预编译的SQL语句
			rs = ps.executeQuery();//执行 SQL 查询，并返回该查询生成的 ResultSet 对象。
			if(rs.next()){
				count =rs.getInt(1);
			}
		} catch (SQLException e) {
			// TODO 自动生成的 catch 块
			e.printStackTrace();
		}finally{
			myco.closeALL(con, ps, rs);
		}
		// TODO 自动生成的方法存根
		return count;
	}
}
