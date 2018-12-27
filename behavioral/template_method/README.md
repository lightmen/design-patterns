#  模板方法模式(TemplateMethod)

## 定义

定义一个操作中的算法的骨架，而将一些步骤延迟到子类， TemplateMethod使得子类可以不改变一个算法的结构即可重定义该算法的某些步骤



# 结构

略



# 参与者

- AbstractClass
  - 定义一个抽象的原语操作，具体的子类将重定义它们以实现一个算法的各个步骤。
  - 实现一个模板方法，定义一个算法的骨架。该模板方法不仅调用原语操作，也调用定义在AbstractClass或者其他对象中的操作
- ConcreteClass
  - ConcreteClass靠AbstractClass来实现算法中不变的步骤