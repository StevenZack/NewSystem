# 如何使用本文档?

例如：

### login

URL : /UserServlet/login

请求参数：

	values{
		openId		:	string类型，用户openId
		nickName	:	string类型，用户昵称
		head		:	string类型，用户头像图片url
	}

返回参数：

	ret{
		status		:	string类型,登录成功则为"OK",失败为"ERR"
		statusInfo	:	string类型，当status为"ERR"时，statusInfo为错误的详细描述
	}

看到以上文档，说明你可以用以下的代码来使用该接口：

``` javascript
api.ajax({
        url: host+'/UserServlet/login',
        method: 'post',
        data: {
            values: {
            	openId:openId,
                nickName: nick,
                file: head
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//登录成功
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

其中host为http://101.200.54.63:8080

---

<br>
<br>
<br>
<br>

# 目录

- [User](#user)
	- [/UserServlet/login](#login-1)
- [Address](#address)
	- [/AddressServlet/addAddress](#addaddress)
	- [/AddressServlet/findAll](#findall)

## User

### login

URL : /UserServlet/login

请求参数：

	values{
		openId		:	string类型，用户openId
		nickName	:	string类型，用户昵称
		head		:	string类型，用户头像图片url
	}

返回参数：

	ret{
		status		:	string类型,登录成功则为"OK",失败为"ERR"
		statusInfo	:	string类型，当status为"ERR"时，statusInfo为错误的详细描述
	}

## Address


### addAddress

URL : /AddressServlet/addAddress

请求参数

	values{
		openId		：	string类型，你要添加地址的用户openId
		address 	：	string类型，你哟啊添加的地址
	}

返回参数

	ret{
		status		:	string类型，成功则返回"OK"，失败返回"ERR"
		statusInfo	:	string类型，若失败，则statusInfo为失败原因
	}

### findAll

URL : /AdressServlet/findAll

请求参数

	values{
		openId		：	string类型，用户的openId
	}

返回参数

	ret{
		status		:	string类型,如果有错误则为"ERR"，无错误则为空字符串，	例如："ERR"
		statusInfo	:	string类型，当status为"ERR"时，statusInfo为错误的详细描述，	例如："该昵称已存在"
		openId		:	string类型，用户的唯一标识符，由服务器随机生成，	例如："fe6e1672cadd4f232e81a68ca17f8144"
		addresses	:	string数组类型，存储了用户的多个地址	例如：["太原理工大学","太原师范学院"]
		nickName	:	string类型，用户的昵称，登录的时候要用这个登录		例如："stevenzack"
		password	:	string类型，密码	例如："123456"
		avatar		:	string类型，用户头像的URL，例如："/pub/avatar/fe6e1672cadd4f232e81a68ca17f8144.png"
	}


