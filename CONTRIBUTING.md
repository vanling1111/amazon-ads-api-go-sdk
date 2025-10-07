# Contributing to Amazon Ads API Go SDK

感谢您考虑为 Amazon Ads API Go SDK 做出贡献！

## 🤝 贡献方式

### 报告 Bug

1. 检查 [Issues](https://github.com/vanling1111/amazon-ads-api-go-sdk/issues) 确认问题未被报告
2. 使用 Bug Report 模板创建新 Issue
3. 提供详细的复现步骤和环境信息

### 提议新功能

1. 先创建 Feature Request Issue 讨论
2. 等待维护者反馈和批准
3. 再开始开发

### 提交代码

1. **Fork 仓库**

```bash
git clone https://github.com/YOUR_USERNAME/amazon-ads-api-go-sdk.git
cd amazon-ads-api-go-sdk
```

2. **创建分支**

```bash
git checkout -b feat/your-feature-name
```

3. **开发和测试**

```bash
# 编写代码
# ...

# 运行测试
go test ./...

# 检查代码质量
golangci-lint run
```

4. **提交代码**

遵循 [Conventional Commits](https://www.conventionalcommits.org/):

```bash
git commit -m "feat(sp-v3): add CreateCampaigns method"
```

5. **推送并创建 PR**

```bash
git push origin feat/your-feature-name
```

然后在 GitHub 上创建 Pull Request。

## 📋 代码规范

### 必须遵守

- ✅ Go 官方代码规范
- ✅ 90%+ 测试覆盖率
- ✅ Google 风格注释（中文）
- ✅ 通过所有 CI 检查

详见 [开发指南](docs/DEVELOPMENT.md)。

## 🔍 Code Review

所有 PR 都需要经过维护者审查才能合并。

**审查标准**:
- 代码质量
- 测试覆盖
- 文档完整性
- 性能影响

## 📄 许可证

贡献的代码将采用项目的双许可证模式（AGPL-3.0 或 Commercial License）。

---

**感谢您的贡献！** 🎉

