# 项目介绍

## 概览
本项目是一个基于 原luban项目（作为 submodule 存放在 `./luban` 目录下）的 golang 版本实现。
它严格遵循 Luban 的设计哲学（类型安全、面向对象、多态等），保持与原版高度一致性的同时，也符合 Golang 的最佳实践。

## 目前已实现的核心模块

在 AI Agent（如 Trae 等工具）的辅助翻译和开发下，本项目目前已经构建了非常完整的配置加载和生成管线，覆盖了从 Schema 解析到数据校验再到多语言目标生成的全流程。

### 1. 核心类型系统 (Core Types & Datas)
- **TType**：定义了静态类型的元数据。支持了基础类型（`int`, `float`, `string`, `bool` 等），集合类型（`array`, `list`, `map`, `set`），以及自定义的结构体（`bean`）和枚举（`enum`），特别增加了本地化文本（`text`）类型。
- **DType**：定义了运行时数据实例结构。与 `TType` 一一对应，承载从数据源中解析出来的数据，并全面实现了访问者模式（Visitor Pattern）以便进行序列化。
- **多态 (Polymorphism)**：通过抽象类型与子类型继承（Parent/Child）实现了面向对象的多态字段读取，能够动态解析 `$type` 列并实例化正确的子类。

### 2. Schema 与装配 (Defs & Schema)
- 实现了 `XmlSchemaLoader`，用于解析与原版 Luban 格式一致的 `__root__.xml` 和其他定义文件。
- 实现了 `DefAssembly`，用于将松散的 Raw 类型定义编译成包含继承树和引用关系的运行时定义。

### 3. 数据加载与解析 (DataLoader)
- **ExcelDataLoader**：通过 `excelize` 库强力支持 Excel 表格解析。
  - 支持表头控制列（`##`, `##var`, `##type`）。
  - 支持行列忽略标记（`#`）。
  - 支持复杂的多行记录合并：可以通过主键留空的方式，将多行数据合并拼装成同一个 Bean 下的 List / Map。
- **CsvDataLoader**：支持解析纯文本 `.csv` 格式，且复用了所有的 Excel 高级控制语义。

### 4. 代码目标生成 (CodeTarget)
完全基于 Go 的 `text/template` 实现了灵活的模板驱动代码生成：
- **TemplateCodeTarget**：通用的外部模板生成器基类。
- **支持语言**：内置支持 Go，并通过加载外部 `.tpl` 文件完美支持生成 C#, C++, Java, Lua, TypeScript 甚至是 Protobuf（`.proto`） 结构定义代码。
- **按组过滤**：完全支持按照 Target Group（如 `client`, `server`）对生成的类字段进行过滤隔离。

### 5. 数据目标导出 (DataTarget)
基于统一的 `IDataVisitor` 实现了多种数据产物格式：
- **JSON**：输出易于阅读的标准 JSON 配置文件。
- **Lua**：导出为 Lua Table 代码文件，便于集成热更游戏前端。
- **Binary**：导出基于小端序压缩的 `.bytes` 二进制文件，以提升加载速度。

### 6. 数据校验 (Validator)
内建了多项核心数据防错校验器，会在数据加载完毕后进行统一校验：
- **RefValidator**：跨表外键引用校验。
- **RangeValidator**：数值范围及字符串长度校验。
- **PathValidator**：本地资源（如美术图片路径）真实存在性校验。

### 7. 多语言本地化 (L10N)
- 支持 `text` 类型，能够在解析表格数据时自动识别、提取需要多语言翻译的文本。
- `L10nManager` 统一管理翻译键值对，对于未设置 Key 的文本能够通过 MD5 自动生成 Key。
- 管道支持单独剥离并导出 `l10n.json` 翻译汇总文件，并在配置数据中替换为对应结构的 Key-Value。

## 运行与 CLI (Pipeline)
提供了精简高效的 `DefaultPipeline` 处理流程，入口在 `cmd/luban/main.go`。
- 可通过 `-input_data_dir` 指定数据和 Schema 目录。
- 可通过 `-code_target` 指定代码语言（`go`, `cs`, `cpp`, `java`, `ts`, `lua`, `pb`）。
- 可通过 `-data_target` 指定数据格式（`json`, `lua`, `bin`）。
- 结合 `-template_dir` 指向项目的 `templates` 文件夹使用即可一键生成双端代码及数据。
