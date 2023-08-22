//表: Queue 
//
// 
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| person_id   | int     |
//| person_name | varchar |
//| weight      | int     |
//| turn        | int     |
//+-------------+---------+
//person_id 是这个表的主键。
//该表展示了所有候车乘客的信息。
//表中 person_id 和 turn 列将包含从 1 到 n 的所有数字，其中 n 是表中的行数。
//turn 决定了候车乘客上巴士的顺序，其中 turn=1 表示第一个上巴士，turn=n 表示最后一个上巴士。
//weight 表示候车乘客的体重，以千克为单位。
// 
//
// 
//
// 有一队乘客在等着上巴士。然而，巴士有1000 千克 的重量限制，所以其中一部分乘客可能无法上巴士。 
//
// 写一条 SQL 查询语句找出 最后一个 上巴士且不超过重量限制的乘客，并报告 person_name 。题目测试用例确保顺位第一的人可以上巴士且不会超重。
// 
//
// 查询结果格式如下所示。 
//
// 
//
// 示例： 
//
// 
//输入：
//Queue 表
//+-----------+-------------+--------+------+
//| person_id | person_name | weight | turn |
//+-----------+-------------+--------+------+
//| 5         | Alice       | 250    | 1    |
//| 4         | Bob         | 175    | 5    |
//| 3         | Alex        | 350    | 2    |
//| 6         | John Cena   | 400    | 3    |
//| 1         | Winston     | 500    | 6    |
//| 2         | Marie       | 200    | 4    |
//+-----------+-------------+--------+------+
//输出：
//+-------------------+
//| person_name       |
//+-------------------+
//| Thomas Jefferson  |
//+-------------------+
//解释：
//为了简化，Queue 表按 turn 列由小到大排序。
//+------+----+-----------+--------+--------------+
//| Turn | ID | Name      | Weight | Total Weight |
//+------+----+-----------+--------+--------------+
//| 1    | 5  | Alice     | 250    | 250          |
//| 2    | 3  | Alex      | 350    | 600          |
//| 3    | 6  | John Cena | 400    | 1000         | (最后一个上巴士)
//| 4    | 2  | Marie     | 200    | 1200         | (无法上巴士)
//| 5    | 4  | Bob       | 175    | ___          |
//| 6    | 1  | Winston   | 500    | ___          |
//+------+----+-----------+--------+--------------+ 
//
// Related Topics 数据库 👍 78 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 自连接
SELECT a.person_name FROM queue a JOIN queue b ON a.turn >= b.turn GROUP BY a.person_id HAVING SUM(b.weight) <= 1000 ORDER BY a.turn DESC LIMIT 1
// 自定义变量
SELECT a.person_name, a.weight FROM (
    SELECT person_name, @pre := @pre+weight AS weight FROM queue, (SELECT @pre := 0) tmp ORDER BY turn
    ) AS a
WHERE a.weight <= 1000 ORDER BY a.weight DESC