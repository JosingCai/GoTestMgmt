#### 依赖
- 语言：golang
- 框架：goadmin
- 数据库：mysql

#### 导入SQL
- doc/sql/init.sql
#### 更新配置文件
- adm.ini
- config.json
#### 文件管理
- 文件保存在/tmp目录
- 执行命令创建子目录：
  - mkdir -p /tmp/{api,test,file}

#### 开启服务
- 命令：go run main.go
- 访问：http://localhost:9033

#### 登录
 - 默认用户：admin/admin

#### 使用示例：
- 新建测试环境TEST
- 在文件管理Swagger子菜单：导入API文件(TEST.json), 导入测试用例文件（TEST.xmin）
- 选中环境，点击Swagger导入API用例，在用例菜单下即可看到对应API数据
- 选中环境，点击Xmind导入测试用例，在用例菜单下即可看到对应测试用例数据


#### API使用
- 在环境设置abnomal模式，其他相关信息，API进行随机测试
- 在测试数据菜单下准备完整的body或urlquery，方便API进行回归测试
- 在公用环境定义API需要的各项数据，API进行组合数据测试

#### 数据统计
- 单个API统计
- 全部API统计

#### 图例说明：
- 见doc/img目录
```
├── allapicount.png
├── apicase.png
├── apicount.png
├── apidetail.png
├── casefile.png
├── hostenv.png
├── resultdetail.png
├── testcase.png
└── testdata.png
```

#### 说明
- 功能持续完善中，部分菜单目前只做了数据展示,主页延用的goadmin，后续会更换
- 使用有问题可以添加wx沟通：liuhuocjx
