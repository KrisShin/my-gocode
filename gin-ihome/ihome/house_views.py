import os

from flask import Blueprint, render_template, request, jsonify, session
from flask_login import login_required, LoginManager

from ihome.models import db, User, House, Facility, Area, HouseImage
from utils.settings import UPLOAD_DIR, SHOW_URL
from utils.statusCode import *

houses = Blueprint('house', __name__)
login_manager = LoginManager()


@login_manager.user_loader
def load_user(user_id):
    return User.query.get(user_id)


@houses.route('/')
def index():
    return render_template('index.html')


@houses.route('/render_my_house/')
@login_required
def render_my_house():
    user = User.query.filter(User.id == session['user_id']).first()
    house_query = House.query.filter(House.user_id == user.id).all()
    house_list = []
    for h in house_query:
        house_list.append(h.to_dict())

    return render_template('myhouse.html', house=house_list)


@houses.route('/render_detail/<id>/')
@login_required
def render_detail(id):
    house = House.query.get(id)
    return render_template('detail.html', house=house.to_full_dict())


@houses.route('/render_new_house/')
@login_required
def render_new_house():
    areas = Area.query.all()
    return render_template('newhouse.html', areas=areas)


@houses.route('/new_house/', methods=['POST'])
@login_required
def new_house():
    data = {}
    for key in request.form:
        if key != 'facility':
            data[key] = request.form.get(key)

    user = User.query.filter(User.id == session['user_id']).first()
    user_id = user.id
    data['user_id'] = user_id
    facilities = Facility.query.filter(Facility.id.in_(request.form.getlist('facility'))).all()
    data['facilities'] = facilities
    index_image_url = request.form.get('index_image_url')
    house = House(**data)
    if index_image_url:
        house.index_image_url = index_image_url
    house.add_update()
    return jsonify({'code': OK, 'house_id': house.id})


@houses.route('/new_house_img/', methods=['POST'])
@login_required
def new_house_img():
    house_id = request.form.get('house_id')
    house = House.query.filter(House.id == house_id).first()
    image = request.files.get('house_image')
    image_url = os.path.join(UPLOAD_DIR, image.filename)
    image.save(image_url)
    house.images = os.path.join('upload', image.filename)
    # 房屋默认图片
    if not house.index_image_url:
        house.index_image_url = os.path.join('upload', image.filename)
        house.add_update()
    house.add_update()
    return jsonify(SUCCESS)
