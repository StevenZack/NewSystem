package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_NewsAdd(t *testing.T) {
	e := addNews(`“分享禁令”能让微信摆脱“管道化”宿命吗？`, `5月18日，微信发布公告称，将在朋友圈禁止未取得证照转播视听节目，这就意味着目前尚未取得视听节目许可证的抖音、快手、梨视频、西瓜、秒拍等，都将无法通过微信朋友圈分享给好友，腾讯视频和微视因为拥有视听节目许可不受限制。

	与此同时，微信对于小游戏分享政策也进一步收紧。5月9日，微信发布公告称小游戏拒绝“分享滥用”，对具有强制分享、利诱分享等行为的小游戏进行能力限制惩罚，情节严重者将被下架。

	微信的这两条“分享禁令”虽然一则对外、一则对内，却都透露出微信拒绝“管道化”的努力。无论让朋友圈成为对手产品获取用户的管道，还是让微信群沦为小游戏病毒传播的管道，都是腾讯不愿意看到的。

	微信沦为管道，戴上“网络中立”紧箍咒

	微信的崛起曾经使运营商沦为“管道”，而一旦微信成为新的基础设施，则面临自身被“管道化”的命运，无论是曾经对网易云音乐、虾米的“封杀”还是对于短视频（艾瑞咨询数据显示，53.5%的用户认知短视频来自于微信／朋友圈）的“禁令”，不过都是不甘心为他人做嫁衣裳（免费输送流量）。

	然而，无论从生态开放，还是从网络中立（曾经适用于运营商的网络中立原则，也成了腾讯头上的“紧箍咒”）的角度，腾讯的屏蔽对手的行为都会被指责为“公器私用”。毕竟，腾讯无权干涉用户自发分享行为。

	腾讯近一年多来之所以力推大王卡（用户量已近1亿），也是因为并不能牺牲自身的“网络中立”去对自家产品做流量倾斜，只能向上一步去破坏运营商的“网络中立”。

	毕竟，腾讯不能向每一家“假道”社交关系链的公司征收“通行税”，也不可能对所有细分市场都投资一遍。腾讯变身投资公司，不过是“管道”分享流量红利的无奈之举。面对有赞这样的生态企业的做大，腾讯也只能通过投资see小电铺来予以制衡。

	碍于“管道”身份，腾讯又不能明目张胆地亲自下场，所以只能通过财务投资、暗中扶持这样“低调”的方式。尽管如此，业内人士还是会默认拼多多为腾讯嫡系，令云集、环球捕手这样的“野生派”艳羡不已。

	然而，即便是钦定“嫡系”，拼多多依然身在微信心思独立，利用微信导流到自家平台（拼多多来自App的订单量已经超过了50%），对此腾讯也只能睁一只眼闭一只眼。另外，腾讯投资的快手却因为没有得到流量倾斜，未能阻击抖音的狂飙崛起，也在心有异志地在打造自己的社交帝国。

	而且，在拼多多、蘑菇街之外的微信电商，微信除了投资see 小电铺这类SAAS工具之外，并不能像阿里那样坐享竞价排名、广告推广等丰厚的收入模式。毕竟，社交电商的流量是去中心化的，腾讯并不掌握流量分发大权。 进一步讲，通过社交分享和公号导流的小程序电商，也不过是把微信关系链当作获客的“管道”。

	这就像微信就为京东、蘑菇街、唯品会开辟了入口，拼多多依然享受最多的微信红利，因为它深植于微信的社交关系链中，对于微信管道的利用率最高，而中心化的官方入口反而效果不佳。

	微信沦为“管道”的另一个表现则是在内容领域，相比于头条在广告领域的盆满钵满，公号运营者们在通过微信吸引粉丝之后，纷纷找到电商、知识付费、定制内容、原生广告等自变现方式，很少选择广点通或互选广告。主要靠社交机制的自发传播，微信并不掌握流量分发大权，所以也很难像微博那样从内容创业者那里收取“流量费” （粉丝头条、粉丝通），只能博得一个“坚持不与开发者分利”的虚名。

	虽然去年曾经传出过订阅号变信息流的传闻，不过微信收回流量分发权必将破坏微信内容生态的根基，所以迟迟未能遂行。只能在潜藏的入口中上线了搜一搜、看一看这样的信息流板块，却因为不能撼动订阅号的入口地位，并未掀起波澜。

	社交体系不仅让微信沦为了内容方获取粉丝的管道，还大大降低了信息获取效率（今日头条的新slogan：信息创造价值简直是针锋相对）。某种程度上，订阅关系和社交传播是一种枷锁。在信息发现、信息消费、信息传播方面，在帮助内容创作者快速获取粉丝方面，微信早早就落于头条、微博后面，进入短视频时代之后更是如此，这或许可以解释头条和抖音的侧翼崛起。`,
		"钛媒体", "互联网", "http://images.tmtpost.com/uploads/images/2018/05/100e7382fecae80d7cac1bcd0f485088_1526787965.jpeg")
	t.Log(e)
	e = addNews(`任正非签发内部文整理：如何为华为延长生命？`, `我们公司所处的ICT行业有这样一个年龄层级分类：


爷爷辈：爱立信、诺基亚都是100多岁，IBM是快100岁，惠普78岁；

大叔辈：这一辈主要是三个70后和两个80后，微软、苹果、Oracle依次是75、76、77年成立的，三个70后，84年成立的思科和87年注册的华为，这两个是80后；

小鲜肉辈：包括Facebook、谷歌、亚马逊，以及中国的BAT，这些都是2000年以后成立的，到现在不到20岁。


我们发现，小鲜肉普遍很值钱，大叔们是渡过中年危机的值钱，爷爷们普遍不值钱。

任总在《一江春水向东流》里面说过一句话：“历史规律就是死亡，而我们责任是要延长生命。”

不管是个人，还是企业，最终都是要死的。我们的努力就是让死晚一点到来，不要过早地夭折。

我们今天讨论的熵减话题就是给我们一个视角，让我们看看如何努力延长企业的寿命，如何让熵减做好人才培养和激励，成为企业的活力根源。

企业之熵

我们先看下表来了解一下熵的概念。


熵增


混乱无效的增加，导致功能减弱失效。

人的衰老、组织的滞怠是自然的熵增,表现为功能逐渐丧失。


熵减


更加有效，导致功能增强。

通过摄入食物、建立效用机制，人和组织可以实现熵减，表现为功能增强。


负熵


带来熵减效应的活性因子。

物质、能量、信息是人的负熵，新成员、新知识、简化管理等是组织的负熵。

理解了熵的理论后，你会发现华为虽然最近这两三年才开始讲熵，经我们回溯发现，华为创立30年以来，从管理哲学阐述到各项政策制定，包括业务战略、人才管理等方面，隐约契合着耗散结构的特征，提倡与外界积极开展物质、能量、信息交换的开放精神。

有次我跟一位高管一起开会，谈到公司新业务时，他说，在老家有一句话，要判断一个家族是不是未来有希望，就要听他家院子中有没有婴儿的啼哭。植物发新芽、家中产婴儿这些都是熵减的现象。

这就给了我们的一个重要启示，也许历史、宇宙、人生的总趋势是熵增，但如果遵循耗散结构的规律，我们就有可能在熵增这个大环境下，构建一个熵减的小环境。

如果能让我们的公司、我们的家庭、我们的人生在一个小环境里实现熵减，对我们来说就是增进活力，延长寿命。

而对企业来说，一个企业正常的生命规律是从创业、萌发，然后到成长、成熟、衰退，最后死亡这样一个过程。所以现在华为面临的一些问题就是中年危机。当然遇到中年危机的也不只华为，所有成功的大公司如IBM、微软等都会遇到。问题只是你能不能应对中年危机。`,
		"钛媒体", "互联网", `http://images.tmtpost.com/uploads/images/2018/05/20180520122836183.jpg`)
	t.Log(e)
}
func addNews(title, content, publisher, typ string, images interface{}) string {
	n := News{}
	n.Title = title
	n.Content = content
	n.Publisher_title = publisher
	n.Type = typ
	if v, ok := images.(string); ok {
		n.Images = append(n.Images, v)
	}
	if v, ok := images.([]string); ok {
		n.Images = v
	}
	// host:="http://101.200.54.63:8080"
	host := "http://101.200.54.63:8080"
	b, e := json.Marshal(&n)
	if e != nil {
		return e.Error()
	}
	br := bytes.NewReader(b)
	rp, e := http.Post(host+"/NewsServlet/add", "application/json", br)
	if e != nil {
		return e.Error()
	}
	back, e := ioutil.ReadAll(rp.Body)
	if e != nil {
		return e.Error()
	}
	return string(back)
}
