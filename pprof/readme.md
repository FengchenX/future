1. 生成对应的文件名 cpu.prof  .prof
2. go tool pprof cpu.prof
3. top 查看样本数据
4. web 生成svg文件


5. 其它 go tool pprof -pdf cpu.prof a.pdf