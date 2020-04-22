package work;

import java.sql.Connection;

import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.PreparedStatement;

public class MyConn {
    /**
     * 创建链接数据库属性
     */
    private static final String DRIVER = "com.mysql.jdbc.Driver";//驱动
    private static final String URL = "jdbc:mysql://localhost:3306/book";//数据路路径
    private static final String USENAME = "root";
    private static final String PASSWORD = "";
    /**
     * 得到链接对象
     */
    public Connection getCon(){
        Connection con = null;//与特定数据库的连接（会话）。在连接上下文中执行 SQL 语句并返回结果
        try {
            Class.forName(DRIVER);//链接数据库查找合适的驱动程序
            con = DriverManager.getConnection(URL, USENAME, PASSWORD);
            //用来处理装载驱动程序并且为创建新的数据库连接提供支持；
        } catch (ClassNotFoundException e) {
            // TODO 自动生成的 catch 块
            e.printStackTrace();
        } catch (SQLException e) {
            // TODO 自动生成的 catch 块
            e.printStackTrace();
        }
        return con;
    }
    /**
     *关闭连接
     * @param con
     * @param ps
     * @param rs
     * PreparedStatement预编译的 SQL 语句的对象。
     * ResultSet数据库结果集的数据表，通常通过执行查询数据库的语句生成
     * Connection与特定数据库的连接（会话）。在连接上下文中执行 SQL 语句并返回结果。
     */
    public void closeALL(Connection con,PreparedStatement ps,ResultSet rs){
        try {
            if(rs!=null){
                rs.close();
            }
            if(ps!=null){
                ps.close();
            }
            if(con!=null){
                con.close();
            }
        } catch (SQLException e) {
            // TODO 自动生成的 catch 块
            e.printStackTrace();
        }
    }
}

