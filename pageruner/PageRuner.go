/**
 * @Author: FLYKO
 * @Description:
 * @File: JomePageRuner
 * @Version: 1.0.0
 * @LOVE LIN FOREVER
 * @Date: 2021/11/09：32
 * @本程序专门用于测试代理模块
 * @Website: www.jomeidc.com
 */

package main

import (
	"github.com/kataras/iris/v12"
)

type picture struct {
	ID      string
	URL     string
	Subject string
}

func main() {
	app := iris.New()
	tmpl := iris.Django("./Template", ".html")
	tmpl.Reload(false)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})
	app.RegisterView(tmpl)

	/*************************************************************************
	 * ///////////////////////////////////////////////////////////////////////
	 *                JomePageRuner V 1.0 Powered By FlyKO
	 * ///////////////////////////////////////////////////////////////////////
	 ************************************************************************/

	app.Get("/{api:string}", func(ctx iris.Context) {
		api := ctx.Params().Get("api")
		if(api == "set"){
			ctx.SetCookieKV("c","OJBK")
			ctx.HTML("API:" + "SetCookies" + "<br>" + "Cookies:")
		}else{
		//ctx.HTML("API:" + "NO!" + "<br>" + "Cookies:"+ctx.GetCookie("c"))
		ctx.HTML(`<body>
<script language="JavaScript" type="text/javascript">
    //设置两个cookie
    document.cookie="userId=u001";
    document.cookie="userName=lazyCat";
    //获取cookie字符串
    var strCookie=document.cookie;
    //将多cookie切割为多个名/值对
    var arrCookie=strCookie.split("; ");
    var userName;
    //遍历cookie数组，处理每个cookie对
    for(var i=0;i<arrCookie.length;i++){
        var arr=arrCookie[i].split("=");
//找到名称为userId的cookie，并返回它的值
        if("userName"==arr[0]){
            userName=arr[1];
            break;
        }
    }
    alert(userName);
</script>
</body>`)
		}
	})

	//INDEX Page STATIC
	app.HandleDir("/js", "./static/js")

	app.Run(iris.Addr(":8899"))
	//app.Run(iris.TLS("0.0.0.0:443","m.crt","m.key"))

}


