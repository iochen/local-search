package local_search_test

import (
	"bytes"
	"testing"

	search "github.com/iochen/local-search"
)

var file =
`---
title: openssl RSA简单使用
date: 2018-02-10 11:44:49
categories:
- Crypto
tags:
- openssl
- 加解密
- RSA
- encrypt
- decrypt
---
最近看到的关于TLS的文章比较多啊，我也来凑凑热闹，来一篇关于非对称加密用法的博客。当然了，这篇文章对不懂的人也没什么用，懂的人也不会看这篇，还是写给自己看吧
<!--more-->
这里的非对称加密用的是**RSA**
# 基本
非对称加密密钥的用法主要是如下几种：`+
"`"+"`"+"`"+`
加密 → 公钥
解密 → 私钥
签名 → 私钥
解密 → 私钥
`+"`"+"`"+"`"+`
# 用法
## 生成密钥
### 生成私钥
比如我要生成一个叫`+"`"+`pr.pem`+"`"+`的私钥文件，直接`+
"`"+"`"+"`"+`
openssl genrsa -out pr.pem
`+"`"+"`"+"`"+`
就可以了，就像这样：
![生成私钥](/img/openssl-RSA-1/f05d8842-7080-4b4c-bcc2-bbc07ceaefc5.png)
### 提取公钥
比如我要从`+`pr.pem`+`里提取公钥，直接`+
"`"+"`"+"`"+`
openssl rsa -pubout -in pr.pem -out pu.pem
`+"`"+"`"+"`"+`
即可，可以看到：
![提取公钥](/img/openssl-RSA-1/6ed6724b-1fac-415e-bd11-8403680fe04c.png)
### 小结
至此，公钥和私钥就生成好了，来看看吧！
私钥：
![私钥](/img/openssl-RSA-1/b567d11b-55c7-4e2e-9629-32a406427301.png)
公钥：
![公钥](/img/openssl-RSA-1/4f6e2b69-034f-4bdd-a02a-52ef9310e0bd.png)
## 加解密
### 加密
我们先创建一个叫的文件，写点什么
😭
`

func TestEntry_Parse(t *testing.T) {
	entry := &search.Entry{}
	reader := bytes.NewReader([]byte(file))
	err := entry.Parse(reader)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v",entry)
}

func TestEntry_Json(t *testing.T) {
	entry := &search.Entry{}
	reader := bytes.NewReader([]byte(file))
	err := entry.Parse(reader)
	if err != nil {
		t.Error(err)
	}
	json, err := entry.Json()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s",string(json))
}