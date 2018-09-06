//{{define "jsfile_withdraw"}}
//<script type="text/javascript">

var giveMoney = 0;
function setMoney(money) {
    giveMoney = money;
    $("#btn_give").html('补卡' + money + '元');
}
$(document).ready(function () {

});
$("#btn_give").click(function () {
    var guid = $("#txt_guid").val();
    console.log(guid);
    if (guid == '' || guid == undefined) {
        alert('请输入guid');
        return false;
    }
    if (giveMoney == 0) {
        alert('请选择金额');
        return false;
    }
    console.log(giveMoney);
    var info = "确定给用户:" + guid + " 补卡 [" + giveMoney + "元] ?";
    if (confirm(info)) {
        ajaxGetJson('withdraw/cardAdd/', { "guid": guid, "money": giveMoney * 100 }, function (data) {
            console.log(data);
            if (data.re) {
                alert('补卡成功!');
            } else {
                alert('补卡失败!');
            }

        }, function (err) {
            alert(err.msg);
        }, true);
    }
});
//</script>
//{{end}}