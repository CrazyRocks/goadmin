//bootstrap-table配置

var baseURL = "http://localhost:8192/";

//工具集合Tools
window.T = {};

// 获取请求参数
// 使用示例
// location.href = http://localhost:8080/index.html?id=123
// T.p('id') --> 123;
var url = function (name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
};
T.p = url;

//全局配置
$.ajaxSetup({
    dataType: "json",
    cache: false
});

//重写alert
window.alert = function (msg, callback) {
    top.swal({
        text: msg,
        icon: "error",
        buttons: false,
		timer:2000,
    }).then((isConfirm) => {
        if (isConfirm) {
            if (typeof (callback) === "function") {
                callback("ok");
            }
        }
    });
}

//重写confirm式样框
window.confirm = function (msg, callback) {
    top.swal({
        text: msg,
        icon: "warning",
        buttons: ["取消", "确定"],
        closeModal: true,
    }).then((isConfirm) => {
        if (isConfirm) {
            if (typeof (callback) === "function") {
                callback("ok");
            }
        }
    });
}

//选择一条记录
function getSelectedRow() {
    var grid = $("#table");
    var rowKey = grid.getGridParam("selrow");
    if (!rowKey) {
        alert("请选择一条记录");
        return;
    }

    var selectedIDs = grid.getGridParam("selarrrow");
    if (selectedIDs.length > 1) {
        alert("只能选择一条记录");
        return;
    }

    return selectedIDs[0];
}

//选择多条记录
//选择多条记录
function getSelectedRows() {
    var selected = $("#table").bootstrapTable('getSelections');
    var ids = new Array();
    for (var i = 0; i < selected.length; i++) {
        ids[i] = selected[i].Id;
    }
    if (ids.length == 0) {
		alert("请选择记录");
    }
    console.log(ids);
    return ids;

}


//判断是否为空
function isBlank(value) {
    return !value || !/\S/.test(value)
}