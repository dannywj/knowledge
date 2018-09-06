//{{define "jsCommon"}}
//<script type="text/javascript">
// Tools
var g_controller_base = 'http://localhost:9999/';

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

function ajaxGetJson(url, params, callback, error_callback, page_loading) {
    if (page_loading) {
        loadingPage();
    }
    $.getJSON(g_controller_base+url, params, function (data) {
        if (page_loading) {
            loadingPage(true);
        }
        if (data.code === 0) {
            callback(data.data);
        } else {
            if (error_callback) {
                error_callback(data);
            } else {
                console.error(data);
                switch (data.status) {
                    case 201:
                        alert("Invalid user identity，please login");
                        window.location = 'login.html';
                        break;
                    case 203:
                        alert("Invalid permission");
                        break;
                    default:
                        alert("get api[{0}] error!".format(s + ' ' + f));
                }
            }
        }
    });
}

function ajaxPostJson(s, f, params, callback, error_callback, page_loading) {
    if (page_loading) {
        loadingPage();
    }
    $.post(g_controller_base + "?s={0}&f={1}".format(s, f), params, function (data) {
        if (page_loading) {
            loadingPage(true);
        }
        if (data.status === 0) {
            callback(data.data);
        } else {
            if (error_callback) {
                error_callback(data);
            } else {
                console.error(data);
                switch (data.status) {
                    case 201:
                        alert("Invalid user identity，please login");
                        window.location = 'login.html';
                        break;
                    case 203:
                        alert("Invalid permission");
                        break;
                    default:
                        alert("post api[{0}] error!".format(s + ' ' + f));
                }
            }
        }
    });
}

function loadingDiv(divSelecter) {
    $(divSelecter).html("<img src='images/loading.gif' style='margin: 0 auto;display: block;'/>");
}

function loadingPage(is_hide) {
    var loading_html = '<section class="loading_shade" id="J_loading_box"> <div class="loading_box"> <div class="loading"></div> <p class="loading_text">Loading...</p> </div> </section>';
    if (is_hide) {
        $(".loading_shade").hide();
    } else {
        $("body").append(loading_html);
    }
}

//</script>
//{{end}}