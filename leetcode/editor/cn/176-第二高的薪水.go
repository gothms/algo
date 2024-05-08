//Employee 表：
//
//
//
//
//+-------------+------+
//| Column Name | Type |
//+-------------+------+
//| id          | int  |
//| salary      | int  |
//+-------------+------+
//在 SQL 中，id 是这个表的主键。
//表的每一行包含员工的工资信息。
//
//
//
//
//
//
// 查询并返回 Employee 表中第二高的薪水 。如果不存在第二高的薪水，查询应该返回 null(Pandas 则返回 None) 。
//
// 查询结果如下例所示。
//
//
//
// 示例 1：
//
//
//输入：
//Employee 表：
//+----+--------+
//| id | salary |
//+----+--------+
//| 1  | 100    |
//| 2  | 200    |
//| 3  | 300    |
//+----+--------+
//输出：
//+---------------------+
//| SecondHighestSalary |
//+---------------------+
//| 200                 |
//+---------------------+
//
//
// 示例 2：
//
//
//输入：
//Employee 表：
//+----+--------+
//| id | salary |
//+----+--------+
//| 1  | 100    |
//+----+--------+
//输出：
//+---------------------+
//| SecondHighestSalary |
//+---------------------+
//| null                |
//+---------------------+
//
//
//
// Related Topics 数据库 👍 1372 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 子查询
//SELECT (SELECT DISTINCT salary FROM employee ORDER BY salary DESC LIMIT 1 OFFSET 1) AS SecondHighestSalary;
//
//SELECT IFNULL((SELECT DISTINCT salary FROM employee ORDER BY salary DESC LIMIT 1 OFFSET 1), NULL) AS SecondHighestSalary;
