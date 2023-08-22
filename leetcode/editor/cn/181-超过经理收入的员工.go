//表：Employee
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| name        | varchar |
//| salary      | int     |
//| managerId   | int     |
//+-------------+---------+
//Id是该表的主键。
//该表的每一行都表示雇员的ID、姓名、工资和经理的ID。
//
//
//
//
// 编写一个SQL查询来查找收入比经理高的员工。
//
// 以 任意顺序 返回结果表。
//
// 查询结果格式如下所示。
//
//
//
// 示例 1:
//
//
//输入:
//Employee 表:
//+----+-------+--------+-----------+
//| id | name  | salary | managerId |
//+----+-------+--------+-----------+
//| 1  | Joe   | 70000  | 3         |
//| 2  | Henry | 80000  | 4         |
//| 3  | Sam   | 60000  | Null      |
//| 4  | Max   | 90000  | Null      |
//+----+-------+--------+-----------+
//输出:
//+----------+
//| Employee |
//+----------+
//| Joe      |
//+----------+
//解释: Joe 是唯一挣得比经理多的雇员。
//
// Related Topics 数据库 👍 637 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 关联子查询
SELECT a.name AS Employee FROM employee AS a WHERE a.managerId IS NOT null && a.salary > (SELECT b.salary FROM employee AS b WHERE a.managerId = b.id);
// 内连接
SELECT a.name AS Employee FROM employee AS a JOIN employee AS b ON a.managerId = b.id AND a.salary > b.salary;
// 内连接
SELECT a.name AS Employee FROM employee AS a JOIN employee AS b ON a.managerId = b.id WHERE a.salary > b.salary;