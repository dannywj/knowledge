<?php

/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
class User {
    private $_di;

    public function __construct($di) {// 构造函数接收传入的容器对象
        $this->_di = $di;
    }

    public function userPay($user_name, $amount) {
        echo $user_name . '-' . $this->_di->uPay->pay($amount);// 根据约定，调用容器对象内的方法，在更换容器内容后也不必修改代码！
//        $obj=$this->_di->uPay;
//        echo $user_name . '-' . $obj()->pay($amount);// 根据约定，调用容器对象内的方法，在更换容器内容后也不必修改代码！
    }
}