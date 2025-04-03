

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