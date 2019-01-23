# 装饰模式(Decorator)



## 定义

动态地给一个对象添加一些额外的职责。 就增加功能来说， Decorator模式相比生成子类更为灵活。



## 参与者

- Component: 抽象构件
  - 定义一个对象接口，可以给这些对象动态的添加职责
- ConcreteComponent: 具体构件
  - 定义一个对象，可以给这个对象添加一些职责
- Decorator: 抽象装饰类
  - 维持一个指向Component对象的指针，并定义一个与Component接口一致的接口
- ConcreteDecorator: 具体装饰类
  - 向组件添加职责