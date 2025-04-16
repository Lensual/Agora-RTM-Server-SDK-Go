# agora_rtm_sdk_cgo

Agora RTM SDK for GO

base on `Agora_RTM_C++_SDK_for_Linux_v2.1.12`

base on [agora_rtm_sdk_c](https://github.com/Lensual/agora_rtm_sdk_c)

需要设置：
在项目目录下面：
export LD_LIBRARY_PATH=./third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk

编译过程：
go build ./cmd/example/

然后执行: ./example <appid> <channelname> <usid> <token_option>