//{{define "jsCommon"}}
//<script type="text/javascript">
// Tools
String.prototype.ReplaceAll = function (f, e) {//吧f替换成e
    var reg = new RegExp(f, "g"); //创建正则RegExp对象   
    return this.replace(reg, e);
}

String.prototype.format = function (args) {
    if (arguments.length > 0) {
        var result = this;
        if (arguments.length == 1 && typeof (args) == "object") {
            for (var key in args) {
                var reg = new RegExp("({" + key + "})", "g");
                result = result.replace(reg, args[key]);
            }
        }
        else {
            for (var i = 0; i < arguments.length; i++) {
                if (arguments[i] == undefined) {
                    return "";
                }
                else {
                    var reg = new RegExp("({[" + i + "]})", "g");
                    result = result.replace(reg, arguments[i]);
                }
            }
        }
        return result;
    }
    else {
        return this;
    }
};

function GetRequest() {
    var url = location.search; //获取url中"?"符后的字串
    var theRequest = new Object();
    if (url.indexOf("?") != -1) {
        var str = url.substr(1);
        strs = str.split("&");
        for (var i = 0; i < strs.length; i++) {
            theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
    }
    return theRequest;
}

function deleteKeyInArray(array, item) {
    array.splice(jQuery.inArray(item, array), 1);
}

function deepCopy(json) {
    if (typeof json == 'number' || typeof json == 'string' || typeof json == 'boolean') {
        return json;
    } else if (typeof json == 'object') {
        if (json instanceof Array) {
            var newArr = [], i, len = json.length;
            for (i = 0; i < len; i++) {
                newArr[i] = arguments.callee(json[i]);
            }
            return newArr;
        } else {
            var newObj = {};
            for (var name in json) {
                newObj[name] = arguments.callee(json[name]);
            }
            return newObj;
        }
    }
}





function isJson(str) {
    try {
        var check_is_json = JSON.parse(str);
        return true;
    }
    catch (e) {
        return false;
    }
}


function scrollToEnd() {
    var h = $(document).height() - $(window).height();
    $(document).scrollTop(h);
}

//</script>
//{{end}}