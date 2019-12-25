标签: openssl

# OpenSSL

---

[toc]

## 安全概述

> NIST(`National Institute of Standards and Technology`)标准
> * 保密性: 数据保密性,隐私性;
> * 完整行: 数据完整性,系统完整性
> * 可用性:
 
> 安全攻击
> * 被动攻击: 窃听
> * 主动攻击: 伪装,重放,消息篡改,拒绝服务

> 安全机制
> * 加密,数字签名,访问控制,数据完整性,认证交换,流量填充,路由控制,公证;

> 安全服务
> * 认证
> * 访问控制
> * 数据保密性: 连接保密性,无连接保密性,选择域保密性,流量保密性
> * 数据完整性
> * 不可否认性

> Linux系统: OpenSSL,gpg(pgp)
	
## 加密算法和协议

> 对称加密: 加密和解密使用同一个密钥;
```
DES: Data Encryption Standard
3DES:
AES: Advanced Encryption Standard(128bits,192bits,258bits,384bits,512bits)
Blowfish
Twofish
IDEA
RC6
CAST5

# 特性: 
  1.加密,解密使用同一个密钥;
  2.将原始数据分割成固定大小的块,逐个进行加密;

# 缺陷: 
  1.密钥过多; 
  2.密钥分发
```

> 公钥加密: 密钥是成对出现
```
公钥: 公开给所有人,pubkey
私钥: 自己留存,必须保证其私密性,secret key
特点: 用公钥加密的数据,只能使用与之配对的私钥解密,反之亦然;
			
数字签名: 主要在于让接收方确认发送方身份;
密钥交换: 发送方用对方的公钥加密一个对称密钥,并发送给对方;
数据加密:
			
算法: RSA,DSA,ELGamal
```

> 单向加密
> * 只能加密,不能解密,提取数据指纹; 
> * 特性: 定长输出,雪崩效应;
> * 功能: 保证数据完整性;
```
md5: 128bits
sha1: 160bits
sha224
sha256
sha384
sha512
```

> 密钥交换: `IKE`
```
公钥加密:
DH(Deffie-Hellman)
    A:p,g
    B:p,g
				
    A:x
        --> p^x%g 
        p^y%g^x
    B:y
        --> p^y%g
        p^x%g^y
```

## PKI 

> PKI: Public Key Infrastructure(公钥基础设施)
> * CA: Certificate Authority,电子商务认证授权机构,负责发放和管理数字证书的机构,验证公钥的合法性; 
> * RA: Registration Authority,数字证书注册中心,负责证书申请者的信息录入,审核及证书发放等工作,是CA功能的一部分;
> * CRL: Certificate Revocation List,证书吊销列表,证书存取库;
> * X.509: 定义了证书的结构以及认证协议标准;
```
版本号
序列号
签名算法ID
发行者名称
有效期限
主体名称(主机证书要和主机名保持一致)
主题公钥
发行者惟一标识
主体的惟一标识
扩展
发行者签名
```
			
## SSL,TLS

> SSL: Secure Socket Layer

> TLS: Transport Layer Security
> * 1995: SSL 2.0,Netscape
> * 1996: SSL 3.0
> * 1999: TLS 1.0
> * 2000: TLS 1.1 RFC 4346
> * 2008: TLS 1.2
> * 2015: TLS 1.3

> 分层设计:
> * 1.最底层: 基础算法原始的实现,aes,rsa,md5
> * 2.向上一层: 各种算法的实现
> * 3.再向上一层: 组合算法实现的半成品
> * 4.用各种组件拼装而成的各种成品密码学协议/软件: tls,ssl
					
## OpenSSL

> 三个组件
> * openssl: 多用途的命令行工具;
> * libcrypto: 公共加密解密库;
> * libssl: 库,实现了ssl及tls;

> openssl命令:
```
openssl version #程序版本号

openssl命令分为3部分: 标准命令,消息摘要命令,加密命令

标准命令: enc,ca,req,...

对称加密:
    工具:openssl enc,gpg
    算法:3des,aes,blowfish,twofish
    enc命令:
      加密: openssl enc -e -des3 -a -salt -in fstab -out fstab.ciphertext
      解密: openssl enc -d -des3 -a -salt -in fstab.ciphertext -out fstab

单向加密:
    工具: md5sum,sha1sum,sha224sum,sha256sum,...,openssl dgst
    dgst命令: openssl dgst -md5 /PATH/TO/SOMEFILE
    md5命令: 
        $ md5sum fstab
        a1d07c4cd1ec01279fd02bd80b53c7a0  fstab
        $ openssl dgst -md5 fstab
        MD5(fstab)= a1d07c4cd1ec01279fd02bd80b53c7a0
    MAC: Message Authentication Code,单向加密的一种延伸应用,用于实现在网络通信中保证所传输的数据的完整性;
        机制:
            CBC-MAC
            HMAC:使用md5或sha1算法;

生成用户密码:
    passwd命令: openssl passwd -1 -salt SALT
    生成随机数: openssl rand -base64|hex NUM
        NUM:表示字节数;-hex时,每个字符4位,出现的字符数为NUM*2;
	
公钥加密:
    加密: 算法:RSA,ELGamal, 工具:gpg,openssl rsautl
    数字签名: 算法:RSA,DSA,ELGamal
        密钥交换: 算法:dh
    DSA:Digital Signature Algorithm
    DSS:Digital Signature Standard
    RSA:
					
生成密钥对: openssl genrsa -out /PATH/TO/PRIVATEKEY.FILE NUM_BITS
    $ (umask 077;openssl genrsa -out key.private 2048)

提取出公钥: 
    $ openssl rsa -in /PATH/TO/PRIVATEKEY -pubout
					
随机数生成器:
    /dev/random: 仅从熵池返回随机数,随机数用尽,阻塞;
    /dev/urandom: 从熵池返回随机数,随机数用尽,会利用软件生成伪随机数,非阻塞;
```

## OpenSSL建立私有CA

> OpenCA: openssl的二次封装,更为强大;

> 证书申请及签署步骤:
> * 1.生成申请请求;
> * 2.RA核验;
> * 3.CA签署;
> * 4.获取证书

> 创建私有CA: openssl的配置文件`/etc/pki/tls/openssl.cnf`
```
(1)创建所需要的文件
    # cd /etc/pki/CA
    # touch index.txt
    # echo 01 > serial

(2)CA自签
    生成密钥: 
      $ (umask 077;openssl genrsa -out private/cakey.pem 2048)
    自签证书:
      $ openssl -req -new -x509 -key private/cakey.pem -days 365 -out cacert.pem
        -new: 生成新证书签署请求;
        -x509: 生成自签证书;
        -key: 生成请求时用到的私钥文件;
        -days: 证书的有效期限;
        -out: 证书的保存路径

(3)CA签证/发证
    (a)用到证书的主机生成证书请求;
        $ (umask 077;openssl genrsa -out /etc/httpd/ssl/httpd.key 2048)
        $ openssl req -new -key /etc/httpd/ssl/httpd.key -days 365 -out /etc/httpd/ssl/httpd.csr
    (b)把请求文件传输给CA;
        $ scp /etc/httpd/ssl/httpd.csr user@CA_IP:/tmp/
    (c)CA签署证书;
        $ openssl ca in /tmp/httpd.csr -out /etc/pki/CA/certs/httpd.crt -days 365
        查看证书中的信息: openssl x509 -in /PATH/FROM/CERT_FILE -noout -text|-subject|-serial

(4)吊销证书
    (a)获取要吊销的证书的serial
        $ openssl x509 -in /PATH/FROM/CERT_FILE -noout -serial -subject
    (b)CA先根据客户提交的serial与subject信息,对比检验是否与index.txt文件中信息一致;
        吊销证书: openssl ca -revoke /etc/pki/CA/newcerts/SERIAL.pem
    (c)生成吊销证书编号(第一次吊销一个证书)
        $ echo 01 >/etc/pki/CA/crlnumber
    (d)更新证书的吊销列表: openssl ca -gencrl -out thisca.crl
       查看crl文件: openssl crl -in /PATH/FROM/CRL_FILE.crl --noout -text
```