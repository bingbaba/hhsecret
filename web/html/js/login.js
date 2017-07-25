$(function () {
    redo = $.getUrlParam("redo");
    //logined
    if((redo == undefined || redo != "1") && checkLogined()) {
        window.location.href = 'sign.html';
    }
});

function loginHandle(){
    var username = $("#username").val();
    var password = $("#password").val();

    var loading = weui.loading('登录中');
    if(login(username, password)){
        loading.hide();
        console.log("login succ!");
        window.location.href = 'sign.html';
    }else {
        loading.hide();
        weui.topTips('登录失败，请检查用户名或密码是否正确', 5000);
    }
}

function login(username, password){
    if(username == undefined || username == "" || password == undefined || password == ""){
        return false;
    }

    var user = {"password": password};
    resp=$.ajax({
        type: 'POST',
        url: "/api/user/"+username+"/login",
        dataType: "json",
        data: JSON.stringify(user),
        async: false
    }).responseJSON;

    console.log(resp);
    if(resp.msg == "OK"){
        localStorage.username = username;
        localStorage.name = resp.data.name;
        return true
    }else {
        weui.topTips(resp.msg, 5000);
        return false
    }
}

function checkLogined() {
    if(localStorage.username == undefined || localStorage.username == ""){
        return false;
    }

    var username = localStorage.username;
    resp=$.ajax({
        type: 'GET',
        dataType: "json",
        url: "/api/user/"+username+"/login",
        async: false
    }).responseJSON;

    if(resp.data.name != undefined && resp.data.name != ""){
        return true
    }else {
        return false
    }
}