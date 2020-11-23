import redis
from utils.functions import get_db_uri
from utils.settings import DATABASE


class Config():

    SQLALCHEMY_DATABASE_URI = get_db_uri(DATABASE)

    SQLALCHEMY_TRACK_MODIFICATIONS = False

    SECRET_KEY = 'secret_key'

    SQLALCHEMY_POOL_SIZE = 100
    # redis数据库使用下方代码
    # SESSION_TYPE = 'redis'
    # SESSION_REDIS = redis.Redis(host='127.0.0.1', port=6379)
    # SESSION_KEY_PREFIX = 'ihome_'
