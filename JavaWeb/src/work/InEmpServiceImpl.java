package work;

import java.util.List;

public class InEmpServiceImpl implements EmpService {
     EmpDao  ed=new IEmpDaoImpl();
	@Override
	public boolean addemp(Emp emp) {
		// TODO Auto-generated method stub
		return ed.addemp(emp);
	}

	@Override
	public boolean updateemp(Emp emp) {
		// TODO Auto-generated method stub
		return ed.updateemp(emp);
	}

	@Override
	public boolean deleteemp(int empid) {
		// TODO Auto-generated method stub
		return ed.deleteemp(empid);
	}

	@Override
	public Emp queryById(int empid) {
		// TODO Auto-generated method stub
		return ed.queryById(empid);
	}

	@Override
	public List<Emp> queryall() {
		// TODO Auto-generated method stub
		return ed.queryall();
	}

	@Override
	public boolean login(int empid, String pwd) {
		// TODO Auto-generated method stub
		return ed.login(empid, pwd);
	}
	@Override
	public List<Emp> queryByLike(String eNanme) {
		// TODO Auto-generated method stub
		return ed.queryByLike(eNanme);
	}

	@Override
	public int getCount() {
		// TODO Auto-generated method stub
		return ed.getCount();
	}

}
