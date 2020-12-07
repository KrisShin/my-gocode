function logout() {
    $.get("/api/logout", function (data) {
        if (200 == data.code) {
            location.href = "/";
        }
    })
}

// $('#to_profile').onclick(function () {
//     $.get(function (data) {
//         if (data.code == 200) {
//             location.href = 'profile.html'
//         }
//     })
// })

$(document).ready(function () {
    var id = $('#auth_status').html()
    if (id != 'None') {
        $('#auth_status').html('已认证')
        $('#auth_status').show()
    }
})