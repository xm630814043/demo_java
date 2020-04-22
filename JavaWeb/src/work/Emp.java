package work;

public class Emp {
	private int empid;
	private String name;
	private int sex;
	private String pwd;
	private String tel;
	
	public Emp() {
		// TODO Auto-generated constructor stub
	}

	public Emp(int empid, String name, int sex, String pwd, String tel) {
		this.empid = empid;
		this.name = name;
		this.sex = sex;
		this.pwd = pwd;
		this.tel = tel;
	}

	public int getEmpid() {
		return empid;
	}

	public String getName() {
		return name;
	}

	public int getSex() {
		return sex;
	}

	public String getPwd() {
		return pwd;
	}

	public String getTel() {
		return tel;
	}

	public void setEmpid(int empid) {
		this.empid = empid;
	}

	public void setName(String name) {
		this.name = name;
	}

	public void setSex(int sex) {
		this.sex = sex;
	}

	public void setPwd(String pwd) {
		this.pwd = pwd;
	}

	public void setTel(String tel) {
		this.tel = tel;
	}
}
