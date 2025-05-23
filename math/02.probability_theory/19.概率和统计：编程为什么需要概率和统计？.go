package _2_probability_theory

/*
数据结构和基础算法体现的大多数都是离散数学的思想
	这些思想更多的时候是给我们提供一种解决问题的思路，在具体指导我们解决问题的时候，我们还需要更多的数学知识
	比如说，在机器学习、数据挖掘等领域，概率统计就发挥着至关重要的作用

率和统计里有哪些需要掌握的概念？
	在第一个模块中，我们认为所有事件都是一分为二的，要么必然发生，要么必然不发生。换句话说，事件的发生只有必然性，没有随机性
		但是现实生活中，我们常常会碰到一些模棱两可的情况
	概率（Probability）就是描述这种可能性的一个数值
		随机变量（Random Variable）来描述事件所有可能出现的状态，并使用概率分布（Probability Distribution）来描述每个状态出现的可能性
		而随机变量又可以分为离散型随机变量（Discrete Random Variable）和连续型随机变量（Continuous Random Variable）
	概率分布
		联合概率（Joint Probability）
		边缘概率（Marginal Probability）
		条件概率
	其实概率论研究的就是这些概率之间相互转化的关系，比如联合概率、条件概率和边缘概率
		通过这些关系，概率论中产生了著名的贝叶斯定理（Bayes’ theorem）
			加上变量的独立性，我们就可以构建朴素贝叶斯（Naive Bayes）分类算法，这个算法在机器学习中的应用非常广泛
		此外，基于概率发展而来的信息论，提出了很多重要的概率
			例如信息熵（Entropy）/ 香农熵（Shannon Entropy）、信息增益（Information Gain）、基尼指数（Gini）等
			这些概念都被运用到了决策树（Decision Tree）的算法中
	概率和统计其实是互逆的
		概率论是对数据产生的过程进行建模，然后研究某种模型所产生的数据有什么特性
		而统计学正好相反，它需要通过已知的数据，来推导产生这些数据的模型是怎样的
			因此统计特别关注数据的各种分布、统计值及其对应的统计意义
			在真实的世界里，我们通常只能观测到一些数据，而无法事先知道，是什么模型产生了这些数据，这时候就要依赖统计学
			所以，海量数据的分析、实验和机器学习，都离不开统计学

概率和统计可以帮我们做什么？
	首先，我还是要提到复杂度分析
		你可以看到，现实中每种情况出现的可能性是不一样的，这也就意味着概率分布是不均匀的
		而不均匀的概率分布，最终会影响平均复杂度的加权平均计算
	除此之外，概率和统计对于机器学习和大数据分析而言更为重要
		机器学习中的监督式学习，就是通过训练样本，估计出模型的参数，最后使用训练得出的模型，对新的数据进行预测
			通过训练样本来估计模型，我们可以交给统计来完成。在机器学习的特征工程步骤中，我们可以使用统计的正态分布，标准化（standardization）不同取值范围的特征，让它们具有可比性
		此外，对机器学习算法进行效果评估时，A/B 测试可以减少不同因素对评测结果的干扰
			为了获得更可靠的结论，我们需要理解统计意义，并为每个 A/B 测试计算相应的统计值
	最后，概率模型从理论上对某些机器学习算法提供了支持
		朴素贝叶斯分类充分利用了贝叶斯定理，使用先验概率推导出后验概率，并通过变量之间相互独立的假设，把复杂的计算进行大幅简化
			简化之后，我们就可以把这个算法运用在海量文本的分类任务上
		而决策树使用了信息熵和信息增益，挑出最具有区分力的条件，构建决策树的结点和分支
			这样构建出的树，不仅分类效率更高，而且更利于人脑的理解
			谷歌的 PageRank 算法利用马尔科夫链的概率转移，有效地刻画了人们浏览互联网的行为，大幅提升了互联网搜索的体验

学习这部分内容，需要做哪些准备？
	这部分内容，有公式是不可避免的，我尽量只保留那些最核心的公式
	机器学习的知识纷繁复杂，涉及广泛，很多问题甚至是跨学科、跨领域的
		我在讲解的时候，尽量给你抽象出最核心的部分，讲清楚来龙去脉，让你了解它整体的运作方式，不影响你对核心知识点的吸收
	当然，你可以适度地补一些概率知识，这样理解起来会更快。我在之前的加餐三中推荐了几本书，你可以找来看看，了解一些基本概念
		另外，你可以准备一些实际工作和项目中的问题。例如，你之前参与的任务，哪些可以使用概率论来解决？碰到的难题有哪些？你是如何解决的？

小结
	最重要就是你耳熟能详的这几个：随机变量、概率分布、联合概率、条件概率和边缘概率
	通过这些概念之间的相互推导，我们可以得到贝叶斯定理，这是朴素贝叶斯等系列算法的核心
		而在概率基础之上发展而来的信息论，定义了信息熵、信息增益和基尼指数等，构成了决策树等系列算法的核心
	概率研究的是模型如何产生数据，统计研究的是如何通过数据来推导其背后的模型
	概率和统计的运用非常多，我这里主要讲了三个方面
		第一，概率可以帮助我们进行更精准的复杂度分析
		第二，概率统计更多是用在机器学习和大数据分析中
		第三，概率统计还可以用在各种机器学习的算法中

思考
	之前你对概率统计的认识是什么样的呢？对这块内容，你觉得最难的是什么？
*/
