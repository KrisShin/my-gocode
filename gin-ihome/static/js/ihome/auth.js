function showSuccessMsg() {
    $('.popup_con').fadeIn('fast', function () {
        setTimeout(function () {
            $('.popup_con').fadeOut('fast', function () {
            });
        }, 1000)
    });
}

$('#form-auth').submit(function (e) {
    e.preventDefault()
    var id_name = $('#real-name').val()
    var id_card = $('#id-card').val()
    $.ajax({
        url: '/users/auth/',
        type: 'POST',
        dataType: 'json',
        data: {'id_name': id_name, 'id_card': id_card},
        success: function (data) {
            if (data.code == 200) {
                alert(333)
                location.href = '/users/render_auth/'
            }
        },
        error: function (data) {
            console.log(data.msg)
        }
    })
})

$(document).ready(function () {
    var id_name = $('#real-name')
    var id_card = $('#id-card')
    if (id_name.val() != 'None' && id_card.val() != 'None') {
        $('#save_btn').css('display', 'none')
        id_name.attr('disabled', 'true')
        id_card.attr('disabled', 'true')
    } else {
        id_name.val('')
        id_card.val('')
    }
})