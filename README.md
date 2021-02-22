# File Server
+ 如仓库名，是一个文件服务器
+ 支持grpc上传、下载文件

## TODO(pending...)
+ 对上传的文件，计算md5，以md5为key查看Redis中是否存在。
    + 不存在则存储文件, 并set key path
    + 存在则对此文件创建一个软连接
+ 支持文件的增量修改

## run
```
    build.sh
    output/fs.out
```
