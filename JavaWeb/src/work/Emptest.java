package work;

import java.util.ArrayList;
import java.util.List;

public class Emptest {
    public static void main(String[] args) {
        /**
         *登陆
         */


//        EmpDao empDao4 = new IEmpDaoImpl();
//        boolean flag = empDao4.login(3, "111");
//        System.out.println(flag);

        /**
         *添加
         */
        Emp emp = new Emp(2, "李四", 1, "111", "110");
        EmpDao empDao = new IEmpDaoImpl();
        boolean emps1 = empDao.addemp(emp);
        System.out.println(emps1);
//
//        /**
//         *修改
//         */
//
//        Emp emp7 = new Emp(2, "王语嫣", 0, "456", "120");
//        EmpDao empDao7 = new IEmpDaoImpl();
//        boolean emps2 = empDao7.updateemp(emp7);
//        System.out.println(emps2);
//
//        /**
//         * 删除
//         */
//
//        EmpDao empDao6 = new IEmpDaoImpl();
//        boolean emps3 = empDao6.deleteemp(2);
//        System.out.println(emps3);
//
//        /**
//         * 统计员工个数
//         */
//
//        EmpDao empDao5 = new IEmpDaoImpl();
//        int count = empDao5.getCount();
//        System.out.println(count);
//
//
//
//        /**
//         * 通过ID查员工
//         */
//        Emp emp3 = new Emp();
//        EmpDao empDao3 = new IEmpDaoImpl();
//        emp3 = empDao3.queryById(3);
//        System.out.println(emp3.getName());
//        /**
//         * 查询所有员工
//         */
//
//        EmpDao empDao1 = new IEmpDaoImpl();
//        List<Emp> list1 = new ArrayList<Emp>();
//        list1 = empDao1.queryall();
//        for (Emp emp1 : list1) {
//            System.out.println(emp1.getPwd());
//        }
//        /**
//         * 模糊查询
//         */
//        EmpDao empDao2 = new IEmpDaoImpl();
//        List<Emp> list2 = new ArrayList<Emp>();
//        list2 = empDao2.queryByLike("段");
//        for (Emp emp2 : list2) {
//            System.out.println(emp2.getName());
//        }
    }
}

