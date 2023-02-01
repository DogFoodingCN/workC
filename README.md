# Work Collector

## Usage

### 1. 安装油猴插件
谷歌浏览器 > 更多工具 > 扩展程序 > Chrome 商店安装插件 > Tampermonkey

### 2. 导入油猴脚本 tampermonkey_scripts.zip

### 3. 安装启动 go 程序

```
go run cmd/main.go
```

### 4. 设置油猴插件

- 油猴插件管理面板

- 编辑脚本

```
// 设置读取页面数
var pages = 10;
// 岗位筛选关键字
var needs = ["php", "java","go","后端", "前端"];
// 设置投递公司黑名单
var blacklist = ["软*动力"];
```

### 5. 获取 boss 直聘链接
关键字搜索、筛选城市、薪资等获取链接
> eg: https://www.zhipin.com/web/geek/job?query=Go&city=101200100&salary=405,406&page=1

### 6. 获取岗位

- 插件开启油猴脚本
- 访问 boss 链接
- 在 go 程序目录获取岗位 excel
> Jobs_2023-XX-XX.xlsx
