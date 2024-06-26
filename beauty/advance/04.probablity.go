package advance

/*
概率统计：如何利用朴素贝叶斯算法过滤垃圾短信？

实现一个简单的垃圾短信过滤功能以及骚扰电话拦截功能，该用什么样的数据结构和算法实现呢？
算法解析
	基于黑名单的过滤器
	基于规则的过滤器
	基于概率统计的过滤器

基于黑名单的过滤器
	黑名单
		原理
			维护一个骚扰电话号码和垃圾短信发送号码的黑名单
		黑名单收集
			公开的网站上下载
			通过类似“360 骚扰电话拦截”的功能，用户自主标记骚扰电话来收集
				被多个用户标记，并且标记个数超过一定阈值的号码，就可以定义为骚扰电话，并将它加入到我们的黑名单中
		存储
			散列表、二叉树等动态数据结构
		内存消耗：
			50万个电话号码，可以接受范围
				把每个号码看作一个字符串，并且假设平均长度是 16 个字节，大约需要 10MB 的内存空间
			500万个电话号码，内存消耗过大
				需要 100MB 的内存空间
				为了实现一个拦截功能，耗费用户如此多的手机内存，这显然有点儿不合理
		弊端
			如果某个垃圾短信发送者的号码并不在黑名单中，那这种方法就没办法拦截了
	布隆过滤器
		把位图大小设置为 10 倍数据大小，也就是 5000 万
		需要使用 5000 万个二进制位（5000 万 bits），换算成字节，也就是不到 7MB 的存储空间
		弊端：
			误判
			重要的电话或者短信，被当成垃圾短信或者骚扰电话拦截了
	服务器：时间换空间
		黑名单存储在服务器端上，把过滤和拦截的核心工作，交给服务器端来做
		手机端只负责将要检查的号码发送给服务器端，服务器端通过查黑名单，判断这个号码是否应该被拦截，并将结果返回给手机端
		弊端：
			网络延迟就会导致处理速度降低
			只有在联网的情况下，才能正常工作
基于规则的过滤器
	原理
		通过短信的内容，来判断某条短信是否是垃圾短信
		预先设定一些规则，如果某条短信符合这些规则，就可以判定它是垃圾短信
	局限性
		规则受人的思维方式局限，规则未免太过简单
		垃圾短信发送者可能会针对规则，精心设计短信，绕过这些规则的拦截
	规则举例
		短信中包含特殊单词（或词语），比如一些非法、淫秽、反动词语等
			如何定义特殊单词？
		短信发送号码是群发号码，非我们正常的手机号码，比如 +60389585
			如何定义什么样的号码是群发号码？
		短信中包含回拨的联系方式，比如手机号码、微信、QQ、网页链接等，因为群发短信的号码一般都是无法回拨的
		短信格式花哨、内容很长，比如包含各种表情、图片、网页链接等
		符合已知垃圾短信的模板。垃圾短信一般都是重复群发，对于已经判定为垃圾短信的短信，我们可以抽象成模板，将获取到的短信与模板匹配，一旦匹配，就可以判定为垃圾短信
	判定规则举例
		如果短信只是满足其中一条规则，就判定为垃圾短信，那会存在比较大的误判的情况
		可以综合多条规则进行判断
			比如，满足 2 条以上才会被判定为垃圾短信
			或者每条规则对应一个不同的得分，满足哪条规则，就累加对应的分数，某条短信的总得分超过某个阈值，才会被判定为垃圾短信
	如何定义特殊单词？
		定义方式
			主观定义：只是自己拍脑袋想，哪些单词属于特殊单词，那势必有比较大的主观性，也很容易漏掉某些单词
			概率统计：基于概率统计的方法，借助计算机强大的计算能力，找出哪些单词最常出现在垃圾短信中，将这些最常出现的单词，作为特殊单词，用来过滤短信
		概率统计
			前提
				有大量的样本数据，即要有大量的短信（比如 1000 万条短信）
				并且还要求，每条短信都做好了标记，它是垃圾短信还是非垃圾短信
			处理
				对这 1000 万条短信，进行分词处理（借助中文或者英文分词算法），去掉“的、和、是”等没有意义的停用词（Stop words），得到 n 个不同的单词
				针对每个单词，统计有多少个垃圾短信出现了这个单词，有多少个非垃圾短信会出现这个单词
					进而求出每个单词出现在垃圾短信中的概率，以及出现在非垃圾短信中的概率
				如果某个单词出现在垃圾短信中的概率，远大于出现在非垃圾短信中的概率，那就把这个单词作为特殊单词，用来过滤垃圾短信
			示例：参见 04.probablity_rule.jpg
朴素贝叶斯算法
	举例：10天的下雨情况和小明上学的情况
		1	2	3	4	5	6	7	8	9	10
		晴天	晴天	下雨	晴天	下雨	晴天	晴天	下雨	晴天	下雨
		上学	没上	上学	上学	没上	上学	上学	没上	上学	上学

		下雨概率：P(B) = 4/10
		不去上学概率：P(A) = 3/10
		下雨天不去上学概率：P(A|B) = 2/4
		因下雨而不去上学概率：P(B|A) = 2/3
	朴素贝叶斯公式
		P(A|B) = P(B|A) * P(A) / P(B)
		P(A|B)：在事件B发生的前提下，事件A发生的概率
		P(B|A)：在事件A发生的前提下，事件B发生的概率
		P(A)：事件A发生的概率
		P(B)：事件B发生的概率
基于概率统计的过滤器
	原理
		基础理论是基于朴素贝叶斯算法
		基于概率统计的过滤器，是基于短信内容来判定是否是垃圾短信
			计算机没办法像人一样理解短信的含义
			所以，需要把短信抽象成一组计算机可以理解并且方便计算的特征项，用这一组特征项代替短信本身，来做垃圾短信过滤
	思路
		1.通过分词算法，把一个短信分割成 n 个单词。这 n 个单词就是一组特征项，全权代表这个短信
			判定一个短信是否是垃圾短信这样一个问题，就变成了，判定同时包含这几个单词的短信是否是垃圾短信
		2.并不像基于规则的过滤器那样，非黑即白，一个短信要么被判定为垃圾短信、要么被判定为非垃圾短息
			而是使用概率，来表征一个短信是垃圾短信的可信程度
			P(短信是垃圾短信|W1,W2,...,Wn同时出现在一条短信中)
		3.错误思路
			尽管我们有大量的短信样本，但是我们没法通过样本数据统计得到这个概率。为什么不可以呢？
			统计同时包含 W1，W2，W3，…，Wn 这 n 个单词的短信有多少个（假设有 x 个）
			再看这里面属于垃圾短信的有几个（假设有 y 个）
			那包含 W1，W2，W3，…，Wn 这 n 个单词的短信是垃圾短信的概率就是 y/x
			问题：
				样本的数量再大，毕竟也是有限的，样本中不会有太多同时包含 W1，W2，W3，…，Wn 的短信的，甚至很多时候，样本中根本不存在这样的短信
				没有样本，也就无法计算概率
	朴素贝叶斯公式
		P(短信是垃圾短信|W1,W2,...,Wn同时出现在一条短信中) =
			P(W1,W2,...,Wn同时出现在一条短信中|短信是垃圾短信) * P(短信是垃圾短信) / P(W1,W2,...,Wn同时出现在一条短信中)
		1.P(W1,W2,...,Wn同时出现在一条短信中|短信是垃圾短信)
			这个概率照样无法通过样本来统计得到
			基于：
				独立事件发生的概率计算公式：P(A*B) = P(A)*P(B)
				如果事件 A 和事件 B 是独立事件，两者的发生没有相关性，事件 A 发生的概率 P(A) 等于 p1，事件 B 发生的概率 P(B) 等于 p2
				那两个同时发生的概率 P(A*B) 就等于 P(A)*P(B)
			分解公式：
				P(W1,W2,...,Wn同时出现在一条短信中|短信是垃圾短信) =
					P(W1出现在短信中|短信是垃圾短信) *
					P(W2出现在短信中|短信是垃圾短信) *
					...							*
					P(Wn出现在短信中|短信是垃圾短信)
				P（Wi 出现在短信中 | 短信是垃圾短信）表示垃圾短信中包含 Wi 这个单词的概率有多大
					这个概率值通过统计样本很容易就能获得
					假设垃圾短信有 y 个，其中包含 Wi 的有 x 个，那这个概率值就等于 x/y
		2.P(短信是垃圾短信)
			把样本中垃圾短信的个数除以总样本短信个数，就是短信是垃圾短信的概率
		3.P(W1,W2,...,Wn同时出现在一条短信中)
			这个概率还是不好通过样本统计得到，原因是样本空间有限
			没必要非得计算这一部分的概率值。为什么呢？
				实际上，我们可以分别计算同时包含 W1,W2,...,Wn 这 n 个单词的短信，是垃圾短信和非垃圾短信的概率
				假设它们分别是 p1 和 p2
				并不需要单纯地基于 p1 值的大小来判断是否是垃圾短信，而是通过对比 p1 和 p2 值的大小，来判断一条短信是否是垃圾短信
				如果 p1 是 p2 的很多倍（比如 10 倍），才确信这条短信是垃圾短信
			p1 = P(短信是垃圾短信|W1,W2,...,Wn同时出现在一条短信中) =
				P(W1,W2,...,Wn同时出现在一条短信中|短信是垃圾短信) * P(短信是垃圾短信) / P(W1,W2,...,Wn同时出现在一条短信中)
			p2 = P(短信是 非 垃圾短信|W1,W2,...,Wn同时出现在一条短信中) =
				P(W1,W2,...,Wn同时出现在一条短信中|短信是 非 垃圾短信) * P(短信是 非 垃圾短信) / P(W1,W2,...,Wn同时出现在一条短信中)
			基于这两个概率的倍数 （p1/p2） 来判断是否是垃圾短信的方法，就可以不用计算 P（W1，W2，W3，…，Wn 同时出现在一条短信中）这一部分的值了

扩展
	上述三种方法，还可以应用到很多类似的过滤、拦截的领域
		如垃圾邮件的过滤
	在黑名单过滤中，使用布隆过滤器，存在误判情况，可能会导致用户投诉
		可以结合三种不同的过滤方式的结果，对同一个短信处理
		如果三者都表明这个短信是垃圾短信，才把它当作垃圾短信拦截过滤
	工程中
		还需要结合具体的场景，以及大量的实验，不断去调整策略
		权衡垃圾短信判定的准确率（是否会把不是垃圾的短信错判为垃圾短信）和召回率（是否能把所有的垃圾短信都找到），来实现我们的需求

思考
	关于垃圾短信过滤和骚扰电话的拦截，还有没有其他方法呢？
*/
