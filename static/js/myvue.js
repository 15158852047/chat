var vue = new Vue({
    delimiters:['[[', ']]'],
    el:"#myappfirst",
    data:{
        showTalk:false,
        userlist:[  //聊天
            {"name":"早安无恙","lastMessage":"我是傻逼！，金少凯牛逼!","time":"下午 2：54","img":"static/images/head/15.jpg","username":"1","notRead":100},
        ],
        listactive:-1,
        activeUser:{},
        ws:null,
        wsChat:null,
        myname:'',
        myusername:'',
        myemail:'',
        users:[//好友列表 //{"A":[{"name":"","message":""},{"name":"","message":""}],"B":[{}]}
        ],
        mlists:[],
        pagecount:1,
        allct:0,
        haveMoreMsg:false,
        showRight:false,
    },
    mounted:function () {
        var _this = this;
        this.showTalk = false;
        this.ws = new WebSocket('ws://'+window.location.host+"/ws");

        this.$http.get("/friendlist").then(function (data) {
            _this.users = data.body;
        });

        this.$http.get("/chatlist").then(function (data) {
            _this.userlist = data.body;
        });

        _this.ws.onmessage = function (data) {
            var d = JSON.parse(data.data);
            if (d.code == 1) {
                _this.myname = d.name;
                _this.myusername = d.username;
                _this.myemail = d.email;
                _this.$notify({
                    title: "尊敬的"+d.name+":",
                    message: d.time,
                    position: 'bottom-right'
                });
            } else if (d.code == 2) {   ////聊天信息
                if (d.from == _this.myusername) {
                    _this.addOneMsg(d.msg,"me","static/images/own_head.jpg");
                    _this.topToList(d);
                } else if (d.from == _this.activeUser.username){
                    _this.addOneMsg(d.msg,"other","static/images/own_head.jpg");
                    _this.topToList(d);
                } else {
                    for (var i=0;i<_this.userlist.length;i++) {
                        if (_this.userlist[i].username == d.from) {
                            var inn = parseInt(_this.userlist[i].notRead);
                            inn++;
                            _this.userlist[i].notRead = ""+inn;
                            _this.topToList(d);
                            break;
                        }
                    }
                }
                _this.allct ++ ;
                _this.scrollToLow(_this.allct);
            }

        }
        _this.ws.onclose = function () {
            _this.ws.send(JSON.stringify({"code":"100","to":'',"msg":''}))
        }
    },
    methods:{
        closeTalk:function () {
            var _this = this;
            _this.showTalk = false;
        },
        showTalks:function () {
            var _this = this;
            _this.showTalk = true;
        },
        showActive:function (index) {
            var _this = this;
            _this.showTalks();
            _this.listactive = index;
            _this.activeUser = _this.userlist[index];
            _this.pagecount = 1;
            _this.getOneMsg();
        },
        sendOneMsg:function () {
            var text = document.getElementById('input_box');
            var s = text.value;
            var _this = this;
            if (text.value == '') {
                showMsg("消息不能为空!","top");
            } else {
                _this.ws.send(JSON.stringify({"code":"1","to":_this.activeUser.username,"msg":s}));
            }
            text.value = '';
        },
        chatWithOne:function (key,index) {
            var _this = this;
            _this.allct = 0 ;
            _this.haveMoreMsg = false;
            var one = _this.users[key][index];
            var a = {"name":one.name,"lastMessage":one.msg,"time":'现在',"img":one.img,"username":one.username};
            var hasChat = false;
            for(var j = 0,len = _this.userlist.length; j < len; j++){
                if(a.username == _this.userlist[j].username) {
                    hasChat = true;
                    document.getElementById('si_1').click();
                    _this.userlist = _this.moveToTop(j,_this.userlist);
                    _this.showActive(0);
                }
            }
            if( !hasChat) {
                var s = [];
                s.push(a);
                s = s.concat(_this.userlist);
                _this.userlist = s;
                document.getElementById('si_1').click();
                _this.showActive(0);
            }
            _this.getOneMsg();
        },
        moveToTop:function (index,arr) {
            var one = arr[index];
            arr.splice(index,1);
            arr.unshift(one);
            return arr
        },
        addOneMsg:function (str,people,img) {
            var _this = this;
            _this.mlists.push({"people":people,"img":img,"message":str});
        },
        getOneMsg:function () {
            var _this = this;
            _this.$http.get("/chatMsg",
                {params: {"username":_this.activeUser.username,"pagecount":_this.pagecount}}).then(function (data){
                    _this.mlists = [];
                    var x = 0;
                    for (var i in data.body.infos) {
                        x++;
                        _this.addOneMsg(data.body.infos[i].msg,data.body.infos[i].people,_this.activeUser.img);
                    }
                    var allcount = 0;
                    if (data.body.haveMore == 1) {
                        _this.haveMoreMsg = true;
                    }
                    if (_this.pagecount > 1) {
                        allcount = _this.pagecount * 10 + x;
                        _this.scrollToLow(allcount);
                    } else {
                        if (x>8) {
                            allcount = x;
                            _this.scrollToLow(allcount);
                        }
                    }
                    _this.allct = allcount;
                    if (_this.allct > 10) {
                        _this.haveMoreMsg = true;
                    }
                    for (var i=0;i<_this.userlist.length;i++) {
                        if (_this.userlist[i].username == _this.activeUser.username) {
                            _this.userlist[i].notRead = "0";
                            _this.activeUser.notRead = "0";
                        }
                    }
            });
        },
        scrollToLow:function (count) {
            var chatbox = document.getElementById("chatbox");
            chatbox.setAttribute("style","top: "+((-count+8)*50+20)+"px; position: absolute;");

        },
        readMore:function () {
            var _this = this;
            _this.pagecount ++ ;
            _this.$http.get("/chatMsg",
                {params: {"username":_this.activeUser.username,"pagecount":_this.pagecount}}).then(function (data) {
                    console.log(data.body);
                var m = [];
                for (var i in data.body.infos) {
                    _this.allct ++ ;
                   // _this.addMoreMsg(data.body.infos[i].msg, data.body.infos[i].people, _this.activeUser.img);
                    m.push({"people":data.body.infos[i].people,"img":_this.activeUser.img,"message":data.body.infos[i].msg});
                }
                _this.addMoreMsg(m);
                if (data.body.haveMore == 0) {
                    _this.haveMoreMsg = false;
                }
            });
        },
        addMoreMsg:function (m) {
            var _this = this;
            for (var i=m.length-1;i>-1;i--) {
                _this.mlists.unshift(m[i]);
            }

        },
        showValue:function (i) {
            var x = parseInt(this.userlist[i].notRead);
            return x;
        },
        canShow:function (i) {
            return this.userlist[i].notRead>0;
        },
        topToList:function (d) {
            var _this = this;
            this.$http.get("/chatlist").then(function (data) {
                _this.userlist = data.body;
            });
        }
    }
});