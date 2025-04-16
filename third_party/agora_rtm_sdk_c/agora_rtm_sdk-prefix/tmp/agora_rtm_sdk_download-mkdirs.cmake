# Distributed under the OSI-approved BSD 3-Clause License.  See accompanying
# file Copyright.txt or https://cmake.org/licensing for details.

cmake_minimum_required(VERSION 3.5)

file(MAKE_DIRECTORY
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download-build"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/tmp"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download-stamp"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src"
  "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download-stamp"
)

set(configSubDirs )
foreach(subDir IN LISTS configSubDirs)
    file(MAKE_DIRECTORY "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download-stamp/${subDir}")
endforeach()
if(cfgdir)
  file(MAKE_DIRECTORY "/home/weihongqin/work/agora_rtm_sdk_c/agora_rtm_sdk_download-prefix/src/agora_rtm_sdk_download-stamp${cfgdir}") # cfgdir has leading slash
endif()
