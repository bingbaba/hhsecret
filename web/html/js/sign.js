var tmpl_sign = '<div class="weui-form-preview sign_info">'+
            '    <div class="weui-form-preview__hd">'+
            '        <div class="weui-form-preview__item">'+
            '            <label class="weui-form-preview__label">{0}时间</label>'+
            '            <em class="weui-form-preview__value">{1}</em>'+
            '        </div>'+
            '    </div>'+
            '    <div class="weui-form-preview__bd">'+
            '        <div class="weui-form-preview__item">'+
            '            <label class="weui-form-preview__label">地点</label>'+
            '            <span class="weui-form-preview__value">{2}</span>'+
            '        </div>'+
            '        <div class="weui-form-preview__item">'+
            '            <label class="weui-form-preview__label">经纬度</label>'+
            '            <span class="weui-form-preview__value">{3},{4}</span>'+
            '        </div>'+
            '    </div>'+
            '    <div class="weui-form-preview__ft">'+
            '    </div>'+
            '</div>';

var signs = [];
function renderListSign(data){
    signs = data.signs;
    var max_i = data.signs.length-1;

    $(".sign_info").remove();
    for (var i = max_i; i >= 0; i--) {
        var signtype, time_str, location, lat, lng;
        if(i==max_i) {
            signtype = "上班";
        }else {
            signtype = "下班";
        }
        time_str = new Date(data.signs[i].datetime).Format("HH:mm");
        location = data.signs[i].featurename;
        lat = data.signs[i].latitude;
        lng = data.signs[i].longitude;

        $("#sign_btn").before(tmpl_sign.format(signtype, time_str, location, lat, lng))
    }
}

function signHandle(){
    var titile_prefix = "";
    if(signs.length > 0){
        titile_prefix = "您今天已打卡"+signs.length+"次"
    }else {
        titile_prefix = "您今天还尚未打卡"
    }

    weui.confirm(titile_prefix+'，请确认是否继续打卡？', {
        title: "打卡操作确认",
        buttons: [{
            label: '取消',
            type: 'default'
        }, {
            label: '坚持打卡',
            type: 'primary',
            onClick: function(){
                
                sign_flag = _sign()
                
                if(sign_flag) {
                    refresh();
                }
            }
        }]
    });
}

function _sign(){
    var loading = weui.loading('正在打卡');
    resp=$.ajax({
        type: 'POST',
        dataType: "json",
        url: "/api/user/"+localStorage.username+"/sign",
        async: false
    }).responseJSON;
    // resp = {msg:"OK"}

    loading.hide();
    if(resp.msg == "OK") {
        weui.toast('打卡成功', 2000);
        return true;
    }else {
        weui.topTips(resp.msg, 5000);
        return false;
    }
}

function refresh(){
    //logined
    if(localStorage.name != undefined) {
        var name = localStorage.name;
        var username = localStorage.username;
        $("#myname").text(name+"("+username+")");

        var loading = weui.loading('加载中');
        $.ajax({
            type: 'GET',
            url: "/api/user/"+username+"/sign",
            dataType: "json",
            success: function(data){
                console.log("list sign suc!");
                
                loading.hide();

                if(data.msg == "user not login"){
                    location = 'login.html?redo=1';
                }else if(data.msg == "OK") {
                    renderListSign(data.data);
                }else {
                    weui.alert(data.msg);
                }
            },
            error: function(){
                loading.hide()
                weui.topTips("获取打卡记录失败", 5000);
            }
        });

        // 
    }else {
        window.location.href = 'login.html';
    }
}

$(function () {
    refresh();
});