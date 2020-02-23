# LearningGo
go语言学习
## learn-go-tips go语音学习demo实例
1. chapter1 hello world
2. chapter2 types类型，添加recover/panic方式处理
3. chapter3 Variables 变量
4. chapter4 Control Structures 
5. chapter5 Arrays,Slices,Map
6. chapter6 Functions 函数
7. chapter7 Structs and Interface 结构体和接口【类似于java的class和Interface】,反射
8. chapter8 package 主要的包（strings,ioutil,list,sort,os,sort,net）
9. chapter9 ioutil/http，包含并发相关demo，协程
10. chapter10
11. chapter11 读取文件夹中的文件，并将其文件名写入另一个文件
12. chapter13 mod管理 
13. chapter14 
14. chapter15 json的使用
15. chapter16 
16. chapter17 搭建基础框架

> 知识点
1. 提供任何类型的指针，*T；
2. 没有构造函数、实例方法、继承层次结构、动态方法查找，只有structs和interfaces。
3. 数组属于数值，使用**slices**作为参数, **slices**是对数组的底层引用。
4. 支持字符串，支持哈希表
5. goroutines单独执行线程；channels线程之间的通信渠道；
6. maps,slices,channels是通过引用而非值传递的。
7. 名称首部大写为public，否则为package-private
8. **error**：遇到程序中的异常类型来代表注入到达文件末尾之类的事件；
9. **panics**：来代表运行时注入试图越界读取数组的运行时错误；
10. **不支持隐式类型**，包含不同类型的表达式需要进行显示转换；
11. **不支持方法重载**。函数和方法在同一作用域内必须有不同的名称；
12. nil代表错误的指针，类似于java的null
