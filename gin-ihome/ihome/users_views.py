import os
import re
from random import choice

from flask import Blueprint, render_template, request, jsonify, session, redirect, url_for
from flask_login import login_user, current_user, logout_user, login_required, LoginManager

from ihome.models import User
from utils.settings import UPLOAD_DIR, SHOW_URL
from utils.statusCode import *

users = Blueprint('users', __name__)
login_manager = LoginManager()


@login_manager.user_loader
def load_user(user_id):
    return User.query.get(user_id)


@users.route('/captcha/', methods=['POST'])
def captcha():
    str = '1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZCXVBNM'
    captcha = ''
    for _ in range(4):
        captcha += choice(str)
    session['code'] = captcha
    return jsonify({'code': OK, 'captcha': captcha})


@users.route('/render_register/')
def render_register():
    if request.method == 'GET':
        return render_template('register.html')


@users.route('/register/', methods=['GET', 'POST'])
def register():
    if request.method == 'POST':
        phone = request.form.get('mobile')
        captcha = request.form.get('imgcode')
        password = request.form.get('password')
        password2 = request.form.get('password2')
        print(phone)
        user = User.query.filter(User.phone == 'phone').first()
        if not user:
            phone_rule = r'^1[345789]\d{9}$'
            if re.fullmatch(phone_rule, phone):
                if captcha == session['code']:
                    if password == password2:
                        user = User()
                        user.phone = phone
                        user.name = phone
                        user.password = password
                        user.add_update()
                        return jsonify(SUCCESS)
                else:
                    return jsonify(USER_REGISTER_WRONG_CONFIRM_PASSWORD)
            else:
                return jsonify(USER_REGISTER_PHONE_WRONG_FORMAT)
        return jsonify(USER_REGISTER_USER_IS_EXISTED)


@users.route('/render_login/')
def render_login():
    if request.method == 'GET':
        return render_template('login.html')


@users.route('/login/', methods=['POST'])
def login():
    if request.method == 'POST':
        phone = request.form.get('mobile')
        password = request.form.get('password')
        user = User.query.filter(User.phone == phone).first()
        if user:
            if user.check_pwd(password):
                login_user(user)
                return jsonify(SUCCESS)
            else:
                return jsonify(USER_LOGIN_WRONG_PASSWORD)
        else:
            return jsonify(USER_LOGIN_USER_NOT_EXIST)


@users.route('/render_my/', methods=['GET', 'POST'])
@login_required
def my_info():
    if request.method == 'GET':
        return render_template('my.html')


@users.route('/logout/')
@login_required
def logout():
    logout_user()
    return render_template('login.html')


@users.route('/render_profile/', methods=['GET'])
@login_required
def render_profile():
    if request.method == 'GET':
        return render_template('profile.html')


@users.route('/profile/', methods=['PATCH'])
@login_required
def profile():
    if request.method == 'PATCH':
        user = User.query.filter(User.id == session['user_id']).first()
        avatar = request.files.get('avatar')
        if avatar:
            avatar_path = os.path.join(UPLOAD_DIR, avatar.filename)
            avatar.save(avatar_path)
            user.avatar = os.path.join(SHOW_URL, avatar.filename)
            user.add_update()
            return jsonify(SUCCESS)
        name = request.form.get('name')
        if User.query.filter(User.name == name).first():
            return
        if name:
            user.name = name
            user.add_update()
            return jsonify(SUCCESS)


@users.route('/render_auth/', methods=['GET'])
@login_required
def render_auth():
    return render_template('auth.html')


@users.route('/auth/', methods=['POST'])
@login_required
def auth():
    user = User.query.filter(User.id == session['user_id']).first()
    id_name = request.form.get('id_name')
    id_card = request.form.get('id_card')
    user2 = User.query.filter(User.id_card == id_card).first()
    if not user2:
        if all([id_name, id_card]):
            user.id_name = id_name
            user.id_card = id_card
            user.add_update()
            return jsonify(SUCCESS)
        else:
            return jsonify(USER_AUTH_INFO_ERROR)
    else:
        return jsonify(USER_AUTH_ID_CARD_DUPLICATE)
