<?php
/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
require_once(__DIR__ . '/DI.php');
require_once(__DIR__ . '/PayMethod.php');
require_once(__DIR__ . '/AliPay.php');
require_once(__DIR__ . '/WechatPay.php');
require_once(__DIR__ . '/User.php');

// 容器
$di = DI::getInstince();
$di->onePay = new AliPay();// 填充容器
echo $di->onePay->pay(100);
echo '<br>';

$di->onePay = new WechatPay();// 填充容器 （更换了对象，但是都继承自同一个抽象类）
echo $di->onePay->pay(200);// 可以调用相同的方法
echo '<br>=========================<br>';

// 依赖注入实例
$di = DI::getInstince();
$di->uPay = new AliPay(); // 可变的容器内容
$user = new User($di); // 将可变的内容传入构造函数中，作为一个对象属性
$user->userPay('danny', 50);
echo '<br>';

$di->uPay = new WechatPay(); // 可变的容器内容(更换内容)
$user2 = new User($di); // 将可变的内容传入构造函数中，作为一个对象属性
$user2->userPay('danny', 60);//调用相同方法，输出不同的值


//$di = DI::getInstince();
//
//$di->uPay = function () {
//    return new AliPay();
//};
//
//var_dump('begin use object');
//$user = new User($di); // 将可变的内容传入构造函数中，作为一个对象属性
//$user->userPay('danny', 50);