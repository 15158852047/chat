var vue = new Vue({
    delimiters:['[[', ']]'],
    el:"#myappfirst",
    data:{
        showTalk:false,
        userlist:[
            {"name":"早安无恙","lastMessage":"我是傻逼！，金少凯牛逼!","time":"下午 2：54","img":"static/images/head/15.jpg"},
            {"name":"夏继涛","lastMessage":"[小程序]","time":"上午 11：03","img":"static/images/head/2.jpg"},
            {"name":"十里老街秋名山车神车队","lastMessage":"乞讨两块交个话费","time":"昨天","img":"static/images/head/3.jpg"},
            {"name":"阿杰","lastMessage":"[动画表情]","time":"昨天","img":"static/images/head/4.jpg"},
        ],
        listactive:-1,
        activeUser:{},
        ws:null,
    },
    mounted:function () {
        this.showTalk = false;
        this.ws = new WebSocket('ws://'+window.location.host+"/ws");
        this.ws.onmessage = function (data) {
            var d = JSON.parse(data.data);
            if (d.code == 1) {
                
            }
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
        }
    }
});