# 如何使用本文档?

例如：

### /UserServlet/register

请求参数：

	values:{
		nickName:str_nick,
		password:str_password
	},
	files:{
		file:image_avatar
	}

返回参数：

	ret{
		status		:	string类型,如果有错误则为"ERR"，无错误则为空字符串，	例如："ERR"
		statusInfo	:	string类型，当status为"ERR"时，statusInfo为错误的详细描述，	例如："该昵称已存在"
		openId		:	string类型，用户的唯一标识符，由服务器随机生成，	例如："fe6e1672cadd4f232e81a68ca17f8144"
		nickName	:	string类型，用户的昵称，登录的时候要用这个登录		例如："stevenzack"
		password	:	string类型，密码	例如："123456"
		avatar		:	string类型，用户头像的URL，例如："/pub/avatar/fe6e1672cadd4f232e81a68ca17f8144.png"
	}

看到以上文档，说明你可以用以下的代码来使用该接口：

``` javascript
api.ajax({
        url: host+'/UserServlet/register',
        method: 'post',
        data: {
            values: {
                nickName: nick,
                password:password
            },
            files: {
                file: image
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	alert(ret.openId) //把用户的openId存储到本地，以后会用到
        	//...
        	alert(ret.avatar) //这是用户的头像URL字符串，把host+ret.avatar插入到img标签的src属性里面即可使用
        } else {
            alert( ret?ret.statusInfo:err.msg );
        }
    });
```

其中host为http://101.200.54.63:8080

---

# 目录

- [User](#user)
-- [/UserServlet/register](#/userservlet/register)

## User

### /UserServlet/register