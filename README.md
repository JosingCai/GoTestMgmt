#### 依赖
- 语言：golang
- 框架：goadmin
- 数据库：mysql
- 第三方工具：xmind2case (pip3 isntall xmind2case)

#### 导入SQL
- doc/sql/init.sql

#### 文件管理
- 文件保存在一个目录下，e.g:/tmp
- 执行命令创建子目录：
  - mkdir -p /tmp/testmgmt/{api,test,file,log}

#### 目录说明
- api:放置swagger api文件，e.g：TEST.json
- test:放置测试用例文件，e.g: TEST_V1.0.0_测试用例.xmind
- file: 放置公用文件，e.g: demo.xmind, smoketest.py ...
- log: 项目自身的日志信息

#### 更新配置文件
- adm.ini
- config.json
  - file_base_path 项内容设置为文件保存的目录
  
#### 开启服务
##### 开发环境：
- 命令：go run main.go
- 访问：http://localhost:9033

##### 生产环境
- linux环境：GOOS=linux GOARCH=amd64 go build -o testmgmt main.go 
- 部署：nohup ./testmgmt &
- 配置nginx: 监听配置文件中配置的端口号(server_port)，默认9033
- 访问：http://10.0.X.X

#### 登录
 - 默认用户：admin/ admin

#### 使用示例：
- 新建测试环境TEST
- 在文件管理Swagger子菜单：导入API文件(TEST.json), 导入测试用例文件（TEST.xmind）
- 选中环境，点击Swagger导入API用例，在用例菜单下即可看到对应API数据
- 选中环境，点击Xmind导入测试用例，在用例菜单下即可看到对应测试用例数据

#### Xmind文件
- demo：doc/file/project_V1.0.0_testcase_demo.xmind
- 字段说明：
  - 以 _ 下线线分割
  - project与测试环境录入的项目名称保持一致
  - 版本V1.0.0用来设置用例所属版本，通过case_id体现
  - demo中的模块名中文名与英文名用中划线-分割，通过case_id体现，未设置，默认置为other



#### API使用
- 在环境设置abnomal模式，其他相关信息，API进行随机测试
- 在测试数据菜单下准备完整的body或urlquery，方便API进行回归测试
- 在公用变量定义API需要的各项数据，API进行组合数据测试

#### 数据统计
- 单个API统计
- 全部API统计
- 测试用例统计

#### 测试结果
- 再次测试

#### 图例说明：
- 见doc/img目录
```
├── allapicount.png
├── apicase.png
├── apicount.png
├── apidetail.png
├── casefile.png
├── homepage.png
├── hostenv.png
├── resultdetail.png
├── testcase.png
└── testdata.png
```

#### 说明
- 使用有问题可以添加wx沟通：liuhuocjx
