1. 将.so文件放到 /usr/lib 或/lib 目录中
2. 在/etc/ld.so.conf中添加/usr/lib 或/lib目录中
eg:
include /etc/ld.so.conf.d/*.conf
/usr/lib

3. 运行sudo ldconfig
