//{{define "jsfile"}}
//<script type="text/javascript">
$(document).ready(function () {

});
// 统计查询
$("#btn_count").click(function () {
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
    $.getJSON("http://localhost:8088/action/count/?begin_date={0}&end_date={1}".format(begin_date, end_date), function (data) {
        if (data.code == 0) {
            console.log(data);
            var actions = data.data;
            var table = ""
            for (var key in actions) {
                table += "<tr><td>{0}</td><td>{1}</td><td>{2}</td></tr>".format(key, actions[key]['total'], actions[key]['uniqueCount']);
            }
            $("#action_result").html(table);
        } else {
            alert(data.msg);
        }
    });
});
//</script>
//{{end}}