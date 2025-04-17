# agora-rtm-sever-sdk-go

Agora RTM SDK fr GO

c api for rtm: build from another one

base on [agora_rtm_sdk_c](https://github.com/Lensual/agora_rtm_sdk_c)

需要设置：
在项目目录下面：
export LD_LIBRARY_PATH=./third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk

编译过程：
go build -o ./bin/ ./cmd/example/

然后执行: ./bin/example <appid> <channelname> <usid> <token_option>
todo list
- [x] 将github 去掉，或者是替换为agora io，看看是否能编译通过？
- [x] 包命令为agrtm
- [x] 将bridge 迁移到agora，不需要bridage，并修改main.go
- [x] 修改scripts下面的build.sh
- [ ] 将目录结构做改变，不用那么复杂。更新后的目录结构为：
- [ ] 更新c API的命名规范，更新为agora_rtm_xxx_