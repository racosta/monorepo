# Blog Renderer using Templating

This is the implementation of the
[Templating](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates)
chapter of [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

This implementation differs because the
[`go-approval-tests`](https://github.com/approvals/go-approval-tests) doesn't
play nicely with Go's module because of some file paths (See
<https://github.com/approvals/go-approval-tests/issues/76>). My implementation
replaces `go-approval-tests` with [`goldie`](https://github.com/sebdah/goldie).
