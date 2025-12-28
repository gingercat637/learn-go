# Go 学习路线 · 完整大纲

目标：从 Go 零基础 → 能独立开发、部署、维护生产级 Go 后端服务

---

## 一、基础语法与语言特性

1. Go 环境安装与基本命令  
   1.1 Go 安装与版本管理  
   1.2 go env 的作用  
   1.3 go run / go build / go install 的区别  
   1.4 Go 程序的执行入口（main 包、main 函数）
   1.5 热更新air安装与配置

2. 基本数据类型与零值  
   2.1 Go 中的类型大致分层次（标量类型与复合类型总览）  
   2.2 布尔、整型、浮点、复数、字符串与别名类型  
   2.3 复合/引用类型概览：数组、切片、结构体、映射、函数、接口、指针、通道  
   2.4 零值机制及其工程意义  
   2.5 常量与 iota

3. 变量声明与作用域  
   3.1 var 声明方式  
   3.2 短变量声明 :=  
   3.3 包级变量与函数级变量  
   3.4 匿名变量 _

4. 控制结构  
   4.1 if 的特殊写法  
   4.2 for（Go 中唯一的循环结构）  
   4.3 switch（值匹配与类型匹配）

5. 函数与返回值  
   5.1 函数定义与调用  
   5.2 多返回值  
   5.3 命名返回值  
   5.4 可变参数  
   5.5 闭包函数

6. defer、panic 与 recover  
   6.1 defer 的执行顺序  
   6.2 panic 的触发条件  
   6.3 recover 的正确使用场景  
   6.4 panic 在业务代码中的边界

7. 错误处理机制  
   7.1 error 接口的本质  
   7.2 显式错误处理模式  
   7.3 错误包装（fmt.Errorf / errors.Is / errors.As）  
   7.4 错误传播的工程实践

8. 指针与值语义  
   8.1 指针的基本使用  
   8.2 值拷贝与引用传递的区别  
   8.3 slice / map / interface 的引用特性

9. 结构体与方法  
   9.1 struct 定义  
   9.2 方法与函数的区别  
   9.3 值接收者与指针接收者  
   9.4 结构体嵌入（embedding）

10. 接口（interface）  
    10.1 接口的定义  
    10.2 隐式实现机制  
    10.3 空接口的使用与风险  
    10.4 最小接口原则

11. 集合类型  
    11.1 array 的特点  
    11.2 slice 的底层结构与扩容  
    11.3 map 的实现与使用限制  
    11.4 range 的使用细节

12. 包与模块  
    12.1 package 的组织方式  
    12.2 import 规则与别名  
    12.3 init 函数的执行顺序  
    12.4 可见性规则（导出标识符）

13. 常用标准库  
    13.1 fmt  
    13.2 os / io  
    13.3 time  
    13.4 encoding/json  
    13.5 net/http  
    13.6 context

---

## 二、Go Module 与工程结构

1. Go Module 基础  
   1.1 go mod init  
   1.2 go mod tidy  
   1.3 依赖版本解析规则  
   1.4 GOPATH 与 Module 的区别

2. 项目结构设计  
   2.1 cmd 目录的职责  
   2.2 internal 的访问限制  
   2.3 pkg 的使用边界  
   2.4 如何拆分业务模块

---

## 三、并发编程模型

1. goroutine  
   1.1 goroutine 的创建  
   1.2 goroutine 的调度模型  
   1.3 goroutine 泄漏问题

2. channel  
   2.1 无缓冲 channel  
   2.2 有缓冲 channel  
   2.3 channel 的关闭规则  
   2.4 单向 channel

3. select  
   3.1 多路复用  
   3.2 超时控制  
   3.3 default 分支的使用场景

4. 同步原语  
   4.1 sync.WaitGroup  
   4.2 sync.Mutex  
   4.3 sync.RWMutex  
   4.4 atomic 操作

5. 并发安全  
   5.1 数据竞争的成因  
   5.2 Go 内存模型概念  
   5.3 race detector 的使用

---

## 四、运行机制与底层原理

1. Go 调度器  
   1.1 G / P / M 模型  
   1.2 抢占式调度

2. 垃圾回收（GC）  
   2.1 三色标记法  
   2.2 STW 的含义  
   2.3 GC 调优的基本思路

3. 内存管理  
   3.1 栈与堆  
   3.2 逃逸分析  
   3.3 sync.Pool 的使用边界

---

## 五、测试、调试与性能分析

1. 单元测试  
   1.1 testing 包  
   1.2 表驱动测试  
   1.3 子测试

2. 基准测试  
   2.1 benchmark 编写方式  
   2.2 性能对比测试

3. 调试  
   3.1 delve 调试器  
   3.2 断点与变量观察

4. 性能分析  
   4.1 CPU profiling  
   4.2 内存 profiling  
   4.3 block / mutex profiling  
   4.4 性能瓶颈定位方法

---

## 六、代码规范与工具链

1. 代码风格  
   1.1 gofmt 的强制性  
   1.2 goimports 的作用  
   1.3 Go 的设计哲学

2. 静态分析  
   2.1 go vet  
   2.2 golangci-lint  
   2.3 staticcheck

3. 工程习惯  
   3.1 不滥用 interface  
   3.2 显式错误处理  
   3.3 context 贯穿请求生命周期

---

## 七、Web 后端开发

1. HTTP 服务  
   1.1 net/http 的核心概念  
   1.2 Handler 与 Middleware

2. Web 框架  
   2.1 Gin  
   2.2 Echo / Fiber（了解）

3. API 设计  
   3.1 RESTful 风格  
   3.2 统一返回结构  
   3.3 错误码设计

4. 鉴权与安全  
   4.1 JWT  
   4.2 HTTPS 与 TLS 基础  
   4.3 常见 Web 安全问题

---

## 八、数据库与存储

1. database/sql  
   1.1 连接池原理  
   1.2 CRUD 操作  
   1.3 事务管理

2. ORM 与增强库  
   2.1 sqlx  
   2.2 gorm  
   2.3 ORM 的使用边界

3. 工程实践  
   3.1 幂等设计  
   3.2 分页与锁机制

---

## 九、RPC、消息与微服务

1. RPC  
   1.1 RPC 与 HTTP 的对比  
   1.2 gRPC 基础  
   1.3 protobuf 定义

2. 消息系统  
   2.1 消息模型  
   2.2 异步处理  
   2.3 消息确认与重试

---

## 十、日志、监控与可观测性

1. 日志系统  
   1.1 结构化日志  
   1.2 日志分级

2. 监控指标  
   2.1 Prometheus  
   2.2 常见指标含义

3. 链路追踪  
   3.1 OpenTelemetry  
   3.2 分布式追踪基本原理

---

## 十一、部署与运维

1. 容器化  
   1.1 Dockerfile 编写  
   1.2 多阶段构建

2. Kubernetes  
   2.1 基本概念  
   2.2 服务部署流程

3. 服务治理  
   3.1 健康检查  
   3.2 优雅关闭  
   3.3 滚动发布

---

## 十二、CI / CD 与交付

1. CI 流程  
   1.1 自动测试  
   1.2 自动 lint

2. 构建与发布  
   2.1 Docker 镜像构建  
   2.2 版本管理（SemVer）  
   2.3 goreleaser 简介

---

## 十三、最终能力目标

1. 能独立完成 Go 后端服务开发  
2. 能正确处理并发与性能问题  
3. 能将服务部署并稳定运行在生产环境  
4. 具备长期维护 Go 服务的工程能力
