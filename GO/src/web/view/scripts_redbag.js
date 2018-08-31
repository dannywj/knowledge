//{{define "jsfile"}}
//<script type="text/javascript">
$(document).ready(function () {

});
// 统计查询
$("#btn_redbag").click(function () {
    var begin_date = $("#txt_begin_date").val();
    var end_date = $("#txt_end_date").val();
    if (!begin_date || !end_date) {
        alert("请选择时间段");
        return false;
    }
    if (begin_date > end_date) {
        alert("开始时间应小于结束时间");
        return false;
    }
    console.log(begin_date.ReplaceAll('-', ''));
    console.log(end_date.ReplaceAll('-', ''));
    begin_date = begin_date.ReplaceAll('-', '');
    end_date = end_date.ReplaceAll('-', '');

    ajaxGetJson('redbag/info/', { begin_date: begin_date, end_date: end_date }, function (data) {
        console.log(data);
        var actions = data;
        var table = ""
        for (var key in actions) {
            table += "<tr><td>{0}</td><td>{1}</td></tr>".format(key, actions[key]['total']);
        }
        $("#action_result").html(table);
    }, function (err) {
        alert(err.msg);
    }, true);
});
//</script>
//{{end}}