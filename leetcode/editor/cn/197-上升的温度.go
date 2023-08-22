//
//
// 表： Weather
//
//
//
//
//+---------------+---------+
//| Column Name   | Type    |
//+---------------+---------+
//| id            | int     |
//| recordDate    | date    |
//| temperature   | int     |
//+---------------+---------+
//在 SQL 中，id 是该表的主键。
//该表包含特定日期的温度信息
//
//
//
// 找出与之前（昨天的）日期相比温度更高的所有日期的 id 。
//
// 返回结果 无顺序要求 。
//
// 结果格式如下例子所示。
//
//
//
// 示例 1：
//
//
//输入：
//Weather 表：
//+----+------------+-------------+
//| id | recordDate | Temperature |
//+----+------------+-------------+
//| 1  | 2015-01-01 | 10          |
//| 2  | 2015-01-02 | 25          |
//| 3  | 2015-01-03 | 20          |
//| 4  | 2015-01-04 | 30          |
//+----+------------+-------------+
//输出：
//+----+
//| id |
//+----+
//| 2  |
//| 4  |
//+----+
//解释：
//2015-01-02 的温度比前一天高（10 -> 25）
//2015-01-04 的温度比前一天高（20 -> 30）
//
//
// Related Topics 数据库 👍 472 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 提交不过
ELECT a.id FROM weather a JOIN weather b ON b.temperature < a.temperature AND DATE(a.recordDate) - DATE(b.recordDate) = 1;
// DATEDIFF
SELECT w1.id FROM Weather w1 JOIN Weather w2 ON w1.temperature > w2.temperature AND DATEDIFF(w1.recordDate, w2.recordDate) = 1;
