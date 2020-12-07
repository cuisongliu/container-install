#!/bin/bash
echo "### Usage" >> Note.md
echo "
\`\`\`sh
# 下载并安装container-install, container-install是个golang的二进制工具，直接下载拷贝到bin目录即可, release页面也可下载
$ wget -c https://cuisongliu.oss-cn-beijing.aliyuncs.com/container-install/latest/linux_amd64/container-install && \\
    chmod +x container-install && mv container-install /usr/bin
\`\`\`
" >> Note.md
echo "### [amd64 下载地址]" >> Note.md
echo "[oss 下载地址](https://${BUCKETNAME:-cuisongliu}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/container-install/${VERSION}/linux_amd64/container-install)" >> Note.md
echo "[latest 版本 oss下载地址](https://${BUCKETNAME:-cuisongliu}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/container-install/${VERSION}/linux_arm64/container-install)" >> Note.md
echo "### [arm64 下载地址]" >> Note.md
echo "[oss 下载地址](https://${BUCKETNAME:-cuisongliu}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/container-install/${VERSION}/linux_arm64/container-install)" >> Note.md
echo "[latest 版本 oss下载地址](https://${BUCKETNAME:-cuisongliu}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/container-install/${VERSION}/linux_arm64/container-install)" >> Note.md

cat Note.md
