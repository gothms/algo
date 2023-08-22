//活动记录表：Activity
//
//
//+---------------+---------+
//| Column Name   | Type    |
//+---------------+---------+
//| user_id       | int     |
//| session_id    | int     |
//| activity_date | date    |
//| activity_type | enum    |
//+---------------+---------+
//该表是用户在社交网站的活动记录。
//该表没有主键，可能包含重复数据。
//activity_type 字段为以下四种值 ('open_session', 'end_session', 'scroll_down', 'send_
//message')。
//每个 session_id 只属于一个用户。
//
//
//
//
// 请写SQL查询出截至 2019-07-27（包含2019-07-27），近 30 天的每日活跃用户数（当天只要有一条活动记录，即为活跃用户）。
//
// 以 任意顺序 返回结果表。
//
// 查询结果示例如下。
//
//
//
// 示例 1:
//
//
//输入：
//Activity table:
//+---------+------------+---------------+---------------+
//| user_id | session_id | activity_date | activity_type |
//+---------+------------+---------------+---------------+
//| 1       | 1          | 2019-07-20    | open_session  |
//| 1       | 1          | 2019-07-20    | scroll_down   |
//| 1       | 1          | 2019-07-20    | end_session   |
//| 2       | 4          | 2019-07-20    | open_session  |
//| 2       | 4          | 2019-07-21    | send_message  |
//| 2       | 4          | 2019-07-21    | end_session   |
//| 3       | 2          | 2019-07-21    | open_session  |
//| 3       | 2          | 2019-07-21    | send_message  |
//| 3       | 2          | 2019-07-21    | end_session   |
//| 4       | 3          | 2019-06-25    | open_session  |
//| 4       | 3          | 2019-06-25    | end_session   |
//+---------+------------+---------------+---------------+
//输出：
//+------------+--------------+
//| day        | active_users |
//+------------+--------------+
//| 2019-07-20 | 2            |
//| 2019-07-21 | 2            |
//+------------+--------------+
//解释：注意非活跃用户的记录不需要展示。
//
// Related Topics 数据库 👍 105 👎 0

package main

func main() {

}

//There is no code of Go type for this problem
// 日期计算
SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users FROM activity WHERE activity_date BETWEEN DATE_SUB('2019-07-27', INTERVAL 29 DAY) AND '2019-07-27' GROUP BY activity_date;
// lc
SELECT
	activity_date AS day,
	COUNT(DISTINCT user_id) AS active_users
FROM
	Activity
WHERE
	DATEDIFF('2019-7-27',activity_date) <30
AND
	activity_date <= '2019-07-27'
GROUP BY
	activity_date
