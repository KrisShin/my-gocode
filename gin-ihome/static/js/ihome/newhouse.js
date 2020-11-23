function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

$(document).ready(function () {
    // $('.popup_con').fadeIn('fast');
    // $('.popup_con').fadeOut('fast');
    $('#form-house-info').submit(function (e) {
        e.preventDefault()
        $(this).ajaxSubmit({
            url: '/house/new_house/',
            type: 'POST',
            dataType: 'json',
            success: function (data) {
                if (data.code == 200) {
                    // location.href = '/house/render_my_house'
                    $('#house-id').val(data.house_id)
                    $('#form-house-image').show()
                    $('#form-house-info').hide()
                }
            },
            error: function (data) {
                console.log(data.msg)
            }
        })
    })
    $('#form-house-image').submit(function (e) {
        e.preventDefault()
        $(this).ajaxSubmit({
            url: '/house/new_house_img/',
            type: 'POST',
            dataType: 'json',
            success: function (data) {
                if (data.code == 200) {
                    location.href = '/house/render_my_house/'
                }
            },
            error: function (data) {
                console.log(data.msg)
            }
        })
    })
})