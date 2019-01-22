# 桥接模式(Bridge)



## 定义

将抽象部分与它的实现部分分离， 使它们都可以独立地变化



## 参与者

- Abstraction：抽象类

  定义抽象类的接口

  维护一个指向Implementor类型对象的指针

- RefinedAbstraction：扩充抽象类

  扩充由Abstraction定义的接口

- Implementor：实现类接口

  定义实现类的接口， 该接口补一个要与Abstraction的接口完全一致

- ConcreteImplementor：具体实现类

  实现Implementor接口并定义它的具体实现

