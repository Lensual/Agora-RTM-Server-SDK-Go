#!/bin/bash

# 设置错误时退出
set -e

# 定义变量
RTM_URL="https://download.agora.io/sdk/release/rtm_agora_sdk.zip"
TEMP_DIR="/tmp/rtm_install_$$"
AGORA_SDK_DIR="./agora_sdk"

echo "开始下载RTM SDK..."

# 创建临时目录
mkdir -p "$TEMP_DIR"

# 下载RTM SDK
echo "正在从 $RTM_URL 下载文件..."
curl -L -o "$TEMP_DIR/rtm_agora_sdk.zip" "$RTM_URL"

# 检查下载是否成功
if [ ! -f "$TEMP_DIR/rtm_agora_sdk.zip" ]; then
    echo "错误：下载失败"
    exit 1
fi

echo "下载完成，开始解压..."

# 解压到临时目录
cd "$TEMP_DIR"
unzip -q rtm_agora_sdk.zip


# 检查解压后的目录结构
if [ ! -d "agora_sdk" ]; then
    echo "错误：解压后的目录结构不正确，未找到 agora_sdk 目录"
    exit 1
fi

echo "解压完成，开始拷贝文件..."

# 回到项目根目录
cd - > /dev/null

# 创建目标目录：如果存在，则不做修改，否则创建
mkdir -p "$AGORA_SDK_DIR"

# 拷贝 rtm_include 目录及其所有 .h 文件
if [ -d "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" ]; then
    echo "拷贝 agora_rtm_sdk_c 目录..."
    cp -r "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" "$AGORA_SDK_DIR/"
else
    echo "警告：未找到 agora_rtm_sdk_c 目录"
fi

# 拷贝所有 .so 文件
echo "拷贝 .so 文件..."
find "$TEMP_DIR/agora_sdk" -name "*.so" -exec cp {} "$AGORA_SDK_DIR/" \;

# 拷贝所有 .dylib 文件
echo "拷贝 .dylib 文件..."
find "$TEMP_DIR/agora_sdk" -name "*.dylib" -exec cp {} "$AGORA_SDK_DIR/" \;

# 清理临时目录
echo "清理临时文件..."
rm -rf "$TEMP_DIR"

echo "RTM SDK 安装完成！"
echo "文件已拷贝到 $AGORA_SDK_DIR 目录下"

