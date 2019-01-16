# 享元模式(Flyweight)



## 定义

运用共享技术有效的支持大量细粒度的对象



## 参与者

- Flyweight: 抽象享元类
  - 描述一个接口， 通过这个接口flyweight可以接受并作用于外部状态
- ConcreteFlyweight: 具体享元类
  - 实现Flyweight接口，并为内部状态（如果有的话）增加存储空间，ConcreteFlyweight对象必须是可共享的。它所存储的状态必须是内部的； 即， 他必须独立于ConcreteFlyweight对象的场景
- UnsharedConcreteFlyweight: 非共享具体享元类
  - 并非所有的Flyweight子类都需要被共享。Flyweight接口使共享成为可能，但它并不是强制共享。在Flyweight对象结构的某些层次，UnsharedConcreteFlyweight对象通常将ConcreteFlyweight对象作为子节点
- FlyweightFactory: 享元工厂类
  - 创建并管理Flyweight对象
  - 确保合理地共享flyweight。当用户请求一个flyweight时，FlyweightFactory对象提供一个已创建的实例或者创建一个（如果不存在的话）。