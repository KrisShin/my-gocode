import os

from flask import Blueprint, render_template, request, jsonify, session
from flask_login import login_required, LoginManager

from ihome.models import User, House, Facility, Area, HouseImage, Order
from utils.statusCode import *

orders = Blueprint('orders', __name__)
login_manager = LoginManager()


@login_manager.user_loader
def load_user(user_id):
    return User.query.get(user_id)


@orders.route('/render_booking/<int:id>/')
@login_required
def render_booking(id):
    house = House.query.get(id)
    return render_template('booking.html', house=house.to_full_dict())


@orders.route('/render_orders/')
@login_required
def render_orders():
    return render_template('orders.html')


@orders.route('/render_lorders/')
@login_required
def render_lorders():
    return render_template('lorders.html')


@orders.route('/orders_info/')
@login_required
def orders_info():
    order_list = Order.query.filter(Order.user_id == session['user_id']).all()
    orders = []
    for order in order_list:
        orders.append(order.to_dict())
    return jsonify({'code': OK, 'orders': orders})


@orders.route('/lorders_info/')
@login_required
def lorders_info():
    order_list = Order.query.filter(Order.house_id == session['user_id']).all()
    orders = []
    for order in order_list:
        orders.append(order.to_dict())
    return jsonify({'code': OK, 'orders': orders})


@orders.route('/booking/', methods=['POST'])
@login_required
def booking():
    data = {}
    data['user_id'] = session['user_id']
    data['begin_date'] = request.form.get('startDate')
    data['end_date'] = request.form.get('endDate')
    data['days'] = request.form.get('days')
    data['house_id'] = request.form.get('house_id')
    data['house_price'] = House.query.get(request.form.get('house_id')).price
    data['amount'] = data['house_price'] * data['days']
    data['status'] = 'WAIT_ACCEPT'
    order = Order(**data)
    order.add_update()
    return jsonify(SUCCESS)
