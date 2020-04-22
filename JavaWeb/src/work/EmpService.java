package work;

import java.util.List;

public interface EmpService {
	/**
	 * 添加员工
	 * @param emp
	 * @return
	 */
	/*
	 * public 是 权限修饰符 java里面的权限修饰符 还有 private protected default
	 * boolean 是java 的一种 数据类型 他的值 只有 true 和 false
	 * Java 一共有 八种基本数据类型（byte int double float short long char boolean） 和 两种引用数据类型。（数组和类）
	 */
	public boolean addemp(Emp emp);
	/**
	 * 修改员工
	 * @param emp
	 * @return
	 */
	public boolean updateemp(Emp emp);
	/**
	 * 删除员工
	 * @param emp
	 * @return
	 */
	public boolean deleteemp(int empid);
	/**
	 * 根据ID查询员工
	 * @param empId
	 * @return
	 * query构造查询对象约束。
	 */
	public Emp queryById(int empid);
	/**
	 * 查询所有员工
	 * @return
	 */
	public List<Emp> queryall();
	/**
	 * 登陆
	 * @param empId
	 * @param pwd
	 * @return
	 * login执行验证。
	 */
	public boolean login(int empid, String pwd);
	/**
	 * 根据名称模糊查询
	 * @param eNanme
	 * @return
	 *query每种约束构造方法仅允许使用适当的类型作为参数。如下例所述，混合调用可以构造约束的随意嵌套：
	 */
	public List<Emp> queryByLike(String eNanme);
	/**
	 * 统计员工个数
	 * @return
	 * getCount返回字符数组中所包含的字符数
	 */
	public int getCount();

}
