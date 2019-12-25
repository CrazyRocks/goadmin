$(function () {

    function addButtonFunc(value, row, index) {
        return '<div id="deleteButton"  class="badge badge-danger m-1"> <i class="fas fa-times mr-1"></i>删除</div>';
    }

    window.operateEvents = {
        'click #deleteButton': function (e, value, row, index) {
            vm.del(row.Id);
        }
    };

    $('#table').bootstrapTable({
        url: baseURL + 'sys/oss/page',
        method: "GET",
        striped: true,
        cache: false,
        pagination: true,
        pageList: [20, 40, 60, 100],
        pageSize: 20,
        pageNumber: 1,
        sortName: "id",
        sortOrder: "desc",
        sidePagination: 'server',
        search: false,
        uniqueId: "Id",
        silent: true,
        classes: "table table-hover",
        paginationHAlign: "left",
        paginationDetailHAlign: "right",
        queryParams: queryParams,
        responseHandler: function (res) {
            return {
                "total": res.data.form.TotalSize,
                "rows": res.data.list
            };

        },
        onLoadSuccess: function () {

        },
        onLoadError: function () {
            alert("数据加载失败！");
        },
        columns: [{
            checkbox: true,
            visible: true
        },
            {
                field: 'Id',
                title: 'ID'
            },
            {
                field: 'Url',
                title: '地址'
            },
            {
                field: 'operate',
                title: '操作',
                events: operateEvents,
                formatter: addButtonFunc
            }
        ]
    });

    function queryParams(params) {
        var temp = {
            offset: params.offset,
            limit: params.limit,
            search: $(".search-input").val(),
            rows: params.limit,
            page: (params.offset / params.limit) + 1,
            sort: params.sort,
            sortOrder: params.order
        };
        return temp;
    };
    $("#search-btn").click(function () {
        $('#table').bootstrapTable('refresh');
    });
    new AjaxUpload('#upload', {
        action: baseURL + "sys/oss/upload",
        name: 'upload-file',
        autoSubmit: true,
        responseType: "json",
        onSubmit: function (file, extension) {
            if (vm.ParamValue == null) {
                alert("云存储配置未配置");
                return false;
            }
            if (!(extension && /^(jpg|jpeg|png|gif)$/.test(extension.toLowerCase()))) {
                alert('只支持jpg、png、gif格式的图片！');
                return false;
            }
        },
        onComplete: function (file, r) {
            if (r.code == 0) {
                vm.reload();
            } else {
                alert(r.msg);
            }
        }
    });
});

var vm = new Vue({
        el: '#rrapp',
        data: {
            showList: true,
            title: null,
            sysoss: {},
            config: {
                ParamValue: {},
            },
            ParamValue: {}
        },
        created: function () {
            this.getConfig();
        },
        methods: {
            query: function () {
                vm.reload();
            },

            edit: function () {
                vm.showList = false;
                vm.title = "修改配置";
                vm.config = {};
                vm.getConfig()
            },
            upload: function () {
                vm.title = "修改配置";
            },
            saveConfig: function (event) {
                $('#btnSaveOrUpdate').button('loading').delay(1000).queue(function () {
                    var url = "sys/oss/config/save";
                    vm.config.ParamValue = JSON.stringify(vm.ParamValue)
                    $.ajax({
                        type: "POST",
                        url: baseURL + url,
                        data: vm.config,
                        success: function (r) {
                            if (r.code === 0) {
                                swal({
                                    text: "操作成功",
                                    icon: "success",
                                    buttons: false,
                                    timer: 2000,
                                });
                                vm.reload();
                                $('#btnSaveOrUpdate').button('reset');
                                $('#btnSaveOrUpdate').dequeue();
                            } else {
                                swal({
                                    text: r.msg,
                                    icon: "error",
                                    buttons: false,
                                    timer: 2000,
                                });
                                $('#btnSaveOrUpdate').button('reset');
                                $('#btnSaveOrUpdate').dequeue();
                            }
                        }
                    })
                    ;
                });
            },
            del: function (Id) {
                var Ids = [Id];
                var lock = false;
                top.swal({
                    title: "确定要删除该记录?",
                    icon: "warning",
                    buttons: ["取消", "确定"],
                    closeModal: true,
                }).then((isConfirm) => {
                    if (isConfirm) {
                        top.swal.close();
                        if (!lock) {
                            lock = true;
                            $.ajax({
                                type: "POST",
                                url: baseURL + "sys/oss/delete",
                                data: {ids: Ids},
                                success: function (r) {
                                    if (r.code == 0) {
                                        swal({
                                            text: "删除成功",
                                            icon: "success",
                                            buttons: false,
                                            timer: 2000,
                                        });
                                        $('#table').bootstrapTable('refresh');
                                    } else {
                                        swal({
                                            text: r.msg,
                                            icon: "error",
                                            buttons: false,
                                            timer: 2000,
                                        });
                                    }
                                }
                            });
                        }
                    } else {
                        top.swal.close();
                    }
                });
            },
            bulkdel: function (event) {
                var Ids = getSelectedRows();
                if (Ids == null || Ids.length == 0) {
                    return;
                }
                var lock = false;
                top.swal({
                    title: "确定要删除该记录?",
                    icon: "warning",
                    buttons: ["取消", "确定"],
                    closeModal: true,
                }).then((isConfirm) => {
                    if (isConfirm) {
                        top.swal.close();
                        if (!lock) {
                            lock = true;
                            $.ajax({
                                type: "POST",
                                url: baseURL + "sys/oss/delete",
                                data: {ids: Ids},
                                success: function (r) {
                                    if (r.code == 0) {
                                        swal({
                                            text: "删除成功",
                                            icon: "success",
                                            buttons: false,
                                            timer: 2000,
                                        });
                                        $('#table').bootstrapTable('refresh');
                                    } else {
                                        swal({
                                            text: r.msg,
                                            icon: "error",
                                            buttons: false,
                                            timer: 2000,
                                        });
                                    }
                                }
                            });
                        }
                    } else {
                        top.swal.close();
                    }
                });
            },
            getConfig: function () {
                $.get(baseURL + "sys/oss/config/get", function (r) {
                    vm.config = r.data
                    vm.ParamValue = JSON.parse(r.data.ParamValue);
                });
            },
            reload: function (event) {
                vm.showList = true;
                $('#table').bootstrapTable('refresh');
            }
        }
    })
;