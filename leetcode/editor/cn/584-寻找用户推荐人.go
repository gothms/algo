//表: Customer
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| name        | varchar |
//| referee_id  | int     |
//+-------------+---------+
//在 SQL 中，id 是该表的主键列。
//该表的每一行表示一个客户的 id、姓名以及推荐他们的客户的 id。
//
// 找出那些 没有被 id = 2 的客户 推荐 的客户的姓名。
//
// 以 任意顺序 返回结果表。
//
// 结果格式如下所示。
//
//
//
// 示例 1：
//
//
//输入：
//Customer 表:
//+----+------+------------+
//| id | name | referee_id |
//+----+------+------------+
//| 1  | Will | null       |
//| 2  | Jane | null       |
//| 3  | Alex | 2          |
//| 4  | Bill | null       |
//| 5  | Zack | 1          |
//| 6  | Mark | 2          |
//+----+------+------------+
//输出：
//+------+
//| name |
//+------+
//| Will |
//| Jane |
//| Bill |
//| Zack |
//+------+
//
// Related Topics 数据库 👍 213 👎 0

package main

func main() {

}

//There is no code of Go type for this problem

SELECT name FROM customer WHERE referee_id <> 2 OR referee_id IS NULL;
