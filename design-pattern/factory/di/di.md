# DI
1. Provide: 获取对象工厂，并且使用一个 map 将对象工厂保存
1. Invoke: 执行入口
1. buildParam: 核心逻辑，构建参数
1. 从容器中获取 provider
1. 递归获取 provider 的参数值
1. 获取到参数之后执行函数
1. 将结果缓存并且返回结果