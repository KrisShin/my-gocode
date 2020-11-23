import os

# 项目根目录
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# static路径
STATIC_DIR = os.path.join(BASE_DIR, 'static')
# templates路径
TEMPLATES_DIR = os.path.join(BASE_DIR, 'templates')
# media路径
MEDIA_DIR = os.path.join(STATIC_DIR, 'media')
# 图片上传绝对路径
UPLOAD_DIR = os.path.join(MEDIA_DIR, 'upload')
# 图片显示相对路径
SHOW_URL = os.path.join('\static', os.path.join('media', 'upload'))

# 数据库配置
DATABASE = {
    # 用户
    'USER': 'root',
    # 密码
    'PASSWORD': '123456',
    # 地址
    'HOST': '127.0.0.1',
    # 端口
    'PORT': '3306',
    # 数据库
    'DB': 'mysql',
    # 驱动
    'DRIVER': 'pymysql',
    # 数据库名称
    'NAME': 'ihome'
}
