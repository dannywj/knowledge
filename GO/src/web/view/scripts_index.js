//{{define "jsfile_index"}}
//<script type="text/javascript">
$(document).ready(function () {
    ajaxGetJson('base/info/', {  }, function (data) {
        console.log(data);
        $("#total_count").html(data.total);
        $("#today_count").html(data.todayCount);
        $("#today_reward_user_count").html(data.todayGetRewardUserCount);
    }, function (err) {
        alert(err.msg);
    }, true);
});
//</script>
//{{end}}