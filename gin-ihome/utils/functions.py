# from flask_session import Session
from flask_sqlalchemy import SQLAlchemy

from ihome.users_views import login_manager

db = SQLAlchemy()


# session = Session()


def init_ext(app):
    login_manager.login_view = 'users.render_login'

    db.init_app(app=app)
    login_manager.init_app(app)
    # session.init_app(app=app)


def get_db_uri(DATABASE):
    user = DATABASE.get('USER')
    password = DATABASE.get('PASSWORD')
    host = DATABASE.get('HOST')
    port = DATABASE.get('PORT')
    name = DATABASE.get('NAME')
    db = DATABASE.get('DB')
    driver = DATABASE.get('DRIVER')

    return db + '+' + driver + '://' + user + ':' + password + '@' + host + ':' + port + '/' + name
