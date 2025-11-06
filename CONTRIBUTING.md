# Contributing

Go SDK 定位：

使用 Go 语言对 C++ SDK 进行包装，提供一个更易用、更安全的接口。

入参、出参、返回值要求为Go的对象，不要传递非托管内存。在 CGO 边界处对对象进行拷贝。

对于出参，尽可能通过多返回值来实现，而不是传入指针。

## VS Code

推荐安装的插件

- region 代码区域高亮 [LincDemo/vscode-region-highlighter](https://github.com/LincDemo/vscode-region-highlighter/releases/tag/1.0.2)
  - 注：该fork解决了`// #region`前方存在`\t`不高亮的问题

推荐的配置

settings.json

```json
{
    "go.toolsEnvVars": {
        "CGO_CFLAGS": "-O0 -g -I${workspaceFolder}/agora_sdk/agora_rtm_sdk_c/include",
        "CGO_LDFLAGS": "-O0 -g -L${workspaceFolder}/agora_sdk",
    },
    "terminal.integrated.env.linux": {
        "LD_LIBRARY_PATH": "${workspaceFolder}/agora_sdk",
    },
    "go.testEnvVars": {
        "LD_LIBRARY_PATH": "${workspaceFolder}/agora_sdk",
    },
    "go.testEnvFile": "${workspaceFolder}/.env",
    "go.testFlags": ["-v"]
}
```

launch.json

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "example",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/example",
            "envFile": "${workspaceFolder}/.env",
            "env": {
                "LD_LIBRARY_PATH": "${env:LD_LIBRARY_PATH}:${workspaceFolder}/agora_sdk",
            },
        }
    ]
}
```
