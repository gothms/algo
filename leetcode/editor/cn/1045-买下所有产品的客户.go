//Customer 表：
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| customer_id | int     |
//| product_key | int     |
//+-------------+---------+
//该表可能包含重复的行。
//customer_id 不为 NULL。
//product_key 是 Product 表的外键(reference 列)。
//
//
// Product 表：
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| product_key | int     |
//+-------------+---------+
//product_key 是这张表的主键（具有唯一值的列）。
//
//
//
//
// 编写解决方案，报告 Customer 表中购买了 Product 表中所有产品的客户的 id。
//
// 返回结果表 无顺序要求 。
//
// 返回结果格式如下所示。
//
//
//
// 示例 1：
//
//
//输入：
//Customer 表：
//+-------------+-------------+
//| customer_id | product_key |
//+-------------+-------------+
//| 1           | 5           |
//| 2           | 6           |
//| 3           | 5           |
//| 3           | 6           |
//| 1           | 6           |
//+-------------+-------------+
//Product 表：
//+-------------+
//| product_key |
//+-------------+
//| 5           |
//| 6           |
//+-------------+
//输出：
//+-------------+
//| customer_id |
//+-------------+
//| 1           |
//| 3           |
//+-------------+
//解释：
//购买了所有产品（5 和 6）的客户的 id 是 1 和 3 。
//
//
// Related Topics 数据库 👍 72 👎 0

package main

func main() {

}

//There is no code of Go type for this problem

SELECT customer_id FROM customer c GROUP BY customer_id HAVING COUNT(DISTINCT product_key) = (SELECT COUNT(*) FROM product);
