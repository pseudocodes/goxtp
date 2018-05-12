# History

## 0.0.1 (2018-05-12)
* OnDepthMarketData 参数修正
	* 修改 Swig 脚本, 添加 directorin 逻辑，适配传入参数

## 0.0.1 (2018-05-11)
* 完善 quote_watch 例子
	* macosx 环境在创建 `QuoteApi` 时会根据传入的 `save_file_path` 创建目录，但是会在 save_file_path 之前加入随机字符，原因不明，之后程序被锁死
	* macosx c++ 程序不会出现该问题
	* linux 环境没有该问题
	* 提前创建好日志目录可以让程序运行下去

## 0.0.1 (2018-05-10)
* 测试 20180227.1.1.16.20 版本的 macosx 库
    * 这个版本的 dylib 库是开发人员用 gcc7 编译的，默认的 clang 环境可以编译通过，but 运行时使用的动态库不对
    * brew 安装 gcc7 之后, 运行时指定链接库的路径，可以把程序跑起来
    * 把行情的例子信息里的注释和运行参数和官方的标准统一

## 0.0.1 (2018-05-09)
* Readme 和 History 文档添加
* 暂时还没有时间完成例子 

## 0.0.1 (2018-05-06)
*   第一次 `Swig` 封装提交
	*   对几个字符串输入参数进行微调
	*   调整 `Swig` 脚本文件中关于 `golang` 类型转化的规则部分
	*   根据 [goctp](https://github.com/pseudcodes/goctp) 的例子写一半，还未具体测试, 未来将提供 `XTP` 项目中同样功能的 Demo
	*   编译输出还报部分 **warning** 以后再处理
