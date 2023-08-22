//选出所有 bonus < 1000 的员工的 name 及其 bonus。
//
// Employee 表单
//
// +-------+--------+-----------+--------+
//| empId |  name  | supervisor| salary |
//+-------+--------+-----------+--------+
//|   1   | John   |  3        | 1000   |
//|   2   | Dan    |  3        | 2000   |
//|   3   | Brad   |  null     | 4000   |
//|   4   | Thomas |  3        | 4000   |
//+-------+--------+-----------+--------+
//empId 是这张表单的主关键字
//
//
// Bonus 表单
//
// +-------+-------+
//| empId | bonus |
//+-------+-------+
//| 2     | 500   |
//| 4     | 2000  |
//+-------+-------+
//empId 是这张表单的主关键字
//
//
// 输出示例：
//
// +-------+-------+
//| name  | bonus |
//+-------+-------+
//| John  | null  |
//| Dan   | 500   |
//| Brad  | null  |
//+-------+-------+
//
//
// Related Topics 数据库 👍 56 👎 0

package main

func main() {

}

//There is no code of Go type for this problem

SELECT name, bonus FROM employee e LEFT JOIN bonus b ON e.empId = b.empId WHERE b.bonus < 1000 OR b.bonus IS NULL;
SELECT name, bonus FROM employee LEFT JOIN bonus USING(empId) WHERE bonus < 1000 OR bonus IS NULL;

