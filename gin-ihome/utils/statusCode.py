# success code 200
OK = 200
SUCCESS = {'code': 200, 'msg': 'success'}
# regiser and login errors code from 1000 to 1100
USER_REGISTER_PARAMS_NOT_COMPLETE = {'code': 1000, 'msg': '请填写所有信息'}
USER_REGISTER_PHONE_WRONG_FORMAT = {'code': 1001, 'msg': '电话号码错误'}
USER_REGISTER_WRONG_CODE = {'code': 1002, 'msg': '验证码错误'}
USER_REGISTER_WRONG_CONFIRM_PASSWORD = {'code': 1003, 'msg': '两次输入密码不一致'}
USER_REGISTER_USER_IS_EXISTED = {'code': 1004, 'msg': '用户已存在,请登录'}

USER_LOGIN_USER_NOT_EXIST = {'code': 1005, 'msg': '用户不存在,请注册'}
USER_LOGIN_WRONG_PASSWORD = {'code': 1006, 'msg': '密码错误'}

# updete my infomation errors code from 1101 to 1200
USER_INFO_IS_NOT_EXIST = {'code': 1101, 'msg': '请填写修改内容'}

# auth my info errors code from 1201 to 1300
USER_AUTH_INFO_ERROR = {'code': 1201, 'msg': '认证信息有误,认证失败'}
USER_AUTH_ID_CARD_DUPLICATE = {'code': 1201, 'msg': '身份证号已注册,认证失败'}

# add house error code from 1301 to 1400
HOUSE_ADD_INFO_NOT_COMPLETE = {'code': 1301, 'msg': '请完整填写房屋信息'}
# HOUSE_ADD_INFO_NOT_COMPLETE = {'code': 1302, 'msg': ''}
