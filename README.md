# agora-rtm-sdk go

支持的os：
linux + mac

备注： mac下建议只做为开发环境，不做运行环境

使用事项：
[x]不要在回调函数中，调用sdk的api
[x]不要在回调中，做耗时的操作


需要设置：
在项目目录下面：
for Linux:
export LD_LIBRARY_PATH=./third_party/agora_rtm_sdk_c/
for mac:
export DYLD_LIBRARY_PATH=./third_party/agora_rtm_sdk_c/

编译过程：
手动编译：
go build -o ./bin/ ./cmd/example/

脚本编译：
./scripts/build.sh

执行: 
./bin/example <appid> <channelname> <usid> <token_option>

版本历史记录：
20250701 release 0.0.5
-- 更新：更新rtm 版本到2.2.5
-- 增加：
   - 增加SendChannelMessage: 直接发送channelmsg，也就是频道消息
   - 增加SendUserMessage: 直接发送usermsg，也就是点对点消息
   - NOTE： 开发者也可以直接调用Publish，但是需要自己实现消息的封装。参考sample
   - 增加：对 *agrtm.LinkStateEvent的go 封装实现。GetGoLinkStateEvent()
-- fix：
   - fix 一个在sample中，申请了optio，但没有做释放的bug
20250428 release 0.0.4
-- 更新：更新rtm 版本到2.2.4.1
-- 修改：修改mac 的LDflags，fix mac下编译错误
20250424 release 0.0.3
-- Add:
   - 增加对mac的支持


20250424 release 0.0.2

## RTM SDK Version

The RTM SDK version has been updated from 2.1.x to 2.2.x.

## Library Files

The library files have been updated to use `.so` files instead of `.a` files.

## Library File Location

All `.so` files are now located in a single directory.
## add 8 callback for logout etc
## add pushmessage in demo to do like echo server
## add chan in demo to get time for each step

最佳实践
1、释放顺序
   - 先释放client，再释放channel
2、回调中，不要调用sdk的api，否则会引起多次销毁
3、回调中，不要做耗时的操作，否则会引起阻塞
4 configure 最后释放




Todo list
-[]回调中，缺少logout的callback
-[]cmd/sample中需要做释放
-[]需要做pprof测试，
   -[]验证cpu、内存耗费
   -[]验证释放有内存泄漏？
-[] 需要修改rtmclient的对象，变更为go中的对象，否则会引起多次销毁
-[] 对mac的支持
-[] 所有的void 都用handle来定义，不要定义各种不同的c_Ixx ,太难理解了
-[] eventhandler: 需要提供一个unreg 方法，或者就是在client的logout的时候，自动来处理？？需要查看rtm c++是否有unreg event handler

- [x] 将github 去掉，或者是替换为agora io，看看是否能编译通过？
- [x] 包命令为agrtm
- [x] 将bridge 迁移到agora，不需要bridage，并修改main.go
- [x] 修改scripts下面的build.sh
- [x] 将目录结构做改变，不用那么复杂。更新后的目录结构为：

1. 整体结构符合 Go 标准项目布局：
   - `cmd/`: 存放主要的应用程序入口，也就是sample 的目录
   - `pkg/`: 存放可以被外部应用程序使用的库代码，也就是rtm的go实现代码
   - `third_party/`: 存放agora_rtm_sdk有关的include和lib文件。其中:
        - `agora_rtm_sdk_c/`：存放的是c api 的include 和.a 文件
        - `agora_rtm_sdk_c/agora_rtm_sdk/`：存放的是rtm_sdk 有关的highlevel API和so文件
   - `scripts/`: 存放构建和部署脚本
   - `bin/`: 存放编译后的可执行文件

2. 主要组件：
   - `cmd/example/`: 示例应用程序
   - `pkg/agora/`: Agora RTM SDK 的核心实现
   - `third_party/agora_rtm_sdk_c/`: C 语言版本的 Agora RTM SDK

3. 代码组织特点：
   - 接口和实现分离：`pkg/agora/` 目录下的文件都以 `I` 开头表示接口
   - 功能模块化：将不同功能（如 Client、Service、Storage 等）分离到不同文件
   - 清晰的依赖关系：通过 `third_party` 目录管理外部依赖

4. 构建系统：
   - 使用 `go.mod` 进行依赖管理
   - 使用 `scripts/build.sh` 进行构建
   - 输出目录为 `bin/`

5. 文档：
   - `README.md` 提供项目说明



建议改进：
1. 可以添加 `internal/` 目录存放私有代码
2. 可以添加 `api/` 目录存放 API 定义
3. 可以添加 `test/` 目录存放测试代码
4. 可以添加 `docs/` 目录存放详细文档
5. 可以添加 `configs/` 目录存放配置文件模板

这个项目结构是一个典型的 Go 项目，遵循了 Go 社区的最佳实践，适合作为其他 Go 项目的参考。

- [ ] 在c api工程中，更新c API的命名规范，更新为agora_rtm_xxx_，

- [ ] c api命名规范更新后，更新go 中的对应c api