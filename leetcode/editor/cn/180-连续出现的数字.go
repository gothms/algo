//表：Logs 
//
// 
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| num         | varchar |
//+-------------+---------+
//在 SQL 中，id 是该表的主键。
//id 是一个自增列。 
//
// 
//
// 找出所有至少连续出现三次的数字。 
//
// 返回的结果表中的数据可以按 任意顺序 排列。 
//
// 结果格式如下面的例子所示： 
//
// 
//
// 示例 1: 
//
// 
//输入：
//Logs 表：
//+----+-----+
//| id | num |
//+----+-----+
//| 1  | 1   |
//| 2  | 1   |
//| 3  | 1   |
//| 4  | 2   |
//| 5  | 1   |
//| 6  | 2   |
//| 7  | 2   |
//+----+-----+
//输出：
//Result 表：
//+-----------------+
//| ConsecutiveNums |
//+-----------------+
//| 1               |
//+-----------------+
//解释：1 是唯一连续出现至少三次的数字。 
//
// Related Topics 数据库 👍 776 👎 0

package main

import (
    "fmt"
)

func main() {

}

//There is no code of Go type for this problem
// 外连接
SELECT DISTINCT a.num AS ConsecutiveNums FROM logs AS a, logs AS b, logs AS c WHERE a.id = b.id+1 AND a.num = b.num AND b.id = c.id+1 AND b.num = c.num;
// 关联子查询
SELECT DISTINCT a.num AS ConsecutiveNums FROM logs AS a WHERE 3 = (SELECT COUNT(*) FROM logs AS b WHERE a.num = b.num AND b.id - a.id BETWEEN 0 AND 2);