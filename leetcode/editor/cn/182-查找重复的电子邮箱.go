//
//
//
// 表: Person
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| email       | varchar |
//+-------------+---------+
//id 是该表的主键列。
//此表的每一行都包含一封电子邮件。电子邮件不包含大写字母。
//
//
//
//
// 编写一个 SQL 查询来报告所有重复的电子邮件。 请注意，可以保证电子邮件字段不为 NULL。
//
// 以 任意顺序 返回结果表。
//
// 查询结果格式如下例。
//
//
//
// 示例 1:
//
//
//输入:
//Person 表:
//+----+---------+
//| id | email   |
//+----+---------+
//| 1  | a@b.com |
//| 2  | c@d.com |
//| 3  | a@b.com |
//+----+---------+
//输出:
//+---------+
//| Email   |
//+---------+
//| a@b.com |
//+---------+
//解释: a@b.com 出现了两次。
//
// Related Topics 数据库 👍 463 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 子查询
SELECT DISTINCT a.email AS Email FROM person AS a WHERE 1 < (SELECT COUNT(*) FROM person AS b WHERE a.email = b.email);
// GROUP BY
SELECT email AS Email FROM person GROUP BY email HAVING COUNT(id) > 1;
// 内连接
SELECT DISTINCT a.email AS Email FROM person a JOIN person b ON a.email = b.email AND a.id <> b.id;
