# skipList
20220601 看了下[fast-skiplist](github.com/sean-public/fast-skiplist)，这个数据结构我都没整明白。。。

20220602 看了篇博客[redis 跳表分析并用 Go 实现](https://mp.weixin.qq.com/s/c3mOGotVOzUrl1P8r-PSxA)，有讲解还是比较好懂的，主要要理清插入和删除操作的操作流程

20220611 又看了两篇[Go 实现跳跃表](https://mp.weixin.qq.com/s/BaDpagOecG7TtLoELhdtOw)，[带你彻底击溃跳表原理及其Golang实现！（内含图解）](https://mp.weixin.qq.com/s/FVghWmqO0BHY3yk-gfTpag)，这两个都和redis里的不太一样，没有backward指针，存的键值对，只比较键，没有对象obj

20220612 看了三篇讲go泛型的[Go官方的泛型简介](https://mp.weixin.qq.com/s/qTGHGRt1aQpgcpm6sbFKFw)，[Go 中的泛型：激动人心的突破](https://mp.weixin.qq.com/s/Zk24GsvpryB64hlSAp06Iw)，[深入浅出Go泛型之泛型使用三步曲](https://mp.weixin.qq.com/s/ieV4ztqu4BR0P1odOZdT5w)

20220613 实现了最基本的skiplist

20220614 [Go io](https://mp.weixin.qq.com/s/A4pXvTr4WboFfP2IiksIPw), [二进制编码](Go 数据存储篇（四）：通过 Gob 包序列化二进制数据)

跳表的基本数据结构和插入删除查找操作[skipList.go](./skipList.go)

将跳表的键值对存入文件和从文件中还原跳表[dumpLoad.go](./dumpLoad.go)

支持一下Mr.Lei的仓库：[SkipList](https://github.com/hustlyhang/SkipList)