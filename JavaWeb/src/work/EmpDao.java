package work;

import java.util.List;

public interface EmpDao {
	/**
	 * 添加员工
	 * @param emp
	 * @return
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
	 * 统计员工个数
	 * @return
	 * getCount返回字符数组中所包含的字符数
	 */
	public int getCount();
	/**
	 * 根据名称模糊查询
	 * @param eNanme
	 * @return
	 *query每种约束构造方法仅允许使用适当的类型作为参数。如下例所述，混合调用可以构造约束的随意嵌套：
	 */
	public List<Emp> queryByLike(String eNanme);
}
