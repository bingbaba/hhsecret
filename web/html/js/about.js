var tmpl_exception_sign = '<div class="weui-cell">'+
'  <div class="weui-cell__bd">'+
'    <p>{0}</p>'+
'  </div>'+
'  <div class="weui-cell__ft redlink">{1}</div>'+
'</div>';


function getCurMonthSign(){
    date = new Date();
    year = date.getYear()+1900;
    month = date.getMonth()+1;
    username = localStorage.username;

    $("#exceptionSigns .weui-cell").remove();
    $.ajax({
        type: 'GET',
        url: "/api/user/"+username+"/sign/month/"+year+"/"+month,
        dataType: "json",
        success: function(data){
            console.log("list month sign suc!");
            if(data.msg == "OK") {
                array = data.data.arr.split(";")
                var execptionNum = 0;
                for (var i = array.length - 1; i >= 0; i--) {
                    if(array[i] == ""){
                        continue;
                    }else if(array[i].indexOf("迟到") >= 0 ||
                        array[i].indexOf("缺勤") >= 0 
                    ){
                        execptionNum++;
                        var day = i+1;
                        var esign = tmpl_exception_sign.format(month+"月"+day+"号", array[i]);
                        $("#exceptionSigns").append(esign);
                    }
                }
            }else {
                weui.alert(data.msg);
            }
        },
        error: function(){
            weui.topTips("获取本月打卡记录失败", 5000);
        }
    });
}

$(function () {
    if(localStorage.name != undefined) {
        $("#avatar").attr("src", localStorage.photoUrl);
        $("#name").text(localStorage.name);
        $("#jobtitle").text(localStorage.jobTitle);
        getCurMonthSign();
    }else {
        window.location.href = 'login.html';
    }
});