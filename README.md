# way_m3u8
由于个人下载需要，基于 orestonce/m3u8d 设计的web端



## Depend on
gin
https://github.com/orestonce/m3u8d


# deploy
前端简单的web端设计好了，就是一个基本实现，没什么特别的

# 编译
安装 golang:1.22.4 然后编译main.go就行，然后放到./bin下去运行就行

# 修改目标
 * [x] 增加.gitignore忽略无关文件
 * [x] 使用embed的形式将static/目录嵌入到最终输出的二进制里面
 * [ ] 默认状态的配置可以直接启动
 * [ ] 使用纯go版本的sqlite，以便跨平台
 * [ ] 使用github action自动编译发布
 * [ ] 配置信息储存在数据库, 无数据库则使用默认状态的配置
 * [ ] 优化界面显示