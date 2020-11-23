$(document).ready(function () {
    $(".auth-warn").show();
    var id = $('#auth_status').html()
    if (!id) {
        $('#houses-list').hide()
    } else {
        $('#auth_warn').hide()
    }

})