package basic

/*
多模式串匹配算法
	Trie
	AC 自动机

	单模式串匹配算法：一个主串中查找一个模式串
	多模式串匹配算法：一个主串中查找多个模式串
基于单模式串和 Trie 树实现的敏感词过滤
	对敏感词字典进行预处理，构建成 Trie 树结构
	如果敏感词字典动态更新了，比如删除、添加了一个敏感词，只需要动态更新一下 Trie 树就可以了
	用户输入的内容作为主串，在 Trie 树中匹配

经典的多模式串匹配算法：AC 自动机，Aho-Corasick 算法
	AC 自动机实际上就是在 Trie 树之上，加了类似 KMP 的 next 数组，只不过此处的 next 数组是构建在树上罢了
	AC 自动机的构建，包含两个操作：
		将多个模式串构建成 Trie 树
		在 Trie 树上构建失败指针（相当于 KMP 中的失效函数 next 数组）
失败指针
	可匹配后缀子串
		某个后缀子串可以匹配某个模式串的前缀，这个后缀子串叫作可匹配后缀子串
	失败指针指向那个最长匹配后缀子串对应的模式串的前缀的最后一个节点
	当要求某个节点的失败指针的时候，我们通过已经求得的、深度更小的那些节点的失败指针来推导
		也就是说，可以逐层依次来求解每个节点的失败指针
		所以，失败指针的构建过程，是一个按层遍历树的过程

	已经求得某个节点 p 的失败指针之后，如何寻找它的子节点的失败指针呢？
		root 的失败指针为 NULL
		如果找到了节点 q 的一个子节点 qc，对应的字符跟节点 pc 对应的字符相同，则将节点 pc 的失败指针指向节点 qc
		如果节点 q 中没有子节点的字符等于节点 pc 包含的字符，则令 q=q->fail（fail 表示失败指针，这里有没有很像 KMP 算法里求 next 的过程？）
			继续上面的查找，直到 q 是 root 为止
		如果还没有找到相同字符的子节点，就让节点 pc 的失败指针指向 root

构建失败指针代码实现
	先构建 Trie 树
	再在 Trie 树上构建失败指针
	// TODO
如何在 AC 自动机上匹配主串？
	AC 自动机从指针 p=root 开始，模式串是 b，主串是 a
	1.如果 p 指向的节点有一个等于 b[i]的子节点 x，我们就更新 p 指向 x
		这个时候我们需要通过失败指针，检测一系列失败指针为结尾的路径是否是模式串
		处理完之后，我们将 i 加一，继续这两个过程
	2.如果 p 指向的节点没有等于 b[i]的子节点，那失败指针就派上用场了，我们让 p=p->fail，然后继续这 2 个过程

如何才能实现一个高性能的敏感词过滤系统呢？
	上述代码实现中，再遍历一遍文本内容（主串），就可以将文本中的所有敏感词替换成“***”
	时间复杂度：构建失败指针
		Trie 树构建的时间复杂度是 O(m*len)，其中 len 表示敏感词的平均长度，m 表示敏感词的个数
		假设 Trie 树中总的节点个数是 k，整个失败指针的构建过程就是 O(k*len)
	时间复杂度：匹配
		O(len)，所以总的匹配的时间复杂度就是 O(n*len)
		实际情况下，可能近似于 O(n)
	对比：
		从时间复杂度上看，AC 自动机匹配的效率跟 Trie 树一样
		实际上，因为失效指针可能大部分情况下都指向 root 节点
			所以绝大部分情况下，在 AC 自动机上做匹配的效率要远高于刚刚计算出的比较宽泛的时间复杂度
		只有在极端情况下，AC 自动机的性能才会退化的跟 Trie 树一样

总结
	AC 自动机是基于 Trie 树的一种改进算法，它跟 Trie 树的关系，就像单模式串中，KMP 算法与 BF 算法的关系一样
	KMP 算法中有一个非常关键的 next 数组，类比到 AC 自动机中就是失败指针
	AC 自动机其实就是 KMP 算法在多模式串上的改造

思考
	各个字符串匹配算法的特点和比较适合的应用场景？
*/
/*
Java 构建失败指针
public void buildFailurePointer() {
  Queue<AcNode> queue = new LinkedList<>();
  root.fail = null;
  queue.add(root);
  while (!queue.isEmpty()) {
    AcNode p = queue.remove();
    for (int i = 0; i < 26; ++i) {	// 开始匹配前缀
      AcNode pc = p.children[i];
      if (pc == null) continue;
      if (p == root) {	// 根节点
        pc.fail = root;
      } else {
        AcNode q = p.fail;	// 已求得 p 的失败指针是 q
        while (q != null) {
          AcNode qc = q.children[pc.data - 'a'];
          if (qc != null) {	// 找到最长匹配前缀
            pc.fail = qc;	// 实际求的是 pc 的失败指针
            break;
          }
          q = q.fail;
        }
        if (q == null) {	// 没有公共前缀
          pc.fail = root;
        }
      }
      queue.add(pc);
    }
  }
}
Java 在 AC 自动机上匹配主串
public void match(char[] text) { // text是主串
  int n = text.length;
  AcNode p = root;
  for (int i = 0; i < n; ++i) {
    int idx = text[i] - 'a';
    while (p.children[idx] == null && p != root) {	// 不匹配，查看失败指针
      p = p.fail; // 失败指针发挥作用的地方
    }	// 2
    p = p.children[idx];
    if (p == null) p = root; // 如果没有匹配的，从root开始重新匹配
    AcNode tmp = p;
    while (tmp != root) { // 打印出可以匹配的模式串
      if (tmp.isEndingChar == true) {
        int pos = i-tmp.length+1;
        System.out.println("匹配起始下标" + pos + "; 长度" + tmp.length);
      }
      tmp = tmp.fail;
    }	// 1
  }
}
*/
