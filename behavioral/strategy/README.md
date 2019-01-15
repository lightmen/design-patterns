# 策略模式(strategy)



## 定义

定义系列的算法， 把它们一个个封装起来， 并且使它们可互相替换， 本模式使得算法可独立于使用它的客户的变化。



## 参与者

- Context: 环境类
  - 用一个ConcreteStrategy对象来配置
  - 维护一个对Strategy对象的引用
  - 可定义一个接口来让Stategy访问它的数据
- Strategy: 抽象策略类
  - 定义所有支持的算法的公共接口。Context使用这个接口来调用某ConcreteStrategy定义的算法
- ConcreteStrategy: 具体策略类
  - 以Strategy接口实现某具体算法