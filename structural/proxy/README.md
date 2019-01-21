# 代理模式(proxy)



## 定义

为其他对象提供一种代理以控制对这个对象的访问



## 参与者

- Subject: 抽象主题角色
  - 定义一个RealSubject和Proxy的共用接口，这样就在任何使用RealSubject的地方都可以使用Proxy
- Proxy: 代理主题角色
  - 保存一个引用使得代理可以访问实体。若RealSubject和Subject的接口相同，Proxy会引用Subject
  - 提供一个与Subject的接口相同的接口，这样代理就可以用来替代实体
  - 控制对实体的存取，并可能负责创建和删除它
  - 其他功能依赖于代理的类型
- RealSubject: 真实主题角色
  - 定义Proxy所代表的实体