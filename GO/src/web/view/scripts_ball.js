//{{define "jsfile_ball"}}
//<script type="text/javascript">
$(document).ready(function () {
   
});
$("#btn_reset").click(function(){
    var guid=$("#txt_guid").val();
    console.log(guid);
    if (guid == '' || guid == undefined) {
        alert('请输入guid');
        return false;
    }
    var info = "确定给用户:" + guid + " 重置能量?" ;
    if (confirm(info)) {
        ajaxGetJson('ball/reset/', { "guid": guid }, function (data) {
            console.log(data);
            if (data.re) {
                alert('重置成功!');
            } else {
                alert('重置失败!');
            }
        }, function (err) {
            alert(err.msg);
        }, true);
    }
});
//</script>
//{{end}}