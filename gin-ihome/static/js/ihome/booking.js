function hrefBack() {
    history.go(-1);
}

function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

function decodeQuery() {
    var search = decodeURI(document.location.search);
    return search.replace(/(^\?)/, '').split('&').reduce(function (result, item) {
        values = item.split('=');
        result[values[0]] = values[1];
        return result;
    }, {});
}

function showErrorMsg() {
    $('.popup_con').fadeIn('fast', function () {
        setTimeout(function () {
            $('.popup_con').fadeOut('fast', function () {
            });
        }, 1000)
    });
}

$(document).ready(function () {
    $(".input-daterange").datepicker({
        format: "yyyy-mm-dd",
        startDate: "today",
        language: "zh-CN",
        autoclose: true
    });
    $(".input-daterange").on("changeDate", function () {

        if (startDate && endDate && startDate > endDate) {
            showErrorMsg();
        } else {
            var startDate = $("#start-date").val();
            var endDate = $("#end-date").val();
            var sd = new Date(startDate);
            var ed = new Date(endDate);
            days = (ed - sd) / (1000 * 3600 * 24) + 1;
            var price = $(".house-text>p>span").html();
            var amount = days * parseFloat(price);
            $(".order-amount>span").html(amount.toFixed(2) + "(共" + days + "晚)");
        }
    });
    $('#submit-btn').bind('click', function () {
        var house_id = $('#house_id').val()
        var startDate = $("#start-date").val();
        var endDate = $("#end-date").val();
        var sd = new Date(startDate);
        var ed = new Date(endDate);
        days = (ed - sd) / (1000 * 3600 * 24) + 1;
        $.ajax({
            url: '/orders/booking/',
            type: 'POST',
            data: {'startDate': startDate, 'endDate': endDate, 'days': days, 'house_id': house_id},
            dataType: 'json',
            success: function (data) {
                if (data.code == 200) {
                    location.href = '/orders/render_orders/'
                }
            },
            error: function (data) {
                console.log(data.msg)
            }
        })
    })
})
