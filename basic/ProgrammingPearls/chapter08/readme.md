保存状态，避免重复计算。算法2和算法4使用了简单的动态规划。通过使用一些空间来保存中间计算结果，我们避免了花时间对其重复计算。

将信息预处理至数据结构中。算法2b中得cumarr结构允许对子向量中的总和进行快速计算。

分治算法。算法3使用了简单地分治算法形式；有关算法设计的教科书介绍了更高级的分治算法形式。

扫描算法。与数组相关的问题经常可以通过思考“如何将x[0..i-1]的解扩展为x[0..i]的解”来解决。算法4通过同时存储已有的答案和一些辅助数据来计算新答案。

累积。算法2b使用了一个累积表，表中第i个元素的值为x中前i个值的总和；这一类表常用于处理有范围限制的问题。例如，业务分析师要确定3月份到10月份的销售额，可以从10月份的本年迄今销售额中减去2月份的销售额。

下界。只有在确定了自己的算法是所有可能算法中最佳的算法以后，算法设计师才可能踏踏实实地睡个好觉。为此，他们必须证明某个相匹配的下界。对问题线性下界的讨论见习题6，更复杂的下界证明可能会十分困难。