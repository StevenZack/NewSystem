# 目录

- [User](#user)
	- /UserServlet/login
- [Address](#address)
	- /AddressServlet/addAddress
	- /AddressServlet/findAll
- [Moment](#moment)
	- /MomentServlet/upload
	- /MomentServlet/findAll
- [News](#news)
	- /NewsServlet/findAllByKind
	- /NewsServlet/findAll
	- /NewsServlet/findPhotoBrowserByNewsId
- [Search](#search)
	- /SearchServlet/findproductByname
	- /SearchServlet/findproductByname_English
	- /SearchServlet/findproductBytype
	- /SearchServlet/fuzzy
- [Order](#order)
	- /OrderServlet/findorderByopenid
	- /OrderServlet/findOrdersByStatus
	- /OrderServlet/addOrder
	- /OrderServlet/payOrder

## User

### /UserServlet/login

``` javascript
api.ajax({
        url: host+'/UserServlet/login',
        method: 'post',
        data: {
            values: {
            	openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830',
                nickName: 'stevenzack',
                file: 'http://thirdqq.qlogo.cn/qqapp/1106895244/CC8FD2EA13FDB7546E2E1DA63FE57903/100'
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

## Address

### /AddressServlet/addAddress

``` javascript
api.ajax({
        url: host+'/AddressServlet/addAddress',
        method: 'post',
        data: {
            values: {
            	openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830',
            	address:'山西省晋中市太原理工大学'
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//添加地址成功
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /AddressServlet/findAll

``` javascript
api.ajax({
        url: host+'/AddressServlet/findAll',
        method: 'post',
        data: {
            values: {
            	openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830',
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var addrs = ret.addresses //这是一个字符串数组
        	//...
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

# Moment

### /MomentServlet/upload

``` javascript
api.ajax({
        url: host+'/MomentServlet/upload',
        method: 'post',
        data: {
            values: {
            	openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830',
            	text:'今天天气真好',
            	address:'山西太原'
            },
            files:{
            	images:['/sdcard/a.png','/sdcard/b.png']
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//发布成功
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /MomentServlet/findAll


``` javascript
api.ajax({
        url: host+'/MomentServlet/findAll',
        method: 'post',
        data: {
            values: {
            	openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830'
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var mos = ret.moments //这是一个Moment对象的数组，
        	//访问案例如下：
        	mos[0].text
        	mos[0].address
        	mos[0].createTime
        	mos[0].images[0] // images是一个字符串数组
        	mos[0].images[1]
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

# News

### /NewsServlet/findAllByKind

``` javascript
api.ajax({
        url: host+'/NewsServlet/findAllByKind',
        method: 'post',
        data: {
            values: {
            	type:'手机'
            }
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var newsArray = ret.newsArray //newsArray是一个News对象的数组
        	//访问案例如下
        	newsArray[0].news_id
        	newsArray[0].title
        	newsArray[0].content
        	newsArray[0].createTime
        	newsArray[0].publisher_title
        	newsArray[0].images[0]
        	newsArray[0].images[1]
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /NewsServlet/findAll

``` javascript
api.ajax({
        url: host+'/NewsServlet/findAll',
        method: 'post'
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var newsArray = ret.newsArray //newsArray是一个News对象的数组
        	//访问案例如下
        	newsArray[0].news_id
        	newsArray[0].title
        	newsArray[0].content
        	newsArray[0].createTime
        	newsArray[0].publisher_title
        	newsArray[0].images[0]
        	newsArray[0].images[1]
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /NewsServlet/findPhotoBrowserByNewsId

``` javascript
api.ajax({
        url: host+'/NewsServlet/findPhotoBrowserByNewsId',
        method: 'post',
        data: {
        	values: {
        		news_id:'3CB9DD5FE6E15DFE3860A0A96A9C3830'
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//ret在此处是一个News对象
        	//访问案例如下
        	ret.news_id
        	ret.title
        	ret.content
        	ret.createTime
        	ret.publisher_title
        	ret.images[0]
        	ret.images[1]
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

# Search

### /SearchServlet/findproductByname

``` javascript
api.ajax({
        url: host+'/SearchServlet/findproductByname',
        method: 'post',
        data: {
        	values: {
        		query:'华为' //这里传如要查询的关键字
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var pros = ret.products //这里pros是一个Product对象的数组
        	//使用案例如下
        	pros[0].productId
        	pros[0].images[0]
        	pros[0].images[1]
        	pros[0].type
        	pros[0].price
        	pros[0].name
        	pros[0].englishName
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /SearchServlet/findproductByname_English

``` javascript
api.ajax({
        url: host+'/SearchServlet/findproductByname_English',
        method: 'post',
        data: {
        	values: {
        		query:'Huawei' //这里传如要查询的关键字
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var pros = ret.products //这里pros是一个Product对象的数组
        	//使用案例如下
        	pros[0].productId
        	pros[0].images[0]
        	pros[0].images[1]
        	pros[0].type
        	pros[0].price
        	pros[0].name
        	pros[0].englishName
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /SearchServlet/findproductBytype

``` javascript
api.ajax({
        url: host+'/SearchServlet/findproductBytype',
        method: 'post',
        data: {
        	values: {
        		query:'手机' //这里传如要查询的类型
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var pros = ret.products //这里pros是一个Product对象的数组
        	//使用案例如下
        	pros[0].productId
        	pros[0].images[0]
        	pros[0].images[1]
        	pros[0].type
        	pros[0].price
        	pros[0].name
        	pros[0].englishName
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /SearchServlet/fuzzy

``` javascript
api.ajax({
        url: host+'/SearchServlet/fuzzy',
        method: 'post',
        data: {
        	values: {
        		query:'华' //这里传如要查询的类型
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var pros = ret.products //这里pros是一个Product对象的数组
        	//使用案例如下
        	pros[0].productId
        	pros[0].images[0]
        	pros[0].images[1]
        	pros[0].type
        	pros[0].price
        	pros[0].name
        	pros[0].englishName
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

# Order

### /OrderServlet/findorderByopenid

``` javascript
api.ajax({
        url: host+'/SearchServlet/findproductBytype',
        method: 'post',
        data: {
        	values: {
        		openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830'
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var orders = ret.orders //这里orders是一个Order对象的数组
        	//orders使用案例如下
        	orders[0].orderId
        	orders[0].orderStatus //订单状态
        	orders[0].createTime
        	orders[0].productId //购买的产品id
        	orders[0].num //购买数量
        	orders[0].price //总价格
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /OrderServlet/findOrdersByStatus

``` javascript
api.ajax({
        url: host+'/SearchServlet/findOrdersByStatus',
        method: 'post',
        data: {
        	values: {
        		openId:'3CB9DD5FE6E15DFE3860A0A96A9C3830',
        		orderStatus:'待付款'
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	var orders = ret.orders //这里orders是一个Order对象的数组
        	//orders使用案例如下
        	orders[0].orderId
        	orders[0].orderStatus //订单状态
        	orders[0].createTime
        	orders[0].productId //购买的产品id
        	orders[0].num //购买数量
        	orders[0].price //总价格
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /OrderServlet/addOrder

``` javascript
api.ajax({
        url: host+'/SearchServlet/addOrder',
        method: 'post',
        data: {
        	values: {
        		buyerId:'3CB9DD5FE6E15DFE3860A0A96A9C3830', //就是你的openId
        		productId:'3CB9DD5FE6E15DFE3860A0A96A9C3830', //你要买的产品的id
        		num:1 //只买一件
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//下单成功
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```

### /OrderServlet/payOrder

``` javascript
api.ajax({
        url: host+'/SearchServlet/addOrder',
        method: 'post',
        data: {
        	values: {
        		orderId:'3CB9DD5FE6E15DFE3860A0A96A9C3830' //你要付款的订单的id
        	}
        }
    },function(ret, err){
        if (ret&&ret.status!='ERR') {
        	//付款成功
        } else {
            alert( ret?('Server err:'+ret.statusInfo):err.msg );
        }
    });
```