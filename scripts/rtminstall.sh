#!/bin/bash

# set error exit
set -e

# define variables
RTM_URL="https://download.agora.io/sdk/release/rtm_agora_sdk_v2.2.5_20251021.zip"
TEMP_DIR="/tmp/rtm_install_$$"
AGORA_SDK_DIR="./agora_sdk"
AGORA_SDK_DIR_MAC="./agora_sdk_mac"

UNAME_S=`uname -s`
OS=unknown



if [[ $UNAME_S == Linux ]]; then
    OS=linux
elif [[ $UNAME_S == Darwin ]]; then
    OS=mac
else
    echo "Unsupported OS: ${UNAME_S}"
    exit 1
fi
echo "OS: ${OS}"

echo "start download RTM SDK..."

# Create temporary directory
mkdir -p "$TEMP_DIR"

# download RTM SDK
echo "downloading from $RTM_URL..."
curl -L -o "$TEMP_DIR/rtm_agora_sdk.zip" "$RTM_URL"

# check download result
if [ ! -f "$TEMP_DIR/rtm_agora_sdk.zip" ]; then
    echo "download failed"
    rm -rf "$TEMP_DIR"
    exit 1
fi

echo "download completed, start unzip..."

# unzip to temporary directory
cd "$TEMP_DIR"
unzip -q rtm_agora_sdk.zip


# check unzip result
if [ ! -d "agora_sdk" ]; then
    echo "unzip failed, agora_sdk directory not found"
    rm -rf "$TEMP_DIR"
    exit 1
fi

echo "unzip completed, start copy files..."

# back to project root directory
cd - > /dev/null

# create target directory based on OS
if [ "$OS" == "mac" ]; then
    # Mac: only create agora_sdk_mac directory
    mkdir -p "$AGORA_SDK_DIR_MAC"
    
    # check write permission for target directory
    if [ ! -w "$AGORA_SDK_DIR_MAC" ]; then
        echo "error: no write permission for $AGORA_SDK_DIR_MAC"
        echo "please check the directory permission or use sudo to run the script"
        rm -rf "$TEMP_DIR"
        exit 1
    fi
    
    # copy agora_rtm_sdk_c directory to agora_sdk_mac
    if [ -d "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" ]; then
        echo "copy agora_rtm_sdk_c directory to $AGORA_SDK_DIR_MAC..."
        cp -r "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" "$AGORA_SDK_DIR_MAC/"
    else
        echo "warning: agora_rtm_sdk_c directory not found"
    fi
    
    # copy .dylib files to agora_sdk_mac directory, not overwrite if exists
    echo "copy .dylib files to $AGORA_SDK_DIR_MAC..."
    DYLIB_COUNT=$(find "$TEMP_DIR/agora_sdk" -name "*.dylib" | wc -l)
    if [ $DYLIB_COUNT -gt 0 ]; then
        find "$TEMP_DIR/agora_sdk" -name "*.dylib" -exec cp -n {} "$AGORA_SDK_DIR_MAC/" \;
        echo "copied $DYLIB_COUNT .dylib files"
    else
        echo "no .dylib files found"
    fi
    
elif [ "$OS" == "linux" ]; then
    # Linux: only create agora_sdk directory
    mkdir -p "$AGORA_SDK_DIR"
    
    # check write permission for target directory
    if [ ! -w "$AGORA_SDK_DIR" ]; then
        echo "error: no write permission for $AGORA_SDK_DIR"
        echo "please check the directory permission or use sudo to run the script"
        rm -rf "$TEMP_DIR"
        exit 1
    fi
    
    # copy agora_rtm_sdk_c directory to agora_sdk
    if [ -d "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" ]; then
        echo "copy agora_rtm_sdk_c directory to $AGORA_SDK_DIR..."
        cp -r "$TEMP_DIR/agora_sdk/agora_rtm_sdk_c" "$AGORA_SDK_DIR/"
    else
        echo "warning: agora_rtm_sdk_c directory not found"
    fi
    
    # copy .so files to agora_sdk directory, not overwrite if exists
    echo "copy .so files to $AGORA_SDK_DIR..."
    SO_COUNT=$(find "$TEMP_DIR/agora_sdk" -name "*.so" | wc -l)
    if [ $SO_COUNT -gt 0 ]; then
        find "$TEMP_DIR/agora_sdk" -name "*.so" -exec cp -n {} "$AGORA_SDK_DIR/" \;
        echo "copied $SO_COUNT .so files"
    else
        echo "no .so files found"
    fi
fi

# 清理临时目录
echo "clean temporary files..."
rm -rf "$TEMP_DIR"

echo "RTM SDK installation completed!"
if [ "$OS" == "mac" ]; then
    echo "Files copied to: $AGORA_SDK_DIR_MAC"
    echo "  - agora_rtm_sdk_c/ (header files)"
    echo "  - *.dylib files"
elif [ "$OS" == "linux" ]; then
    echo "Files copied to: $AGORA_SDK_DIR"
    echo "  - agora_rtm_sdk_c/ (header files)"
    echo "  - *.so files"
fi

