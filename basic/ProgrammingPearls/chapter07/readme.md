在计算机科学中，二叉树（英语：Binary tree）是每个节点最多只有两个分支(不存在分支度大于2的节点)的树结构。通常分支被称作“左子树”和“右子树”。二叉树的分支具有左右次序，不能颠倒。

二叉树的第i层至多拥有 2^{i-1}个节点数；k的二叉树至多总共有 2^(k+1)-1 个节点数（定义根节点所在深度 k0=0，而总计拥有节点数匹配的，称为“满二叉树”；深度为k有n个节点的二叉树，当且仅当其中的每一节点，都可以和同样深度k的满二叉树，序号为1到n的节点一对一对应时，称为“完全二叉树”。；对任何一棵非空的二叉树T，如果其叶片(终端节点)数为n0，分支度为2的节点数为n2，则n0 = n2 + 1。

与普通树不同，普通树的节点个数至少为1，而二叉树的节点个数可以为0；普通树节点的最大分支度没有限制，而二叉树节点的最大分支度为2；普通树的节点无左、右次序之分，而二叉树的节点有左、右次序之分。
二叉树通常作为数据结构应用，典型用法是对节点定义一个标记函数，将一些值与每个节点相关系。这样标记的二叉树就可以实现二叉查找树和二元堆积，并应用于高效率的搜索和排序


“对于现代计算机来说，将循环展开则有助于避免管道阻减少分支、增加指令级并行性。”


Litte定律：系统中物体的平均数量等于物体离开系统的平均速率和每个物体在系统中停留的平均时间的乘积。即：队列中物体的平均数量 = 进入速率 * 平均停留时间