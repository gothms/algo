//表： Employee
//
//
//+--------------+---------+
//| 列名          | 类型    |
//+--------------+---------+
//| id           | int     |
//| name         | varchar |
//| salary       | int     |
//| departmentId | int     |
//+--------------+---------+
//在 SQL 中，id是此表的主键。
//departmentId 是 Department 表中 id 的外键（在 Pandas 中称为 join key）。
//此表的每一行都表示员工的 id、姓名和工资。它还包含他们所在部门的 id。
//
//
//
//
// 表： Department
//
//
//+-------------+---------+
//| 列名         | 类型    |
//+-------------+---------+
//| id          | int     |
//| name        | varchar |
//+-------------+---------+
//在 SQL 中，id 是此表的主键列。
//此表的每一行都表示一个部门的 id 及其名称。
//
//
//
//
// 查找出每个部门中薪资最高的员工。 按 任意顺序 返回结果表。 查询结果格式如下例所示。
//
//
//
// 示例 1:
//
//
//输入：
//Employee 表:
//+----+-------+--------+--------------+
//| id | name  | salary | departmentId |
//+----+-------+--------+--------------+
//| 1  | Joe   | 70000  | 1            |
//| 2  | Jim   | 90000  | 1            |
//| 3  | Henry | 80000  | 2            |
//| 4  | Sam   | 60000  | 2            |
//| 5  | Max   | 90000  | 1            |
//+----+-------+--------+--------------+
//Department 表:
//+----+-------+
//| id | name  |
//+----+-------+
//| 1  | IT    |
//| 2  | Sales |
//+----+-------+
//输出：
//+------------+----------+--------+
//| Department | Employee | Salary |
//+------------+----------+--------+
//| IT         | Jim      | 90000  |
//| Sales      | Henry    | 80000  |
//| IT         | Max      | 90000  |
//+------------+----------+--------+
//解释：Max 和 Jim 在 IT 部门的工资都是最高的，Henry 在销售部的工资最高。
//
// Related Topics 数据库 👍 667 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 子查询 很慢
SELECT d.name AS Department, e.name AS Employee, e.salary AS Salary FROM employee e, department d WHERE e.salary >= (SELECT MAX(salary) FROM employee ee GROUP BY departmentId HAVING e.departmentId = ee.departmentId) AND e.departmentId = d.id;
// IN 快多了
SELECT d.name AS Department, e.name AS Employee, e.salary AS Salary FROM employee e, department d WHERE e.departmentId = d.id AND (e.departmentId, salary) IN (SELECT departmentId, MAX(salary) FROM employee GROUP BY departmentId);
