# go-craft

基于 OpenAPI 规范快速构建和管理 Web API的工具

许可授权：<a href="https://sm.mit-license.org" target="_blank"><img src="https://img.shields.io/badge/license-MIT-green.svg" /></a></p>

## 目录

1. [介绍](#chapter-001)
2. [技术方案](#chapter-002)
3. [安装](#chapter-003)<br>
  3.1 [环境依赖](#chapter-0031)<br>
  3.2 [从源码构建](#chapter-0032)<br>
  3.3 [直接下载](#chapter-0033)<br>
  3.4 [运行 go-craft](#chapter-0034)
4. [测试](#chapter-004)
5. [文档](#chapter-005)
6. [社区](#chapter-006)
7. [参与贡献](#chapter-007)
8. [授权](#chapter-008)
999. [码云特技](#chapter-999)

## 1. 介绍 <a id="chapter-001"></a>

go-craft 是一个快速构建和管理 Web API 的工具，可以用来生成客户端代码和服务端代码。
go-craft 最初的构想来自于 [think-builder](https://github.com/goldeagle/think-builder)，
为[thinkphp](https://github.com/top-think/think)用户提供一个可以快速生成CRUD脚手架代码的工具。随着 think-builder 的开发，
发现我可能需要一个更加强力的工具，可以管理 web 应用的方方面面。

后来进行过多种的尝试，包括通过扩展 thinkphp 官方 console 命令工具、使用 php 编写独立的工具等等，最后决定改为提供一个“更原始”的工具。
也就是说这个工具不需要依赖于任何初始环境就可以开始构建 Web API的整个工作流，包括但是不局限于：创建目录结构、管理应用配置、管理应用依赖、
管理路由、管理事件、管理视图……

go-craft 开始主要面向 thinkphp 框架，但是设计时已经考虑可以扩展为其他语言或者框架。
这样的一个工具，不能依赖于预先提供的语言环境（例如 php-cli）或者 runtime（例如 php-fpm）就可以运行。
另外我考虑需要用一个实际项目来学习掌握 Go 语言开发，所以我选择了使用 Go 语言来重构这一切。

在参考了 go-ctrl、openapi-generator 等多个代码生成工具之后，我决定基于 OpenAPI 3.0/3.1 标准来开发 go-craft。

go-craft 计划支持的生成的客户端语言（版本）如下：

- Android Java/Kotlin（7/10/11）
- C （14/17/20）
- C++ （14/17/20）
- Deno JavaScript/TypeScript （1.7）
- Go （1.15）
- Java （8/11/17）
- Node JavaScript/TypesSript （14.15/15.7）
- Kotlin（1.4）
- PHP（7.4/8.0）
- Rust（1.49）
- Swift（5）

go-craft 计划支持的生成的服务端语言（框架）如下：

- C++（Drogon v1.3.x）
- Go （Goframe v1.15.x，Gin v1.6.x）
- Node（Express v4.17.x/5.0.x）
- PHP（ThinkPHP v6.x, Lumen v8.x，Laravel v8.x）
- Rust（Actix-web v3.x，Tide v0.16.x）

go-craft 计划支持生成文档的格式如下：

- markdown
- html
- pdf
- hoppscotch json
- postman json

go-craft 计划支持生成支持的 schema 如下：

- mysql （5.7.x/8.0.x）
- mariadb （10.5.x）
- tidb （4.0.x/5.0.x）
- graphql （june2018)
- protobuf （3.14.x)
- tdengine (2.0.x)

因此，go-craft 具有以下特性：

- 高性能、高并发性的代码
- 核心程序以 CLI 命令行方式提供，并提供web界面
- 尽量低的内存等资源占用
- 尽量高的性能和低延迟
- 支持 x86_64 和 aarch64 CPU 架构
- 支持 Linux、MacOS、Windows 操作系统
- 支持 generator、editor、document server 多种模式

go-craft 支持的功能：

- 生成项目目录（允许多种语言）
- 生成项目 Makefile
- 生成 API 服务端代码
- 生成 API 客户端代码（SDK）
- 生成多种数据库结构
- 生成协议文件
- 生成 API 文档
- 生成 Hoppscotch 配置 json
- 生成 Apache / Nginx / Traefik 等网关配置文件
- 生成 Dockerfile

## 2. 软件架构 <a id="chapter-002"></a>

go-craft 希望提供一个快速、轻型、安全的 Web API构建工具，所以选择了可靠且高效的 [Go](https://golang.google.cn) 语言。
并使用 [MIT](./LICENSE) 授权，用户可以根据自己的需求进行二次开发。

go-craft 可以作为本地命令或远程命令运行，通过 ansible 进行前后端交互，并在此基础上提供官方 GUI 提升操作建议性。

## 3. 安装 <a id="chapter-003"></a>

### 3.1 环境依赖 <a id="chapter-0031"></a>

go-craft 需要使用 go 1.16 以上版本进行编译构建。
可以去 https://golang.google.cn/dl/ 获取最新的 go 版本。

相关文档可以参考 https://golang.google.cn/doc/install
  
这些都装好了之后，就可以通过源代码构建 go-craft 了。

### 3.2 从源代码构建 <a id="chapter-0032"></a>

```bash
# 下载 go-craft 源代码
$ git clone https://gitee.com/go-craft/go-craft
$ cd go-craft

# 编译
$ make
```

编译完成后，可以直接通过 `./go-craft` 目录来运行 go-craft。

### 3.3 直接安装 <a id="chapter-0033"></a>

- 手工下载

暂不提供

- deb 安装

未来会为 “debian家族” （例如：`debian` 、`kali` 或 `ubuntu`） 操作系统提供 deb 安装包和 ppa 源。


## 添加 submodules
```bash
git submodule add https://github.com/go-craft/doc doc
git submodule add https://github.com/go-craft/example example
git submodule add https://github.com/go-craft/tmpl-client-cpp-qt5 tmpl/client/cpp/qt5
git submodule add https://github.com/go-craft/tmpl-client-java-android tmpl/client/java/android
git submodule add https://github.com/go-craft/tmpl-client-rust-reqwest tmpl/client/rust/reqwest
git submodule add https://github.com/go-craft/tmpl-schema-mysql tmpl/schema/mysql
git submodule add https://github.com/go-craft/tmpl-schema-pgsql tmpl/schema/pgsql
git submodule add https://github.com/go-craft/tmpl-schema-tdengine tmpl/schema/tdengine
git submodule add https://github.com/go-craft/tmpl-server-node-express tmpl/server/node/exporess
git submodule add https://github.com/go-craft/tmpl-server-php-thinkphp tmpl/server/php/thinkphp
git submodule add https://github.com/go-craft/tmpl-server-php-laravel tmpl/server/php/laravel
git submodule add https://github.com/go-craft/tmpl-server-go-goframe tmpl/server/go/goframe
git submodule add https://github.com/go-craft/tmpl-server-go-gin tmpl/server/go/gin
git submodule add https://github.com/go-craft/webapp-thinkphp webapp/thinkphp
git submodule add https://github.com/go-craft/webapp-goframe webapp/goframe
git submodule add https://github.com/go-craft/ui-vue ui/vue
git submodule init
```

### 3.4 运行 go-craft <a id="chapter-0034"></a>

#### 命令模式

- 通过源码编译后，可以直接在源码目录运行命令：

```bash
$ ./go-craft <command> <option> <target>
```

- 通过安装工具安装了 go-craft，可以运行命令：

```bash
$ /usr/loca/bin/go-craft <command> <option> <target>
```

如果 `/usr/local/bin` 已经在 `PATH` 环境变量中，则可以直接：

```bash
$ go-craft <command> <option> <target>
```

#### 守护模式

- 通过源码编译后，可以直接在源码目录运行命令：

```bash
$ ./target/release/go-craft -d
``` 

- 通过安装工具安装了 go-craft，可以运行命令：

```bash
$ /usr/loca/bin/go-craft -d
```

如果 `/usr/local/bin` 已经在 `PATH` 环境变量中，则可以直接：

```bash
$ go-craft -d
```

- 通过 `systemd` 运行守护模式

1. 复制代码目录下的 `./scripts/go-craft.service` 到 `systemd` 用户目录 (例如 `~/.config/systemd/user`)
2. 启动服务

```bash
$ sudo systemctl start go-craft
```

## 4. 测试 <a id="chapter-004"></a>

  ```
  $ go test
  ```

## 5. 文档 <a id="chapter-005"></a>

项目官方网站： http://go-craft.enet51.com

官方手册资源库： https://github.com/go-craft/doc

## 6. 社区 <a id="chapter-006"></a>

QQ 群： 348077414

## 7. 参与贡献 <a id="chapter-007"></a>

1.  Fork 本仓库
2.  新建 feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

## 8. 授权 <a id="chapter-008"></a>

[授权](./LICENSE)