from flask import Flask

from ihome.house_views import houses
from ihome.orders_views import orders
from ihome.users_views import users
from utils.functions import init_ext
from utils.settings import STATIC_DIR, TEMPLATES_DIR
from utils.config import Config


def create_app():
    app = Flask(__name__, static_folder=STATIC_DIR, template_folder=TEMPLATES_DIR)
    app.register_blueprint(blueprint=houses, url_prefix='/house')
    app.register_blueprint(blueprint=users, url_prefix='/users')
    app.register_blueprint(blueprint=orders, url_prefix='/orders')

    app.config.from_object(Config)

    init_ext(app)

    return app
