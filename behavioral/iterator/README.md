# 迭代器模式(Iterator)



## 定义

提供一种方法顺序访问一个聚合对象中各个元素， 而又不需要暴露该对象的内部表示



## 参与者

- Iterator(迭代器)
  - 迭代器定义访问和遍历元素的接口
- ConcreteIterator（具体迭代器）
  - 具体迭代器实现迭代器接口
  - 对该聚合遍历时跟踪当前位置
- Aggregate（聚合）
  - 聚合定义创建相应迭代器对象的接口
- ConcreteAggregate（具体聚合）
  - 具体聚合实现创建相应聚合中的当前对象，并能够返回ConcreteIterator的一个适当的实例